package painter

import (
	"math"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Vehicle struct {
	name           string
	texture        *sdl.Texture
	X, Y           int32
	W, H           int32
	XSpeed, YSpeed int32
}

func newVehicle(name string, filepath string, r *sdl.Renderer, errChannel chan error) Vehicle {
	texture, err := img.LoadTexture(r, filepath)
	errChannel <- err
	var X, Y, W, H int32
	var YSpeed int32
	switch name {
	case "yellowCar":
		X = 450
		Y = -300
		W = 100
		H = 200
		YSpeed = 6
	case "greyCar":
		X = 50
		Y = -300
		W = 120
		H = 200
		YSpeed = 3
	case "redCar":
		X = 250
		Y = -300
		W = 100
		H = 200
		YSpeed = 4
	case "motorbike":
		W = 70
		H = 170
	}
	return Vehicle{texture: texture, X: X, Y: Y, W: W, H: H, YSpeed: YSpeed, name: name}
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
		if v.checkOutOfBounds() && v.Y > 0 && v.name != "motorbike" {
			v.Y = -300
		}
	}
}

func copyInRenderer(r *sdl.Renderer, errChannel chan error, vehicles ...*Vehicle) {
	for _, v := range vehicles {
		vehicleRect := &sdl.Rect{X: v.X, Y: v.Y, W: v.W, H: v.H}
		errChannel <- r.Copy(v.texture, nil, vehicleRect)
	}
}

func (v *Vehicle) checkOutOfBounds() bool {
	if v.X < -50 || v.X > 550 || v.Y < -50 || v.Y > 1050 {
		return true
	}
	return false
}

func (v *Vehicle) checkIntersect(vehicles ...*Vehicle) bool {
	for _, ve := range vehicles {
		horizontalDistance := math.Abs(float64(v.X-ve.X)) * 1.2
		verticalDistance := math.Abs(float64(v.Y-ve.Y)) * 1.2
		combinedWidth := float64(v.W/2 + ve.W/2)
		combinedHeight := float64(v.H/2 + ve.H/2)
		if horizontalDistance < combinedWidth && verticalDistance < combinedHeight {
			return true
		}
	}
	return false
}
