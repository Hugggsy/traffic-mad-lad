package painter

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

//Scene : contains all elements necesssary to paint the scene
type Scene struct {
	time      int32
	font      *ttf.Font
	renderer  *sdl.Renderer
	bg        *sdl.Texture
	yellowCar *Vehicle
	greyCar   *Vehicle
	redCar    *Vehicle
	motorbike *Vehicle
}

type Vehicle struct {
	texture         *sdl.Texture
	speedMultiplier int32
}

func newVehicle(filepath string, speedMultiplier float32, r *sdl.Renderer, errChannel chan error) Vehicle {
	texture, err := img.LoadTexture(r, filepath)
	errChannel <- err
	return Vehicle{texture: texture}
}

//to be continued
func NewScene(r *sdl.Renderer, errChannel chan error) Scene {
	font, err := ttf.OpenFont("./resources/font/valuoldcaps.ttf", 20)
	errChannel <- err
	bg, err := img.LoadTexture(r, "resources/img/road.png")

	yellowCar := newVehicle("resources/img/yellowcar.png", 0.5, r, errChannel)
	redCar := newVehicle("resources/img/redcar.jpg", 1, r, errChannel)
	greyCar := newVehicle("resources/img/greycar.png", 1.5, r, errChannel)
	motorbike := newVehicle("resources/img/motorbike.png", 1, r, errChannel)
	return Scene{renderer: r, font: font, bg: bg, yellowCar: &yellowCar, redCar: &redCar, greyCar: &greyCar, motorbike: &motorbike}
}

func (s *Scene) DrawTitle(errChannel chan error) {
	s.renderer.Clear()

	titleColor := sdl.Color{R: 255, G: 0, B: 100, A: 255}
	title, err := s.font.RenderUTF8Solid("Traffic MAD LAD", titleColor)
	errChannel <- err
	defer title.Free()

	t, err := s.renderer.CreateTextureFromSurface(title)
	errChannel <- err
	defer t.Destroy()

	rect := &sdl.Rect{X: 50, Y: 20, W: 500, H: 200}
	errChannel <- s.renderer.Copy(t, nil, rect)
	s.renderer.Present()
}

func (s *Scene) Paint(errChannel chan error) {
	s.time++
	fmt.Println(s.time)
	s.renderer.Clear()
	errChannel <- s.renderer.Copy(s.bg, nil, nil)
	carRect := &sdl.Rect{X: 300, Y: 200 + s.time, W: 120, H: 200}
	errChannel <- s.renderer.Copy(s.greyCar.texture, nil, carRect)
	carRect2 := &sdl.Rect{X: 450, Y: 0 + 2*s.time, W: 100, H: 200}
	errChannel <- s.renderer.Copy(s.yellowCar.texture, nil, carRect2)
	bikeRect := &sdl.Rect{X: 300, Y: 800 - 2*s.time, W: 70, H: 170}
	errChannel <- s.renderer.Copy(s.motorbike.texture, nil, bikeRect)
	s.renderer.Present()
}
