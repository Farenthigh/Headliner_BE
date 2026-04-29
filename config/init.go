package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Initenv() {
    // 1. ลองหาไฟล์ .env ใน Folder ปัจจุบันก่อน
    err := godotenv.Load()
    if err != nil {
        // 2. ถ้าหาไม่เจอ ไม่ต้องสั่ง Fatal! 
        // แค่ Log บอกไว้เผื่อเราลืมตั้งค่าที่หน้า Cloud
        log.Println("Note: .env file not found, will use system environment variables")
    }
}
