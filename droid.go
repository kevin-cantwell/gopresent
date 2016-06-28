package main

import (
	"image/png"
	"os"
)

// START OMIT
import "github.com/kevin-cantwell/dotmatrix"

func main() {
	f, _ := os.Open("droid.png")
	img, _ := png.Decode(f)
	dotmatrix.Encode(os.Stdout, img)
}

// END OMIT
