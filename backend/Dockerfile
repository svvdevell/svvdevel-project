FROM golang:1.22-bullseye as build

WORKDIR /app
COPY . .

# Устанавливаем SQLite и зависимости
RUN apt-get update && apt-get install -y \
    sqlite3 \
    libsqlite3-dev \
    gcc \
    && rm -rf /var/lib/apt/lists/*

# Настраиваем переменные для статической сборки
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

RUN go mod download
RUN go build -ldflags '-linkmode external -extldflags "-static"' -a -o main .

FROM debian:bullseye-slim

WORKDIR /app

# Устанавливаем минимальные зависимости
RUN apt-get update && apt-get install -y \
    ca-certificates \
    sqlite3 \
    && rm -rf /var/lib/apt/lists/*

COPY --from=build /app/main .

# Создаем необходимые директории
RUN mkdir -p /app/uploads /app/data

# Устанавливаем права
RUN chmod +x /app/main

VOLUME ["/app/data", "/app/uploads"]

EXPOSE 8080

CMD ["/app/main"]
