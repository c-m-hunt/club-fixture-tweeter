package img

import (
	_ "embed"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//go:embed fonts/Mangonel.ttf
var mangonelFontBytes []byte

func GetMangonelFont(points float64) (font.Face, error) {
	f, err := truetype.Parse(mangonelFontBytes)
	if err != nil {
		return nil, err
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size: points,
		// Hinting: font.HintingFull,
	})
	return face, nil
}


