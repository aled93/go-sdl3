/*
 * Logic implementation of the Snake game. It is designed to efficiently
 * represent the state of the game in memory.
 *
 * This code is public domain. Feel free to use it for any purpose!
 */
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sdl3/pkg/sdl"
)

const (
	STEP_RATE_IN_MILLISECONDS  = 125
	SNAKE_BLOCK_SIZE_IN_PIXELS = 24
	SDL_WINDOW_WIDTH           = (SNAKE_BLOCK_SIZE_IN_PIXELS * SNAKE_GAME_WIDTH)
	SDL_WINDOW_HEIGHT          = (SNAKE_BLOCK_SIZE_IN_PIXELS * SNAKE_GAME_HEIGHT)

	SNAKE_GAME_WIDTH  = 24
	SNAKE_GAME_HEIGHT = 18
	SNAKE_MATRIX_SIZE = (SNAKE_GAME_WIDTH * SNAKE_GAME_HEIGHT)

	THREE_BITS = 0x7 /* ~CHAR_MAX >> (CHAR_BIT - SNAKE_CELL_MAX_BITS) */
)

func SHIFT(x, y int) int {
	return (x + (y * SNAKE_GAME_WIDTH)) * SNAKE_CELL_MAX_BITS
}

var joystick sdl.Joystick

type SnakeCell int32

const (
	SnakeCell_Nothing SnakeCell = iota
	SnakeCell_SRight
	SnakeCell_SUp
	SnakeCell_SLeft
	SnakeCell_SDown
	SnakeCell_Food
)

const SNAKE_CELL_MAX_BITS = 3 /* floor(log2(SNAKE_CELL_FOOD)) + 1 */

type SnakeDirection int32

const (
	SnakeDir_Right SnakeDirection = iota
	SnakeDir_Up
	SnakeDir_Left
	SnakeDir_Down
)

type SnakeContext struct {
	Cells           [(SNAKE_MATRIX_SIZE * SNAKE_CELL_MAX_BITS) / 8]uint8
	HeadXPos        int8
	HeadYPos        int8
	TailXPos        int8
	TailYPos        int8
	NextDir         SnakeDirection
	InhibitTailStep uint8
	OccupiedCells   uint
}

type AppState struct {
	Window   sdl.Window
	Renderer sdl.Renderer
	SnakeCtx SnakeContext
	LastStep uint64
}

func (ctx *SnakeContext) SnakeCellAt(x, y int8) SnakeCell {
	shift := SHIFT(int(x), int(y))
	rng := uint16(ctx.Cells[shift/8])
	if (shift/8)+1 < len(ctx.Cells) {
		rng |= uint16(ctx.Cells[(shift/8)+1]) << 8
	}
	return SnakeCell((rng >> (shift % 8)) & THREE_BITS)
}

func SetRectXY(r *sdl.FRect, x, y int16) {
	r.X = float32(x * SNAKE_BLOCK_SIZE_IN_PIXELS)
	r.Y = float32(y * SNAKE_BLOCK_SIZE_IN_PIXELS)
}

func (ctx *SnakeContext) PutCellAt(x, y int8, ct SnakeCell) {
	shift := SHIFT(int(x), int(y))
	adjust := shift % 8
	pos := ctx.Cells[shift/8:]
	rng := uint16(pos[0]) | uint16(pos[1])<<8
	rng &= ^(THREE_BITS << adjust) /* clear bits */
	rng |= (uint16(ct) & THREE_BITS) << adjust
	pos[0] = uint8(rng)
	pos[1] = uint8(rng >> 8)
}

func (ctx *SnakeContext) AreCellsFull() bool {
	return ctx.OccupiedCells == SNAKE_GAME_WIDTH*SNAKE_GAME_HEIGHT
}

func (ctx *SnakeContext) NewFoodPos() {
	for {
		x := int8(rand.Intn(SNAKE_GAME_WIDTH))
		y := int8(rand.Intn(SNAKE_GAME_HEIGHT))
		if ctx.SnakeCellAt(x, y) == SnakeCell_Nothing {
			ctx.PutCellAt(x, y, SnakeCell_Food)
			break
		}
	}
}

func (ctx *SnakeContext) SnakeInitialize() {
	println("Starting game")
	clear(ctx.Cells[:])
	ctx.HeadXPos = SNAKE_GAME_WIDTH / 2
	ctx.TailXPos = SNAKE_GAME_WIDTH / 2
	ctx.HeadYPos = SNAKE_GAME_HEIGHT / 2
	ctx.TailYPos = SNAKE_GAME_HEIGHT / 2
	ctx.NextDir = SnakeDir_Right
	ctx.InhibitTailStep = 4
	ctx.OccupiedCells = 4
	ctx.OccupiedCells--
	ctx.PutCellAt(ctx.TailXPos, ctx.TailYPos, SnakeCell_SRight)
	for i := 0; i < 4; i++ {
		ctx.NewFoodPos()
		ctx.OccupiedCells++
	}
}

func (ctx *SnakeContext) SnakeRedir(dir SnakeDirection) {
	ct := ctx.SnakeCellAt(ctx.HeadXPos, ctx.HeadYPos)
	if (dir == SnakeDir_Right && ct != SnakeCell_SLeft) ||
		(dir == SnakeDir_Up && ct != SnakeCell_SDown) ||
		(dir == SnakeDir_Left && ct != SnakeCell_SRight) ||
		(dir == SnakeDir_Down && ct != SnakeCell_SUp) {
		ctx.NextDir = dir
	}
}

func WrapAround(val *int8, max int8) {
	if *val < 0 {
		*val = max - 1
	} else if *val > max-1 {
		*val = 0
	}
}

func (ctx *SnakeContext) SnakeStep() {
	dir_as_cell := SnakeCell(int32(ctx.NextDir) + 1)
	/* Move tail forward */
	ctx.InhibitTailStep--
	if ctx.InhibitTailStep == 0 {
		ctx.InhibitTailStep++
		ct := ctx.SnakeCellAt(ctx.TailXPos, ctx.TailYPos)
		ctx.PutCellAt(ctx.TailXPos, ctx.TailYPos, SnakeCell_Nothing)
		switch ct {
		case SnakeCell_SRight:
			ctx.TailXPos++
			break
		case SnakeCell_SUp:
			ctx.TailYPos--
			break
		case SnakeCell_SLeft:
			ctx.TailXPos--
			break
		case SnakeCell_SDown:
			ctx.TailYPos++
			break
		default:
			break
		}
		WrapAround(&ctx.TailXPos, SNAKE_GAME_WIDTH)
		WrapAround(&ctx.TailYPos, SNAKE_GAME_HEIGHT)
	}
	/* Move head forward */
	prev_xpos := ctx.HeadXPos
	prev_ypos := ctx.HeadYPos
	switch ctx.NextDir {
	case SnakeDir_Right:
		ctx.HeadXPos++
		break
	case SnakeDir_Up:
		ctx.HeadYPos--
		break
	case SnakeDir_Left:
		ctx.HeadXPos--
		break
	case SnakeDir_Down:
		ctx.HeadYPos++
		break
	default:
		break
	}
	WrapAround(&ctx.HeadXPos, SNAKE_GAME_WIDTH)
	WrapAround(&ctx.HeadYPos, SNAKE_GAME_HEIGHT)
	/* Collisions */
	ct := ctx.SnakeCellAt(ctx.HeadXPos, ctx.HeadYPos)
	if ct != SnakeCell_Nothing && ct != SnakeCell_Food {
		ctx.SnakeInitialize()
		return
	}
	ctx.PutCellAt(prev_xpos, prev_ypos, dir_as_cell)
	ctx.PutCellAt(ctx.HeadXPos, ctx.HeadYPos, dir_as_cell)
	if ct == SnakeCell_Food {
		if ctx.AreCellsFull() {
			ctx.SnakeInitialize()
			return
		}
		ctx.NewFoodPos()
		ctx.InhibitTailStep++
		ctx.OccupiedCells++
	}
}

func (ctx *SnakeContext) HandleKeyEvent(key_code sdl.Scancode) bool {
	switch key_code {
	/* Quit. */
	case sdl.ScancodeEscape:
	case sdl.ScancodeQ:
		return false
	/* Restart the game as if the program was launched. */
	case sdl.ScancodeR:
		ctx.SnakeInitialize()
		break
	/* Decide new direction of the snake. */
	case sdl.ScancodeRight:
		ctx.SnakeRedir(SnakeDir_Right)
		break
	case sdl.ScancodeUp:
		ctx.SnakeRedir(SnakeDir_Up)
		break
	case sdl.ScancodeLeft:
		ctx.SnakeRedir(SnakeDir_Left)
		break
	case sdl.ScancodeDown:
		ctx.SnakeRedir(SnakeDir_Down)
		break
	default:
		break
	}
	return true
}

func (ctx *SnakeContext) HandleHatEvent(hat uint8) {
	switch sdl.Hat(hat) {
	case sdl.HAT_Right:
		ctx.SnakeRedir(SnakeDir_Right)
		break
	case sdl.HAT_Up:
		ctx.SnakeRedir(SnakeDir_Up)
		break
	case sdl.HAT_Left:
		ctx.SnakeRedir(SnakeDir_Left)
		break
	case sdl.HAT_Down:
		ctx.SnakeRedir(SnakeDir_Down)
		break
	default:
		break
	}
}

func (app *AppState) Update() {
	ctx := &app.SnakeCtx
	now := sdl.GetTicks()

	// run game logic if we're at or past the time to run it.
	// if we're _really_ behind the time to run it, run it
	// several times.
	for (now - app.LastStep) >= STEP_RATE_IN_MILLISECONDS {
		ctx.SnakeStep()
		app.LastStep += STEP_RATE_IN_MILLISECONDS
	}

	r := sdl.FRect{
		W: SNAKE_BLOCK_SIZE_IN_PIXELS,
		H: SNAKE_BLOCK_SIZE_IN_PIXELS,
	}
	sdl.SetRenderDrawColor(app.Renderer, 0, 0, 0, sdl.ALPHA_OPAQUE)
	sdl.RenderClear(app.Renderer)
	for i := 0; i < SNAKE_GAME_WIDTH; i++ {
		for j := 0; j < SNAKE_GAME_HEIGHT; j++ {
			ct := ctx.SnakeCellAt(int8(i), int8(j))
			if ct == SnakeCell_Nothing {
				continue
			}
			SetRectXY(&r, int16(i), int16(j))
			if ct == SnakeCell_Food {
				sdl.SetRenderDrawColor(app.Renderer, 80, 80, 255, sdl.ALPHA_OPAQUE)
			} else /* body */ {
				sdl.SetRenderDrawColor(app.Renderer, 0, 128, 0, sdl.ALPHA_OPAQUE)
			}
			sdl.RenderFillRect(app.Renderer, &r)
		}
	}
	sdl.SetRenderDrawColor(app.Renderer, 255, 255, 0, sdl.ALPHA_OPAQUE) /*head*/
	SetRectXY(&r, int16(ctx.HeadXPos), int16(ctx.HeadYPos))
	sdl.RenderFillRect(app.Renderer, &r)
	sdl.RenderPresent(app.Renderer)
}

var extended_metadata = map[sdl.PropertyName]string{
	sdl.PROP_APP_METADATA_URL_STRING:       "https://examples.libsdl.org/SDL3/demo/01-snake/",
	sdl.PROP_APP_METADATA_CREATOR_STRING:   "SDL team",
	sdl.PROP_APP_METADATA_COPYRIGHT_STRING: "Placed in the public domain",
	sdl.PROP_APP_METADATA_TYPE_STRING:      "game",
}

func main() {
	runtime.LockOSThread()

	if err := sdl.LoadLibrary(); err != nil {
		panic(err)
	}

	if !sdl.SetAppMetadata("Example Snake game", "1.0", "com.example.Snake") {
		panic("failed set app metadata: " + sdl.GetError())
	}

	for propName, propValue := range extended_metadata {
		if !sdl.SetAppMetadataProperty(propName, propValue) {
			panic("failed set app metadata " + string(propName) + ": " + sdl.GetError())
		}
	}

	if !sdl.Init(sdl.InitVideo | sdl.InitJoystick) {
		panic("Couldn't initialize SDL: " + sdl.GetError())
	}

	app := AppState{}

	if !sdl.CreateWindowAndRenderer("examples/demo/snake", SDL_WINDOW_WIDTH, SDL_WINDOW_HEIGHT, 0, &app.Window, &app.Renderer) {
		panic("failed create window and renderer: " + sdl.GetError())
	}

	app.SnakeCtx.SnakeInitialize()

	app.LastStep = sdl.GetTicks()

main_loop:
	for {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			switch event.Type {
			case sdl.EventQuit, sdl.EventWindowCloseRequested:
				break main_loop /* end the program, reporting success to the OS. */

			case sdl.EventJoystickAdded:
				if joystick == 0 {
					joystick = sdl.OpenJoystick(event.AsJoyDeviceEvent().Which)
					if joystick == 0 {
						fmt.Printf("Failed to open joystick ID %d: %s\n", event.AsJoyDeviceEvent().Which, sdl.GetError())
					}
				}
				break
			case sdl.EventJoystickRemoved:
				if joystick != 0 && (sdl.GetJoystickID(joystick) == event.AsJoyDeviceEvent().Which) {
					sdl.CloseJoystick(joystick)
					joystick = 0
				}
				break
			case sdl.EventJoystickHatMotion:
				app.SnakeCtx.HandleHatEvent(event.AsJoyHatEvent().Value)
			case sdl.EventKeyDown:
				if !app.SnakeCtx.HandleKeyEvent(event.AsKeyboardEvent().Scancode) {
					break main_loop
				}
			}
		}

		app.Update()
	}

	if joystick != 0 {
		sdl.CloseJoystick(joystick)
	}

	sdl.DestroyRenderer(app.Renderer)
	sdl.DestroyWindow(app.Window)
}
