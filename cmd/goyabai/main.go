package main

import (
	"fmt"

	yabai "github.com/irth/goyabai"
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

	w, err := y.Windows()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", w)

	display, err := y.Display(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", display)

	space, err := y.Space(4)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", space)
	space.Destroy()

	//	windows, err := y.Window(w[0].ID)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("%+v\n", windows)

}
