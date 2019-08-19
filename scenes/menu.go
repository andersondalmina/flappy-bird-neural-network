package scenes

import (
	"image/color"

	"github.com/andersondalmina/flappy-bird-neural-network/components"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var colorMenu = colornames.Gold

// Option is an option on the menu
type Option struct {
	text  string
	hover bool
	goTo  Scene
}

type menu struct {
	options []Option
}

// CreateMenuOption creates an option to the menu
func CreateMenuOption(text string, scene Scene) Option {
	return Option{
		text:  text,
		hover: false,
		goTo:  scene,
	}
}

// CreateMenuScene create the game menu
func CreateMenuScene() Scene {
	sp := CreateSinglePlayerScene()
	sia := CreateIAScene(1)

	var o []Option
	o = append(o, CreateMenuOption("Single Player", sp))
	o = append(o, CreateMenuOption("AI", sia))

	o[0].hover = true

	return &menu{
		options: o,
	}
}

func (s *menu) Run(win *pixelgl.Window) Scene {
	win.Clear(colornames.Skyblue)

	drawBackground(win)
	drawFloor(win)

	var text []components.Text
	var color color.RGBA

	for _, o := range s.options {
		color = colornames.White
		if o.hover == true {
			color = colorMenu
		}
		text = append(text, components.CreateTextLine(o.text, color))
	}

	mat := pixel.IM.Scaled(win.Bounds().Center(), 5).Moved(pixel.V(0, 50))
	components.WriteText(text, colorMenu, win, mat)

	if win.JustPressed(pixelgl.KeyEnter) {
		return s.options[s.getOptionHover()].goTo
	}

	if win.JustPressed(pixelgl.KeyDown) {
		nh := s.getOptionHover()

		if nh >= 0 && nh < len(s.options)-1 {
			s.options[nh].hover = false
			s.options[nh+1].hover = true
		}
	}

	if win.JustPressed(pixelgl.KeyUp) {
		nh := s.getOptionHover()

		if nh > 0 {
			s.options[nh].hover = false
			s.options[nh-1].hover = true
		}
	}

	return s
}

func (s *menu) getOptionHover() int {
	for i, o := range s.options {
		if o.hover == true {
			return i
		}
	}

	return -1
}
