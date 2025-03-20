package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/obzva/gato"
)

func main() {
	// if verbose, log processing time at the end of the program
	startTime := time.Now()

	// flags
	wPtr := flag.Int("w", 0, "desired width of output image, defaults to keep the ratio of the original image when omitted (at least one of two, width or height, is required)")
	hPtr := flag.Int("h", 0, "desired height of output image, defaults to keep the ratio of the original image when omitted (at least one of two, width or height, is required)")
	methodPtr := flag.String("m", "bilinear", "desired interpolation method, defaults to bilinear (options: nearest-neighbor, bilinear, and bicubic)")
	outputPtr := flag.String("o", "", "desired output filename, defaults to '[input filename]-[method name].jpg'")
	verbosePtr := flag.Bool("v", false, "verbose mode to log processing time")
	flag.Parse()

	// non flag arguments: input image filename
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("input image filename is required")
	}
	input := args[0]

	// compile regex for filenames
	re, err := regexp.Compile(`^(.+)\.([^.]+)$`)
	if err != nil {
		log.Fatal(err)
	}

	// check input filename
	matches := re.FindStringSubmatch(input)
	if len(matches) != 3 {
		log.Fatalf("invalid input filename: %s", input)
	}
	inputName := matches[0]

	// open image
	img, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()

	// create gato.Data
	data, err := gato.NewData(input, img)
	if err != nil {
		log.Fatal(err)
	}

	// create gato.Processor
	prc, err := gato.NewProcessor(gato.Instruction{
		Width:         *wPtr,
		Height:        *hPtr,
		Interpolation: *methodPtr,
	})
	if err != nil {
		log.Fatal(err)
	}

	// process image
	res, err := prc.Process(data)
	if err != nil {
		log.Fatal(err)
	}

	// determine output format and filename
	var outFormat string // jpeg, png
	if *outputPtr == "" {
		*outputPtr = fmt.Sprintf("%s-%s.jpg", inputName, *methodPtr)
		outFormat = data.Format
	} else {
		matches := re.FindStringSubmatch(*outputPtr)
		if len(matches) != 3 {
			log.Fatalf("invalid output filename: %s", *outputPtr)
		}
		outFormat = matches[2]
	}

	// create output file
	out, err := os.Create(*outputPtr)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// encode image
	switch outFormat {
	case "jpeg":
		if err := jpeg.Encode(out, res, nil); err != nil {
			log.Fatal(err)
		}
	case "png":
		if err := png.Encode(out, res); err != nil {
			log.Fatal(err)
		}
	}

	// if verbose, print info
	if *verbosePtr {
		log.Printf("interpolation method: %s\n", *methodPtr)
		log.Printf("processing time: %v\n", time.Since(startTime))
	}
}
