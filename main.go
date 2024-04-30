package main

import (
	"fmt"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	winWidth, winHeight int     = 800, 600
	paddleOffsetX       float32 = 30
	paddleOffsetY       float32 = 30
)

type gameState int

const (
	start gameState = iota
	play
)

type color struct {
	r, g, b byte
}

type position struct {
	x, y float32
}

type paddle struct {
	position
	width, height, speed int
	color                color
}

type ball struct {
	position
	radius float32
	xVel   float32
	yVel   float32
	color  color
}

func main() {
	window, err := sdl.CreateWindow("Pong", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tex.Destroy()

	pixels := make([]byte, winWidth*winHeight*4)

	paddleOne := paddle{position: position{x: paddleOffsetX, y: paddleOffsetY}, width: 30, height: 200, speed: 3, color: color{r: 255, g: 255, b: 255}}
	paddleTwo := paddle{position: position{x: float32(winWidth) - (paddleOffsetX * 2), y: paddleOffsetY}, width: 30, height: 200, speed: 3, color: color{r: 255, g: 255, b: 255}}

	// ball := ball{position: position{x: 51, y: 50}, radius: 30, xVel: 4, yVel: 4, color: color{r: 255, g: 255, b: 255}}

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		paddleOne.draw(pixels)
		paddleTwo.draw(pixels)
		// ball.draw(pixels)

		tex.Update(nil, unsafe.Pointer(&pixels[0]), int(winWidth)*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()
		sdl.Delay(3000)
	}
}

func (paddle *paddle) draw(pixels []byte) {
	for x := paddle.x; x < paddle.x; x++ {
		for y := paddle.y; y < paddle.y; y++ {
			setPixel(int(x), int(y), paddle.color, pixels)
		}
	}
}

func (ball *ball) draw(pixels []byte) {
	for x := -ball.radius; x < ball.radius; x++ {
		for y := -ball.radius; y < ball.radius; y++ {
			if x*x+y*y < ball.radius*ball.radius {
				setPixel(int(ball.x+x), int(ball.y+y), ball.color, pixels)
			}
		}
	}
}

func setPixel(x, y int, c color, pixels []byte) {
	index := (y*winWidth + x) * 4

	if (index < 0) || (index > len(pixels)-4) {
		return
	}
	pixels[index] = c.r
	pixels[index+1] = c.g
	pixels[index+2] = c.b
}
