package yabai

import (
	"os/exec"
)

type Yabai struct {
	Binary string
}

func (y *Yabai) binary() string {
	if y.Binary != "" {
		return y.Binary
	}
	return "yabai"
}

func (y *Yabai) call(args ...string) (string, error) {
	cmd := exec.Command(y.binary(), args...)

	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(stdout), nil
}
