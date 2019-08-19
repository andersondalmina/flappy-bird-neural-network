package main

import (
	"fmt"
	"math/rand"
	"time"

	_ "image/png"

	"github.com/andersondalmina/flappy-bird-neural-network/components"
	"github.com/andersondalmina/flappy-bird-neural-network/scenes"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	rand.Seed(time.Now().UTC().UnixNano())

	t := "Flappy Bird Neural Network!"
	win, err := components.CreateWindow(t)
	if err != nil {
		panic(err)
	}

	s := scenes.CreateMenuScene()

	var (
		frames = 0
		second = time.Tick(time.Second)
		last   = time.Now()
	)

	for !win.Closed() {
		if win.JustPressed(pixelgl.KeyEscape) {
			return
		}

		components.Delta = time.Since(last).Seconds()
		last = time.Now()

		s = s.Run(win)

		win.Update()

		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", t, frames))
			frames = 0
		default:
		}
	}
}
