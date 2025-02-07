// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate ./volcano

package plotter_test

import (
	"image/color"

	"github.com/emptywe/plot"
	"github.com/emptywe/plot/palette"
	"github.com/emptywe/plot/plotter"
	"github.com/emptywe/plot/vg"
	"github.com/emptywe/plot/vg/draw"
	"gonum.org/v1/gonum/mat"
)

type deciGrid struct{ mat.Matrix }

func (g deciGrid) Dims() (c, r int)   { r, c = g.Matrix.Dims(); return c, r }
func (g deciGrid) Z(c, r int) float64 { return g.Matrix.At(r, c) }
func (g deciGrid) X(c int) float64 {
	_, n := g.Matrix.Dims()
	if c < 0 || c >= n {
		panic("index out of range")
	}
	return 10 * float64(c)
}
func (g deciGrid) Y(r int) float64 {
	m, _ := g.Matrix.Dims()
	if r < 0 || r >= m {
		panic("index out of range")
	}
	return 10 * float64(r)
}

func Example_volcano() {
	var levels []float64
	for l := 100.5; l < mat.Max(volcano.Matrix.(*mat.Dense)); l += 5 {
		levels = append(levels, l)
	}
	c := plotter.NewContour(volcano, levels, palette.Rainbow(len(levels), (palette.Yellow+palette.Red)/2, palette.Blue, 1, 1, 1))
	quarterStyle := draw.LineStyle{
		Color:  color.Black,
		Width:  vg.Points(0.5),
		Dashes: []vg.Length{0.2, 0.4},
	}
	halfStyle := draw.LineStyle{
		Color:  color.Black,
		Width:  vg.Points(0.5),
		Dashes: []vg.Length{5, 2, 1, 2},
	}
	c.LineStyles = append(c.LineStyles, quarterStyle, halfStyle, quarterStyle)

	h := plotter.NewHeatMap(volcano, palette.Heat(len(levels)*2, 1))

	p := plot.New()
	p.Title.Text = "Maunga Whau Volcano"

	p.Add(h)
	p.Add(c)

	p.X.Padding = 0
	p.Y.Padding = 0
	_, p.X.Max, _, p.Y.Max = h.DataRange()

	if err := p.Save(10*vg.Centimeter, 10*vg.Centimeter, "testdata/volcano.png"); err != nil {
		panic(err)
	}
}
