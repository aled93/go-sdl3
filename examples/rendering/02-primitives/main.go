/*
 * This example creates an SDL window and renderer, and then draws some lines,
 * rectangles and points to it every frame.
 *
 * This code is public domain. Feel free to use it for any purpose!
 */
package main

import (
	"math/rand"
	"runtime"
	"sdl3/pkg/sdl"
)

/* We will use this renderer to draw into this window every frame. */
var window sdl.Window
var renderer sdl.Renderer
var points [500]sdl.FPoint

func main() {
	runtime.LockOSThread()

	if err := sdl.LoadLibrary(); err != nil {
		panic(err)
	}

	sdl.SetAppMetadata("Example Renderer Primitives", "1.0", "com.example.renderer-primitives")

	if !sdl.Init(sdl.InitVideo) {
		panic("Couldn't initialize SDL: " + sdl.GetError())
	}

	if !sdl.CreateWindowAndRenderer("examples/renderer/primitives", 640, 480, 0, &window, &renderer) {
		panic("Couldn't create window/renderer: " + sdl.GetError())
	}

	sdl.SetRenderVSync(renderer, sdl.VSyncAdaptive)

	/* set up some random points */
	for i := range points {
		points[i].X = (rand.Float32() * 440.0) + 100.0
		points[i].Y = (rand.Float32() * 280.0) + 100.0
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

		for i := range points {
			points[i].X = (rand.Float32() * 440.0) + 100.0
			points[i].Y = (rand.Float32() * 280.0) + 100.0
		}

		var rect sdl.FRect

		/* as you can see from this, rendering draws over whatever was drawn before it. */
		sdl.SetRenderDrawColor(renderer, 33, 33, 33, sdl.ALPHA_OPAQUE) /* dark gray, full alpha */
		sdl.RenderClear(renderer)                                      /* start with a blank canvas. */

		/* draw a filled rectangle in the middle of the canvas. */
		sdl.SetRenderDrawColor(renderer, 0, 0, 255, sdl.ALPHA_OPAQUE) /* blue, full alpha */
		rect.X = 100
		rect.Y = 100
		rect.W = 440
		rect.H = 280
		sdl.RenderFillRect(renderer, &rect)

		/* draw some points across the canvas. */
		sdl.SetRenderDrawColor(renderer, 255, 0, 0, sdl.ALPHA_OPAQUE) /* red, full alpha */
		sdl.RenderPoints(renderer, points[:], int32(len(points)))

		/* draw a unfilled rectangle in-set a little bit. */
		sdl.SetRenderDrawColor(renderer, 0, 255, 0, sdl.ALPHA_OPAQUE) /* green, full alpha */
		rect.X += 30
		rect.Y += 30
		rect.W -= 60
		rect.H -= 60
		sdl.RenderRect(renderer, &rect)

		/* draw two lines in an X across the whole canvas. */
		sdl.SetRenderDrawColor(renderer, 255, 255, 0, sdl.ALPHA_OPAQUE) /* yellow, full alpha */
		sdl.RenderLine(renderer, 0, 0, 640, 480)
		sdl.RenderLine(renderer, 0, 480, 640, 0)

		sdl.RenderPresent(renderer) /* put it all on the screen! */
	}
}
