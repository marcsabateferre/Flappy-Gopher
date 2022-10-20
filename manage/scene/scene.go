package scene

import (
	"app/entity"
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type scene struct {
	bg         *sdl.Texture
	birdEntity *entity.Bird
}

func NewScene(r *sdl.Renderer) (*scene, error) {
	bg, error := img.LoadTexture(r, "resources/png/bg.png")
	if error != nil {
		return nil, fmt.Errorf("could not load background image: %v", error)
	}
	var bird *entity.Bird

	bird, error = entity.CreateBird(r)
	if error != nil {
		return nil, error
	}
	return &scene{bg: bg, birdEntity: bird}, nil
}

func (s *scene) handleEvent(event sdl.Event) bool {
	switch event.(type) {
	case *sdl.QuitEvent:
		return true
	case *sdl.MouseButtonEvent:
		s.birdEntity.Jump()
	default:

	}
	return false
}

func (s *scene) Run(events <-chan sdl.Event, r *sdl.Renderer) <-chan error {
	errc := make(chan error)

	go func() {
		defer close(errc)
		tick := time.Tick(10 * time.Millisecond)
		for {
			select {
			case e := <-events:
				if done := s.handleEvent(e); done {
					return
				}
			case <-tick:
				s.updateScene()

				if s.birdEntity.IsDead() {
					s.birdEntity.RestartBird()
					s.birdEntity.UpdateBird()
				}
				if err := s.paint(r); err != nil {
					errc <- err
				}
			}
		}
	}()

	return errc
}

func (s *scene) paint(r *sdl.Renderer) error {
	r.Clear()
	if error := r.Copy(s.bg, nil, nil); error != nil {
		return fmt.Errorf("could not copy background: %v", error)
	}
	if error := s.birdEntity.Paint(r); error != nil {
		return error
	}

	r.Present()
	return nil
}

func (s *scene) Destroy() {
	s.bg.Destroy()
}

func (s *scene) updateScene() {
	s.birdEntity.UpdateBird()
}
