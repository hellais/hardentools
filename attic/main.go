// Hardentools
// Copyright (C) 2017  Security Without Borders
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.


// go get golang.org/x/image/font/inconsolata
// go get golang.org/x/mobile/gl

package main

import (
	_ "os"
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/exp/shiny/driver"
	_ "golang.org/x/exp/shiny/imageutil"
	"golang.org/x/exp/shiny/screen"

	"golang.org/x/exp/shiny/unit"
	"golang.org/x/exp/shiny/widget"
	"golang.org/x/exp/shiny/widget/node"
	"golang.org/x/exp/shiny/widget/theme"
)



// XXX include this with something like gobindata
//var hardenLogo, _ := walk.NewBitmapFromFile("assets/hardenlogo.png")

type EventPool struct {
	events []string
}

func (e *EventPool) AppendText(s string) {
	//
}

var events = EventPool{}
var darkGrey = color.RGBA{74, 74, 74, 0xff}
var lightGrey = color.RGBA{155, 155, 155, 0xff}
var lighterGrey = color.RGBA{216, 216, 216, 0xff}

var px = unit.Pixels

func colorPatch(c color.Color, w, h unit.Value) *widget.Sizer {
	return widget.NewSizer(w, h, widget.NewUniform(theme.StaticColor(c), nil))
}


func main() {
	t := theme.Default

	driver.Main(func(s screen.Screen) {
		hf := widget.NewFlow(widget.AxisHorizontal,
			widget.NewLabel("Cyan:"),
			widget.WithLayoutData(
				colorPatch(color.RGBA{0x00, 0x7f, 0x7f, 0xff}, px(0), px(20)),
				widget.FlowLayoutData{AlongWeight: 1, ExpandAlong: true},
			),
			widget.NewLabel("Magenta:"),
			widget.WithLayoutData(
				colorPatch(color.RGBA{0x7f, 0x00, 0x7f, 0xff}, px(0), px(30)),
				widget.FlowLayoutData{AlongWeight: 2, ExpandAlong: true},
			),
			widget.NewLabel("Yellow:"),
			widget.WithLayoutData(
				colorPatch(color.RGBA{0x7f, 0x7f, 0x00, 0xff}, px(0), px(40)),
				widget.FlowLayoutData{AlongWeight: 3, ExpandAlong: true},
			),

		)
		vf := widget.NewFlow(widget.AxisVertical,
			colorPatch(color.RGBA{0xff, 0x00, 0x00, 0xff}, px(80), px(40)),
			colorPatch(color.RGBA{0x00, 0xff, 0x00, 0xff}, px(50), px(50)),
			colorPatch(color.RGBA{0x00, 0x00, 0xff, 0xff}, px(20), px(60)),
			widget.WithLayoutData(
				hf,
				widget.FlowLayoutData{ExpandAcross: true},
			),
			widget.NewLabel(fmt.Sprintf(
				"The black rectangle is 1.5 inches x 1 inch when viewed at %v DPI.", t.GetDPI())),
				widget.NewPadder(widget.AxisBoth, unit.Pixels(8),
				colorPatch(color.Black, unit.Inches(1.5), unit.Inches(1)),
			),
		)
		var width, height = 640, 480
		rgba := image.NewRGBA(image.Rect(0, 0, width, height))
		draw.Draw(rgba, rgba.Bounds(), t.GetPalette().Neutral(), image.Point{}, draw.Src)
		vf.Measure(t, width, height)
		vf.Rect = rgba.Bounds()
		vf.Layout(t)
		vf.PaintBase(&node.PaintBaseContext{
			Theme: t,
			Dst:   rgba,
		}, image.Point{})

		widget.RunWindow(s, vf, &widget.RunWindowOptions{
			NewWindowOptions: screen.NewWindowOptions{
				Title: "WidgetGallery Shiny Example",
			},
		})
	})
}
