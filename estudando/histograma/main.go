package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

func png2c(fileName string) {
	var colorList = make(map[string]int)
	var colorCount = 0
	var bitmapArray = "\t"

	var definePrefix = strings.Replace(fileName, ".", "_", -1)
	definePrefix = strings.Replace(definePrefix, "/", "_", -1)

	imgfile, err := os.Open(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v file not found!\n", fileName)
		os.Exit(1)
	}

	defer imgfile.Close()

	imgCfg, _, err := image.DecodeConfig(imgfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	width := imgCfg.Width
	height := imgCfg.Height

	imgfile.Seek(0, 0)

	img, _, err := image.Decode(imgfile)
	if err != nil {
		fmt.Printf("error trying to decode %v. %v!\n", fileName, err)
		os.Exit(1)
	}

	var lineCount = 6
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			argb := fmt.Sprintf("0x%02X%02X%02X%02X",
				uint8(a),
				uint8(r),
				uint8(g),
				uint8(b))

			var patern = "0x%02X, "
			if y == height-1 && x == width-1 {
				patern = "0x%02X"
			}

			if _, ok := colorList[argb]; ok {
				//fmt.Println("element found")
				bitmapArray += fmt.Sprintf(patern, colorList[argb])
			} else {
				colorList[argb] = colorCount
				bitmapArray += fmt.Sprintf(patern, colorCount)
				colorCount++
				//fmt.Println("element not found")
			}

			if lineCount > 70 {
				bitmapArray += "\n\t"
				lineCount = 0
			}

			lineCount += 6
			//fmt.Printf("[%d:%d] %s\n", x, y, argb)
		}
	}

	ncolors := len(colorList)

	fmt.Printf("const uint32_t %s_width = %d;\n", definePrefix, width)
	fmt.Printf("const uint32_t %s_height = %d;\n", definePrefix, height)
	fmt.Printf("const uint32_t %s_ncolors = %d;\n", definePrefix, ncolors)
	fmt.Printf("const uint32_t %s_color_index[] = { \n", definePrefix)

	i := 0
	for k := range colorList {
		i++
		if i == ncolors {
			fmt.Printf("\t%s\n", k)
		} else {
			fmt.Printf("\t%s,\n", k)
		}
	}

	fmt.Printf("};\n")

	fmt.Printf("const uint8_t %s_bitmap[] = {\n", definePrefix)
	fmt.Printf(bitmapArray)
	fmt.Printf("\n};\n")
}

func main() {
	var file string

	flag.StringVar(&file, "file", "", "filename.png")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: png2c -file=filename.png\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	if file == "" {
		flag.Usage()
	}

	png2c(file)
}
