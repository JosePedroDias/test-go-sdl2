// based on: https://github.com/veandco/go-sdl2/blob/master/examples/render.go
// author: José Pedro Dias
// reference: http://wiki.libsdl.org/APIByCategory
// reference: ~/go/src/github.com/veandco/go-sdl2/sdl/*.go

package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

var winTitle string = "testing go sdl 2"
var winWidth, winHeight int = 800, 600

func main() {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var rect sdl.Rect
	//var points []sdl.Point
	//var rects []sdl.Rect

	window = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if window == nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", sdl.GetError())
		os.Exit(1)
	}

	//window.SetFullscreen(sdl.WINDOW_FULLSCREEN)

	renderer = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if renderer == nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", sdl.GetError())
		os.Exit(2)
	}

	renderer.SetDrawColor(64, 64, 64, 255)
	rect = sdl.Rect{0, 0, int32(winWidth), int32(winHeight)}
	renderer.FillRect(&rect)

	renderer.SetDrawColor(255, 0, 0, 255) // set color opaque red
	renderer.DrawPoint(150, 300)          // set pixel at 150, 300

	/*	renderer.SetDrawColor(255, 255, 255, 255)
		renderer.DrawPoint(150, 300)

		renderer.SetDrawColor(0, 0, 255, 255)
		renderer.DrawLine(0, 0, 200, 200)

		points = []sdl.Point{{0, 0}, {100, 300}, {100, 300}, {200, 0}}
		renderer.SetDrawColor(255, 255, 0, 255)
		renderer.DrawLines(points)

		rect = sdl.Rect{300, 0, 200, 200}
		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.DrawRect(&rect)

		rects = []sdl.Rect{{400, 400, 100, 100}, {550, 350, 200, 200}}
		renderer.SetDrawColor(0, 255, 255, 255)
		renderer.DrawRects(rects)

		rect = sdl.Rect{250, 250, 200, 200}
		renderer.SetDrawColor(0, 255, 0, 255)
		renderer.FillRect(&rect)

		rects = []sdl.Rect{{500, 300, 100, 100}, {200, 300, 200, 200}}
		renderer.SetDrawColor(255, 0, 255, 255)
		renderer.FillRects(rects)*/

	renderer.Present()

	sdl.Delay(2000)

	renderer.Destroy()
	window.Destroy()
}
