package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"os"
	"time"
)

const (
	CHARS    = "      .:-=+*#%@"
	REQ_SIZE = 64
)

// reduce the size of the gif to 100x100
func resizeFrames(frames []*image.Paletted, width, height int) []*image.Paletted {
	hScale := float64(height) / float64(frames[0].Bounds().Dy())
	wScale := float64(width) / float64(frames[0].Bounds().Dx())

	resizedFrames := make([]*image.Paletted, len(frames))

	for i, frame := range frames {
		newWidth := int(float64(frame.Bounds().Dx()) * wScale)
		newHeight := int(float64(frame.Bounds().Dy()) * hScale)

		resizedFrames[i] = image.NewPaletted(image.Rect(0, 0, newWidth, newHeight), frame.Palette)

		for x := 0; x < newWidth; x++ {
			for y := 0; y < newHeight; y++ {
				resizedFrames[i].Set(x, y, frame.At(int(float64(x)/wScale), int(float64(y)/hScale)))
			}

		}

	}

	return resizedFrames
}

func readGif(fileName string) *gif.GIF {
	file, err := os.Open("gifs/tom.gif")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Error closing file: %s", err)
		}
	}(file)

	gifImage, err := gif.DecodeAll(file)
	if err != nil {
		log.Fatalf("Error decoding file: %s", err)
	}

	return gifImage
}

func readFrames(fileName string) []*image.Paletted {
	gifImage := readGif(fileName)
	// resize the gif
	gifImage.Image = resizeFrames(gifImage.Image, REQ_SIZE, REQ_SIZE)

	return gifImage.Image
}

// convert color to grayscale
func grayscaleChar(color color.Color) int {
	r, g, b, _ := color.RGBA()
	val := int(0.299*float64(r)+0.587*float64(g)+0.114*float64(b)) / 256

	val = val * (len(CHARS) - 1) / 255

	return val
}

func clearConsole() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	fileName := "gifs/tom.gif"
	gifImage := readFrames(fileName)

	frameValues := make([][]int, len(gifImage))

	for i, frame := range gifImage {
		frameValues[i] = make([]int, frame.Bounds().Dx()*frame.Bounds().Dy())
		for x := 0; x < frame.Bounds().Dx(); x++ {
			for y := 0; y < frame.Bounds().Dy(); y++ {
				frameValues[i][x*frame.Bounds().Dy()+y] = grayscaleChar(frame.At(x, y))
			}
		}
	}

	for {
		for i, frame := range frameValues {
			// roate the image 90 degrees
			for y := 0; y < gifImage[i].Bounds().Dy(); y++ {
				for x := 0; x < gifImage[i].Bounds().Dx(); x++ {
					char := CHARS[frame[x*gifImage[i].Bounds().Dy()+y]]
					fmt.Printf("%c%c%c", char, char, char)
				}
				fmt.Printf("\n")
			}

			fmt.Printf("\n\n")

			time.Sleep(30 * time.Millisecond)
			clearConsole()
		}

		time.Sleep(100 * time.Millisecond)
	}
}
