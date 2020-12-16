// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmpimg

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"path/filepath"
	"testing"
)

const wantDiffEncoded = `iVBORw0KGgoAAAANSUhEUgAAAZAAAAEzEAIAAADAxR6YAAAHBklEQVR4nOzYjW3kRABA4RWiC0QXUMami2xRSRfrMu76gDKQZQ3+2907xEMg8X3SXZKxdzwZ29JTfrgAAJASWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAMYEFABATWAAAsR//7QX8X12v8//T9E/NvPje+R+tZjvP9882PjV/naZvf+J8lb96xb+/l+uar9d5zfPX8+rnY/PocvTZ9ebj4+j5TmyPjauNM8f8YyXrGh7NM2ZaPjXmHj8dZx4j+zUsax2zr0fHjGOG9frrOfs9WM/bjmx3bnxmP368/2MHzmtfj61j49/2+tv1bO/Ho1UdzzyO7898NsNx/NU+bHfkeMXZzz+Nnz4/3m/rPXi/He/B6r55Tt+evk3btZ73ffsGfH6MPfn8+PL1cvnydVnB779dLrfb/c/fa+zT9imaXhw9P3fH53H9fbdvwH6/lvdv/zSvT/B4e/dzL2e/7d7bZ98vOzqPvG3mWPZ42d375g1f5tzfgePI8Uk8P3GP1nAc+fWXeWT+f5rmOzH2dn5K9nt0/+b8j97ddYXnt+31vq2zHD+/3onz0XGH3m/P3tBnb8szr2Z5vdOvZzyv6Di+fDdNt9t5Dn/BAgCICSwAgJjAAgCICSwAgJjAAgCICSwAgJjAAgCICSwAgJjAAgCICSwAgJjAAgAAAOC/zV+wAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAICawAABiAgsAIPZHAAAA///KAlB2mCuyaAAAAABJRU5ErkJggg==`

func TestDiff(t *testing.T) {
	got, err := ioutil.ReadFile(filepath.FromSlash("./testdata/failed_input.png"))
	if err != nil {
		t.Fatalf("failed to read failed file: %v", err)
	}
	want, err := ioutil.ReadFile(filepath.FromSlash("./testdata/good_golden.png"))
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	v1, _, err := image.Decode(bytes.NewReader(got))
	if err != nil {
		t.Fatalf("unexpected error decoding failed file: %v", err)
	}
	v2, _, err := image.Decode(bytes.NewReader(want))
	if err != nil {
		t.Fatalf("unexpected error decoding golden file: %v", err)
	}

	dst := image.NewRGBA64(v1.Bounds().Union(v2.Bounds()))
	rect := Diff(dst, v1, v2)
	if rect != dst.Bounds() {
		t.Errorf("unexpected bound for diff: got:%+v want:%+v", rect, dst.Bounds())
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, dst)
	if err != nil {
		t.Fatalf("failed to encode difference png: %v", err)
	}
	gotDiff := base64.StdEncoding.EncodeToString(buf.Bytes())
	if gotDiff != wantDiffEncoded {
		t.Errorf("unexpected encoded diff value:\ngot:%s\nwant:%s", gotDiff, wantDiffEncoded)
	}
}

func TestEqual(t *testing.T) {
	got, err := ioutil.ReadFile("testdata/approx_got_golden.png")
	if err != nil {
		t.Fatal(err)
	}

	ok, err := Equal("png", got, got)
	if err != nil {
		t.Fatalf("could not compare images: %+v", err)
	}
	if !ok {
		t.Fatalf("same image does not compare equal")
	}
}

func TestEqualApprox(t *testing.T) {
	got, err := ioutil.ReadFile("testdata/approx_got_golden.png")
	if err != nil {
		t.Fatal(err)
	}

	want, err := ioutil.ReadFile("testdata/approx_want_golden.png")
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		delta uint8
		ok    bool
	}{
		{0, false},
		{1, false},
		{2, false},
		{3, false},
		{4, false},
		{5, false},
		{6, false},
		{7, true},
		{8, true},
		{9, true},
		{10, true},
		{255, true},
	} {
		t.Run(fmt.Sprintf("delta=%d", int(tc.delta)), func(t *testing.T) {
			ok, err := EqualApprox("png", got, want, tc.delta)
			if err != nil {
				t.Fatalf("could not compare images: %+v", err)
			}
			if ok != tc.ok {
				t.Fatalf("got=%v, want=%v", ok, tc.ok)
			}
		})
	}
}
