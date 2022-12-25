package logic

import (
	"errors"
	"fmt"
	"strings"
)

type frame struct {
	size             int
	backgroundSymbol string
}

// Frame builder
type frameBuilder struct {
	size             int
	backgroundSymbol string
}

// Frame's interface
type FrameI interface {
	SizeMap(size int) FrameI
	BackgroundSymbol(s string) FrameI

	CreateMap() (*frame, error)
}

// Create new builder of map
func NewFrameBuilder() FrameI {
	return frameBuilder{}
}

// Get size of map
func (m *frame) GetSize() int {
	return m.size
}

// Get backgroud symbol of map
func (m *frame) GetBackSymbol() string {
	return m.backgroundSymbol
}

// Draw map in the console
func (fr *frame) DrawMap(s *snake, f *food) {

	s.CheckSnakeTail()

	for rows := 0; rows < fr.GetSize(); rows++ {
		for cols := 0; cols < fr.GetSize(); cols++ {
			if cols == 0 || rows == 0 || rows == fr.GetSize()-1 || cols == fr.GetSize()-1 {
				fmt.Print("#")
			} else if rows == s.coordinates.y && cols == s.coordinates.x {
				fmt.Print(s.snakeSymbol)
			} else if rows == f.coordinate.y && cols == f.coordinate.x {
				fmt.Print(f.foodSymbol)
			} else {
				var print bool = false

				for k := 0; k < s.length; k++ {
					if s.tailCoordinates.x[k] == cols && s.tailCoordinates.y[k] == rows {
						print = true
						fmt.Print(strings.ToLower(s.snakeSymbol))
					}
				}
				if !print {
					fmt.Print(fr.backgroundSymbol)
				}
			}
		}
		fmt.Println()
	}
}

// Create the size of map
func (m frameBuilder) SizeMap(size int) FrameI {
	m.size = size
	return m
}

// Set the symbol of background map
func (m frameBuilder) BackgroundSymbol(s string) FrameI {
	m.backgroundSymbol = s
	return m
}

// Create map
func (f frameBuilder) CreateMap() (*frame, error) {

	if f.backgroundSymbol == "" {
		f.backgroundSymbol = " "
	}

	if f.size < 20 {
		return &frame{}, errors.New("size of frame should be more than 20")
	} else if f.size == 0 {
		f.size = 20
	}

	return &frame{
		size:             f.size,
		backgroundSymbol: f.backgroundSymbol,
	}, nil
}
