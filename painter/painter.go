package painter

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

//Scene : contains all elements necesssary to paint the scene
type Scene struct {
	time      int32
	renderer  *sdl.Renderer
	bg        *sdl.Texture
	yellowCar *Vehicle
	greyCar   *Vehicle
	redCar    *Vehicle
	motorbike *Vehicle
}

//to be continued
func NewScene(r *sdl.Renderer, errChannel chan error) Scene {
	bg, err := img.LoadTexture(r, "resources/img/road.png")
	errChannel <- err
	yellowCar := newVehicle("yellowCar", "resources/img/yellowcar.png", 0.5, r, errChannel)
	redCar := newVehicle("redCar", "resources/img/redcar.jpg", 1, r, errChannel)
	greyCar := newVehicle("greyCar", "resources/img/greycar.png", 1.5, r, errChannel)
	motorbike := newVehicle("motorbike", "resources/img/motorbike.png", 1, r, errChannel)
	return Scene{renderer: r, bg: bg, yellowCar: &yellowCar, redCar: &redCar, greyCar: &greyCar, motorbike: &motorbike}
}

func (s *Scene) DrawTitle(errChannel chan error) {
	s.renderer.Clear()

	title := createTextTexture("Traffic MAD lad", s.renderer, errChannel)
	copyTextInRendered(title, 50, 50, 500, 200, s.renderer, errChannel)

	subtitle := createTextTexture("Vivant... Je suis vivant!", s.renderer, errChannel)
	copyTextInRendered(subtitle, 100, 300, 400, 75, s.renderer, errChannel)

	subsubtitle := createTextTexture("(Joe Bar Team ref ;-D)", s.renderer, errChannel)
	copyTextInRendered(subsubtitle, 150, 450, 300, 50, s.renderer, errChannel)

	s.renderer.Present()
}

func (s *Scene) Paint(errChannel chan error) {
	s.time++
	s.renderer.Clear()
	errChannel <- s.renderer.Copy(s.bg, nil, nil)
	s.yellowCar.copyInRenderer(s.renderer, errChannel)
	s.redCar.copyInRenderer(s.renderer, errChannel)
	s.greyCar.copyInRenderer(s.renderer, errChannel)
	s.motorbike.copyInRenderer(s.renderer, errChannel)

	s.renderer.Present()
}
