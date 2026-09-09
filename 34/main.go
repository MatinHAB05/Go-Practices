package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// https://developers.google.com/recaptcha/docs/v3
// https://www.google.com/recaptcha/admin/ پدرت در میاد اینو پیدا کنی خخخخخخ

const (
	RecaptchaSecretKey = "...." // از توی لینک دوم پیداش کن
	RecaptchaVerifyURL = "https://www.google.com/recaptcha/api/siteverify"
	ScoreThreshold     = 0.5 // امتیاز بین 0.0 (ربات) تا 1.0 (انسان)
)

// ساختار داده‌های ارسالی از فرانت‌اند
type LoginRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	RecaptchaToken string `json:"recaptcha_token"`
}

// ساختار پاسخ دریافتی از API گوگل
// توی داک ش اشاره کرده به این
type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

// تابع اعتبارسنجی توکن با سرورهای گوگل
func verifyRecaptcha(token string) (*RecaptchaResponse, error) {
	// ارسال پارامترها به صورت Form Data
	// توی داک اینو اشاره کرده اش
	resp, err := http.PostForm(RecaptchaVerifyURL, url.Values{
		"secret":   {RecaptchaSecretKey},
		"response": {token},
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var recaptchaRes RecaptchaResponse
	if err := json.NewDecoder(resp.Body).Decode(&recaptchaRes); err != nil {
		return nil, err
	}

	return &recaptchaRes, nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "فرمت درخواست نامعتبر است"}`, http.StatusBadRequest)
		return
	}

	// ۱. استعلام توکن از گوگل
	captchaResult, err := verifyRecaptcha(req.RecaptchaToken)
	if err != nil || !captchaResult.Success {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"msg":    "اعتبارسنجی کپچا ناموفق بود",
			"detail": captchaResult,
		})
		return
	}

	// ۲. بررسی امتیاز (Score) کاربر
	if captchaResult.Score < ScoreThreshold {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "error",
			"msg":    "رفتار شما شبیه به ربات تشخیص داده شد!",
			"score":  captchaResult.Score,
		})
		return
	}

	// ۳. پردازش لاگین در صورت تایید انسانی بودن
	fmt.Printf("ورود موفقیت‌آمیز برای کاربر: %s | امتیاز گوگل: %.2f\n", req.Username, captchaResult.Score)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"msg":    "ورود موفقیت‌آمیز بود",
		"score":  captchaResult.Score,
	})
}

func main() {
	http.HandleFunc("/api/login", loginHandler)

	fmt.Println("سرور Go روی پورت 8080 فعال شد...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
