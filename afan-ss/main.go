package main

import (
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	fmt.Println("Mengambil screenshot layar...")

	n := screenshot.NumActiveDisplays()
	fmt.Printf("Jumlah monitor aktif: %d\n", n)

	for i := 0; i < n; i++ {
		// Ambil ukuran layar monitor ke-i
		bounds := screenshot.GetDisplayBounds(i)

		// Capture layar sesuai ukuran bounds
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			fmt.Printf("Gagal capture display %d: %v\n", i, err)
			continue
		}

		// Buat nama file dengan timestamp supaya tidak tertimpa
		filename := fmt.Sprintf("screenshot_%d_%s.png", i, time.Now().Format("20060102_150405"))

		// Buat file PNG
		file, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Gagal buat file: %v\n", err)
			continue
		}

		// Encode image ke format PNG dan tulis ke file
		if err := png.Encode(file, img); err != nil {
			fmt.Printf("Gagal encode PNG: %v\n", err)
			file.Close()
			continue
		}

		file.Close()
		fmt.Printf("Tersimpan: %s (%dx%d px)\n", filename, bounds.Dx(), bounds.Dy())
	}
}
