package main

import (
	"log"
	"sync"
	"time"

	"github.com/Hugggsy/traffic-mad-lad/painter"
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
	//defer sdl.Quit()

	errChannel <- ttf.Init()
	//defer ttf.Quit()

	_, renderer, err := sdl.CreateWindowAndRenderer(600, 1000, sdl.WINDOW_SHOWN)
	errChannel <- err
	//defer window.Destroy()

	scene := painter.NewScene(renderer, errChannel)
	scene.DrawTitle(errChannel)
	time.Sleep(2 * time.Second)

	go func() {
		for {
			scene.Paint(errChannel)
			time.Sleep(10 * time.Millisecond)
		}
	}()

}
