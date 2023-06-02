# Menggunakan Go image sebagai base image
FROM golang:1.18-alpine

# Menentukan working directory di dalam kontainer
WORKDIR /app

# Copy file go.mod dan go.sum ke dalam working directory
COPY ../../go.mod ./
COPY ../../go.sum ./
COPY ../../main.go ./

# Menjalankan perintah go mod download untuk mendownload dependencies
RUN go mod download

# Copy seluruh file proyek ke dalam working directory
COPY . .

# Build aplikasi Go
RUN go build -o api

# Menjalankan aplikasi saat kontainer dimulai
CMD ["go", "run", "main.go"]
