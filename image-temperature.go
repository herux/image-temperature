package imagetemperature

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

type ImageColorAdjuster struct {
	imageInFile      *os.File
	image            image.Image
	ImageInFilename  string
	ImageOutFilename string
}

func New(Infilename, Outfilename string) *ImageColorAdjuster {
	imageInFile, err := os.Open(Infilename)
	if err != nil {
		fmt.Println(err)
	}
	defer imageInFile.Close()

	image, err := jpeg.Decode(imageInFile)

	return &ImageColorAdjuster{
		imageInFile:      imageInFile,
		image:            image,
		ImageInFilename:  Infilename,
		ImageOutFilename: Outfilename,
	}
}

func (ica *ImageColorAdjuster) ByTemperature(temperature float64) image.Image {
	bounds := ica.image.Bounds()
	newImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originColor := ica.image.At(x, y)
			r, g, b, _ := originColor.RGBA()

			newR := uint8(float64(r) * temperature)
			newG := uint8(float64(g))
			newB := uint8(float64(b))

			newR = ica.min(ica.max(newR, 0), 255)
			newB = ica.min(ica.max(newB, 0), 255)
			newG = ica.min(ica.max(newG, 0), 255)

			newColor := color.RGBA{
				newR, newG, newB, 255,
			}
			newImg.Set(x, y, newColor)
		}
	}

	return newImg
}

func (ica *ImageColorAdjuster) SaveToFile(newImage image.Image) error {
	imageOutFile, err := os.Create(ica.ImageOutFilename)
	if err != nil {
		return err
	}
	defer imageOutFile.Close()
	return jpeg.Encode(imageOutFile, newImage, nil)
}

func (ica *ImageColorAdjuster) min(a, b uint8) uint8 {
	if a < b {
		return a
	}
	return b
}

func (ica *ImageColorAdjuster) max(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}
