package painter

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func createTextTexture(s string, r *sdl.Renderer, errChannel chan error) *sdl.Texture {
	font, err := ttf.OpenFont("./resources/font/valuoldcaps.ttf", 20)
	errChannel <- err
	textColor := sdl.Color{R: 150, G: 0, B: 100, A: 255}
	text, err := font.RenderUTF8Solid(s, textColor)
	errChannel <- err
	defer text.Free()
	t, err := r.CreateTextureFromSurface(text)
	errChannel <- err
	return t
}

func copyTextInRendered(t *sdl.Texture, X, Y, W, H int32, r *sdl.Renderer, errChannel chan error) {
	rect := &sdl.Rect{X: X, Y: Y, W: W, H: H}
	errChannel <- r.Copy(t, nil, rect)
}
