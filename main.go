package main

import (
	"log"
	"runtime"
	"time"

	"github.com/Hugggsy/traffic-mad-lad/painter"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	errChannel := make(chan error)
	events := make(chan sdl.Event)
	go handleErrors(errChannel)
	go becomeMadLad(events, errChannel)
	//event handler has to be locked to main thread by the library's requirement
	runtime.LockOSThread()
	for {
		events <- sdl.WaitEvent()
	}
}

func handleErrors(errChannel chan error) {
	for {
		err := <-errChannel
		if err != nil {
			log.Fatal(err)
		}
	}
}

//Launches actual game
func becomeMadLad(events chan sdl.Event, errChannel chan error) {
	errChannel <- sdl.Init(sdl.INIT_EVERYTHING)

	errChannel <- ttf.Init()

	_, renderer, err := sdl.CreateWindowAndRenderer(600, 1000, sdl.WINDOW_SHOWN)
	errChannel <- err

	scene := painter.NewScene(renderer, errChannel)
	//Title screen
	scene.DrawTitle(errChannel)
	time.Sleep(5 * time.Second)

	//Runs game
	go scene.Run(events, errChannel)
}
