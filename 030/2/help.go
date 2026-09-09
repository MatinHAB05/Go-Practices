package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func initConfig() Config {
	// Retrieve configuration via Viper
	viper.AddConfigPath("./../")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("ERR: Failed to read config file: ", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("ERR: Failed to unmarshal config: ", err)
	}
	return cfg
}

// پرچم (?s) باعث می‌شود . شامل نیولاین (\n) هم بشود
var jsonExtractor = regexp.MustCompile(`(?s)(\{.*\}|\[.*\])`)

// فرمت‌دهی و زیباسازی JSON
func formatTelegramJSON(rawJSON string) string {
	var obj any
	err := json.Unmarshal([]byte(rawJSON), &obj)
	if err != nil {
		return rawJSON // اگر JSON معتبر نبود، متن خام بازگردانده می‌شود
	}

	prettyBytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return rawJSON
	}

	return string(prettyBytes)
}

func initLogger(cmdLog bool) (func(format string, args ...any), func()) {
	// ۱. ساخت پوشه logs در صورت عدم وجود
	logsDir := "logs"
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		log.Fatalf("ERR: Failed to create logs directory: %v", err)
	}

	timeStamp := time.Now().Format("2006-01-02_15-04-05")

	// ۲. ساخت فایل لاگ خام (Raw)
	rawFileName := filepath.Join(logsDir, fmt.Sprintf("%s_raw.log", timeStamp))
	rawFile, err := os.OpenFile(rawFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("ERR: Failed to create raw log file: %v", err)
	}

	// ۳. ساخت فایل لاگ تمیز (Clean)
	cleanFileName := filepath.Join(logsDir, fmt.Sprintf("%s_clean.log", timeStamp))
	cleanFile, err := os.OpenFile(cleanFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		rawFile.Close()
		log.Fatalf("ERR: Failed to create clean log file: %v", err)
	}

	// تعریف Loggerها
	rawLogger := log.New(rawFile, "", log.LstdFlags|log.Lmicroseconds)
	cleanLogger := log.New(cleanFile, "", log.LstdFlags|log.Lmicroseconds)
	termLogger := log.New(os.Stdout, "\033[36m[TELEGRAM]\033[0m ", log.Ltime)

	// ۴. تابع اصلی لاگر
	customDebugHandler := func(format string, args ...any) {
		fullMsg := fmt.Sprintf(format, args...)

		// الف) ذخیره مستقیم و بدون دستکاری در فایل RAW
		rawLogger.Println(fullMsg)

		// ب) پردازش و تمیزسازی برای فایل CLEAN و ترمینال
		cleanMsg := strings.ReplaceAll(fullMsg, "\r\n", "\n")
		cleanMsg = strings.ReplaceAll(cleanMsg, "\r", "")

		match := jsonExtractor.FindString(cleanMsg)

		if match != "" {
			prefix := cleanMsg[:strings.Index(cleanMsg, match)]
			prefix = strings.TrimSpace(prefix)

			formattedJSON := formatTelegramJSON(match)

			// ذخیره در فایل CLEAN
			cleanLogger.Printf("%s\nPayload:\n%s\n%s\n", prefix, formattedJSON, strings.Repeat("-", 80))

			if cmdLog {
				// چاپ در ترمینال
				termLogger.Printf("\033[33m%s\033[0m\n%s\n%s\n", prefix, formattedJSON, strings.Repeat("-", 50))
			}
		} else {
			cleanLogger.Println(cleanMsg)
			if cmdLog {
				termLogger.Println(cleanMsg)
			}
		}
	}

	// تابع جهت بستن هردو فایل هنگام خروج
	cleanup := func() {
		rawFile.Close()
		cleanFile.Close()
	}

	return customDebugHandler, cleanup
}
