# Выберите базовый образ с Golang
FROM golang:1.19-alpine as builder

# Установите инструменты для компиляции
RUN apk add --no-cache git build-base pkgconfig alsa-utils alsa-utils-doc alsa-lib alsa-lib-dev alsaconf alsa-ucm-conf

# Установите рабочую директорию
WORKDIR /app

# Копируйте файлы с исходным кодом проекта
COPY . .

# Установите зависимости
RUN go mod download

# Соберите проект для Linux amd64
RUN GOOS=linux go build -o cringecast-client-linux main.go

# Создайте новый образ на базе Linux alpine для запуска приложения
FROM alpine:latest

ENV MQTT_URL=""
ENV MQTT_TOPIC=""
ENV MQTT_USERNAME=""
ENV MQTT_PASSWORD=""

# Установите зависимости, необходимые для работы приложения
RUN apk --no-cache add ca-certificates

# Установите рабочую директорию
WORKDIR /app/

# Копируйте исполняемый файл из образа builder
COPY --from=builder /app/cringecast-client-linux .
COPY --from=builder /app/static static

EXPOSE 80

# Укажите команду для запуска приложения
CMD  ./cringecast-client-linux --mqtt-url=$MQTT_URL --mqtt-topic=$MQTT_TOPIC --mqtt-username=$MQTT_USERNAME --mqtt-password=$MQTT_PASSWORD