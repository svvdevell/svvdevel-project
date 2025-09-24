FROM golang:1.22 as build

WORKDIR /app
COPY . .
RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev
RUN go mod download
RUN go build -o main .

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=build /app/main .

RUN apt-get update && apt-get install -y ca-certificates sqlite3 && rm -rf /var/lib/apt/lists/*
RUN mkdir -p /app/uploads /app/data

VOLUME /app/data
VOLUME /app/uploads
EXPOSE 8080
CMD ["/app/main"]
