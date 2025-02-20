// Copyright ©2020 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vggio_test

import (
	"image/color"
	"math"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"

	"github.com/emptywe/plot"
	"github.com/emptywe/plot/plotter"
	"github.com/emptywe/plot/vg"
	"github.com/emptywe/plot/vg/draw"
	"github.com/emptywe/plot/vg/vggio"
)

func ExampleCanvas() {
	const (
		w   = 20 * vg.Centimeter
		h   = 15 * vg.Centimeter
		dpi = 96
	)
	go func(w, h vg.Length) {
		win := app.NewWindow(
			app.Title("Gonum"),
			app.Size(
				unit.Px(float32(w.Dots(dpi))),
				unit.Px(float32(h.Dots(dpi))),
			),
		)
		defer win.Close()

		done := time.NewTimer(2 * time.Second)
		defer done.Stop()

		for e := range win.Events() {
			switch e := e.(type) {
			case system.FrameEvent:
				p := plot.New()
				p.Title.Text = "My title"
				p.X.Label.Text = "X"
				p.Y.Label.Text = "Y"

				quad := plotter.NewFunction(func(x float64) float64 {
					return x * x
				})
				quad.Color = color.RGBA{B: 255, A: 255}

				exp := plotter.NewFunction(func(x float64) float64 {
					return math.Pow(2, x)
				})
				exp.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
				exp.Width = vg.Points(2)
				exp.Color = color.RGBA{G: 255, A: 255}

				sin := plotter.NewFunction(func(x float64) float64 {
					return 10*math.Sin(x) + 50
				})
				sin.Dashes = []vg.Length{vg.Points(4), vg.Points(5)}
				sin.Width = vg.Points(4)
				sin.Color = color.RGBA{R: 255, A: 255}

				p.Add(quad, exp, sin)
				p.Legend.Add("x^2", quad)
				p.Legend.Add("2^x", exp)
				p.Legend.Add("10*sin(x)+50", sin)
				p.Legend.ThumbnailWidth = 0.5 * vg.Inch

				p.X.Min = 0
				p.X.Max = 10
				p.Y.Min = 0
				p.Y.Max = 100

				p.Add(plotter.NewGrid())

				gtx := layout.NewContext(new(op.Ops), e)
				cnv := vggio.New(gtx, w, h, vggio.UseDPI(dpi))
				p.Draw(draw.New(cnv))

				e.Frame(cnv.Paint())

			case key.Event:
				switch e.Name {
				case "Q", key.NameEscape:
					win.Close()
				}

			case system.DestroyEvent:
				os.Exit(0)
			}
		}
	}(w, h)

	app.Main()
}
