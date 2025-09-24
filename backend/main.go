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

    _ "github.com/mattn/go-sqlite3"
    "github.com/gorilla/mux"
)

type Response struct {
    Message string `json:"message"`
    Status  string `json:"status"`
    Version string `json:"version"`
}

type CarRequest struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    CarBrand    string    `json:"carBrand"`
    Phone       string    `json:"phone"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"createdAt"`
}

// Структуры для API ответов
type AdminRequestResponse struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    CarBrand    string    `json:"carBrand"`
    Phone       string    `json:"phone"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"createdAt"`
    ImageCount  int       `json:"imageCount"`
}

type CarImageResponse struct {
    ID           int       `json:"id"`
    CarRequestID int       `json:"carRequestId"`
    FileName     string    `json:"fileName"`
    FileURL      string    `json:"fileUrl"`
    FileSize     int       `json:"fileSize"`
    CreatedAt    time.Time `json:"createdAt"`
}

type RequestDetailResponse struct {
    ID          int                `json:"id"`
    Name        string             `json:"name"`
    CarBrand    string             `json:"carBrand"`
    Phone       string             `json:"phone"`
    Description string             `json:"description"`
    CreatedAt   time.Time          `json:"createdAt"`
    Images      []CarImageResponse `json:"images"`
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
    r.HandleFunc("/api/admin/requests/{id}", getAdminRequestDetailHandler).Methods("GET")
    r.HandleFunc("/api/admin/requests/{id}/images", getRequestImagesHandler).Methods("GET")
    
    // Статические файлы для изображений
    r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads/"))))
    
    log.Println("🚀 Go API server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
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
            phone VARCHAR(20) NOT NULL,
            description TEXT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )`,
        `CREATE TABLE IF NOT EXISTS car_images (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            car_request_id INTEGER REFERENCES car_requests(id) ON DELETE CASCADE,
            file_name VARCHAR(255) NOT NULL,
            file_path VARCHAR(500) NOT NULL,
            file_size INTEGER,
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

    // Ограничиваем размер загружаемых файлов (50MB)
    r.ParseMultipartForm(50 << 20)

    // Получаем данные из формы
    name := strings.TrimSpace(r.FormValue("name"))
    carBrand := strings.TrimSpace(r.FormValue("carBrand"))
    phone := strings.TrimSpace(r.FormValue("phone"))
    description := strings.TrimSpace(r.FormValue("description"))

    // Валидация обязательных полей
    if name == "" || carBrand == "" || phone == "" {
        http.Error(w, `{"error": "Обязательные поля: name, carBrand, phone"}`, http.StatusBadRequest)
        return
    }

    // Вставляем заявку в базу данных SQLite
    query := `INSERT INTO car_requests (name, car_brand, phone, description) 
              VALUES (?, ?, ?, ?)`
    
    result, err := db.Exec(query, name, carBrand, phone, description)
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

    // Обработка загруженных изображений
    uploadedFiles := []string{}
    if files := r.MultipartForm.File["images"]; len(files) > 0 {
        for i, fileHeader := range files {
            if i >= 10 { // Ограничиваем максимум 10 файлов
                break
            }

            file, err := fileHeader.Open()
            if err != nil {
                log.Printf("Error opening uploaded file: %v", err)
                continue
            }
            defer file.Close()

            // Проверяем размер файла (максимум 5MB)
            if fileHeader.Size > 5<<20 {
                log.Printf("File too large: %s", fileHeader.Filename)
                continue
            }

            // Генерируем уникальное имя файла
            ext := filepath.Ext(fileHeader.Filename)
            fileName := fmt.Sprintf("car_%d_image_%d_%d%s", carRequestID, i+1, time.Now().Unix(), ext)
            filePath := filepath.Join("uploads", fileName)

            // Создаем файл на диске
            dst, err := os.Create(filePath)
            if err != nil {
                log.Printf("Error creating file: %v", err)
                continue
            }
            defer dst.Close()

            // Копируем содержимое
            if _, err := io.Copy(dst, file); err != nil {
                log.Printf("Error saving file: %v", err)
                os.Remove(filePath) // Удаляем файл при ошибке
                continue
            }

            // Сохраняем информацию о файле в базе данных
            imageQuery := `INSERT INTO car_images (car_request_id, file_name, file_path, file_size) 
                          VALUES (?, ?, ?, ?)`
            _, err = db.Exec(imageQuery, carRequestID, fileHeader.Filename, filePath, fileHeader.Size)
            if err != nil {
                log.Printf("Error saving image info to database: %v", err)
                os.Remove(filePath) // Удаляем файл при ошибке сохранения в БД
                continue
            }

            uploadedFiles = append(uploadedFiles, fileName)
        }
    }

    // Возвращаем успешный ответ
    response := map[string]interface{}{
        "message": "Заявка успешно создана",
        "status": "success",
        "data": map[string]interface{}{
            "id": carRequestID,
            "name": name,
            "carBrand": carBrand,
            "phone": phone,
            "description": description,
            "uploadedImages": len(uploadedFiles),
        },
    }

    log.Printf("New car request created: ID=%d, Name=%s, Brand=%s, Images=%d", 
               carRequestID, name, carBrand, len(uploadedFiles))

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
    query := `SELECT id, name, car_brand, phone, description, created_at 
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
        err := rows.Scan(&req.ID, &req.Name, &req.CarBrand, &req.Phone, &req.Description, &req.CreatedAt)
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
        SELECT cr.id, cr.name, cr.car_brand, cr.phone, cr.description, cr.created_at,
               COUNT(ci.id) as image_count
        FROM car_requests cr
        LEFT JOIN car_images ci ON cr.id = ci.car_request_id
        GROUP BY cr.id, cr.name, cr.car_brand, cr.phone, cr.description, cr.created_at
        ORDER BY cr.created_at DESC
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
        err := rows.Scan(&req.ID, &req.Name, &req.CarBrand, &req.Phone, 
                        &req.Description, &req.CreatedAt, &req.ImageCount)
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

// Получение детальной информации о заявке с изображениями
func getAdminRequestDetailHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    vars := mux.Vars(r)
    requestID := vars["id"]
    
    // Получаем основную информацию о заявке
    var req RequestDetailResponse
    query := "SELECT id, name, car_brand, phone, description, created_at FROM car_requests WHERE id = ?"
    err := db.QueryRow(query, requestID).Scan(&req.ID, &req.Name, &req.CarBrand, 
                                             &req.Phone, &req.Description, &req.CreatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, `{"error": "Заявка не найдена"}`, http.StatusNotFound)
        } else {
            log.Printf("Error getting request detail: %v", err)
            http.Error(w, `{"error": "Ошибка получения данных"}`, http.StatusInternalServerError)
        }
        return
    }
    
    // Получаем изображения
    imageQuery := "SELECT id, car_request_id, file_name, file_path, file_size, created_at FROM car_images WHERE car_request_id = ?"
    imageRows, err := db.Query(imageQuery, requestID)
    if err != nil {
        log.Printf("Error getting request images: %v", err)
    } else {
        defer imageRows.Close()
        
        for imageRows.Next() {
            var img CarImageResponse
            var filePath string
            err := imageRows.Scan(&img.ID, &img.CarRequestID, &img.FileName, 
                                 &filePath, &img.FileSize, &img.CreatedAt)
            if err != nil {
                log.Printf("Error scanning image row: %v", err)
                continue
            }
            
            // Создаем URL для изображения
            fileName := filepath.Base(filePath)
            img.FileURL = "/uploads/" + fileName
            
            req.Images = append(req.Images, img)
        }
    }
    
    response := map[string]interface{}{
        "status": "success",
        "data":   req,
    }
    
    json.NewEncoder(w).Encode(response)
}

// Получение только изображений для конкретной заявки
func getRequestImagesHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    vars := mux.Vars(r)
    requestID := vars["id"]
    
    // Проверяем что заявка существует
    var exists int
    err := db.QueryRow("SELECT COUNT(*) FROM car_requests WHERE id = ?", requestID).Scan(&exists)
    if err != nil || exists == 0 {
        http.Error(w, `{"error": "Заявка не найдена"}`, http.StatusNotFound)
        return
    }
    
    // Получаем изображения
    query := "SELECT id, car_request_id, file_name, file_path, file_size, created_at FROM car_images WHERE car_request_id = ? ORDER BY created_at ASC"
    rows, err := db.Query(query, requestID)
    if err != nil {
        log.Printf("Error getting images: %v", err)
        http.Error(w, `{"error": "Ошибка получения изображений"}`, http.StatusInternalServerError)
        return
    }
    defer rows.Close()
    
    var images []CarImageResponse
    for rows.Next() {
        var img CarImageResponse
        var filePath string
        err := rows.Scan(&img.ID, &img.CarRequestID, &img.FileName, 
                        &filePath, &img.FileSize, &img.CreatedAt)
        if err != nil {
            log.Printf("Error scanning image row: %v", err)
            continue
        }
        
        // Создаем URL для изображения
        fileName := filepath.Base(filePath)
        img.FileURL = "/uploads/" + fileName
        
        images = append(images, img)
    }
    
    response := map[string]interface{}{
        "status": "success",
        "data":   images,
        "count":  len(images),
    }
    
    json.NewEncoder(w).Encode(response)
}