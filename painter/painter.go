package painter

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

//Scene : contains all elements necesssary to paint the scene
type Scene struct {
	font      *ttf.Font
	bg        *sdl.Texture
	yellowCar *sdl.Texture
	redCar    *sdl.Texture
}

//to be continued
func NewScene() Scene {

	return
}

func drawTitle(r *sdl.Renderer, errChannel chan error) {
	r.Clear()

	font, err := ttf.OpenFont("./resources/font/valuoldcaps.ttf", 20)
	errChannel <- err
	defer font.Close()

	titleColor := sdl.Color{R: 0, G: 200, B: 250, A: 255}
	s, err := font.RenderUTF8Solid("Traffic MAD LAD", titleColor)
	errChannel <- err
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	errChannel <- err
	defer t.Destroy()

	errChannel <- r.Copy(t, nil, nil)
	r.Present()
}

func drawRoad(r *sdl.Renderer, errChannel chan error) {
	r.Clear()
	texture, err := img.LoadTexture(r, "resources/img/road.png")
	errChannel <- err
	defer texture.Destroy()
	errChannel <- r.Copy(texture, nil, nil)

	r.Present()
}
