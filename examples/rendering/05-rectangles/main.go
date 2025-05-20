/*
 * This example creates an SDL window and renderer, and then draws some
 * rectangles to it every frame.
 *
 * This code is public domain. Feel free to use it for any purpose!
 */
package main

import (
	"runtime"
	"sdl3/pkg/sdl"
)

/* We will use this renderer to draw into this window every frame. */
var window sdl.Window
var renderer sdl.Renderer

const WINDOW_WIDTH = 640
const WINDOW_HEIGHT = 480

func main() {
	runtime.LockOSThread()

	if err := sdl.LoadLibrary(); err != nil {
		panic(err)
	}

	sdl.SetAppMetadata("Example Renderer Rectangles", "1.0", "com.example.renderer-rectangles")

	if !sdl.Init(sdl.InitVideo) {
		panic("Couldn't initialize SDL: " + sdl.GetError())
	}

	if !sdl.CreateWindowAndRenderer("examples/renderer/rectangles", WINDOW_WIDTH, WINDOW_HEIGHT, 0, &window, &renderer) {
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

		var rects [16]sdl.FRect
		now := sdl.GetTicks()

		/* we'll have the rectangles grow and shrink over a few seconds. */
		var direction float32

		if (now % 2000) >= 1000 {
			direction = 1.0
		} else {
			direction = -1.0
		}

		scale := (float32((int(now%1000))-500) / 500.0) * direction

		/* as you can see from this, rendering draws over whatever was drawn before it. */
		sdl.SetRenderDrawColor(renderer, 0, 0, 0, sdl.ALPHA_OPAQUE) /* black, full alpha */
		sdl.RenderClear(renderer)                                   /* start with a blank canvas. */

		/* Rectangles are comprised of set of X and Y coordinates, plus width and
		height. (0, 0) is the top left of the window, and larger numbers go
		down and to the right. This isn't how geometry works, but this is
		pretty standard in 2D graphics. */

		/* Let's draw a single rectangle (square, really). */
		rects[0].X = 100
		rects[0].Y = 100
		rects[0].W = 100 + (100 * scale)
		rects[0].H = 100 + (100 * scale)
		sdl.SetRenderDrawColor(renderer, 255, 0, 0, sdl.ALPHA_OPAQUE) /* red, full alpha */
		sdl.RenderRect(renderer, &rects[0])

		/* Now let's draw several rectangles with one function call. */
		for i := 0; i < 3; i++ {
			var size float32 = float32(i+1) * 50.0
			rects[i].W = size + (size * scale)
			rects[i].H = size + (size * scale)
			rects[i].X = (WINDOW_WIDTH - rects[i].W) / 2  /* center it. */
			rects[i].Y = (WINDOW_HEIGHT - rects[i].H) / 2 /* center it. */
		}
		sdl.SetRenderDrawColor(renderer, 0, 255, 0, sdl.ALPHA_OPAQUE) /* green, full alpha */
		sdl.RenderRects(renderer, rects[:], 3)                        /* draw three rectangles at once */

		/* those were rectangle _outlines_, really. You can also draw _filled_ rectangles! */
		rects[0].X = 400
		rects[0].Y = 50
		rects[0].W = 100 + (100 * scale)
		rects[0].H = 50 + (50 * scale)
		sdl.SetRenderDrawColor(renderer, 0, 0, 255, sdl.ALPHA_OPAQUE) /* blue, full alpha */
		sdl.RenderFillRect(renderer, &rects[0])

		/* ...and also fill a bunch of rectangles at once... */
		for i := 0; i < len(rects); i++ {
			w := float32(WINDOW_WIDTH / len(rects))
			h := float32(i) * 8.0
			rects[i].X = float32(i) * w
			rects[i].Y = WINDOW_HEIGHT - h
			rects[i].W = w
			rects[i].H = h
		}
		sdl.SetRenderDrawColor(renderer, 255, 255, 255, sdl.ALPHA_OPAQUE) /* white, full alpha */
		sdl.RenderFillRects(renderer, rects[:], int32(len(rects)))

		sdl.RenderPresent(renderer) /* put it all on the screen! */
	}
}
