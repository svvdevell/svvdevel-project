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
    ADMIN_PASSWORD_HASH = "$2a$10$rEkN9X2vB5aH.Ux8gVzKxOqP7Y6ZC3mJ8nQwRtDfLsVpKhGjMuNxS"
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
// Авторизація з детальним логуванням для дебагу
func loginHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

    log.Println("🔍 Login attempt received")
    log.Printf("Content-Type: %s", r.Header.Get("Content-Type"))
    
    var req LoginRequest
    
    // Перевіряємо Content-Type
    contentType := r.Header.Get("Content-Type")
    
    if strings.Contains(contentType, "application/x-www-form-urlencoded") {
        log.Println("📋 Parsing form data...")
        if err := r.ParseForm(); err != nil {
            log.Printf("❌ Form parse error: %v", err)
            http.Error(w, `{"status":"error","message":"Невірний формат даних"}`, http.StatusBadRequest)
            return
        }
        
        req.Username = r.FormValue("username")
        req.Password = r.FormValue("password")
        log.Printf("📝 Form data - Username: '%s', Password length: %d", req.Username, len(req.Password))
    } else {
        log.Println("📋 Parsing JSON data...")
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            log.Printf("❌ JSON parse error: %v", err)
            http.Error(w, `{"status":"error","message":"Невірний формат даних"}`, http.StatusBadRequest)
            return
        }
        log.Printf("📝 JSON data - Username: '%s', Password length: %d", req.Username, len(req.Password))
    }

    req.Username = strings.TrimSpace(req.Username)
    req.Password = strings.TrimSpace(req.Password)

    // Валідація
    if req.Username == "" || req.Password == "" {
        log.Println("❌ Empty username or password")
        http.Error(w, `{"status":"error","message":"Логін та пароль обов'язкові"}`, http.StatusBadRequest)
        return
    }

    // Перевірка username
    log.Printf("🔍 Comparing usernames: received='%s', expected='%s'", req.Username, ADMIN_USERNAME)
    if req.Username != ADMIN_USERNAME {
        log.Println("❌ Username mismatch")
        time.Sleep(2 * time.Second)
        http.Error(w, `{"status":"error","message":"Невірний логін або пароль"}`, http.StatusUnauthorized)
        return
    }

    // Перевірка пароля
    log.Println("🔐 Checking password hash...")
    log.Printf("Password to check: '%s'", req.Password)
    log.Printf("Hash in constant: %s", ADMIN_PASSWORD_HASH)
    
    err := bcrypt.CompareHashAndPassword([]byte(ADMIN_PASSWORD_HASH), []byte(req.Password))
    if err != nil {
        log.Printf("❌ Password hash mismatch: %v", err)
        time.Sleep(2 * time.Second)
        http.Error(w, `{"status":"error","message":"Невірний логін або пароль"}`, http.StatusUnauthorized)
        return
    }

    log.Println("✅ Password verified successfully")

    token, err := generateToken()
    if err != nil {
        log.Printf("❌ Token generation error: %v", err)
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
    log.Printf("✅ Successful login for %s, token: %s", req.Username, token[:10]+"...")

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
