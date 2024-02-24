package main

import (
	"flag"
	"fmt"

	imagetemperature "herux.com/image-temperature"
)

func main() {
	temperature := flag.Float64("temperature", 1.0, "Temperature adjustment value (positive for warmer, negative for cooler)")
	imageInFilename := flag.String("input", "", "Input JPEG image file path")
	imageOutFilename := flag.String("output", "", "Output JPEG image file path")
	flag.Parse()

	if *imageInFilename == "" {
		fmt.Println("Image input filename is required")
	}

	if *imageOutFilename == "" {
		fmt.Println("Image output filename is required")
	}

	ica := imagetemperature.New(*imageInFilename, *imageOutFilename)
	newImage := ica.ByTemperature(*temperature)
	ica.SaveToFile(newImage)
}
