# Menggunakan Go image sebagai base image
FROM golang:1.18-alpine

# Install air
RUN go install github.com/cosmtrek/air@latest

# Menentukan working directory di dalam kontainer
WORKDIR /app

# Copy file go.mod dan go.sum ke dalam working directory
COPY ../../go.mod ./
COPY ../../go.sum ./
COPY ../../air.toml ./
COPY ../../main.go ./

# Menjalankan perintah go mod download untuk mendownload dependencies
RUN go mod download

# Copy seluruh file proyek ke dalam working directory
COPY . .

# Build aplikasi Go dan memberikan izin eksekusi pada file api
RUN go build -o api && chmod +x ./api


# Menjalankan aplikasi saat kontainer dimulai
CMD ["air", "-c", "air.toml"]




