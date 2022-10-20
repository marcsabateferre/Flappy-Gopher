package main

import (
	sceneManage "app/manage/scene"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"runtime"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error {
	error := sdl.Init(sdl.INIT_EVERYTHING)
	if error != nil {
		return fmt.Errorf("Sdl could not init: %v", error)
	}
	defer sdl.Quit()

	w, r, error := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if error != nil {
		return fmt.Errorf("could not create window: %v", error)
	}
	defer w.Destroy()

	time.Sleep(1 * time.Second)

	s, error := sceneManage.CreateScene(r)
	if error != nil {
		return fmt.Errorf("could not create scene: %v", error)
	}
	defer s.Destroy()

	events := make(chan sdl.Event)
	errc := s.Run(events, r)

	runtime.LockOSThread()
	for {
		select {
		case events <- sdl.WaitEvent():
		case err := <-errc:
			return err
		}
	}
}
