package main

import (
	"fmt"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int32 = 800, 600

type gameState int

type color struct {
	r, g, b byte
}

const (
	start gameState = iota
	play
)

func main() {
	window, err := sdl.CreateWindow("Pong", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)

	if err != nil {
		fmt.Println(err)
		return
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		fmt.Println(err)
		return
	}

	pixels := make([]byte, winWidth*winHeight*4)
	pixels[0] = 255

	// tex.Update(nil, pixels, winWidth*4)
	tex.Update(nil, unsafe.Pointer(&pixels[0]), int(winWidth)*4)
	fmt.Println(start)
	fmt.Println(play)
}
