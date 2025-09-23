package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "database/sql"
    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
)

type Response struct {
    Message string `json:"message"`
    Status  string `json:"status"`
    Version string `json:"version"`
}

var db *sql.DB

func main() {
    // Подключение к БД
    dbHost := os.Getenv("DB_HOST")
    if dbHost == "" {
        dbHost = "localhost"
    }
    
    dbUser := os.Getenv("DB_USER")
    if dbUser == "" {
        dbUser = "developer"
    }
    
    dbPassword := os.Getenv("DB_PASSWORD")
    if dbPassword == "" {
        dbPassword = "your_strong_password"
    }
    
    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "site1_db"
    }

    connStr := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"
    var err error
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Printf("Database connection error: %v", err)
    } else {
        log.Println("Database connected successfully")
    }

    r := mux.NewRouter()
    
    r.HandleFunc("/health", healthHandler).Methods("GET")
    r.HandleFunc("/api/status", statusHandler).Methods("GET")
    
    log.Println("🚀 Go API server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
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
    }
    json.NewEncoder(w).Encode(response)
}
