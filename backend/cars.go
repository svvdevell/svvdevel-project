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

    "github.com/gorilla/mux"
)

// Структуры для автомобилей в продаже
type Car struct {
    ID           int       `json:"id"`
    Brand        string    `json:"brand"`
    Model        string    `json:"model"`
    Year         int       `json:"year"`
    Color        string    `json:"color"`
    Fuel         string    `json:"fuel"`
    Transmission string    `json:"transmission"`
    Drive        string    `json:"drive"`
    Mileage      int       `json:"mileage"`
    Price        int       `json:"price"`
    Status       string    `json:"status"`
    Description  string    `json:"description"`
    CreatedAt    time.Time `json:"createdAt"`
    UpdatedAt    time.Time `json:"updatedAt"`
    ImageCount   int       `json:"imageCount"`
}

type CarImage struct {
    ID        int       `json:"id"`
    CarID     int       `json:"carId"`
    FileName  string    `json:"fileName"`
    FileURL   string    `json:"fileUrl"`
    FileSize  int       `json:"fileSize"`
    CreatedAt time.Time `json:"createdAt"`
}

type CarWithImages struct {
    Car
    Images []CarImage `json:"images"`
}

// Создание таблиц для автомобилей
func createCarTables() {
    queries := []string{
        `CREATE TABLE IF NOT EXISTS cars (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            brand VARCHAR(100) NOT NULL,
            model VARCHAR(100) NOT NULL,
            year INTEGER NOT NULL,
            color VARCHAR(50),
            fuel VARCHAR(50) NOT NULL,
            transmission VARCHAR(50) NOT NULL,
            drive VARCHAR(50) NOT NULL,
            mileage INTEGER NOT NULL,
            price INTEGER NOT NULL,
            status VARCHAR(20) DEFAULT 'active',
            description TEXT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )`,
        `CREATE TABLE IF NOT EXISTS car_images_sale (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            car_id INTEGER REFERENCES cars(id) ON DELETE CASCADE,
            file_name VARCHAR(255) NOT NULL,
            file_path VARCHAR(500) NOT NULL,
            file_size INTEGER,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )`,
    }

    for _, query := range queries {
        if _, err := db.Exec(query); err != nil {
            log.Printf("Error creating car table: %v", err)
        }
    }

    // Добавляем колонки если их нет (для существующих БД)
    db.Exec(`ALTER TABLE cars ADD COLUMN color VARCHAR(50)`)
    db.Exec(`ALTER TABLE cars ADD COLUMN status VARCHAR(20) DEFAULT 'active'`)
    db.Exec(`ALTER TABLE cars ADD COLUMN price INTEGER DEFAULT 0`)

    log.Println("Car tables initialized")
}

// Регистрация роутов для автомобилей
func registerCarRoutes(r *mux.Router) {
    r.HandleFunc("/api/cars-sale", createCarHandler).Methods("POST")
    r.HandleFunc("/api/cars-sale", getCarsHandler).Methods("GET")
    r.HandleFunc("/api/cars-sale/{id}", getCarDetailHandler).Methods("GET")
    r.HandleFunc("/api/cars-sale/{id}", updateCarHandler).Methods("PUT")
    r.HandleFunc("/api/cars-sale/{id}", deleteCarHandler).Methods("DELETE")
}

// Создание автомобиля
func createCarHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    r.ParseMultipartForm(50 << 20)

    brand := strings.TrimSpace(r.FormValue("brand"))
    model := strings.TrimSpace(r.FormValue("model"))
    yearStr := strings.TrimSpace(r.FormValue("year"))
    color := strings.TrimSpace(r.FormValue("color"))
    fuel := strings.TrimSpace(r.FormValue("fuel"))
    transmission := strings.TrimSpace(r.FormValue("transmission"))
    drive := strings.TrimSpace(r.FormValue("drive"))
    mileageStr := strings.TrimSpace(r.FormValue("mileage"))
    priceStr := strings.TrimSpace(r.FormValue("price"))
    status := strings.TrimSpace(r.FormValue("status"))
    description := strings.TrimSpace(r.FormValue("description"))

    if status == "" {
        status = "active"
    }

    if brand == "" || model == "" || yearStr == "" || fuel == "" || 
       transmission == "" || drive == "" || mileageStr == "" || priceStr == "" {
        http.Error(w, `{"error": "Все поля обязательны кроме описания и цвета"}`, http.StatusBadRequest)
        return
    }

    year, err := strconv.Atoi(yearStr)
    if err != nil || year < 1900 || year > time.Now().Year()+1 {
        http.Error(w, `{"error": "Некорректный год"}`, http.StatusBadRequest)
        return
    }

    mileage, err := strconv.Atoi(mileageStr)
    if err != nil || mileage < 0 {
        http.Error(w, `{"error": "Некорректный пробег"}`, http.StatusBadRequest)
        return
    }

    price, err := strconv.Atoi(priceStr)
    if err != nil || price < 0 {
        http.Error(w, `{"error": "Некорректная цена"}`, http.StatusBadRequest)
        return
    }

    query := `INSERT INTO cars (brand, model, year, color, fuel, transmission, drive, mileage, price, status, description) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
    
    result, err := db.Exec(query, brand, model, year, color, fuel, transmission, drive, mileage, price, status, description)
    if err != nil {
        log.Printf("Error inserting car: %v", err)
        http.Error(w, `{"error": "Ошибка сохранения автомобиля"}`, http.StatusInternalServerError)
        return
    }

    carID, err := result.LastInsertId()
    if err != nil {
        log.Printf("Error getting last insert ID: %v", err)
        http.Error(w, `{"error": "Ошибка получения ID"}`, http.StatusInternalServerError)
        return
    }

    uploadedFiles := []string{}
    if files := r.MultipartForm.File["images"]; len(files) > 0 {
        for i, fileHeader := range files {
            if i >= 10 {
                break
            }

            file, err := fileHeader.Open()
            if err != nil {
                log.Printf("Error opening uploaded file: %v", err)
                continue
            }
            defer file.Close()

            if fileHeader.Size > 5<<20 {
                log.Printf("File too large: %s", fileHeader.Filename)
                continue
            }

            ext := filepath.Ext(fileHeader.Filename)
            fileName := fmt.Sprintf("car_sale_%d_image_%d_%d%s", carID, i+1, time.Now().Unix(), ext)
            filePath := filepath.Join("uploads", fileName)

            dst, err := os.Create(filePath)
            if err != nil {
                log.Printf("Error creating file: %v", err)
                continue
            }
            defer dst.Close()

            if _, err := io.Copy(dst, file); err != nil {
                log.Printf("Error saving file: %v", err)
                os.Remove(filePath)
                continue
            }

            imageQuery := `INSERT INTO car_images_sale (car_id, file_name, file_path, file_size) 
                          VALUES (?, ?, ?, ?)`
            _, err = db.Exec(imageQuery, carID, fileHeader.Filename, filePath, fileHeader.Size)
            if err != nil {
                log.Printf("Error saving image info to database: %v", err)
                os.Remove(filePath)
                continue
            }

            uploadedFiles = append(uploadedFiles, fileName)
        }
    }

    response := map[string]interface{}{
        "message": "Автомобиль успешно добавлен",
        "status":  "success",
        "data": map[string]interface{}{
            "id":             carID,
            "brand":          brand,
            "model":          model,
            "year":           year,
            "color":          color,
            "status":         status,
            "uploadedImages": len(uploadedFiles),
        },
    }

    log.Printf("New car added: ID=%d, %s %s %d", carID, brand, model, year)
    json.NewEncoder(w).Encode(response)
}

// Получение списка автомобилей
func getCarsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

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

    query := `
        SELECT c.id, c.brand, c.model, c.year, c.color, c.fuel, c.transmission, c.drive, 
               c.mileage, c.price, c.status, c.description, c.created_at, c.updated_at,
               COUNT(ci.id) as image_count
        FROM cars c
        LEFT JOIN car_images_sale ci ON c.id = ci.car_id
        GROUP BY c.id
        ORDER BY c.created_at DESC
        LIMIT ? OFFSET ?
    `

    rows, err := db.Query(query, limit, offset)
    if err != nil {
        log.Printf("Error querying cars: %v", err)
        http.Error(w, `{"error": "Ошибка получения данных"}`, http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var cars []map[string]interface{}
    for rows.Next() {
        var car Car
        var color, status sql.NullString
        
        err := rows.Scan(&car.ID, &car.Brand, &car.Model, &car.Year, &color, &car.Fuel,
                        &car.Transmission, &car.Drive, &car.Mileage, &car.Price , &status, &car.Description,
                        &car.CreatedAt, &car.UpdatedAt, &car.ImageCount)
        if err != nil {
            log.Printf("Error scanning car row: %v", err)
            continue
        }

        // Получаем первое изображение для каждой машины
        var images []CarImage
        imageQuery := `SELECT id, car_id, file_name, file_path, file_size, created_at 
                       FROM car_images_sale WHERE car_id = ? ORDER BY created_at ASC LIMIT 1`
        
        imageRows, err := db.Query(imageQuery, car.ID)
        if err == nil {
            defer imageRows.Close()
            for imageRows.Next() {
                var img CarImage
                var filePath string
                if err := imageRows.Scan(&img.ID, &img.CarID, &img.FileName, &filePath, &img.FileSize, &img.CreatedAt); err == nil {
                    fileName := filepath.Base(filePath)
                    img.FileURL = "/uploads/" + fileName
                    images = append(images, img)
                }
            }
        }

        carMap := map[string]interface{}{
            "id":           car.ID,
            "brand":        car.Brand,
            "model":        car.Model,
            "year":         car.Year,
            "color":        color.String,
            "fuel":         car.Fuel,
            "transmission": car.Transmission,
            "drive":        car.Drive,
            "mileage":      car.Mileage,
            "price":        car.Price,
            "status":       status.String,
            "description":  car.Description,
            "createdAt":    car.CreatedAt,
            "updatedAt":    car.UpdatedAt,
            "imageCount":   car.ImageCount,
            "images":       images,
        }
        
        cars = append(cars, carMap)
    }

    var total int
    err = db.QueryRow("SELECT COUNT(*) FROM cars").Scan(&total)
    if err != nil {
        log.Printf("Error counting cars: %v", err)
        total = 0
    }

    response := map[string]interface{}{
        "status": "success",
        "data":   cars,
        "pagination": map[string]interface{}{
            "page":  page,
            "limit": limit,
            "total": total,
            "pages": (total + limit - 1) / limit,
        },
    }

    json.NewEncoder(w).Encode(response)
}

// Получение детальной информации об автомобиле
func getCarDetailHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    vars := mux.Vars(r)
    carID := vars["id"]

    var car Car
    var color, status sql.NullString
    
    query := `SELECT id, brand, model, year, color, fuel, transmission, drive, mileage, price, 
                     status, description, created_at, updated_at 
              FROM cars WHERE id = ?`
    
    err := db.QueryRow(query, carID).Scan(&car.ID, &car.Brand, &car.Model, &car.Year,
                                         &color, &car.Fuel, &car.Transmission, &car.Drive, &car.Mileage, &car.Price,
                                         &status, &car.Description, &car.CreatedAt, &car.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, `{"error": "Автомобиль не найден"}`, http.StatusNotFound)
        } else {
            log.Printf("Error getting car detail: %v", err)
            http.Error(w, `{"error": "Ошибка получения данных"}`, http.StatusInternalServerError)
        }
        return
    }

    car.Color = color.String
    car.Status = status.String

    var images []CarImage
    imageQuery := `SELECT id, car_id, file_name, file_path, file_size, created_at 
                   FROM car_images_sale WHERE car_id = ? ORDER BY created_at ASC`
    
    imageRows, err := db.Query(imageQuery, carID)
    if err != nil {
        log.Printf("Error getting car images: %v", err)
    } else {
        defer imageRows.Close()
        
        for imageRows.Next() {
            var img CarImage
            var filePath string
            err := imageRows.Scan(&img.ID, &img.CarID, &img.FileName, 
                                 &filePath, &img.FileSize, &img.CreatedAt)
            if err != nil {
                log.Printf("Error scanning image row: %v", err)
                continue
            }
            
            fileName := filepath.Base(filePath)
            img.FileURL = "/uploads/" + fileName
            images = append(images, img)
        }
    }

    result := CarWithImages{
        Car:    car,
        Images: images,
    }

    response := map[string]interface{}{
        "status": "success",
        "data":   result,
    }

    json.NewEncoder(w).Encode(response)
}

// Обновление автомобиля
func updateCarHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    vars := mux.Vars(r)
    carID := vars["id"]

    // Проверяем существование автомобиля
    var exists int
    err := db.QueryRow("SELECT COUNT(*) FROM cars WHERE id = ?", carID).Scan(&exists)
    if err != nil || exists == 0 {
        http.Error(w, `{"error": "Автомобиль не найден"}`, http.StatusNotFound)
        return
    }

    // Ограничиваем размер загружаемых файлов (50MB)
    r.ParseMultipartForm(50 << 20)

    // Получаем данные из формы
    brand := strings.TrimSpace(r.FormValue("brand"))
    model := strings.TrimSpace(r.FormValue("model"))
    yearStr := strings.TrimSpace(r.FormValue("year"))
    color := strings.TrimSpace(r.FormValue("color"))
    fuel := strings.TrimSpace(r.FormValue("fuel"))
    transmission := strings.TrimSpace(r.FormValue("transmission"))
    drive := strings.TrimSpace(r.FormValue("drive"))
    mileageStr := strings.TrimSpace(r.FormValue("mileage"))
    priceStr := strings.TrimSpace(r.FormValue("price"))
    status := strings.TrimSpace(r.FormValue("status"))
    description := strings.TrimSpace(r.FormValue("description"))

    // Валидация обязательных полей
    if brand == "" || model == "" || yearStr == "" || fuel == "" || 
       transmission == "" || drive == "" || mileageStr == "" || priceStr == "" {
        http.Error(w, `{"error": "Все поля обязательны кроме описания и цвета"}`, http.StatusBadRequest)
        return
    }

    // Конвертация строк в числа
    year, err := strconv.Atoi(yearStr)
    if err != nil || year < 1900 || year > time.Now().Year()+1 {
        http.Error(w, `{"error": "Некорректный год"}`, http.StatusBadRequest)
        return
    }

    mileage, err := strconv.Atoi(mileageStr)
    if err != nil || mileage < 0 {
        http.Error(w, `{"error": "Некорректный пробег"}`, http.StatusBadRequest)
        return
    }

    price, err := strconv.Atoi(priceStr)
    if err != nil || price < 0 {
        http.Error(w, `{"error": "Некорректная цена"}`, http.StatusBadRequest)
        return
    }
    // Обновляем данные автомобиля
    updateQuery := `UPDATE cars SET brand = ?, model = ?, year = ?, color = ?, fuel = ?, 
                    transmission = ?, drive = ?, mileage = ?, price = ?, status = ?, description = ?, 
                    updated_at = CURRENT_TIMESTAMP 
                    WHERE id = ?`
    
    _, err = db.Exec(updateQuery, brand, model, year, color, fuel, transmission, 
                     drive, mileage, price, status, description, carID)
    if err != nil {
        log.Printf("Error updating car: %v", err)
        http.Error(w, `{"error": "Ошибка обновления автомобиля"}`, http.StatusInternalServerError)
        return
    }

    // Обработка удаления изображений
    deleteImagesStr := r.FormValue("deleteImages")
    if deleteImagesStr != "" {
        imageIDs := strings.Split(deleteImagesStr, ",")
        for _, idStr := range imageIDs {
            imageID := strings.TrimSpace(idStr)
            if imageID == "" {
                continue
            }

            // Получаем путь к файлу перед удалением из БД
            var filePath string
            err := db.QueryRow("SELECT file_path FROM car_images_sale WHERE id = ? AND car_id = ?", 
                              imageID, carID).Scan(&filePath)
            if err == nil {
                // Удаляем файл с диска
                os.Remove(filePath)
                // Удаляем запись из БД
                db.Exec("DELETE FROM car_images_sale WHERE id = ? AND car_id = ?", imageID, carID)
                log.Printf("Deleted image: %s", filePath)
            }
        }
    }

    // Обработка новых загруженных изображений
    uploadedFiles := []string{}
    if files := r.MultipartForm.File["images"]; len(files) > 0 {
        // Проверяем текущее количество изображений
        var currentImageCount int
        db.QueryRow("SELECT COUNT(*) FROM car_images_sale WHERE car_id = ?", carID).Scan(&currentImageCount)

        for i, fileHeader := range files {
            if currentImageCount+i >= 10 { // Максимум 10 файлов
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
            fileName := fmt.Sprintf("car_sale_%s_image_%d_%d%s", carID, currentImageCount+i+1, time.Now().Unix(), ext)
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
                os.Remove(filePath)
                continue
            }

            // Сохраняем информацию о файле в базе данных
            imageQuery := `INSERT INTO car_images_sale (car_id, file_name, file_path, file_size) 
                          VALUES (?, ?, ?, ?)`
            _, err = db.Exec(imageQuery, carID, fileHeader.Filename, filePath, fileHeader.Size)
            if err != nil {
                log.Printf("Error saving image info to database: %v", err)
                os.Remove(filePath)
                continue
            }

            uploadedFiles = append(uploadedFiles, fileName)
        }
    }

    response := map[string]interface{}{
        "message": "Автомобиль успешно обновлен",
        "status":  "success",
        "data": map[string]interface{}{
            "id":             carID,
            "brand":          brand,
            "model":          model,
            "year":           year,
            "newImagesAdded": len(uploadedFiles),
        },
    }

    log.Printf("Car updated: ID=%s, %s %s %d", carID, brand, model, year)
    json.NewEncoder(w).Encode(response)
}

// Удаление автомобиля
func deleteCarHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    vars := mux.Vars(r)
    carID := vars["id"]

    // Проверяем существование автомобиля
    var exists int
    err := db.QueryRow("SELECT COUNT(*) FROM cars WHERE id = ?", carID).Scan(&exists)
    if err != nil || exists == 0 {
        http.Error(w, `{"error": "Автомобиль не найден"}`, http.StatusNotFound)
        return
    }

    // Получаем все изображения автомобиля
    rows, err := db.Query("SELECT file_path FROM car_images_sale WHERE car_id = ?", carID)
    if err != nil {
        log.Printf("Error querying car images: %v", err)
    } else {
        defer rows.Close()
        
        // Удаляем файлы изображений с диска
        for rows.Next() {
            var filePath string
            if err := rows.Scan(&filePath); err == nil {
                if err := os.Remove(filePath); err != nil {
                    log.Printf("Error deleting file %s: %v", filePath, err)
                } else {
                    log.Printf("Deleted file: %s", filePath)
                }
            }
        }
    }

    // Удаляем записи об изображениях из БД
    _, err = db.Exec("DELETE FROM car_images_sale WHERE car_id = ?", carID)
    if err != nil {
        log.Printf("Error deleting car images from DB: %v", err)
    }

    // Удаляем сам автомобиль
    _, err = db.Exec("DELETE FROM cars WHERE id = ?", carID)
    if err != nil {
        log.Printf("Error deleting car: %v", err)
        http.Error(w, `{"error": "Ошибка удаления автомобиля"}`, http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "message": "Автомобиль успешно удален",
        "status":  "success",
    }

    log.Printf("Car deleted: ID=%s", carID)
    json.NewEncoder(w).Encode(response)
}