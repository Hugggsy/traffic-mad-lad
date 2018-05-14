package main

import (
	"log"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	errChannel := make(chan error)
	go handleErrors(errChannel)
	var wg sync.WaitGroup
	wg.Add(1)
	go becomeMadLad(&wg, errChannel)
	wg.Wait()
}

func handleErrors(errChannel chan error) {
	for {
		err := <-errChannel
		if err != nil {
			log.Fatal(err)
		}
	}
}

func becomeMadLad(wg *sync.WaitGroup, errChannel chan error) {
	errChannel <- sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	errChannel <- ttf.Init()
	defer ttf.Quit()

	window, renderer, err := sdl.CreateWindowAndRenderer(1000, 200, sdl.WINDOW_SHOWN)
	errChannel <- err
	defer window.Destroy()

	drawTitle(renderer, errChannel)

	time.Sleep(2 * time.Second)
	wg.Done()
}

func drawTitle(r *sdl.Renderer, errChannel chan error) {
	r.Clear()

	font, err := ttf.OpenFont("./resources/font/valuoldcaps.ttf", 20)
	errChannel <- err
	defer font.Close()

	titleColor := sdl.Color{R: 200, G: 0, B: 200, A: 255}
	s, err := font.RenderUTF8Solid("Traffic MAD LAD", titleColor)
	errChannel <- err
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	errChannel <- err
	defer t.Destroy()

	errChannel <- r.Copy(t, nil, nil)
	r.Present()
}
