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

type Space struct {
	ID                 int    `json:"id"`
	UUID               string `json:"uuid"`
	Index              int    `json:"index"`
	Label              string `json:"label"`
	Type               string `json:"type"` // TODO: enum
	Display            int    `json:"display"`
	Windows            []int  `json:"windows"`
	FirstWindow        int    `json:"first-window"`
	LastWindow         int    `json:"last-window"`
	HasFocus           bool   `json:"has-focus"`
	IsVisible          bool   `json:"is-visible"`
	IsNativeFullscreen bool   `json:"is-native-fullscreen"`
}

func (y *Yabai) Spaces() ([]Space, error) {
	var spaces []Space
	err := y.query(&spaces, "spaces")
	return spaces, err
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
