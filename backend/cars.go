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
    Fuel         string    `json:"fuel"`
    Transmission string    `json:"transmission"`
    Drive        string    `json:"drive"`
    Mileage      int       `json:"mileage"`
    Description  string    `json:"description"`
    CreatedAt    time.Time `json:"createdAt"`
    UpdatedAt    time.Time `json:"updatedAt"`
    ImageCount   int       `json:"imageCount"`
}

type CarImage struct {
    ID       int       `json:"id"`
    CarID    int       `json:"carId"`
    FileName string    `json:"fileName"`
    FileURL  string    `json:"fileUrl"`
    FileSize int       `json:"fileSize"`
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
            fuel VARCHAR(50) NOT NULL,
            transmission VARCHAR(50) NOT NULL,
            drive VARCHAR(50) NOT NULL,
            mileage INTEGER NOT NULL,
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

    // Ограничиваем размер загружаемых файлов (50MB)
    r.ParseMultipartForm(50 << 20)

    // Получаем данные из формы
    brand := strings.TrimSpace(r.FormValue("brand"))
    model := strings.TrimSpace(r.FormValue("model"))
    yearStr := strings.TrimSpace(r.FormValue("year"))
    fuel := strings.TrimSpace(r.FormValue("fuel"))
    transmission := strings.TrimSpace(r.FormValue("transmission"))
    drive := strings.TrimSpace(r.FormValue("drive"))
    mileageStr := strings.TrimSpace(r.FormValue("mileage"))
    description := strings.TrimSpace(r.FormValue("description"))

    // Валидация обязательных полей
    if brand == "" || model == "" || yearStr == "" || fuel == "" || 
       transmission == "" || drive == "" || mileageStr == "" {
        http.Error(w, `{"error": "Все поля обязательны кроме описания"}`, http.StatusBadRequest)
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

    // Вставляем автомобиль в базу данных
    query := `INSERT INTO cars (brand, model, year, fuel, transmission, drive, mileage, description) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
    
    result, err := db.Exec(query, brand, model, year, fuel, transmission, drive, mileage, description)
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

    // Обработка загруженных изображений
    uploadedFiles := []string{}
    if files := r.MultipartForm.File["images"]; len(files) > 0 {
        for i, fileHeader := range files {
            if i >= 10 { // Максимум 10 файлов
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
            fileName := fmt.Sprintf("car_sale_%d_image_%d_%d%s", carID, i+1, time.Now().Unix(), ext)
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
        "message": "Автомобиль успешно добавлен",
        "status":  "success",
        "data": map[string]interface{}{
            "id":             carID,
            "brand":          brand,
            "model":          model,
            "year":           year,
            "fuel":           fuel,
            "transmission":   transmission,
            "drive":          drive,
            "mileage":        mileage,
            "description":    description,
            "uploadedImages": len(uploadedFiles),
        },
    }

    log.Printf("New car added: ID=%d, %s %s %d", carID, brand, model, year)
    json.NewEncoder(w).Encode(response)
}

// Получение списка автомобилей
func getCarsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Параметры пагинации
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

    // Запрос автомобилей с количеством изображений
    query := `
        SELECT c.id, c.brand, c.model, c.year, c.fuel, c.transmission, c.drive, 
               c.mileage, c.description, c.created_at, c.updated_at,
               COUNT(ci.id) as image_count
        FROM cars c
        LEFT JOIN car_images_sale ci ON c.id = ci.car_id
        GROUP BY c.id, c.brand, c.model, c.year, c.fuel, c.transmission, c.drive, 
                 c.mileage, c.description, c.created_at, c.updated_at
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

    var cars []Car
    for rows.Next() {
        var car Car
        err := rows.Scan(&car.ID, &car.Brand, &car.Model, &car.Year, &car.Fuel,
                        &car.Transmission, &car.Drive, &car.Mileage, &car.Description,
                        &car.CreatedAt, &car.UpdatedAt, &car.ImageCount)
        if err != nil {
            log.Printf("Error scanning car row: %v", err)
            continue
        }
        cars = append(cars, car)
    }

    // Подсчет общего количества
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

    // Получаем основную информацию об автомобиле
    var car Car
    query := `SELECT id, brand, model, year, fuel, transmission, drive, mileage, 
                     description, created_at, updated_at 
              FROM cars WHERE id = ?`
    
    err := db.QueryRow(query, carID).Scan(&car.ID, &car.Brand, &car.Model, &car.Year,
                                         &car.Fuel, &car.Transmission, &car.Drive, &car.Mileage,
                                         &car.Description, &car.CreatedAt, &car.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, `{"error": "Автомобиль не найден"}`, http.StatusNotFound)
        } else {
            log.Printf("Error getting car detail: %v", err)
            http.Error(w, `{"error": "Ошибка получения данных"}`, http.StatusInternalServerError)
        }
        return
    }

    // Получаем изображения
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

// Обновление автомобиля (заглушка)
func updateCarHandler(w http.ResponseWriter, r *http.Request) {
    http.Error(w, `{"error": "Функция в разработке"}`, http.StatusNotImplemented)
}

// Удаление автомобиля (заглушка)
func deleteCarHandler(w http.ResponseWriter, r *http.Request) {
    http.Error(w, `{"error": "Функция в разработке"}`, http.StatusNotImplemented)
}