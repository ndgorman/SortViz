package main

import (
	"bytes"
	"fmt"
	"github.com/wcharczuk/go-chart"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/png"
	"log"
	"math/rand"
	"os"
)

func createShuffledNumberLine(length int) []float64 {
	numberLine := createNumberLine(length)
	rand.Shuffle(len(numberLine), func(a, b int) {
		numberLine[a], numberLine[b] = numberLine[b], numberLine[a]
	})

	return numberLine
}

func createNumberLine(length int) []float64 {
	numberLine := make([]float64, length, length)
	for i := range numberLine {
		numberLine[i] = float64(i)
	}
	return numberLine
}

func makeChart(yValues *[]float64) chart.Chart {
	x := createNumberLine(len(*yValues))

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
					Show:        true,
				},
				XValues: x,
				YValues: *yValues,
			},
		},
	}
	return graph
}

func createGif(steps int, images []*image.Paletted) gif.GIF {
	delays := make([]int, 0, steps)
	g := gif.GIF{Image: images, Delay: delays}
	return g
}

func chartToPaletted(c *chart.Chart) (*image.Paletted, error) {
	graph := *c
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		fmt.Printf("Error rendering chart: %v", err)
		return nil, err
	}
	img, err := png.Decode(buffer)
	if err != nil {
		fmt.Printf("Error decoding png: %v", err)
		return nil, err
	}
	pi := image.NewPaletted(img.Bounds(), palette.Plan9)
	draw.Draw(pi, img.Bounds(), img, image.Point{}, 3)

	return pi, err
}

func main() {
	initial := createShuffledNumberLine(100)
	stepChan := make(chan []float64, 0)
	pi := make([]*image.Paletted, 0, 20)

	go func() {
		quickSort(&initial, 0, len(initial)-1, stepChan)
		close(stepChan)
	}()
	for i := range stepChan {
		c := makeChart(&i)
		imgP, err := chartToPaletted(&c)
		if err != nil {
			log.Fatal(err)
		}
		pi = append(pi, imgP)
	}
	delays := make([]int, len(pi))
	for s := range delays {
		delays[s] = 200
	}
	g := gif.GIF{Image: pi, Delay: delays}
	file, err := os.Create("sort.gif")
	if err != nil {
		log.Fatal(err)
	}
	err = gif.EncodeAll(file, &g)
	if err != nil {
		log.Fatal(err)
	}

}
