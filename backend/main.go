package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "time"
    "bytes"

    _ "github.com/mattn/go-sqlite3"
    "github.com/gorilla/mux"
)

// Структуры для Telegram API
type TelegramMessage struct {
    ChatID    string `json:"chat_id"`
    Text      string `json:"text"`
    ParseMode string `json:"parse_mode"`
}

type Response struct {
    Message string `json:"message"`
    Status  string `json:"status"`
    Version string `json:"version"`
}

type CarRequest struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    CarBrand    string    `json:"carBrand"`
    CarModel    string    `json:"carModel"`
    CarYear     string    `json:"carYear"`
    CarTrans    string    `json:"carTrans"`
    Phone       string    `json:"phone"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"createdAt"`
}

// Структуры для API ответов
type AdminRequestResponse struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    CarBrand    string    `json:"carBrand"`
    CarModel    string    `json:"carModel"`
    CarYear     string    `json:"carYear"`
    CarTrans    string    `json:"carTrans"`
    Phone       string    `json:"phone"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"createdAt"`
}

type RequestDetailResponse struct {
    ID          int                `json:"id"`
    Name        string             `json:"name"`
    CarBrand    string             `json:"carBrand"`
    CarModel    string             `json:"carModel"`
    CarYear     string             `json:"carYear"`
    CarTrans    string             `json:"carTrans"`
    Phone       string             `json:"phone"`
    Description string             `json:"description"`
    CreatedAt   time.Time          `json:"createdAt"`
}

var db *sql.DB

func main() {
    // Подключение к SQLite БД
    dbPath := os.Getenv("DB_PATH")
    if dbPath == "" {
        dbPath = "./data/database.db"
    }
    
    // Создаем папку для БД если её нет
    if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
        log.Printf("Error creating database directory: %v", err)
    }

    var err error
    db, err = sql.Open("sqlite3", dbPath)
    if err != nil {
        log.Printf("Database connection error: %v", err)
    } else {
        log.Printf("SQLite database connected successfully: %s", dbPath)
        // Создаем таблицы при подключении
        createTables()
        createCarTables()
    }

    // Создаем папку для загрузок если её нет
    if err := os.MkdirAll("uploads", 0755); err != nil {
        log.Printf("Error creating uploads directory: %v", err)
    }

    r := mux.NewRouter()
    
    // Добавляем CORS middleware
    r.Use(corsMiddleware)
    
    // Основные роуты
    r.HandleFunc("/health", healthHandler).Methods("GET")
    r.HandleFunc("/api/status", statusHandler).Methods("GET")
    r.HandleFunc("/api/cars", createCarRequestHandler).Methods("POST")
    r.HandleFunc("/api/cars", getCarRequestsHandler).Methods("GET")
    
    // Админ API роуты
    r.HandleFunc("/api/admin/requests", getAdminRequestsHandler).Methods("GET")
    registerCarRoutes(r)
    // Статические файлы для изображений
    r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads/"))))
    
    log.Println("🚀 Go API server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

var (
    telegramBotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
    telegramChatID   = os.Getenv("TELEGRAM_CHAT_ID")
)

// Функция отправки текстового сообщения в Telegram
func sendTelegramMessage(message string) error {
    if telegramBotToken == "" || telegramChatID == "" {
        log.Println("Telegram not configured - skipping notification")
        return nil
    }

    telegramURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramBotToken)
    
    payload := TelegramMessage{
        ChatID:    telegramChatID,
        Text:      message,
        ParseMode: "HTML",
    }
    
    jsonData, err := json.Marshal(payload)
    if err != nil {
        return fmt.Errorf("error marshaling telegram message: %v", err)
    }
    
    resp, err := http.Post(telegramURL, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return fmt.Errorf("error sending telegram message: %v", err)
    }
    defer resp.Body.Close()
    
    if resp.StatusCode != 200 {
        body, _ := io.ReadAll(resp.Body)
        return fmt.Errorf("telegram API error: %s", string(body))
    }
    
    log.Println("Telegram notification sent successfully")
    return nil
}

// Функция формирования красивого сообщения о новой заявке
func formatCarRequestMessage(name, carBrand, carModel, carYear, carTrans, phone, description string) string {
    cleanPhone := strings.ReplaceAll(phone, " ", "")
    cleanPhone = strings.ReplaceAll(cleanPhone, "(", "")
    cleanPhone = strings.ReplaceAll(cleanPhone, ")", "")
    cleanPhone = strings.ReplaceAll(cleanPhone, "-", "")
    
    message := fmt.Sprintf(`🚗 <b>Новая заявка на выкуп</b>

👤 <b>Имя:</b> %s
🚙 <b>Марка:</b> %s  
🚗 <b>Модель:</b> %s 
🚗 <b>Год:</b> %s 
🚗 <b>Трансмиссия:</b> %s 
📞 <b>Телефон:</b> <a href="tel:%s">%s</a>`, name, carBrand, carModel, carYear, carTrans, cleanPhone, phone)

    if description != "" {
        message += fmt.Sprintf(`
📝 <b>Описание:</b> %s`, description)
    }

    message += fmt.Sprintf(`

⏰ <b>Время:</b> %s`, time.Now().Format("02.01.2006 15:04:05"))

    return message
}

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}

// Создание таблиц для SQLite
func createTables() {
    queries := []string{
        `CREATE TABLE IF NOT EXISTS car_requests (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name VARCHAR(255) NOT NULL,
            car_brand VARCHAR(255) NOT NULL,
            car_model VARCHAR(255) NOT NULL,
            car_year VARCHAR(255) NOT NULL,
            car_trans VARCHAR(255) NOT NULL,
            phone VARCHAR(20) NOT NULL,
            description TEXT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )`,
    }

    for _, query := range queries {
        if _, err := db.Exec(query); err != nil {
            log.Printf("Error creating table: %v", err)
        }
    }
    log.Println("SQLite database tables initialized")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response := Response{
        Message: "API is running",
        Status:  "ok",
        Version: "1.0.0",
    }
    json.NewEncoder(w).Encode(response)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    dbStatus := "connected"
    if db == nil {
        dbStatus = "disconnected"
    } else if err := db.Ping(); err != nil {
        dbStatus = "error: " + err.Error()
    }
    
    response := map[string]interface{}{
        "api": "running",
        "database": dbStatus,
        "version": "1.0.0",
        "db_type": "SQLite",
    }
    json.NewEncoder(w).Encode(response)
}

// Обработчик создания заявки на авто
func createCarRequestHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Парсим только обычные поля формы
    if err := r.ParseForm(); err != nil {
        http.Error(w, `{"error": "Ошибка парсинга формы"}`, http.StatusBadRequest)
        return
    }

    // Получаем данные из формы
    name := strings.TrimSpace(r.FormValue("name"))
    carBrand := strings.TrimSpace(r.FormValue("carBrand"))
    carModel := strings.TrimSpace(r.FormValue("carModel"))
    carYear := strings.TrimSpace(r.FormValue("carYear"))
    carTrans := strings.TrimSpace(r.FormValue("carTrans"))
    phone := strings.TrimSpace(r.FormValue("phone"))
    description := strings.TrimSpace(r.FormValue("description"))

    // Валидация обязательных полей
    if name == "" || carBrand == "" || phone == "" || carModel == "" || carYear == "" || carTrans == "" {
        http.Error(w, `{"error": "Обязательные поля: name, carBrand, carModel, carYear, carTrans, phone"}`, http.StatusBadRequest)
        return
    }

    // Вставляем заявку в базу данных SQLite
    query := `INSERT INTO car_requests (name, car_brand, car_model, car_year, car_trans, phone, description) 
              VALUES (?, ?, ?, ?, ?, ?, ?)`
    
    result, err := db.Exec(query, name, carBrand, carModel, carYear, carTrans, phone, description)
    if err != nil {
        log.Printf("Error inserting car request: %v", err)
        http.Error(w, `{"error": "Ошибка сохранения данных"}`, http.StatusInternalServerError)
        return
    }

    // Получаем ID вставленной записи
    carRequestID, err := result.LastInsertId()
    if err != nil {
        log.Printf("Error getting last insert ID: %v", err)
        http.Error(w, `{"error": "Ошибка получения ID записи"}`, http.StatusInternalServerError)
        return
    }

    // Отправляем уведомление в Telegram (БЕЗ ФОТО)
    go func() {
        message := formatCarRequestMessage(name, carBrand, carModel, carYear, carTrans, phone, description)
        if err := sendTelegramMessage(message); err != nil {
            log.Printf("Error sending telegram message: %v", err)
        }
    }()

    // Возвращаем успешный ответ
    response := map[string]interface{}{
        "message": "Заявка успешно создана",
        "status": "success",
        "data": map[string]interface{}{
            "id": carRequestID,
            "name": name,
            "carBrand": carBrand,
            "carModel": carModel,
            "carYear": carYear,
            "carTrans": carTrans,
            "phone": phone,
            "description": description,
        },
    }

    log.Printf("New car request created: ID=%d, Name=%s, Brand=%s, Model=%s, Year=%s, Trans=%s", 
            carRequestID, name, carBrand, carModel, carYear, carTrans)

    json.NewEncoder(w).Encode(response)
}

// Получение всех заявок
func getCarRequestsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Получаем параметры пагинации
    page := 1
    limit := 20

    if p := r.URL.Query().Get("page"); p != "" {
        if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
            page = parsed
        }
    }

    if l := r.URL.Query().Get("limit"); l != "" {
        if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 100 {
            limit = parsed
        }
    }

    offset := (page - 1) * limit

    // Запрос заявок с пагинацией для SQLite
    query := `SELECT id, name, car_brand, car_model, car_year, car_trans, phone, description, created_at 
              FROM car_requests 
              ORDER BY created_at DESC 
              LIMIT ? OFFSET ?`

    rows, err := db.Query(query, limit, offset)
    if err != nil {
        log.Printf("Error querying car requests: %v", err)
        http.Error(w, `{"error": "Ошибка получения данных"}`, http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var requests []CarRequest
    for rows.Next() {
        var req CarRequest
        err := rows.Scan(&req.ID, &req.Name, &req.CarBrand, &req.CarModel, &req.CarYear, &req.CarTrans, &req.Phone, &req.Description, &req.CreatedAt)
        if err != nil {
            log.Printf("Error scanning row: %v", err)
            continue
        }
        requests = append(requests, req)
    }

    // Подсчет общего количества записей
    var total int
    err = db.QueryRow("SELECT COUNT(*) FROM car_requests").Scan(&total)
    if err != nil {
        log.Printf("Error counting records: %v", err)
        total = 0
    }

    response := map[string]interface{}{
        "status": "success",
        "data": requests,
        "pagination": map[string]interface{}{
            "page": page,
            "limit": limit,
            "total": total,
            "pages": (total + limit - 1) / limit,
        },
    }

    json.NewEncoder(w).Encode(response)
}

// Получение списка всех заявок для админки
func getAdminRequestsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    // Параметры пагинации
    page := 1
    limit := 50
    
    if p := r.URL.Query().Get("page"); p != "" {
        if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
            page = parsed
        }
    }
    
    if l := r.URL.Query().Get("limit"); l != "" {
        if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 100 {
            limit = parsed
        }
    }
    
    offset := (page - 1) * limit
    
    // Запрос заявок с количеством изображений
    query := `
        SELECT id, name, car_brand, car_model, car_year, car_trans, phone, description, created_at
        FROM car_requests
        ORDER BY created_at DESC
        LIMIT ? OFFSET ?
    `
    
    rows, err := db.Query(query, limit, offset)
    if err != nil {
        log.Printf("Error querying admin requests: %v", err)
        http.Error(w, `{"error": "Ошибка получения данных"}`, http.StatusInternalServerError)
        return
    }
    defer rows.Close()
    
    var requests []AdminRequestResponse
    for rows.Next() {
        var req AdminRequestResponse
        err := rows.Scan(&req.ID, &req.Name, &req.CarBrand, &req.CarModel, &req.CarYear, &req.CarTrans, &req.Phone, 
                        &req.Description, &req.CreatedAt)
        if err != nil {
            log.Printf("Error scanning admin request row: %v", err)
            continue
        }
        requests = append(requests, req)
    }
    
    // Подсчет общего количества записей
    var total int
    err = db.QueryRow("SELECT COUNT(*) FROM car_requests").Scan(&total)
    if err != nil {
        log.Printf("Error counting admin requests: %v", err)
        total = 0
    }
    
    response := map[string]interface{}{
        "status": "success",
        "data": requests,
        "pagination": map[string]interface{}{
            "page":  page,
            "limit": limit,
            "total": total,
            "pages": (total + limit - 1) / limit,
        },
    }
    
    json.NewEncoder(w).Encode(response)
}