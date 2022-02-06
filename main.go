package main

import (
	"github.com/onerciller/portm/command"
	"github.com/onerciller/portm/ui"
)

func main() {
	c := command.New()
	u := ui.New(*c)
	u.Render()
}
