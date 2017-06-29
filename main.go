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

package main

import (
	"os"
	"github.com/lxn/walk"
	"golang.org/x/sys/windows/registry"
	. "github.com/lxn/walk/declarative"
)

var window *walk.MainWindow
var events *walk.TextEdit
var progress *walk.ProgressBar

const harden_key_path = "SOFTWARE\\Security Without Borders\\"

func check_status() bool {
	key, err := registry.OpenKey(registry.CURRENT_USER, harden_key_path, registry.READ)
	if err != nil {
		return false
	}

	value, _, err := key.GetIntegerValue("Harden")
	if err != nil {
		return false
	}

	if value == 1 {
		return true
	} else {
		return false
	}
}

func mark_status(hardened bool) {
	key, _, err := registry.CreateKey(registry.CURRENT_USER, harden_key_path, registry.WRITE)
	if err != nil {
		panic(err)
	}

	if hardened {
		key.SetDWordValue("Harden", 1)
	} else {
		key.SetDWordValue("Harden", 0)
	}
}

func harden_all() {
	/*trigger_all(true)
	mark_status(true)
	*/
	walk.MsgBox(window, "Done!", "I have hardened all risky features!\nFor all changes to take effect please restart Windows.", walk.MsgBoxIconInformation)
	os.Exit(0)
}

func restore_all() {
	/*
	trigger_all(false)
	mark_status(false)
	*/

	walk.MsgBox(window, "Done!", "I have restored all risky features!\nFor all changes to take effect please restart Windows.", walk.MsgBoxIconExclamation)
	os.Exit(0)
}

func trigger_all(harden bool) {
	trigger_wsh(harden)
	trigger_ole(harden)
	trigger_macro(harden)
	trigger_activex(harden)
	trigger_pdf_js(harden)
	trigger_pdf_objects(harden)
	trigger_autorun(harden)
	trigger_powershell(harden)
	trigger_uac(harden)
	trigger_fileassoc(harden)
	progress.SetValue(100) 
}

func main() {
	var topBar *walk.Composite
    var sideBar *walk.Composite
    var mainArea *walk.Composite
    var middleArea *walk.Composite

	// XXX include this with something like gobindata
	hardenLogo, _ := walk.NewBitmapFromFile("assets/hardenlogo.png")
	// XXX YOLO errors
	darkGrey, _ := walk.NewSolidColorBrush(walk.RGB(74, 74, 74))
	lightGrey, _ := walk.NewSolidColorBrush(walk.RGB(155, 155, 155))
	lighterGrey, _ := walk.NewSolidColorBrush(walk.RGB(216, 216, 216))

	defer darkGrey.Dispose()
	defer lightGrey.Dispose()
	defer lighterGrey.Dispose()

	/*
	var label_text, button_text, events_text string
	var button_func func()

	if check_status() == false {
		button_text = "Harden!"
		button_func = harden_all
		label_text = "Ready to harden some features of your system?"
	} else {
		button_text = "Restore..."
		button_func = restore_all
		label_text = "We have already hardened some risky features, do you want to restore them?"
	}
	*/

	MainWindow{
		AssignTo: &window,
		Title: "Harden - Security Without Borders",
		MinSize: Size{800, 500},
		Layout: Grid{MarginsZero: true, SpacingZero: true, Rows: 2},
		Children: []Widget{
			Composite{
				Row: 0,
                AssignTo: &topBar,
				Layout: Grid{MarginsZero: true, SpacingZero: true, Columns: 6},
				MinSize: Size{Height: 85},
				MaxSize: Size{Height: 85},
                Children: []Widget{
					ImageView{
						// XXX this crap doesn't work
						// Background: darkGrey,
						Column: 0,
						MaxSize: Size{Width: 64},
						ColumnSpan: 1,
						Image: hardenLogo,
					},
					Composite{
						Column: 1,
						ColumnSpan: 1,
						Layout:   Grid{MarginsZero: true, SpacingZero: true, Rows: 2},
						Children: []Widget{
							Label{
								Row: 0,
								MaxSize: Size{Height: 40},
								Text: "Harden",
								Font: Font{PointSize: 24},
							},
							Label{
								Row: 1,
								MaxSize: Size{Height: 40},
								Text: "v1.1.1",
								Font: Font{PointSize: 16},
							},
						},
					},
					Label{
						Column: 3,
						Text: "Scan",
						Font: Font{PointSize: 32},
					},
					HSpacer{
						Column: 2,
						ColumnSpan: 3,
					},
                },
            },
			Composite{
				Row: 1,
				AssignTo: &middleArea,
				Layout: Grid{MarginsZero: true, SpacingZero: true, Columns: 6},
                Children: []Widget{
					Composite{
						Column: 0,
						AssignTo: &sideBar,
						Layout:   VBox{MarginsZero: true},
					},
					Composite{
						Column: 1,
						ColumnSpan: 5,
						AssignTo: &mainArea,
						Layout:   VBox{MarginsZero: true},
					},
                },
            },
			/*
			Label{Text: label_text},
			PushButton{
				Text: button_text,
				OnClicked: button_func,
			},
			ProgressBar{
				AssignTo: &progress,
			},
			TextEdit{
				AssignTo: &events,
				Text: events_text,
				ReadOnly: true,
			},
			*/
		},
	}.Create()

	topBar.SetBackground(darkGrey)
	sideBar.SetBackground(lightGrey)

	window.Run()
}
