/*
 * This example creates an SDL window and renderer, and then draws some
 * textures to it every frame.
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

	sdl.SetAppMetadata("Example Renderer Textures", "1.0", "com.example.renderer-textures")

	if !sdl.Init(sdl.InitVideo) {
		panic("Couldn't initialize SDL: " + sdl.GetError())
	}

	if !sdl.CreateWindowAndRenderer("examples/renderer/textures", WINDOW_WIDTH, WINDOW_HEIGHT, 0, &window, &renderer) {
		panic("Couldn't create window/renderer: " + sdl.GetError())
	}

	/* Textures are pixel data that we upload to the video hardware for fast drawing. Lots of 2D
	   engines refer to these as "sprites." We'll do a static texture (upload once, draw many
	   times) with data from a bitmap file. */

	/* sdl.Surface is pixel data the CPU can access. sdl.Texture is pixel data the GPU can access.
	   Load a .bmp into a surface, move it to a texture from there. */
	bmp_path := sdl.GetBasePath() + "sample.bmp"
	surface := sdl.LoadBMP(bmp_path)
	if surface == nil {
		panic("Couldn't load bitmap: " + sdl.GetError())
	}

	texture_width := surface.W
	texture_height := surface.H

	texture := sdl.CreateTextureFromSurface(renderer, surface)
	if texture == nil {
		panic("Couldn't create static texture: " + sdl.GetError())
	}

	sdl.DestroySurface(surface) /* done with this, the texture has a copy of the pixels now. */

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

		var dst_rect sdl.FRect
		now := sdl.GetTicks()

		/* we'll have some textures move around over a few seconds. */
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

		/* Just draw the static texture a few times. You can think of it like a
		stamp, there isn't a limit to the number of times you can draw with it. */

		/* top left */
		dst_rect.X = (100.0 * scale)
		dst_rect.Y = 0.0
		dst_rect.W = float32(texture_width)
		dst_rect.H = float32(texture_height)
		sdl.RenderTexture(renderer, texture, nil, &dst_rect)

		/* center this one. */
		dst_rect.X = (float32(WINDOW_WIDTH - texture_width)) / 2.0
		dst_rect.Y = (float32(WINDOW_HEIGHT - texture_height)) / 2.0
		dst_rect.W = float32(texture_width)
		dst_rect.H = float32(texture_height)
		sdl.RenderTexture(renderer, texture, nil, &dst_rect)

		/* bottom right. */
		dst_rect.X = (float32(WINDOW_WIDTH - texture_width)) - (100.0 * scale)
		dst_rect.Y = float32(WINDOW_HEIGHT - texture_height)
		dst_rect.W = float32(texture_width)
		dst_rect.H = float32(texture_height)
		sdl.RenderTexture(renderer, texture, nil, &dst_rect)

		sdl.RenderPresent(renderer) /* put it all on the screen! */
	}
}
