// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter_test

import (
	"log"

	"github.com/emptywe/plot"
	"github.com/emptywe/plot/plotter"
	"github.com/emptywe/plot/vg"
)

// Draw the plot logo.
func Example() {
	p := plot.New()

	plotter.DefaultLineStyle.Width = vg.Points(1)
	plotter.DefaultGlyphStyle.Radius = vg.Points(3)

	p.Y.Tick.Marker = plot.ConstantTicks([]plot.Tick{
		{Value: 0, Label: "0"}, {Value: 0.25, Label: ""}, {Value: 0.5, Label: "0.5"}, {Value: 0.75, Label: ""}, {Value: 1, Label: "1"},
	})
	p.X.Tick.Marker = plot.ConstantTicks([]plot.Tick{
		{Value: 0, Label: "0"}, {Value: 0.25, Label: ""}, {Value: 0.5, Label: "0.5"}, {Value: 0.75, Label: ""}, {Value: 1, Label: "1"},
	})

	pts := plotter.XYs{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0.5, Y: 1}, {X: 0.5, Y: 0.6}, {X: 0, Y: 0.6}}
	line, err := plotter.NewLine(pts)
	if err != nil {
		log.Panic(err)
	}
	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		log.Panic(err)
	}
	p.Add(line, scatter)

	pts = plotter.XYs{{X: 1, Y: 0}, {X: 0.75, Y: 0}, {X: 0.75, Y: 0.75}}
	line, err = plotter.NewLine(pts)
	if err != nil {
		log.Panic(err)
	}
	scatter, err = plotter.NewScatter(pts)
	if err != nil {
		log.Panic(err)
	}
	p.Add(line, scatter)

	pts = plotter.XYs{{X: 0.5, Y: 0.5}, {X: 1, Y: 0.5}}
	line, err = plotter.NewLine(pts)
	if err != nil {
		log.Panic(err)
	}
	scatter, err = plotter.NewScatter(pts)
	if err != nil {
		log.Panic(err)
	}
	p.Add(line, scatter)

	err = p.Save(100, 100, "testdata/plotLogo.png")
	if err != nil {
		log.Panic(err)
	}
}
