package components

import (
	"image"
	"os"

	"github.com/faiface/pixel"
)

// Sprites teste
var Sprites = loadSprites()

// LoadSprites load all imagens and return a sprite to each one
func loadSprites() map[string]*pixel.Sprite {
	result := make(map[string]*pixel.Sprite)
	picBackground, err := loadPicture("./assets/background.png")
	if err != nil {
		panic(err)
	}
	result["background"] = pixel.NewSprite(picBackground, picBackground.Bounds())

	picFloor, err := loadPicture("./assets/floor.png")
	if err != nil {
		panic(err)
	}
	result["floor"] = pixel.NewSprite(picFloor, picFloor.Bounds())

	picBird, err := loadPicture("./assets/bird10.png")
	if err != nil {
		panic(err)
	}
	result["bird10"] = pixel.NewSprite(picBird, picBird.Bounds())

	picBird, err = loadPicture("./assets/bird11.png")
	if err != nil {
		panic(err)
	}
	result["bird11"] = pixel.NewSprite(picBird, picBird.Bounds())

	picBird, err = loadPicture("./assets/bird12.png")
	if err != nil {
		panic(err)
	}
	result["bird12"] = pixel.NewSprite(picBird, picBird.Bounds())

	picBird, err = loadPicture("./assets/bird13.png")
	if err != nil {
		panic(err)
	}
	result["bird13"] = pixel.NewSprite(picBird, picBird.Bounds())

	picBird, err = loadPicture("./assets/bird14.png")
	if err != nil {
		panic(err)
	}
	result["bird14"] = pixel.NewSprite(picBird, picBird.Bounds())

	picBird, err = loadPicture("./assets/bird15.png")
	if err != nil {
		panic(err)
	}
	result["bird15"] = pixel.NewSprite(picBird, picBird.Bounds())

	picBird, err = loadPicture("./assets/bird16.png")
	if err != nil {
		panic(err)
	}
	result["bird16"] = pixel.NewSprite(picBird, picBird.Bounds())

	picUp, err := loadPicture("./assets/pipeup.png")
	if err != nil {
		panic(err)
	}
	result["pipeUp"] = pixel.NewSprite(picUp, picUp.Bounds())

	picDown, err := loadPicture("./assets/pipedown.png")
	if err != nil {
		panic(err)
	}
	result["pipeDown"] = pixel.NewSprite(picDown, picDown.Bounds())

	pic, err := loadPicture("./assets/wall.png")
	if err != nil {
		panic(err)
	}
	result["wall"] = pixel.NewSprite(pic, pic.Bounds())

	return result
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
