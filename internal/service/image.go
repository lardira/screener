package service

import (
	"bytes"
	_ "embed"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/img.png
var imgBytes []byte

var Img *ebiten.Image

func init() {
	var err error
	Img, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(imgBytes))
	if err != nil {
		log.Fatal(err)
	}
}
