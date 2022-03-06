/*
Copyright © 2022 Maria Petrova marycool674@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package converter

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/spf13/viper"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
)

type converterConfig struct {
	inputFile  string
	outputFile string
	invert     bool
	width      uint
	height     uint
}

type customImage struct {
	image  image.Image
	width  int
	height int
}

var (
	asciiArray  []string
	ascii       = []string{"#", "@", "&", "X", "x", ";", ":", ",", ".", " "}
	asciiInvert = []string{" ", ".", ",", ":", ";", "x", "X", "&", "@", "#"}
)

// GrayToChar pixel color value to ascii-symbol value mapper
func GrayToChar(num uint8) string {
	if num <= 25 {
		return asciiArray[0]
	} else if num > 25 && num <= 50 {
		return asciiArray[1]
	} else if num > 50 && num <= 75 {
		return asciiArray[2]
	} else if num > 75 && num <= 100 {
		return asciiArray[3]
	} else if num > 100 && num <= 125 {
		return asciiArray[4]
	} else if num > 125 && num <= 150 {
		return asciiArray[5]
	} else if num > 150 && num <= 175 {
		return asciiArray[6]
	} else if num > 175 && num <= 200 {
		return asciiArray[7]
	} else if num > 200 && num <= 225 {
		return asciiArray[8]
	}
	return asciiArray[9]
}

// initConfig initialise config
func initConfig(inputFile string) converterConfig {
	conf := converterConfig{
		inputFile:  inputFile,
		outputFile: viper.GetString("outputFile"),
		invert:     viper.GetBool("invert"),
		width:      viper.GetUint("width"),
		height:     viper.GetUint("height"),
	}
	if conf.invert {
		asciiArray = asciiInvert
	} else {
		asciiArray = ascii
	}
	return conf
}

// getImageData get the image.Image instance, as well as it's width and height
func getImageData(cfg converterConfig) (customImage, error) {
	existingImageFile, err := os.Open(cfg.inputFile)
	if err != nil {
		return customImage{}, err
	}
	imageData, _, err := image.Decode(existingImageFile)
	if err != nil {
		return customImage{}, err
	}
	if cfg.width == 0 && cfg.height == 0 {
		cfg.width = 300
	} else if cfg.width != 0 && cfg.height != 0 {
		cfg.height = 0
	}
	imageToConvert := resize.Resize(cfg.width, cfg.height, imageData, resize.Lanczos3)
	existingImageFile.Seek(0, 0)
	bounds := imageToConvert.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	return customImage{
		image:  imageToConvert,
		width:  w,
		height: h,
	}, nil
}

// getAsciiSymbols convert the image.Image to 2D ascii-symbols array
func getAsciiSymbols(handledImage customImage) ([]string, error) {
	var pixels []string
	grayScale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{handledImage.width, handledImage.height}})
	for y := 0; y < handledImage.height; y++ {
		var av string
		for x := 0; x < handledImage.width; x++ {
			imageColor := handledImage.image.At(x, y)
			rr, gg, bb, _ := imageColor.RGBA()
			r := math.Pow(float64(rr), 2.2)
			g := math.Pow(float64(gg), 2.2)
			b := math.Pow(float64(bb), 2.2)
			m := math.Pow(0.2125*r+0.7154*g+0.0721*b, 1/2.2)
			Y := uint16(m + 0.5)
			grayColor := color.Gray{uint8(Y >> 8)}
			av = av + GrayToChar(grayColor.Y) + GrayToChar(grayColor.Y)
			grayScale.Set(x, y, grayColor)
		}
		pixels = append(pixels, av)
	}
	return pixels, nil
}

// saveAsciiToFile save ascii-symbols to a file
func saveAsciiToFile(cfg converterConfig, pixels []string) error {
	f, err := os.Create(cfg.outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	for i := 0; i < len(pixels); i++ {
		_, err = f.WriteString(fmt.Sprintf("%v\n", pixels[i]))
		if err != nil {
			return err
		}
	}
	return nil
}

// ConvertToAscii the main function in here
func ConvertToAscii(args []string) {
	cfg := initConfig(args[0])
	handledImage, err := getImageData(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	pixels, err := getAsciiSymbols(handledImage)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = saveAsciiToFile(cfg, pixels)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ASCII-converted image file generated:", cfg.outputFile)
}
