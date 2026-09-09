package main

import (
	"crypto/sha512"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func DrawBinaryImageFromHash(width, height int, hashBytes [64]byte, name string) {

	maxPixels := 64
	if width*height > maxPixels {
		log.Fatalf("خطا: تعداد پیکسل‌ها (%d * %d = %d) از حداکثر تعداد مجاز (%d) بیشتر است.\n", width, height, width*height, maxPixels)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	byteIndex := 0

	black := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var pixelColor color.RGBA

			if byteIndex < len(hashBytes) {

				currentByteValue := hashBytes[byteIndex]

				if currentByteValue < 128 {
					pixelColor = black
				} else {
					pixelColor = white
				}
				byteIndex++
			} else {

				pixelColor = black
			}

			img.Set(x, y, pixelColor)
		}
	}

	file, err := os.Create(fmt.Sprintf("%s.png", name))
	if err != nil {
		log.Fatalf("خطا در ایجاد فایل %s.png: %v\n", name, err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Fatalf("خطا در کد کردن تصویر به PNG: %v\n", err)
	}
	log.Printf("تصویر سیاه و سفید '%s.png' با موفقیت ساخته شد.\n", name)
}

func main() {
	var name = "Matin Hasanali Baki"

	hasher := sha512.New()
	hasher.Write([]byte(name))
	hashBytesArray := hasher.Sum(nil)

	var hashBytes [64]byte
	copy(hashBytes[:], hashBytesArray)

	width := 8
	height := 8
	imageName := "binary_image_from_hash"

	DrawBinaryImageFromHash(width, height, hashBytes, imageName)
}
