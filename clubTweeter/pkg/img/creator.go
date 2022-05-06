package img

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)


type Layer interface {
	AddLayerToContext(*gg.Context) error
}

type ImgLayer struct {
	Image string
	X int
	Y int
	Scale float32
}

func (l ImgLayer) AddLayerToContext(dc *gg.Context) error {
	img, err := gg.LoadImage(l.Image)
	if err != nil {
		return err
	}

	imgResized := resize.Resize(
		uint(float32(img.Bounds().Dx()) * l.Scale), 
		uint(float32(img.Bounds().Dy())* l.Scale), 
		img,
		resize.Lanczos3,
	)

	dc.DrawImage(imgResized, l.X, l.Y)
	return nil
}

type TextLayer struct {
	Text string
	X int
	Y int
	Color color.Color
	Rotate float64
	Size float64
	Centered bool
}

func (l TextLayer) AddLayerToContext(dc *gg.Context) error {
	font, _ := GetMangonelFont(l.Size)
	dc.Push()
	dc.Rotate(l.Rotate)
	dc.SetFontFace(font)
	dc.SetColor(color.White)
	if l.Centered {
		dc.DrawStringAnchored(l.Text, float64(l.X), float64(l.Y), 0.5, 0.5)
	} else {
		dc.DrawString(l.Text, float64(l.X), float64(l.Y))
	}
	
	dc.Pop()
	return nil
}


type Layers []Layer

func NewImgLayer(image string) ImgLayer {
	return ImgLayer{
		image, 200, 200, 1.0,
	}
}

func CreateLayeredImg(background string, layers Layers, output string) error {
	imgBack, err := gg.LoadImage(background)
	if err != nil {
		return err
	}

	dc := gg.NewContextForImage(imgBack)

    dc.DrawImage(imgBack, 0, 0)

	for _, layer := range layers {
		layer.AddLayerToContext(dc)
	}

	return dc.SavePNG(output)
}