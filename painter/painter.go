package painter

import (
	"fmt"
	"time"

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

//NewScene : Loads all textures and fonts once
func NewScene(r *sdl.Renderer, errChannel chan error) Scene {
	bg, err := img.LoadTexture(r, "resources/img/road.png")
	errChannel <- err
	yellowCar := newVehicle("yellowCar", "resources/img/yellowcar.png", r, errChannel)
	redCar := newVehicle("redCar", "resources/img/redcar.jpg", r, errChannel)
	greyCar := newVehicle("greyCar", "resources/img/greycar.png", r, errChannel)
	motorbike := newVehicle("motorbike", "resources/img/motorbike.png", r, errChannel)
	motorbike.setPosition(270, 800)
	return Scene{renderer: r, bg: bg, yellowCar: &yellowCar, redCar: &redCar, greyCar: &greyCar, motorbike: &motorbike}
}

//DrawTitle : prints ele;ents of texts on title screen
func (s *Scene) DrawTitle(errChannel chan error) {
	s.renderer.Clear()

	title := createTextTexture("Traffic MAD lad", s.renderer, errChannel)
	copyTextInRendered(title, 50, 100, 500, 150, s.renderer, errChannel)

	tips := createTextTexture("Use the arrow keys to dodge incoming cars", s.renderer, errChannel)
	copyTextInRendered(tips, 50, 300, 500, 60, s.renderer, errChannel)

	subtitle := createTextTexture("Vivant... Je suis vivant!", s.renderer, errChannel)
	copyTextInRendered(subtitle, 100, 600, 400, 60, s.renderer, errChannel)

	subsubtitle := createTextTexture("(Joe Bar Team ref ;-D)", s.renderer, errChannel)
	copyTextInRendered(subsubtitle, 150, 700, 300, 30, s.renderer, errChannel)

	s.renderer.Present()
}

//drawGameOver : Draw game over screen when called
func (s *Scene) drawGameOver(errChannel chan error) {
	s.renderer.Clear()

	title := createTextTexture("GAME OVER", s.renderer, errChannel)
	copyTextInRendered(title, 50, 300, 500, 150, s.renderer, errChannel)

	subtitle := createTextTexture(fmt.Sprintf("YOUR GAME LASTED %ds", s.time/100), s.renderer, errChannel)
	copyTextInRendered(subtitle, 50, 500, 500, 60, s.renderer, errChannel)

	s.renderer.Present()
}

func (s *Scene) paint(errChannel chan error) {
	s.time++
	s.renderer.Clear()
	errChannel <- s.renderer.Copy(s.bg, nil, nil)
	updateVehiclePositions(s.yellowCar, s.redCar, s.greyCar, s.motorbike)
	copyInRenderer(s.renderer, errChannel, s.yellowCar, s.redCar, s.greyCar, s.motorbike)

	s.renderer.Present()
}

func (s *Scene) reset(errChannel chan error) {
	s.time = 0
	s.motorbike.setPosition(270, 800)
	s.yellowCar.setPosition(450, -300)
	s.redCar.setPosition(250, -300)
	s.greyCar.setPosition(50, -300)
}

func (s *Scene) handleEvent(event sdl.Event, errChannel chan error) {
	switch e := event.(type) {
	case *sdl.QuitEvent:
		errChannel <- fmt.Errorf("User closed window")
	case *sdl.KeyboardEvent:
		s.handleKeyPress(e)
	}
}

func (s *Scene) handleKeyPress(k *sdl.KeyboardEvent) {
	pressed := (k.Type == 768)
	switch k.Keysym.Scancode {
	case 79:
		s.motorbike.modifySpeedValue(pressed, "right")
	case 80:
		s.motorbike.modifySpeedValue(pressed, "left")
	case 81:
		s.motorbike.modifySpeedValue(pressed, "down")
	case 82:
		s.motorbike.modifySpeedValue(pressed, "up")
	default:
		fmt.Println(k.Keysym.Scancode)
	}
}

//Run : launches game, dispatches events and actions to other actrors
func (s *Scene) Run(events <-chan sdl.Event, errChannel chan error) {
	ticker := time.Tick(10 * time.Millisecond)
	gameover := false
	go func() {
		for {
			select {
			case e := <-events:
				s.handleEvent(e, errChannel)
			case <-ticker:
				if !gameover {
					s.paint(errChannel)
					gameover = s.motorbike.checkOutOfBounds() || s.motorbike.checkIntersect(s.yellowCar, s.redCar, s.greyCar)
				} else {
					s.drawGameOver(errChannel)
					time.Sleep(2 * time.Second)
					s.reset(errChannel)
					gameover = false
				}
			}
		}
	}()
}
