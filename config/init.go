package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func Initenv() string {
    _, filename, _, _ := runtime.Caller(0)
    currentFileDir := filepath.Dir(filename)

    var configPath string
    for {
        configPath = filepath.Join(currentFileDir, ".env")
        if _, err := os.Stat(configPath); !os.IsNotExist(err) {
            // เจอไฟล์แล้ว โหลดเลย
            _ = godotenv.Load(configPath)
            return configPath
        }
		
        parentDir := filepath.Dir(currentFileDir)
        if parentDir == currentFileDir {
            // --- แก้ตรงนี้: แทนที่จะ log.Fatalf ให้แค่ return ค่าว่าง ---
            log.Println("Warning: Reached root directory, no .env file found.")
            return "" 
        }
        currentFileDir = parentDir
    }
}
