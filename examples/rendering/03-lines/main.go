/*
 * This example creates an SDL window and renderer, and then draws some lines
 * to it every frame.
 *
 * This code is public domain. Feel free to use it for any purpose!
 */
package main

import (
	"math"
	"math/rand"
	"runtime"
	"sdl3/pkg/sdl"
)

/* We will use this renderer to draw into this window every frame. */
var window sdl.Window
var renderer sdl.Renderer

func main() {
	runtime.LockOSThread()

	if err := sdl.LoadLibrary(); err != nil {
		panic(err)
	}

	sdl.SetAppMetadata("Example Renderer Lines", "1.0", "com.example.renderer-lines")

	if !sdl.Init(sdl.InitVideo) {
		panic("Couldn't initialize SDL: " + sdl.GetError())
	}

	if !sdl.CreateWindowAndRenderer("examples/renderer/lines", 640, 480, 0, &window, &renderer) {
		panic("Couldn't create window/renderer: " + sdl.GetError())
	}

main_loop:
	for {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			switch event.Type {
			case sdl.EventQuit, sdl.EventWindowCloseRequested:
				break main_loop /* end the program, reporting success to the OS. */

			case sdl.EventKeyDown:
				keyEv := event.AsKeyboardEvent()

				switch keyEv.Scancode {
				case sdl.ScancodeEscape:
					break main_loop
				}
			}
		}

		/* Lines (line segments, really) are drawn in terms of points: a set of
		X and Y coordinates, one set for each end of the line.
		(0, 0) is the top left of the window, and larger numbers go down
		and to the right. This isn't how geometry works, but this is pretty
		standard in 2D graphics. */
		line_points := []sdl.FPoint{
			{100, 354}, {220, 230}, {140, 230}, {320, 100}, {500, 230},
			{420, 230}, {540, 354}, {400, 354}, {100, 354},
		}

		/* as you can see from this, rendering draws over whatever was drawn before it. */
		sdl.SetRenderDrawColor(renderer, 100, 100, 100, sdl.ALPHA_OPAQUE) /* grey, full alpha */
		sdl.RenderClear(renderer)                                         /* start with a blank canvas. */

		/* You can draw lines, one at a time, like these brown ones... */
		sdl.SetRenderDrawColor(renderer, 127, 49, 32, sdl.ALPHA_OPAQUE)
		sdl.RenderLine(renderer, 240, 450, 400, 450)
		sdl.RenderLine(renderer, 240, 356, 400, 356)
		sdl.RenderLine(renderer, 240, 356, 240, 450)
		sdl.RenderLine(renderer, 400, 356, 400, 450)

		/* You can also draw a series of connected lines in a single batch... */
		sdl.SetRenderDrawColor(renderer, 0, 255, 0, sdl.ALPHA_OPAQUE)
		sdl.RenderLines(renderer, line_points, int32(len(line_points)))

		/* here's a bunch of lines drawn out from a center point in a circle. */
		/* we randomize the color of each line, so it functions as animation. */
		for i := 0; i < 360; i++ {
			var size float32 = 30.0
			var x float32 = 320.0
			var y float32 = 95.0 - (size / 2.0)
			sdl.SetRenderDrawColor(renderer, uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), sdl.ALPHA_OPAQUE)
			sdl.RenderLine(renderer, x, y, x+float32(math.Sin(float64(i)))*size, y+float32(math.Cos(float64(i)))*size)
		}

		sdl.RenderPresent(renderer) /* put it all on the screen! */
	}
}
