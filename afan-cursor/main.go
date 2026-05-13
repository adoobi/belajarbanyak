package main

import (
	"fmt"
	"math"
	"time"

	"github.com/go-vgo/robotgo"
)

// moveTo menggerakkan kursor ke posisi x, y dengan jeda
func moveTo(x, y int, delay time.Duration) {
	robotgo.Move(x, y)
	time.Sleep(delay)
}

// drawLines menggambar garis dari titik-titik yang diberikan
func drawLines(points [][2]int, delay time.Duration) {
	for _, p := range points {
		moveTo(p[0], p[1], delay)
	}
}

// drawR menggambar huruf R dengan pusat di (cx, cy)
// size menentukan ukuran huruf
func drawR(cx, cy, size int, delay time.Duration) {
	fmt.Println("Menggambar huruf R...")

	step := delay

	// Titik-titik huruf R:
	// Garis vertikal (kiri atas ke kiri bawah)
	vertical := [][2]int{}
	for i := 0; i <= 10; i++ {
		vertical = append(vertical, [2]int{cx, cy + i*size/10})
	}

	// Kembali ke atas untuk mulai lengkung
	topCurve := [][2]int{{cx, cy}}

	// Lengkung atas R (setengah lingkaran kanan)
	curve := [][2]int{}
	for deg := -90; deg <= 90; deg += 10 {
		rad := float64(deg) * math.Pi / 180
		px := cx + int(float64(size/2)*math.Cos(rad))
		py := cy + size/4 + int(float64(size/4)*math.Sin(rad))
		curve = append(curve, [2]int{px, py})
	}

	// Titik tengah (perut R)
	middle := [][2]int{{cx, cy + size/2}}

	// Kaki kanan R (diagonal ke kanan bawah)
	leg := [][2]int{}
	for i := 0; i <= 5; i++ {
		px := cx + i*size/5
		py := cy + size/2 + i*size/10
		leg = append(leg, [2]int{px, py})
	}

	// Jalankan semua gerakan
	drawLines(vertical, step)
	drawLines(topCurve, step)
	drawLines(curve, step)
	drawLines(middle, step)
	drawLines(leg, step)
}

// drawHeart menggambar bentuk hati dengan pusat di (cx, cy)
func drawHeart(cx, cy, size int, delay time.Duration) {
	fmt.Println("Menggambar hati...")

	// Rumus parametrik hati:
	// x = size * 16 * sin^3(t)
	// y = size * (13cos(t) - 5cos(2t) - 2cos(3t) - cos(4t))
	points := [][2]int{}
	for deg := 0; deg <= 360; deg += 3 {
		t := float64(deg) * math.Pi / 180
		scale := float64(size) / 17.0

		px := cx + int(scale*16*math.Pow(math.Sin(t), 3))
		py := cy - int(scale*(13*math.Cos(t)-
			5*math.Cos(2*t)-
			2*math.Cos(3*t)-
			math.Cos(4*t)))

		points = append(points, [2]int{px, py})
	}

	drawLines(points, delay)
}

func main() {
	// Jeda 3 detik supaya ada waktu pindah ke area yang tepat
	fmt.Println("Mulai dalam 3 detik... pindahkan kursor ke tengah layar")
	time.Sleep(3 * time.Second)

	// Ambil posisi tengah layar sebagai pusat gambar
	cx, cy := robotgo.Location()
	fmt.Printf("Pusat gambar: (%d, %d)\n", cx, cy)

	delay := 30 * time.Millisecond

	// Pilih yang mau digambar: drawR atau drawHeart
	drawHeart(cx, cy, 80, delay)

	time.Sleep(500 * time.Millisecond)

	drawR(cx+200, cy, 80, delay)

	fmt.Println("Selesai!")
}

// dah
