# --- Stage 1: Build Stage ---
FROM golang:1.25-alpine AS builder

WORKDIR /app

# ติดตั้ง dependencies ก่อน (เพื่อใช้ Layer Cache)
COPY go.mod go.sum ./
RUN go mod download

# คัดลอกโค้ดทั้งหมดแล้ว Build เป็น Binary
COPY . .
# -ldflags="-w -s" ช่วยลดขนาดไฟล์ Binary ลงไปอีก
RUN go build -ldflags="-w -s" -o main .

# --- Stage 2: Run Stage ---
# ใช้ Alpine ธรรมดา หรือ Distroless เพื่อความเบาและปลอดภัย
FROM alpine:latest  

# ติดตั้ง ca-certificates กรณีแอปต้องคุยกับ HTTPS (เช่น Firebase/Railway DB)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# ก๊อปปี้เฉพาะไฟล์ Binary มาจาก Stage builder
COPY --from=builder /app/main .

# อย่าลืมก๊อปปี้โฟลเดอร์พวก static หรือ template ถ้ามี (เช่น โฟลเดอร์ที่เก็บ key firebase)
# COPY --from=builder /app/serviceAccountKey.json ./ 

# รับค่า PORT จาก Cloud
ENV PORT=8080
EXPOSE 8080

# รันแอป
CMD ["./main"]