package img

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"

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

    imgBackFile,err := os.Open(background)
	if err != nil {
		return err
	}
    
	imgBack, err := png.Decode(imgBackFile)
	if err != nil {
		return err
	}
    defer imgBackFile.Close()

    b := imgBack.Bounds()
    imgOut := image.NewRGBA(b)
    draw.Draw(imgOut, b, imgBack, image.ZP, draw.Src)

	imgOutFile,err := os.Create(output)
	if err != nil {
		return err
	}
	defer imgOutFile.Close()

	for _, layer := range layers {
		imgLayerFile,err := os.Open(layer.Image)
		if err != nil {
			return err
		}
		imgLayer,err := png.Decode(imgLayerFile)
		if err != nil {
			return err
		}
		defer imgLayerFile.Close()

		offset := image.Pt(layer.X, layer.Y)

		fmt.Print(imgLayer.Bounds().Dx())

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