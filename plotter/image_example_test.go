// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter_test

import (
	"image/png"
	"log"
	"os"

	"github.com/emptywe/plot"
	"github.com/emptywe/plot/plotter"
	"github.com/emptywe/plot/vg"
)

// An example of embedding an image in a plot.
func ExampleImage() {
	p := plot.New()
	p.Title.Text = "A Logo"

	// load an image
	f, err := os.Open("testdata/image_plot_input.png")
	if err != nil {
		log.Fatalf("error opening image file: %v\n", err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Fatalf("error decoding image file: %v\n", err)
	}

	p.Add(plotter.NewImage(img, 100, 100, 200, 200))

	const (
		w = 5 * vg.Centimeter
		h = 5 * vg.Centimeter
	)

	err = p.Save(w, h, "testdata/image_plot.png")
	if err != nil {
		log.Fatalf("error saving image plot: %v\n", err)
	}
}

// An example of embedding an image in a plot with non-linear axes.
func ExampleImage_log() {
	p := plot.New()
	p.Title.Text = "A Logo"

	// load an image
	f, err := os.Open("testdata/gopher.png")
	if err != nil {
		log.Fatalf("error opening image file: %v\n", err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Fatalf("error decoding image file: %v\n", err)
	}

	p.Add(plotter.NewImage(img, 100, 100, 10000, 10000))

	// Transform axes.
	p.X.Scale = plot.LogScale{}
	p.Y.Scale = plot.LogScale{}
	p.X.Tick.Marker = plot.LogTicks{Prec: -1}
	p.Y.Tick.Marker = plot.LogTicks{Prec: 3}

	const (
		w = 5 * vg.Centimeter
		h = 5 * vg.Centimeter
	)

	err = p.Save(w, h, "testdata/image_plot_log.png")
	if err != nil {
		log.Fatalf("error saving image plot: %v\n", err)
	}
}
