package scene

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"time"
)

type scene struct {
	bg *sdl.Texture
}

func NewScene(r *sdl.Renderer) (*scene, error) {
	bg, error := img.LoadTexture(r, "resources/png/bg.png")
	if error != nil {
		return nil, fmt.Errorf("could not load background image: %v", error)
	}

	return &scene{bg: bg}, nil
}

func (s *scene) handleEvent(event sdl.Event) bool {
	switch event.(type) {
	case *sdl.QuitEvent:
		return true
	case *sdl.MouseMotionEvent, *sdl.WindowEvent, *sdl.TouchFingerEvent, *sdl.CommonEvent:
	default:
		log.Printf("unknown event %T", event)
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
	r.Present()
	return nil
}

func (s *scene) Destroy() {
	s.bg.Destroy()
}
