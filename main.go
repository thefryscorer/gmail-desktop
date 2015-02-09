package main

import (
	"fmt"
	"os"

	"gopkg.in/qml.v1"
)

var webview string
var addr string

func main() {
	webview = `
	import QtQuick 2.0
	import QtWebKit 3.0
	import Material 0.1

	ApplicationWindow {
		title: "Gmail"
		id: window
		flags: "FramelessWindowHint"
	    	width: 460
    		height: 640
    		visible: true


		theme {
			primaryColor: "#FF1744"
			accentColor: "#607D8B"
		}

		initialPage: page

		Page {
			id: page

			actions: [
				Action {
					iconName: "navigation/close"
					name: "Close"
					onTriggered: ctrl.quit()
				}
			]

			title: "Gmail"
			WebView {
				anchors.fill: parent
				url: "http://mail.google.com/mail/mu"
			}
		}
	}`
	qml.Run(run)
}

func run() error {

	engine := qml.NewEngine()
	component, err := engine.LoadString("webview.qml", webview)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	ctrl := &Control{
		win: component.CreateWindow(nil),
		quit: func() {
			os.Exit(0)
		},
	}

	engine.Context().SetVar("ctrl", ctrl)
	ctrl.win.Wait()
	return nil
}

func (ctrl *Control) Quit() {
	os.Exit(0)
}

type Control struct {
	win  *qml.Window
	quit func()
}
