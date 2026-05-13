package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func main() {

	// Mengambil layar monitor pertama
	bounds := screenshot.GetDisplayBounds(0)

	// Capture layar
	img, err := screenshot.CaptureRect(bounds)

	if err != nil {
		panic(err)
	}

	// Membuat file hasil screenshot
	file, err := os.Create("hasil_screenshot.png")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Simpan gambar ke file PNG
	png.Encode(file, img)

	fmt.Println("Screenshot berhasil disimpan")
}
