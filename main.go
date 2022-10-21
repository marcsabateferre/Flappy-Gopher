package main

import (
	sceneManage "app/manage"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
	"runtime"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(2)
	}
}

func run() error {
	error := sdl.Init(sdl.INIT_EVERYTHING)
	if error != nil {
		return fmt.Errorf("Sdl could not init: %s", error)
	}
	defer sdl.Quit()
	if error := ttf.Init(); error != nil {
		return fmt.Errorf("could not initialize TTF: %s", error)
	}
	defer ttf.Quit()

	w, r, error := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if error != nil {
		return fmt.Errorf("could not create window: %s", error)
	}
	defer w.Destroy()

	time.Sleep(1 * time.Second)

	s, error := sceneManage.CreateScene(r)
	if error != nil {
		return fmt.Errorf("could not create scene: %s", error)
	}
	defer s.Destroy()

	events := make(chan sdl.Event)
	errorchan := s.Run(events, r)

	runtime.LockOSThread()
	for {
		select {
		case events <- sdl.WaitEvent():
		case err := <-errorchan:
			return err
		}
	}
}
