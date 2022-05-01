package img

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

type CreateLayeredImgOptions struct {
	OverlayX int
	OverlayY int
}

func NewCreateLayeredImgOptions() CreateLayeredImgOptions {
	return CreateLayeredImgOptions{
		200, 200,
	}
}

type ImgLayer struct {
	Image string
	X int
	Y int
	Scale float32
}

type ImgLayers []ImgLayer

func NewImgLayer(image string) ImgLayer {
	return ImgLayer{
		image, 200, 200, 1.0,
	}
}

func CreateLayeredImg(background string, layers ImgLayers, output string, options CreateLayeredImgOptions) error {
    
	imgBack, err := gg.LoadImage(background)
	if err != nil {
		return err
	}

    b := imgBack.Bounds()
    imgOut := image.NewRGBA(b)
    draw.Draw(imgOut, b, imgBack, image.ZP, draw.Src)

	imgOutFile,err := os.Create(output)
	if err != nil {
		return err
	}
	defer imgOutFile.Close()

	for _, layer := range layers {
		imgLayer,err := gg.LoadImage(layer.Image)
		if err != nil {
			return err
		}

		offset := image.Pt(layer.X, layer.Y)

		imgLayerResized := resize.Resize(
			uint(float32(imgLayer.Bounds().Dx()) * layer.Scale), 
			uint(float32(imgLayer.Bounds().Dy())* layer.Scale), 
			imgLayer,
			resize.Lanczos3,
		)

		draw.Draw(imgOut, imgLayerResized.Bounds().Add(offset), imgLayerResized, image.ZP, draw.Over)
	}

	jpeg.Encode(imgOutFile, imgOut, &jpeg.Options{jpeg.DefaultQuality})
	return nil
}