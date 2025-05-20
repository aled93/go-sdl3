/*
 * This example creates an SDL window and renderer, and then draws some points
 * to it every frame.
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
var last_time uint64 = 0

const WINDOW_WIDTH = 640
const WINDOW_HEIGHT = 480

const NUM_POINTS = 500
const MIN_PIXELS_PER_SECOND = 30 /* move at least this many pixels per second. */
const MAX_PIXELS_PER_SECOND = 60 /* move this many pixels per second at most. */

/* (track everything as parallel arrays instead of a array of structs,
   so we can pass the coordinates to the renderer in a single function call.) */

/*
Points are plotted as a set of X and Y coordinates.

	(0, 0) is the top left of the window, and larger numbers go down
	and to the right. This isn't how geometry works, but this is pretty
	standard in 2D graphics.
*/
var points [NUM_POINTS]sdl.FPoint
var point_speeds [NUM_POINTS]float32

func main() {
	runtime.LockOSThread()

	if err := sdl.LoadLibrary(); err != nil {
		panic(err)
	}

	sdl.SetAppMetadata("Example Renderer Points", "1.0", "com.example.renderer-points")

	if !sdl.Init(sdl.InitVideo) {
		panic("Couldn't initialize SDL: " + sdl.GetError())
	}

	if !sdl.CreateWindowAndRenderer("examples/renderer/points", WINDOW_WIDTH, WINDOW_HEIGHT, 0, &window, &renderer) {
		panic("Couldn't create window/renderer: " + sdl.GetError())
	}

	/* set up the data for a bunch of points. */
	for i := 0; i < len(points); i++ {
		points[i].X = rand.Float32() * (float32(WINDOW_WIDTH))
		points[i].Y = rand.Float32() * (float32(WINDOW_HEIGHT))
		point_speeds[i] = MIN_PIXELS_PER_SECOND + (rand.Float32() * (MAX_PIXELS_PER_SECOND - MIN_PIXELS_PER_SECOND))
	}

	last_time = sdl.GetTicks()

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

		now := sdl.GetTicks()
		elapsed := (float32((now - last_time))) / 1000.0 /* seconds since last iteration */

		/* let's move all our points a little for a new frame. */
		for i := 0; i < len(points); i++ {
			distance := elapsed * point_speeds[i]
			points[i].X += distance
			points[i].Y += distance
			if (points[i].X >= WINDOW_WIDTH) || (points[i].Y >= WINDOW_HEIGHT) {
				/* off the screen restart it elsewhere! */
				if rand.Intn(2) == 0 {
					points[i].X = rand.Float32() * float32(WINDOW_WIDTH)
					points[i].Y = 0.0
				} else {
					points[i].X = 0.0
					points[i].Y = rand.Float32() * float32(WINDOW_HEIGHT)
				}
				point_speeds[i] = MIN_PIXELS_PER_SECOND + (rand.Float32() * (MAX_PIXELS_PER_SECOND - MIN_PIXELS_PER_SECOND))
			}
		}

		last_time = now

		/* as you can see from this, rendering draws over whatever was drawn before it. */
		sdl.SetRenderDrawColor(renderer, 0, 0, 0, sdl.ALPHA_OPAQUE)       /* black, full alpha */
		sdl.RenderClear(renderer)                                         /* start with a blank canvas. */
		sdl.SetRenderDrawColor(renderer, 255, 255, 255, sdl.ALPHA_OPAQUE) /* white, full alpha */
		sdl.RenderPoints(renderer, points[:], int32(len(points)))         /* draw all the points! */

		/* You can also draw single points with sdl.RenderPoint(), but it's
		cheaper (sometimes significantly so) to do them all at once. */

		sdl.RenderPresent(renderer) /* put it all on the screen! */
	}
}
