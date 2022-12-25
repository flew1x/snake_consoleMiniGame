package logic

import (
	"errors"
	"os"
	"os/exec"
	"time"
)

type snake struct {
	coordinates     *snakeCoordinate
	tailCoordinates *tailCoordinates
	snakeSymbol     string
	eDirection      int
	length          int
	isDiy           bool
	isMaxLength     bool
	score           int
	maxLength       int
}

// Snake coordinates now
type snakeCoordinate struct {
	x, y int
}

// Tail coordinates now
type tailCoordinates struct {
	x, y []int
}

// Get snake coordinates
func (s *snake) GetSnakeCoordinates() *snakeCoordinate {
	return s.coordinates
}

func (s *snake) GetStateOfSnake() bool {
	return s.isDiy
}

// Get tails coordinates
func (s *snake) GetTailCoordinates() *tailCoordinates {
	return s.tailCoordinates
}

// Get snake's symbol
func (s *snake) GetSymbol() string {
	return s.snakeSymbol
}

// Interface which describe snake's struct
type SnakeI interface {
	CreateSnake() (*snake, error)

	StartPos(x, y int) SnakeI
	SnakeSymbol(str string) SnakeI
	MaxLength(n int) SnakeI
}

// Create a new builder for snake
func NewSnakeBuilder() SnakeI {
	return snakeBuilder{}
}

// Struct builder for snake
type snakeBuilder struct {
	xStart int
	yStart int

	maxLength int

	snakeSymbol string
}

// Set start position of snake
func (s snakeBuilder) StartPos(x, y int) SnakeI {
	s.xStart = x
	s.yStart = y
	return s
}

// Set snake's symbol
func (s snakeBuilder) SnakeSymbol(str string) SnakeI {
	s.snakeSymbol = str
	return s
}

func (s snakeBuilder) MaxLength(n int) SnakeI {
	s.maxLength = n
	return s
}

// Create a new snake instance
func (s snakeBuilder) CreateSnake() (*snake, error) {

	if s.maxLength > 100000 {
		return &snake{}, errors.New("warning! max length of snake should be less than 100000")
	} else if s.maxLength == 0 {
		s.maxLength = 100
	} else if s.maxLength <= 1 {
		return &snake{}, errors.New("warning! max length of snake should be more than 1")
	}

	var maxTailLengthX = make([]int, s.maxLength)
	var maxTailLengthY = make([]int, s.maxLength)

	if s.snakeSymbol == "" {
		s.snakeSymbol = "O"
	} else if s.snakeSymbol == " " {
		return &snake{}, errors.New("warning! snake's symbol cannot be space")
	}

	return &snake{
		coordinates:     &snakeCoordinate{x: s.xStart, y: s.yStart},
		snakeSymbol:     s.snakeSymbol,
		tailCoordinates: &tailCoordinates{x: maxTailLengthX, y: maxTailLengthY},
		length:          1,
		maxLength:       s.maxLength,
	}, nil
}

// Move the snake
func (s *snake) GoSnake(fr *frame, f *food) {
	go readKey(s)
	go s.checkSnakeState(fr, f)
	for {
		t := time.NewTimer(400 * time.Millisecond)

		switch s.eDirection {
		case 1:
			s.coordinates.y--
		case 2:
			s.coordinates.y++
		case 3:
			s.coordinates.x--
		case 4:
			s.coordinates.x++
		}
		<-t.C
	}
}

// Check a snake tail and move whose position depend where snake
func (s *snake) CheckSnakeTail() {
	var prevX = s.tailCoordinates.x[0]
	var prevY = s.tailCoordinates.y[0]
	var prev2X, prev2Y int
	s.tailCoordinates.x[0] = s.coordinates.x
	s.tailCoordinates.y[0] = s.coordinates.y

	for i := 1; i < s.length; i++ {
		prev2X = s.tailCoordinates.x[i]
		prev2Y = s.tailCoordinates.y[i]
		s.tailCoordinates.x[i] = prevX
		s.tailCoordinates.y[i] = prevY
		prevX = prev2X
		prevY = prev2Y
	}
}

// Check snake state and doing something with game states
func (s *snake) checkSnakeState(fr *frame, f *food) {
	for {
		t := time.NewTimer(200 * time.Millisecond)

		if s.coordinates.x == 0 || s.coordinates.x >= fr.GetSize()-1 || s.coordinates.y == 0 || s.coordinates.y >= fr.GetSize()-1 {
			s.isDiy = true
		}

		if s.coordinates.x == f.GetFoodCoordinates().x {
			if s.coordinates.y == f.GetFoodCoordinates().y {
				s.score += 10
				s.length += 1
				f.setFoodRandomCoordinates()
			}
		}

		if s.length == s.maxLength {
			s.isMaxLength = true
		}

		for i := 1; i < s.length; i++ {
			if s.coordinates.x == s.tailCoordinates.x[i] && s.coordinates.y == s.tailCoordinates.y[i] {
				s.isDiy = true
			}
		}
		<-t.C
	}
}

// Read keys
func readKey(s *snake) {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var key []byte = make([]byte, 1)
	for {
		os.Stdin.Read(key)
		switch string(key) {
		case "w":
			s.eDirection = 1
		case "s":
			s.eDirection = 2
		case "a":
			s.eDirection = 3
		case "d":
			s.eDirection = 4
		}
	}
}
