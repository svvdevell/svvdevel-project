# 🚗 Elegance Auto - Документация проекта

## 📋 Общее описание
Веб-приложение для автосалона "Elegance Auto" с функционалом:
- Прием заявок на выкуп автомобилей
- Каталог автомобилей в продаже
- Административная панель для управления
- Уведомления в Telegram

**Домен:** https://eleganceauto.od.ua  
**Сервер:** VPS на Ubuntu 24, IP: 91.99.231.3  
**Расположение:** `/root/svvdevel-project`

---

## 🏗️ Архитектура проекта

### Структура директорий
```
/root/svvdevel-project/
├── backend/              # Go API сервер
│   ├── main.go          # Основной файл (роуты, заявки)
│   ├── cars.go          # Логика автомобилей на продажу
│   ├── auth.go          # Аутентификация и сессии
│   ├── data/            # База данных SQLite
│   │   └── database.db  # Файл БД (сохраняется между деплоями)
│   └── uploads/         # Загруженные изображения
├── frontend/            # Vue.js приложение
│   ├── src/
│   ├── package.json
│   ├── vite.config.js
│   ├── Dockerfile
│   └── nginx.conf       # Nginx для фронтенд контейнера
├── .github/
│   └── workflows/
│       └── deploy.yml   # CI/CD конфигурация
├── docker-compose.yml   # Оркестрация контейнеров
├── nginx.conf           # Главный Nginx (проксирование)
└── .env                 # Переменные окружения (не в Git)
```

---

## 🔧 Технологический стек

### Backend (Go)
- **Язык:** Go 1.20+
- **Фреймворк:** Gorilla Mux (роутинг)
- **База данных:** SQLite3
- **Библиотеки:**
  - `github.com/gorilla/mux` - HTTP роутер
  - `github.com/mattn/go-sqlite3` - SQLite драйвер
  - `golang.org/x/crypto/bcrypt` - Хеширование паролей
- **Порт контейнера:** 8080 (маппится на 8001 хоста)

### Frontend (Vue.js)
- **Фреймворк:** Vue 3.3
- **Сборщик:** Vite 4.0
- **State Management:** Pinia 3.0
- **Роутинг:** Vue Router 4.5
- **Стилизация:** SASS 1.93
- **Библиотеки:**
  - `@vitejs/plugin-vue` - Vue plugin для Vite
- **Порт контейнера:** 80 (маппится на 3001 хоста)

### DevOps
- **Контейнеризация:** Docker + Docker Compose
- **Web-сервер:** Nginx Alpine
- **SSL:** Let's Encrypt (certbot)
- **CI/CD:** GitHub Actions
- **Деплой:** SSH с помощью `appleboy/ssh-action`

---

## 🗄️ База данных (SQLite)

### Таблицы

#### 1. `car_requests` - Заявки на выкуп
```sql
id              INTEGER PRIMARY KEY AUTOINCREMENT
name            VARCHAR(255) - Имя клиента
car_brand       VARCHAR(255) - Марка автомобиля
car_model       VARCHAR(255) - Модель
car_year        VARCHAR(255) - Год выпуска
car_trans       VARCHAR(255) - Трансмиссия
phone           VARCHAR(20) - Телефон
description     TEXT - Дополнительное описание
created_at      DATETIME DEFAULT CURRENT_TIMESTAMP
```

#### 2. `cars` - Автомобили в продаже
```sql
id              INTEGER PRIMARY KEY AUTOINCREMENT
brand           VARCHAR(100) - Марка
model           VARCHAR(100) - Модель
year            INTEGER - Год
color           VARCHAR(50) - Цвет
fuel            VARCHAR(50) - Тип топлива
transmission    VARCHAR(50) - Трансмиссия
drive           VARCHAR(50) - Привод
mileage         INTEGER - Пробег
price           INTEGER - Цена
volume          INTEGER - Объем двигателя
status          VARCHAR(20) DEFAULT 'active' - Статус
description     TEXT - Описание
created_at      DATETIME DEFAULT CURRENT_TIMESTAMP
updated_at      DATETIME DEFAULT CURRENT_TIMESTAMP
```

#### 3. `car_images_sale` - Фото автомобилей на продажу
```sql
id              INTEGER PRIMARY KEY AUTOINCREMENT
car_id          INTEGER REFERENCES cars(id) ON DELETE CASCADE
file_name       VARCHAR(255) - Оригинальное имя файла
file_path       VARCHAR(500) - Путь к файлу
file_size       INTEGER - Размер файла
created_at      DATETIME DEFAULT CURRENT_TIMESTAMP
```

#### 4. `car_images` - Фото из заявок (таблица существует, но не используется)

#### 5. `admin_sessions` - Сессии администраторов
```sql
token           VARCHAR(255) PRIMARY KEY
username        VARCHAR(100)
created_at      DATETIME DEFAULT CURRENT_TIMESTAMP
expires_at      DATETIME
```

**Важно:** База данных хранится в `/root/svvdevel-project/backend/data/database.db` и **сохраняется между деплоями** благодаря Docker volumes.

---

## 🌐 API Endpoints

### Публичные роуты
- `GET /health` - Health check
- `GET /api/status` - Статус API
- `POST /api/cars` - Создать заявку на выкуп
- `GET /api/cars` - Получить заявки (с пагинацией)
- `GET /api/cars-sale` - Список авто на продажу
- `GET /api/cars-sale/{id}` - Детали авто

### Защищенные роуты (требуют Bearer token)
- `POST /api/cars-sale` - Добавить авто на продажу
- `PUT /api/cars-sale/{id}` - Обновить авто
- `DELETE /api/cars-sale/{id}` - Удалить авто

### Аутентификация
- `POST /api/auth/login` - Вход в систему
- `POST /api/auth/logout` - Выход
- `GET /api/auth/verify` - Проверка токена

### Админ панель
- `GET /api/admin/requests` - Все заявки для админки

### Статические файлы
- `/uploads/*` - Загруженные изображения

---

## 🔐 Аутентификация

### Учетные данные администратора
- **Username:** `admin`
- **Password:** `eleganceautoadminpass777!`
- **Hash:** `$2b$10$gytUP2mQ2vEuGCZ04JmX7O1uv4JM5/dhKR.QwHnhy4PmiWIp2tyNm`

### Механизм
- Bearer Token в заголовке `Authorization: Bearer <token>`
- Сессии хранятся в памяти (map) на время работы контейнера
- Время жизни токена: 24 часа
- Автоматическая очистка истекших сессий каждые 10 минут

---

## 📬 Telegram интеграция

### Переменные окружения
```bash
TELEGRAM_BOT_TOKEN=<токен бота>
TELEGRAM_CHAT_ID=<ID чата>
```

### Функционал
При создании заявки на выкуп автоматически отправляется уведомление в Telegram с:
- Именем клиента
- Маркой и моделью авто
- Годом и трансмиссией
- Телефоном (кликабельная ссылка)
- Описанием
- Временем создания

---

## 🐳 Docker конфигурация

### Контейнеры

#### 1. Backend (`svvdevel_backend`)
```yaml
build: ./backend
ports: "8001:8080"
environment:
  - DB_PATH=/app/data/database.db
  - TELEGRAM_BOT_TOKEN=${TELEGRAM_BOT_TOKEN}
  - TELEGRAM_CHAT_ID=${TELEGRAM_CHAT_ID}
volumes:
  - ./backend/data:/app/data         # ✅ БД сохраняется
  - ./backend/uploads:/app/uploads   # ✅ Файлы сохраняются
```

#### 2. Frontend (`svvdevel_frontend`)
```yaml
build: ./frontend
ports: "3001:80"
```

#### 3. Nginx (`svvdevel_nginx`)
```yaml
image: nginx:alpine
ports: 
  - "80:80"
  - "443:443"
volumes:
  - ./nginx.conf:/etc/nginx/nginx.conf
  - /etc/letsencrypt:/etc/letsencrypt:ro
```

### Сеть
- Тип: Bridge
- Имя: `app-network`
- Все контейнеры в одной сети для взаимодействия

---

## 🔄 CI/CD Pipeline (GitHub Actions)

### Триггеры
- Push в ветку `main`
- Ручной запуск (`workflow_dispatch`)

### Процесс деплоя
1. **Checkout code** - Клонирование репозитория
2. **SSH Deploy:**
   - Подключение к серверу через SSH
   - `git reset --hard origin/main` - Жесткий сброс на последний коммит
   - Экспорт переменных окружения из GitHub Secrets
   - `docker-compose down` - Остановка контейнеров
   - Удаление старых образов
   - `docker-compose build --no-cache --pull` - Пересборка
   - `docker-compose up -d` - Запуск
   - Ожидание 15 секунд
   - Вывод логов и статуса контейнеров
   - Очистка неиспользуемых образов
3. **Verify deployment** - Проверка доступности сайта

### GitHub Secrets
```
SSH_HOST=91.99.231.3
SSH_USER=root
SSH_PRIVATE_KEY=<приватный SSH ключ>
TELEGRAM_BOT_TOKEN=<токен>
TELEGRAM_CHAT_ID=<ID>
```

---

## 🌐 Nginx конфигурация

### Корневой Nginx (главный прокси)
```nginx
# HTTP -> HTTPS редирект
server {
    listen 80;
    return 301 https://$server_name$request_uri;
}

# HTTPS сервер
server {
    listen 443 ssl http2;
    server_name eleganceauto.od.ua;
    
    # SSL сертификаты Let's Encrypt
    ssl_certificate /etc/letsencrypt/live/eleganceauto.od.ua/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/eleganceauto.od.ua/privkey.pem;
    
    # Проксирование /api/* на backend:8080
    location /api/ {
        proxy_pass http://backend:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
    
    # Проксирование /uploads/* на backend:8080
    location /uploads/ {
        proxy_pass http://backend:8080;
    }
    
    # Проксирование / на frontend:80
    location / {
        proxy_pass http://frontend:80;
        proxy_set_header Host $host;
    }
}
```

### Frontend Nginx (внутри контейнера)
```nginx
server {
    listen 80;
    root /usr/share/nginx/html;
    index index.html;
    
    # Кеширование статики
    location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg|woff|woff2|ttf|eot)$ {
        expires 1y;
        add_header Cache-Control "public, immutable";
    }
    
    # SPA fallback
    location / {
        try_files $uri $uri/ /index.html;
    }
}
```

---

## 🛡️ Безопасность

### Защита от спама
- **Rate Limiting:** Максимум 5 заявок с одного IP за 10 минут
- **Time-based validation:** Форма должна быть заполнена минимум за 3 секунды
- **IP tracking:** Хранение истории запросов по IP

### CORS
```go
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
Access-Control-Allow-Headers: Content-Type, Authorization
```

### SSL/TLS
- Сертификаты Let's Encrypt
- TLS 1.2, 1.3
- Автоматический редирект с HTTP на HTTPS

---

## 📦 Загрузка файлов

### Ограничения
- **Максимум файлов:** 10 изображений на автомобиль
- **Размер файла:** До 5MB на изображение
- **Общий размер формы:** До 50MB
- **Форматы:** JPEG, PNG, GIF, WEBP и другие

### Хранение
- Папка: `/root/svvdevel-project/backend/uploads/`
- Именование: `car_sale_{car_id}_image_{index}_{timestamp}.{ext}`
- Привязка к БД через таблицу `car_images_sale`
- При удалении авто - файлы автоматически удаляются с диска

---

## 🚀 Команды для работы с проектом

### На сервере

#### Просмотр логов
```bash
cd /root/svvdevel-project
docker-compose logs -f backend    # Логи бэкенда
docker-compose logs -f frontend   # Логи фронтенда
docker-compose logs -f nginx      # Логи nginx
docker-compose logs --tail=50     # Последние 50 строк всех логов
```

#### Управление контейнерами
```bash
docker-compose ps                 # Статус контейнеров
docker-compose up -d              # Запуск
docker-compose down               # Остановка
docker-compose restart            # Перезапуск
docker-compose build --no-cache   # Пересборка без кеша
```

#### Работа с базой данных
```bash
# Подключение к базе
docker exec -it svvdevel_backend sh -c "sqlite3 /app/data/database.db"

# Внутри sqlite3:
.tables                           # Список таблиц
.schema cars                      # Структура таблицы
SELECT * FROM cars;               # Все записи
SELECT COUNT(*) FROM car_requests; # Количество заявок
.quit                             # Выход

# Или одной командой:
docker exec -it svvdevel_backend sh -c "sqlite3 /app/data/database.db 'SELECT COUNT(*) FROM cars;'"
```

#### Очистка данных
```bash
# Очистить таблицу cars
docker exec -it svvdevel_backend sh -c "sqlite3 /app/data/database.db 'DELETE FROM cars;'"

# Очистить изображения
docker exec -it svvdevel_backend sh -c "sqlite3 /app/data/database.db 'DELETE FROM car_images_sale;'"

# Сбросить автоинкремент
docker exec -it svvdevel_backend sh -c "sqlite3 /app/data/database.db 'DELETE FROM sqlite_sequence WHERE name=\"cars\";'"

# Удалить файлы изображений
rm -rf backend/uploads/*
```

#### Проверка сайта
```bash
curl -I https://eleganceauto.od.ua/        # Проверка HTTPS
curl http://backend:8080/health            # Health check (внутри сети)
docker exec -it svvdevel_nginx curl -I http://backend:8080/api/status
```

### Git операции
```bash
git status                        # Статус изменений
git add .                         # Добавить все файлы
git commit -m "message"           # Коммит
git push origin main              # Отправить на GitHub (запустит деплой)
git pull origin main              # Обновить код
```

---

## ⚠️ Важные моменты

### База данных сохраняется
- ✅ При деплое данные НЕ удаляются
- ✅ База лежит на хосте в `./backend/data/database.db`
- ✅ Docker volume маппит её в контейнер
- ✅ Все заявки и авто сохраняются между обновлениями

### Переменные окружения
- Не хранятся в Git (`.env` в `.gitignore`)
- Передаются из GitHub Secrets через CI/CD
- Экспортируются перед `docker-compose up`

### Сетевая конфигурация
- Фронтенд доступен только через основной Nginx
- Бэкенд доступен только через основной Nginx
- Прямой доступ к контейнерам снаружи закрыт

---

## 🐛 Troubleshooting

### Проблема: 404 на /api/*
**Причина:** Nginx не проксирует на бэкенд  
**Решение:** Проверить корневой `nginx.conf`, должна быть секция `location /api/`

### Проблема: База очищается при деплое
**Причина:** Удаление файла `database.db` в deploy скрипте  
**Решение:** Убрать `rm backend/data/database.db` из `deploy.yml`

### Проблема: Форма не отправляется
**Причина:** CORS или проблема с nginx проксированием  
**Решение:** 
1. Проверить логи: `docker-compose logs backend`
2. Проверить CORS в `main.go`
3. Проверить nginx конфигурацию

### Проблема: SSL не работает
**Причина:** Сертификаты не обновлены или неправильные пути  
**Решение:**
```bash
certbot renew                     # Обновить сертификаты
ls -la /etc/letsencrypt/live/eleganceauto.od.ua/
docker-compose restart nginx      # Перезапустить nginx
```

### Проблема: Контейнер падает
**Решение:**
```bash
docker-compose logs <container>   # Посмотреть логи
docker-compose restart <container> # Перезапустить
docker system prune -a            # Очистить Docker
```

---

## 📝 TODO
- ~~delete description input from datatable~~ (см. комментарий в main.go:22)

---

## 📞 Контакты и ресурсы

- **Репозиторий:** GitHub (приватный)
- **Сервер:** 91.99.231.3
- **Домен:** eleganceauto.od.ua
- **Админ панель:** https://eleganceauto.od.ua/admin 

---

**Последнее обновление:** 24 октября 2025  
**Версия документации:** 1.0