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

func (pipe *Pipe) paint(r *sdl.Renderer) error {
	rect := &sdl.Rect{X: pipe.x, Y: 600 - pipe.height, W: pipe.width, H: pipe.height}
	flip := sdl.FLIP_NONE
	if pipe.rotated {
		rect.Y = 0
		flip = sdl.FLIP_VERTICAL
	}

	if err := r.CopyEx(pipe.texture, nil, rect, 0, nil, flip); err != nil {
		return fmt.Errorf("background not copy: %s", err)
	}

	return nil
}

func (pipes *Pipes) UpdatePipes() {
	var totalPipes []*Pipe
	for _, p := range pipes.pipes {
		p.x -= 2
		if p.x+p.width > 0 {
			totalPipes = append(totalPipes, p)
		}
	}

	pipes.pipes = totalPipes
}

func (pipe *Pipe) UpdatePipe() {
	pipe.x -= 2
}

func (pipes *Pipes) Destroy() {
	for _, p := range pipes.pipes {
		p.texture.Destroy()
	}
}

func (pipes *Pipes) RestartPipes() {
	pipes.pipes = nil
}

func (pipes *Pipes) CheckCollisions(bird *Bird) {
	for _, p := range pipes.pipes {
		p.checkCollision(bird)
	}
}

func (pipe *Pipe) checkCollision(bird *Bird) {
	if bird.x > pipe.x && bird.x < pipe.width+bird.width {
		//Bird passes the pipe

	}
	if bird.x > pipe.x && pipe.x < bird.width {
		//TODO: fix the vertical checking
		if !pipe.rotated && bird.y < pipe.height {
			//bird.dead = true
		}

		if pipe.rotated && bird.y > pipe.height {
			//bird.dead = true
		}

	}
}
