# Tahap 1: Build stage
# Menggunakan official Go image sebagai base untuk build
FROM golang:1.23.2-alpine AS builder

# Mengatur working directory di dalam container
WORKDIR /app

# Copy semua file proyek ke working directory di container
COPY . .

# Download dependencies
RUN go mod tidy

# Build aplikasi
RUN go build -o server .

# Tahap 2: Final stage
# Menggunakan base image yang ringan
FROM alpine:latest

# Mengatur working directory untuk stage final
WORKDIR /root/

# Copy hasil build dari stage 1
COPY --from=builder /app/server .

# Expose port 8080 untuk aplikasi
EXPOSE 8080

# Menjalankan aplikasi
CMD ["./server"]
