// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package palette_test

import (
	"log"
	"strconv"

	"github.com/emptywe/plot"
	"github.com/emptywe/plot/palette"
	"github.com/emptywe/plot/palette/moreland"
	"github.com/emptywe/plot/plotter"
)

// This example creates a color bar and a second color bar where the
// direction of the colors are reversed.
func ExampleReverse() {
	p := plot.New()
	l := &plotter.ColorBar{ColorMap: moreland.Kindlmann()}
	l2 := &plotter.ColorBar{ColorMap: palette.Reverse(moreland.Kindlmann())}
	l.ColorMap.SetMin(0.5)
	l.ColorMap.SetMax(2.5)
	l2.ColorMap.SetMin(2.5)
	l2.ColorMap.SetMax(4.5)

	p.Add(l, l2)
	p.HideY()
	p.X.Padding = 0
	p.Title.Text = "A ColorMap and its Reverse"

	if err := p.Save(300, 48, "testdata/reverse.png"); err != nil {
		log.Panic(err)
	}
}

// This example creates a color palette from a reversed ColorMap.
func ExampleReverse_palette() {
	p := plot.New()
	thumbs := plotter.PaletteThumbnailers(palette.Reverse(moreland.Kindlmann()).Palette(10))
	for i, t := range thumbs {
		p.Legend.Add(strconv.Itoa(i), t)
	}
	p.HideAxes()
	p.X.Padding = 0
	p.Y.Padding = 0

	if err := p.Save(35, 120, "testdata/reverse_palette.png"); err != nil {
		log.Panic(err)
	}
}
