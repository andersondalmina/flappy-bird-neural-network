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

const ghostCountdown = 1300

var gravity = 20.0
var maxSpeed = gravity * 30

// Bird component
type Bird struct {
	X              float64
	Y              float64
	speed          float64
	sprite         *pixel.Sprite
	Points         int64
	Dead           bool
	Ghost          bool
	GhostCountdown int
}

// NewBird creates a new bird component
func NewBird(x float64, sprite *pixel.Sprite) *Bird {
	return &Bird{
		X:              x,
		Y:              WindowHeight / 2,
		speed:          0,
		sprite:         sprite,
		Points:         0,
		Dead:           false,
		Ghost:          false,
		GhostCountdown: ghostCountdown,
	}
}

// Kill set bird as dead
func (b *Bird) Kill() {
	b.Dead = true
}

// Update run all operations on bird
func (b *Bird) Update() {
	if b.GhostCountdown > 0 {
		b.GhostCountdown--
	}

	if b.speed < maxSpeed {
		b.speed += gravity
	}

	if b.Dead {
		b.X -= GameXSpeed * Delta
	}

	if b.Y > 80 || b.speed < 0 {
		b.Y -= b.speed * Delta
	}
}

// Draw the bird on window
func (b *Bird) Draw(win *pixelgl.Window) {
	mat := pixel.IM.Rotated(pixel.V(0, 0), Min(15*math.Pi/180, -b.speed/10*math.Pi/180))
	mat = mat.Moved(pixel.V(b.X, b.Y))

	sprite := b.sprite
	if b.Ghost {
		sprite = Sprites["bird16"]
	}

	sprite.Draw(win, mat)
}

// Jump change speed of bird to jump
func (b *Bird) Jump() {
	b.speed = -gravity * 20
}

// IncreasePoint run all inscreases
func (b *Bird) IncreasePoint() {
	b.Points++
}

// UseGhost enable ghost power on bird
func (b *Bird) UseGhost() {
	if b.GhostCountdown == 0 {
		b.GhostCountdown = ghostCountdown
		b.Ghost = true

		time.AfterFunc(2*time.Second, func() {
			b.Ghost = false
		})
	}
}
