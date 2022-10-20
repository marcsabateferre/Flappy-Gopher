package entity

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

type Pipe struct {
	x       int32
	width   int32
	height  int32
	rotated bool
	texture *sdl.Texture
}

func CreatePipe(r *sdl.Renderer) (*Pipe, error) {
	var image = "resources/png/pipe.png"
	texture, error := img.LoadTexture(r, image)
	if error != nil {
		return nil, fmt.Errorf("could not load bird image: %v", error)
	}
	return &Pipe{
		x:       400,
		height:  100 + int32(rand.Intn(300)),
		width:   50,
		texture: texture,
		rotated: rand.Float32() > 0.4}, nil
}

func (p *Pipe) Paint(r *sdl.Renderer) error {

	rect := &sdl.Rect{X: p.x, Y: 600 - p.height, W: p.width, H: p.height}
	flip := sdl.FLIP_NONE
	if p.rotated {
		rect.Y = 0
		flip = sdl.FLIP_VERTICAL
	}

	if err := r.CopyEx(p.texture, nil, rect, 0, nil, flip); err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}

	return nil
}
