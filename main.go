package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	if err := becomeMadLad(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(2)
	}

}
func becomeMadLad() error {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("Failed to init sdl: %v", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("Failed to initialise ttf: %v", err)
	}
	defer ttf.Quit()

	window, renderer, err := sdl.CreateWindowAndRenderer(1000, 200, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("Failed to create window: %v", err)
	}
	defer window.Destroy()

	if err := drawTitle(renderer); err != nil {
		return fmt.Errorf("Failed not draw title %v", err)
	}

	time.Sleep(5 * time.Second)

	return nil
}

func drawTitle(r *sdl.Renderer) error {
	r.Clear()

	font, err := ttf.OpenFont("valuoldcaps.ttf", 20)
	if err != nil {
		return fmt.Errorf("Failed to load font: %v", err)
	}
	defer font.Close()

	titleColor := sdl.Color{R: 255, G: 0, B: 0, A: 255}
	s, err := font.RenderUTF8Solid("Traffic MAD LAD", titleColor)
	if err != nil {
		return fmt.Errorf("Failed to render font: %v", err)
	}
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("Failed to create texture from surface: %v", err)
	}
	defer t.Destroy()

	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("Failed to copy texture")
	}

	r.Present()
	return nil
}
