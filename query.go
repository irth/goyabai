package yabai

import (
	"encoding/json"
	"fmt"
)

func (y *Yabai) query(out any, module string, args ...string) error {
	queryArgs := []string{"-m", "query", fmt.Sprintf("--%s", module)}
	queryArgs = append(queryArgs, args...)
	stdout, err := y.call(queryArgs...)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(stdout), out)
}

type Frame struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W float64 `json:"w"`
	H float64 `json:"h"`
}

type Display struct {
	ID     int    `json:"id"`
	UUID   string `json:"uuid"`
	Index  int    `json:"index"`
	Frame  Frame  `json:"frame"`
	Spaces []int  `json:"spaces"`
}

func (y *Yabai) Displays() ([]Display, error) {
	var displays []Display
	err := y.query(&displays, "displays")
	return displays, err
}

func (y *Yabai) Display(id int) (Display, error) {
	var display Display
	err := y.query(&display, "displays", "--display", fmt.Sprint(id))
	return display, err
}

type Window struct {
	ID                 int     `json:"id"`
	PID                int     `json:"pid"`
	App                string  `json:"app"`
	Title              string  `json:"title"`
	Frame              Frame   `json:"frame"`
	Role               string  `json:"role"`
	Subrole            string  `json:"subrole"`
	Display            int     `json:"display"`
	Space              int     `json:"space"`
	Level              int     `json:"level"`
	Opacity            float64 `json:"opacity"`
	SplitType          string  `json:"split-type"`
	SplitChild         string  `json:"split-child"`
	StackIndex         int     `json:"stack-index"`
	CanMove            bool    `json:"can-move"`
	CanResize          bool    `json:"can-resize"`
	HasFocus           bool    `json:"has-focus"`
	HasShadow          bool    `json:"has-shadow"`
	HasBorder          bool    `json:"has-border"`
	HasParentZoom      bool    `json:"has-parent-zoom"`
	HasFullscreenZoom  bool    `json:"has-fullscreen-zoom"`
	IsNativeFullscreen bool    `json:"is-native-fullscreen"`
	IsVisible          bool    `json:"is-visible"`
	IsMinimized        bool    `json:"is-minimized"`
	IsHidden           bool    `json:"is-hidden"`
	IsFloating         bool    `json:"is-floating"`
	IsSticky           bool    `json:"is-sticky"`
	IsTopmost          bool    `json:"is-topmost"`
	IsGrabbed          bool    `json:"is-grabbed"`
}

func (y *Yabai) Windows() ([]Window, error) {
	var windows []Window
	err := y.query(&windows, "windows")
	return windows, err
}

func (y *Yabai) Window(id int) (Window, error) {
	var window Window
	err := y.query(&window, "windows", "--window", fmt.Sprint(id))
	return window, err
}
