package yabai

import "fmt"

type Space struct {
	y                  *Yabai
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
	if err != nil {
		return nil, err
	}

	for _, space := range spaces {
		space.y = y
	}

	return spaces, nil
}

func (y *Yabai) Space(id int) (Space, error) {
	var space Space
	space.y = y
	err := y.query(&space, "spaces", "--space", fmt.Sprint(id))
	return space, err
}

func (s *Space) Destroy() error {
	_, err := s.y.call("-m", "space", "--destroy", fmt.Sprint(s.Index))
	return err
}
