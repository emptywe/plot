// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recorder

import (
	"image"
	"image/color"
	"strings"
	"testing"

	"github.com/emptywe/plot/font"
	"github.com/emptywe/plot/vg"
)

func TestRecorder(t *testing.T) {
	tr := font.Font{Typeface: "Liberation", Variant: "Serif"}

	var rec Canvas
	rec.Actions = append(rec.Actions, &FillString{Font: tr, Size: 12, Point: vg.Point{X: 0, Y: 10}, String: "Text"})
	rec.Comment("End of preamble")
	rec.Scale(1, 2)
	rec.Rotate(0.72)
	rec.KeepCaller = true
	rec.Stroke(vg.Path{{Type: vg.MoveComp, Pos: vg.Point{X: 3, Y: 4}}})
	rec.Push()
	rec.Pop()
	rec.Translate(vg.Point{X: 3, Y: 4})
	rec.KeepCaller = false
	rec.SetLineWidth(100)
	rec.SetLineDash([]font.Length{2, 5}, 6)
	rec.SetColor(color.RGBA{R: 0x65, G: 0x23, B: 0xf2})
	rec.Fill(vg.Path{{Type: vg.MoveComp, Pos: vg.Point{X: 3, Y: 4}}, {Type: vg.LineComp, Pos: vg.Point{X: 2, Y: 3}}, {Type: vg.CloseComp}})
	rec.DrawImage(
		vg.Rectangle{
			Min: vg.Point{X: 0, Y: 0},
			Max: vg.Point{X: 10, Y: 10},
		},
		img,
	)
	if len(rec.Actions) != len(want) {
		t.Fatalf("unexpected number of actions recorded: got:%d want:%d", len(rec.Actions), len(want))
	}
	for i, a := range rec.Actions {
		if got := a.Call(); !strings.HasSuffix(got, want[i]) {
			t.Errorf("unexpected action:\n\tgot: %#v\n\twant: %#v", got, want[i])
		}
	}

	var replay Canvas
	err := rec.ReplayOn(&replay)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	for i, a := range rec.Actions {
		got := replay.Actions[i].Call()
		want := a.Call()
		if !strings.HasSuffix(want, got) {
			t.Errorf("unexpected action:\n\tgot: %#v\n\twant: %#v", got, want)
		}
	}

	replay.Reset()
	rec.Actions = append(rec.Actions, &FillString{Font: font.Font{Typeface: "Foo"}, Size: 12, Point: vg.Point{X: 0, Y: 10}, String: "Bar"})
	err = rec.ReplayOn(&replay)
	switch {
	case err == nil:
		t.Errorf("expected an error")
	case !strings.HasPrefix(err.Error(), "unknown font: Foo"):
		t.Errorf("unexpected error: %v", err)
	}
}

var img image.Image = image.NewGray(image.Rect(0, 0, 20, 20))

var want = []string{
	`FillString("LiberationSerif-Regular", 12, 0, 10, "Text")`,
	`Comment("End of preamble")`,
	`Scale(1, 2)`,
	`Rotate(0.72)`,
	`github.com/emptywe/plot/vg/recorder/recorder_test.go:26 Stroke(vg.Path{vg.PathComp{Type:0, Pos:vg.Point{X:3, Y:4}, Control:[]vg.Point(nil), Radius:0, Start:0, Angle:0}})`,
	`github.com/emptywe/plot/vg/recorder/recorder_test.go:27 Push()`,
	`github.com/emptywe/plot/vg/recorder/recorder_test.go:28 Pop()`,
	`github.com/emptywe/plot/vg/recorder/recorder_test.go:29 Translate(3, 4)`,
	`SetLineWidth(100)`,
	`SetLineDash([]font.Length{2, 5}, 6)`,
	`SetColor(color.RGBA{R:0x65, G:0x23, B:0xf2, A:0x0})`,
	`Fill(vg.Path{vg.PathComp{Type:0, Pos:vg.Point{X:3, Y:4}, Control:[]vg.Point(nil), Radius:0, Start:0, Angle:0}, vg.PathComp{Type:1, Pos:vg.Point{X:2, Y:3}, Control:[]vg.Point(nil), Radius:0, Start:0, Angle:0}, vg.PathComp{Type:4, Pos:vg.Point{X:0, Y:0}, Control:[]vg.Point(nil), Radius:0, Start:0, Angle:0}})`,
	`DrawImage(vg.Rectangle{Min:vg.Point{X:0, Y:0}, Max:vg.Point{X:10, Y:10}}, {image.Rectangle{Min:image.Point{X:0, Y:0}, Max:image.Point{X:20, Y:20}}, IMAGE:iVBORw0KGgoAAAANSUhEUgAAABQAAAAUCAAAAACo4kLRAAAAFElEQVR4nGJiwAJGBQeVICAAAP//JBgAKeMueQ8AAAAASUVORK5CYII=})`,
}
