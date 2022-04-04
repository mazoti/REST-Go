package rest_server

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"testing"
)

func TestInvalidIPaddress(t *testing.T) {
	if err := Start("127.0.0.111111:8080"); err != nil {
		if err.Code() != BIND_ERROR {
			t.Errorf("FAILED with code %d, expected %d", err.Code(), BIND_ERROR)
		}
	}
}

func TestInvalidPort(t *testing.T) {
	if err := Start("127.0.0.1:111111"); err != nil {
		if err.Code() != BIND_ERROR {
			t.Errorf("FAILED with code %d, expected %d", err.Code(), BIND_ERROR)
		}
	}
}

func TestToASCII(t *testing.T) {
	handler := Server{}
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))

	whitePixel := color.RGBA{255, 255, 255, 255}
	redPixel := color.RGBA{255, 0, 0, 0}
	greenPixel := color.RGBA{0, 255, 0, 0}
	bluePixel := color.RGBA{0, 0, 255, 0}

	img.Set(0, 0, redPixel)
	img.Set(0, 1, greenPixel)
	img.Set(1, 0, bluePixel)
	img.Set(1, 1, whitePixel)

	if handler.toASCII(img) != "d&\n? \n" {
		t.Errorf("FAILED")
	}
}

func BenchmarkToASCII(b *testing.B) {
	handler := Server{}

	img1 := image.NewRGBA(image.Rect(0, 0, 100, 100))
	img2 := image.NewRGBA(image.Rect(0, 0, 100, 100))
	img3 := image.NewRGBA(image.Rect(0, 0, 100, 100))
	img4 := image.NewRGBA(image.Rect(0, 0, 100, 100))

	whitePixel := color.RGBA{255, 255, 255, 255}
	redPixel := color.RGBA{255, 0, 0, 0}
	greenPixel := color.RGBA{0, 255, 0, 0}
	bluePixel := color.RGBA{0, 0, 255, 0}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			img1.Set(i, j, whitePixel)
			img2.Set(i, j, redPixel)
			img3.Set(i, j, greenPixel)
			img4.Set(i, j, bluePixel)
		}
	}

	handler.toASCII(img1)
	handler.toASCII(img2)
	handler.toASCII(img3)
	handler.toASCII(img4)
}
