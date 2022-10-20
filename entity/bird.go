package entity

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

var gravity = 0.1

type Bird struct {
	x        int32
	y        int32
	time     int
	width    int32
	height   int32
	speed    float64
	dead     bool
	textures []*sdl.Texture
}

func CreateBird(r *sdl.Renderer) (*Bird, error) {
	var textures []*sdl.Texture
	for i := 1; i < 5; i++ {
		var image = "resources/png/frame-" + strconv.Itoa(i) + ".png"
		texture, error := img.LoadTexture(r, image)
		if error != nil {
			return nil, fmt.Errorf("could not load background image: %v", error)
		}
		textures = append(textures, texture)
	}

	return &Bird{
		textures: textures,
		x:        10,
		y:        300,
		width:    50,
		height:   43}, nil
}

func (b *Bird) Paint(r *sdl.Renderer) error {

	rect := &sdl.Rect{X: 10, Y: 600 - b.y - b.height/2, W: b.width, H: b.height}

	i := b.time / 10 % len(b.textures)
	if error := r.Copy(b.textures[i], nil, rect); error != nil {
		return fmt.Errorf("could not copy background: %v", error)
	}
	return nil
}

func (b *Bird) UpdateBird() {
	b.time++
	b.y -= int32(b.speed)
	if b.y < 0 {
		b.dead = true
	}
	b.speed += gravity
}

func (b *Bird) IsDead() bool {
	return b.dead
}

func (b *Bird) RestartBird() {
	b.y = 300
	b.speed = 0
	b.dead = false
}
