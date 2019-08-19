package components

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const BirdHeight = 46
const BirdWidth = 54
const gravity = 20.0
const BirdX = WindowWidth/2 - WindowWidth*0.2

var maxSpeed = gravity / 5

// Bird component
type Bird struct {
	x      float64
	y      float64
	speed  float64
	sprite *pixel.Sprite
	points int64
	death  bool
	enable bool
}

// NewBird creates a new bird component
func NewBird(x float64, sprite *pixel.Sprite) *Bird {
	return &Bird{
		x:      x,
		y:      WindowHeight / 2,
		speed:  2,
		sprite: sprite,
		points: 0,
		death:  false,
		enable: true,
	}
}

func (b *Bird) GetX() float64 {
	return b.x
}

func (b *Bird) GetY() float64 {
	return b.y
}

func (b *Bird) GetPoints() int64 {
	return b.points
}

func (b *Bird) Enable() {
	b.enable = true
}

func (b *Bird) IsDeath() bool {
	return b.death
}

func (b *Bird) Death() {
	b.death = true
}

func (b *Bird) Update() {
	if b.death == true || b.enable == true {
		if b.speed < maxSpeed {
			b.speed += gravity * Delta
		}
	}

	if b.death == true {
		b.x -= XSpeed * Delta
		if b.y > 80 {
			b.y -= b.speed
		}
	} else {
		if b.enable && b.y > 80 {
			b.y -= b.speed
		} else {
			if b.y >= WindowHeight/2+20 {
				b.speed = 2
			} else if b.y < WindowHeight/2-20 {
				b.speed = -2
			}

			b.y -= b.speed
		}
	}
}

func (b *Bird) Draw(win *pixelgl.Window) error {
	mat := pixel.IM

	if b.enable && b.speed > 0 {
		if b.speed > -5 && b.speed < 3 {
			mat = mat.Rotated(pixel.V(0, 0), b.speed*-15*math.Pi/180)
		} else {
			mat = mat.Rotated(pixel.V(0, 0), -45*math.Pi/180)
		}
	} else if b.enable == false {
		mat = mat.Rotated(pixel.V(0, 0), -18*math.Pi/180)
	}

	mat = mat.Moved(pixel.V(b.x, b.y))

	b.sprite.Draw(win, mat)

	return nil
}

func (b *Bird) Jump() {
	b.speed = -gravity * Delta * 23
}

func (b *Bird) IncreasePoint() {
	b.points++
}
