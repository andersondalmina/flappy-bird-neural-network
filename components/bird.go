package components

import (
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// BirdHeight is the height of the bird sprite
const BirdHeight = 46

// BirdWidth is the width of the bird sprite
const BirdWidth = 54

// BirdX is the inicial position X of the bird
const BirdX = WindowWidth/2 - WindowWidth*0.2
const gravity = 20.0

var maxSpeed = gravity / 4

// Bird component
type Bird struct {
	x           float64
	y           float64
	speed       float64
	sprite      *pixel.Sprite
	points      int64
	death       bool
	enableGhost bool
	ghost       bool
}

// NewBird creates a new bird component
func NewBird(x float64, sprite *pixel.Sprite) *Bird {
	return &Bird{
		x:           x,
		y:           WindowHeight / 2,
		speed:       2,
		sprite:      sprite,
		points:      0,
		death:       false,
		enableGhost: true,
		ghost:       false,
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

func (b *Bird) IsDeath() bool {
	return b.death
}

func (b *Bird) Death() {
	b.death = true
}

func (b *Bird) Update() {
	if b.speed < maxSpeed {
		b.speed += gravity * Delta
	}

	if b.death == true {
		b.x -= XSpeed * Delta

		if b.y > 80 {
			b.y -= b.speed
		}

		return
	}

	b.y -= b.speed
}

func (b *Bird) Draw(win *pixelgl.Window) {
	mat := pixel.IM.Rotated(pixel.V(0, 0), Min(15*math.Pi/180, b.speed*-8*math.Pi/180))
	mat = mat.Moved(pixel.V(b.x, b.y))

	sprite := b.sprite
	if b.ghost {
		sprite = Sprites["bird16"]
	}

	sprite.Draw(win, mat)
}

func (b *Bird) Jump() {
	b.speed = -gravity * Delta * 23
}

func (b *Bird) IncreasePoint() {
	b.points++
}

func (b *Bird) Ghost() bool {
	return b.ghost
}

func (b *Bird) IsEnableGhost() bool {
	return b.enableGhost
}

func (b *Bird) UseGhost() {
	if b.enableGhost {
		b.ghost = true
		b.enableGhost = false

		time.AfterFunc(2*time.Second, func() {
			b.ghost = false
		})

		time.AfterFunc(7*time.Second, func() {
			b.enableGhost = true
		})
	}
}
