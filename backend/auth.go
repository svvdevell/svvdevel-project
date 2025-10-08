package main

import (
    "crypto/rand"
    "encoding/base64"
    "encoding/json"
    "log"
    "net/http"
    "strings"
    "time"

    "golang.org/x/crypto/bcrypt"
)

// Учетные данные администратора (в реальном проекте хранить в БД)
const (
    ADMIN_USERNAME      = "admin"
    // Пароль: eleganceautoadminpass777!
    ADMIN_PASSWORD_HASH = "$2a$10$X8L5vZjKQJ3YHx5FqP6xfOrYN4gFqPxV6K5cPZrF7uXnGZQ2VrQea"
)

// Структуры
type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Token   string `json:"token,omitempty"`
}

type Session struct {
    Token     string
    Username  string
    CreatedAt time.Time
    ExpiresAt time.Time
}

// Хранилище сессий
var sessions = make(map[string]*Session)

// Создание таблицы для хранения сессий
func createAuthTables() {
    query := `CREATE TABLE IF NOT EXISTS admin_sessions (
        token VARCHAR(255) PRIMARY KEY,
        username VARCHAR(100) NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        expires_at DATETIME NOT NULL
    )`

    if _, err := db.Exec(query); err != nil {
        log.Printf("Error creating sessions table: %v", err)
    }
    log.Println("Auth tables initialized")
}

// Генерация токена
func generateToken() (string, error) {
    b := make([]byte, 32)
    if _, err := rand.Read(b); err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(b), nil
}

// Middleware проверки токена
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, `{"error": "Invalid token format"}`, http.StatusUnauthorized)
            return
        }

        token := parts[1]
        session, exists := sessions[token]
        if !exists {
            http.Error(w, `{"error": "Invalid token"}`, http.StatusUnauthorized)
            return
        }

        if time.Now().After(session.ExpiresAt) {
            delete(sessions, token)
            http.Error(w, `{"error": "Token expired"}`, http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    }
}

// Авторизация
func loginHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

    var req LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, `{"status":"error","message":"Невірний формат даних"}`, http.StatusBadRequest)
        return
    }

    if req.Username != ADMIN_USERNAME {
        time.Sleep(2 * time.Second)
        http.Error(w, `{"status":"error","message":"Невірний логін або пароль"}`, http.StatusUnauthorized)
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(ADMIN_PASSWORD_HASH), []byte(req.Password)); err != nil {
        time.Sleep(2 * time.Second)
        http.Error(w, `{"status":"error","message":"Невірний логін або пароль"}`, http.StatusUnauthorized)
        return
    }

    token, err := generateToken()
    if err != nil {
        http.Error(w, `{"status":"error","message":"Помилка сервера"}`, http.StatusInternalServerError)
        return
    }

    session := &Session{
        Token:     token,
        Username:  req.Username,
        CreatedAt: time.Now(),
        ExpiresAt: time.Now().Add(24 * time.Hour),
    }

    sessions[token] = session
    log.Printf("✅ Successful login for %s", req.Username)

    json.NewEncoder(w).Encode(LoginResponse{
        Status:  "success",
        Message: "Успішний вхід",
        Token:   token,
    })
}

// Выход
func logoutHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

    authHeader := r.Header.Get("Authorization")
    if authHeader != "" {
        parts := strings.Split(authHeader, " ")
        if len(parts) == 2 && parts[0] == "Bearer" {
            token := parts[1]
            delete(sessions, token)
        }
    }

    json.NewEncoder(w).Encode(map[string]string{
        "status":  "success",
        "message": "Вихід успішний",
    })
}

// Проверка токена
func verifyTokenHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        http.Error(w, `{"error":"Unauthorized"}`, http.StatusUnauthorized)
        return
    }

    parts := strings.Split(authHeader, " ")
    if len(parts) != 2 || parts[0] != "Bearer" {
        http.Error(w, `{"error":"Invalid token format"}`, http.StatusUnauthorized)
        return
    }

    token := parts[1]
    session, exists := sessions[token]
    if !exists {
        http.Error(w, `{"error":"Invalid token"}`, http.StatusUnauthorized)
        return
    }

    if time.Now().After(session.ExpiresAt) {
        delete(sessions, token)
        http.Error(w, `{"error":"Token expired"}`, http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(map[string]interface{}{
        "status":   "success",
        "username": session.Username,
        "expires":  session.ExpiresAt.Format(time.RFC3339),
    })
}

// Очистка просроченных сессий
func cleanupExpiredSessions() {
    go func() {
        for {
            time.Sleep(10 * time.Minute)
            for token, session := range sessions {
                if time.Now().After(session.ExpiresAt) {
                    delete(sessions, token)
                    log.Printf("🧹 Deleted expired session for %s", session.Username)
                }
            }
        }
    }()
}
