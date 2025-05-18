/*
 * This example code creates an SDL window and renderer, and then clears the
 * window to a different color every frame, so you'll effectively get a window
 * that's smoothly fading between colors.
 *
 * This code is public domain. Feel free to use it for any purpose!
 */
package main

import (
	"log"
	"math"
	"runtime"
	"sdl3/pkg/sdl"
)

/* We will use this renderer to draw into this window every frame. */
var window *sdl.Window
var renderer *sdl.Renderer

func main() {
	runtime.LockOSThread()

	if err := sdl.LoadLibrary(); err != nil {
		panic(err)
	}

	sdl.SetAppMetadata("Example Renderer Clear", "1.0", "com.example.renderer-clear")

	if !sdl.Init(sdl.IF_Video) {
		log.Printf("Couldn't initialize SDL: %s\n", sdl.GetError())
		return
	}
	defer sdl.Quit()

	if !sdl.CreateWindowAndRenderer("examples/renderer/clear", 640, 480, 0, &window, &renderer) {
		log.Printf("Couldn't create window/renderer: %s\n", sdl.GetError())
		return
	}

main_loop:
	for {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			switch event.Type {
			case sdl.ET_Quit, sdl.ET_WindowCloseRequested:
				break main_loop /* end the program, reporting success to the OS. */

			case sdl.ET_KeyDown:
				keyEv := event.AsKeyboardEvent()

				switch keyEv.Scancode {
				case sdl.SC_Escape:
					break main_loop
				}
			}
		}

		now := float64(sdl.GetTicks()) / 1000.0 /* convert from milliseconds to seconds. */
		/* choose the color for the frame we will draw. The sine wave trick makes it fade between colors smoothly. */
		red := float32(0.5 + 0.5*math.Sin(now))
		green := float32(0.5 + 0.5*math.Sin(now+math.Pi*2/3))
		blue := float32(0.5 + 0.5*math.Sin(now+math.Pi*4/3))
		sdl.SetRenderDrawColorFloat(renderer, red, green, blue, sdl.ALPHA_OPAQUE_FLOAT) /* new color, full alpha. */

		/* clear the window to the draw color. */
		sdl.RenderClear(renderer)

		/* put the newly-cleared rendering on the screen. */
		sdl.RenderPresent(renderer)
	}
}
