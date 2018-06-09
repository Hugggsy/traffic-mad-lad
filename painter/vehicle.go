package painter

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Vehicle struct {
	texture         *sdl.Texture
	speedMultiplier float32
	X, Y            int32
	W, H            int32
}

func newVehicle(name string, filepath string, speedMultiplier float32, r *sdl.Renderer, errChannel chan error) Vehicle {
	texture, err := img.LoadTexture(r, filepath)
	errChannel <- err
	var W, H int32
	switch name {
	case "yellowCar":
		W = 100
		H = 200
	case "greyCar":
		W = 120
		H = 200
	case "redCar":
		W = 100
		H = 200
	case "motorcycle":
		W = 70
		H = 170
	}
	return Vehicle{texture: texture, W: W, H: H}
}

func (v Vehicle) copyInRenderer(r *sdl.Renderer, errChannel chan error) {
	carRect := &sdl.Rect{X: v.X, Y: v.Y, W: v.W, H: v.H}
	errChannel <- r.Copy(v.texture, nil, carRect)
}
