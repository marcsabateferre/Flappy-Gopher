package entity

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"strconv"
)

var gravity = 0.1

type Bird struct {
	x        int32
	y        int32
	frame    int
	width    int32
	height   int32
	speed    float64
	dead     bool
	textures []*sdl.Texture
	points   int
}

func CreateBird(r *sdl.Renderer) (*Bird, error) {
	var textures []*sdl.Texture
	for i := 1; i < 5; i++ {
		var image = "resources/png/frame-" + strconv.Itoa(i) + ".png"
		texture, error := img.LoadTexture(r, image)
		if error != nil {
			return nil, fmt.Errorf("could not load bird image: %s", error)
		}
		textures = append(textures, texture)
	}

	return &Bird{
		textures: textures,
		x:        10,
		y:        300,
		width:    50,
		points:   0,
		height:   43}, nil
}

func (b *Bird) Paint(r *sdl.Renderer) error {

	rect := &sdl.Rect{X: 10, Y: 600 - b.y - b.height/2, W: b.width, H: b.height}

	i := b.frame / 10 % len(b.textures)
	if error := r.Copy(b.textures[i], nil, rect); error != nil {
		return fmt.Errorf("could not copy background: %s", error)
	}
	return nil
}

func (b *Bird) UpdateBird() {
	b.frame++
	b.y -= int32(b.speed)
	if b.y < 0 {
		b.dead = true
	}
	if b.y > 600 {
		b.dead = true
	}
	b.speed += gravity
}

func (b *Bird) IsDead() bool {
	return b.dead
}

func (b *Bird) RestartBird() bool {
	if b.dead {
		b.y = 300
		b.speed = 0
		b.dead = false
		b.points = 0
		b.UpdateBird()
		return true
	}

	return false
}

func (b *Bird) Jump() {
	b.speed = -5
}
func (bird *Bird) PaintPoints(r *sdl.Renderer) error {
	f, err := ttf.OpenFont("resources/fonts/OpenSans-Bold.ttf", 10)
	if err != nil {
		return fmt.Errorf("could not load font: %s", err)
	}
	defer f.Close()

	c := sdl.Color{R: 0, G: 0, B: 0, A: 255}
	var pointsTotal = bird.points
	s, err := f.RenderUTF8Solid("Points: "+strconv.Itoa(pointsTotal), c)
	if err != nil {
		return fmt.Errorf("could not render points text: %s", err)
	}
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("could not render texture: %s", err)
	}
	defer t.Destroy()
	rect := &sdl.Rect{X: 280, Y: 500, W: 200, H: 100}
	if err := r.Copy(t, nil, rect); err != nil {
		return fmt.Errorf("could not copy texture: %s", err)
	}

	return nil
}

func (bird *Bird) updatePoints() {
	if bird.IsDead() {
		bird.points = 0
	}

	bird.points += 1
}
