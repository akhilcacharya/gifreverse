// gifreverse project main.go
package main

import (
	"fmt"
	"image"
	"image/gif"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) > 2 || len(args) == 0 {
		showUsage()
		return
	}

	input := args[0]

	fmt.Println("Reading file", input)
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	img, err := gif.DecodeAll(file)

	if err != nil {
		fmt.Println(err)
		return
	}

	reversed := new(gif.GIF)

	reversed.Delay = img.Delay
	reversed.LoopCount = img.LoopCount
	reversed.Image = reverse(img.Image)

	fmt.Println("Reversing..")

	var output string

	if len(args) == 2 {
		output = strings.TrimSuffix(args[1], filepath.Ext(input)) + ".gif"
	} else {
		output = strings.TrimSuffix(input, filepath.Ext(input)) + "_reversed.gif"
	}

	result, err := os.Create(output)

	if err != nil {
		fmt.Println(err)
		return
	}

	gif.EncodeAll(result, reversed)

	fmt.Println("✓ Saved", output)
}

func showUsage() {
	fmt.Println("gifreverse usage:")
	fmt.Println("==> gifreverse normal.gif reversed.gif")
}

func reverse(images []*image.Paletted) []*image.Paletted {
	result := []*image.Paletted{}
	for i := len(images) - 1; i >= 0; i-- {
		result = append(result, images[i])
	}
	return result
}
