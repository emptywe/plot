// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter_test

import (
	"testing"

	"github.com/emptywe/plot/cmpimg"
)

func TestLogScale(t *testing.T) {
	cmpimg.CheckPlot(Example_logScale, t, "logscale.png")
}
