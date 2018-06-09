package painter

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Vehicle struct {
	texture        *sdl.Texture
	X, Y           int32
	W, H           int32
	XSpeed, YSpeed int32
}

func newVehicle(name string, filepath string, r *sdl.Renderer, errChannel chan error) Vehicle {
	texture, err := img.LoadTexture(r, filepath)
	errChannel <- err
	var W, H int32
	var YSpeed int32
	switch name {
	case "yellowCar":
		W = 100
		H = 200
		YSpeed = 4
	case "greyCar":
		W = 120
		H = 200
		YSpeed = 2
	case "redCar":
		W = 100
		H = 200
		YSpeed = 4
	case "motorbike":
		W = 70
		H = 170
	}
	return Vehicle{texture: texture, W: W, H: H, YSpeed: YSpeed}
}
func (v *Vehicle) setPosition(X, Y int32) {
	v.X = X
	v.Y = Y
}

func (v *Vehicle) modifySpeedValue(pressed bool, key string) {
	switch key {
	case "up":
		if pressed {
			v.YSpeed = -5
		} else {
			v.YSpeed = 0
		}
	case "down":
		if pressed {
			v.YSpeed = +5
		} else {
			v.YSpeed = 0
		}
	case "right":
		if pressed {
			v.XSpeed = +3
		} else {
			v.XSpeed = 0
		}
	case "left":
		if pressed {
			v.XSpeed = -3
		} else {
			v.XSpeed = 0
		}
	}
}

func updateVehiclePositions(vehicles ...*Vehicle) {
	for _, v := range vehicles {
		v.X = v.X + v.XSpeed
		v.Y = v.Y + v.YSpeed
	}
}

func (v *Vehicle) copyInRenderer(r *sdl.Renderer, errChannel chan error) {
	vehicleRect := &sdl.Rect{X: v.X, Y: v.Y, W: v.W, H: v.H}
	errChannel <- r.Copy(v.texture, nil, vehicleRect)
}

func (v *Vehicle) checkOutOfBounds() bool {
	if v.X < 0 || v.X > 600 || v.Y < 0 || v.Y > 1050 {
		return true
	}
	return false
}
