package entity

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"time"
)

type Pipes struct {
	pipes []*Pipe
}

type Pipe struct {
	x       int32
	width   int32
	height  int32
	rotated bool
	texture *sdl.Texture
}

func CreatePipes(r *sdl.Renderer) (*Pipes, error) {
	pipes := &Pipes{}
	go func() {
		for {
			var pipe = createPipe(r)
			if pipe != nil {
				pipes.pipes = append(pipes.pipes, pipe)
				time.Sleep(time.Second)
			}
		}
	}()

	return pipes, nil
}

func createPipe(r *sdl.Renderer) *Pipe {
	var image = "resources/png/pipe.png"
	texture, error := img.LoadTexture(r, image)
	if error != nil {
		return nil
	}
	return &Pipe{
		x:       800,
		height:  100 + int32(rand.Intn(300)),
		width:   50,
		texture: texture,
		rotated: rand.Float32() > 0.4}
}

func (pipes *Pipes) Paint(r *sdl.Renderer) error {
	for _, p := range pipes.pipes {
		if err := p.paint(r); err != nil {
			return err
		}
	}
	return nil
}

func (p *Pipe) paint(r *sdl.Renderer) error {

	rect := &sdl.Rect{X: p.x, Y: 600 - p.height, W: p.width, H: p.height}
	flip := sdl.FLIP_NONE
	if p.rotated {
		rect.Y = 0
		flip = sdl.FLIP_VERTICAL
	}

	if err := r.CopyEx(p.texture, nil, rect, 0, nil, flip); err != nil {
		return fmt.Errorf("background not copy: %v", err)
	}

	return nil
}

func (pipes *Pipes) UpdatePipes() {
	var remaining []*Pipe
	for _, p := range pipes.pipes {
		p.x -= 2
		if p.x+p.width > 0 {
			remaining = append(remaining, p)
		}
	}
	pipes.pipes = remaining
}

func (p *Pipe) UpdatePipe() {
	p.x -= 2
}

func (pipes *Pipes) Destroy() {
	for _, p := range pipes.pipes {
		p.texture.Destroy()
	}
}

func (pipes *Pipes) RestartPipes() {
	pipes.pipes = nil
}
