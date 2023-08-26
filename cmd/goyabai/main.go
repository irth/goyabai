package main

import (
	"fmt"

	"github.com/irth/goyabai"
)

func main() {
	y := &yabai.Yabai{}

    d, err := y.Displays()
    if err != nil {
      panic(err)
    }
    fmt.Printf("%+v\n", d)

    s, err := y.Spaces()
    if err != nil {
      panic(err)
    }
    fmt.Printf("%+v\n", s)
}
