package logic

import (
	"errors"
	"math/rand"
	"time"
)

type food struct {
	foodSymbol string
	coordinate *foodCoordinates

	frame *frame
}

// This struct include coordinates of food now
type foodCoordinates struct {
	x int
	y int
}

// Create food instance
func CreateFood(s string, fr *frame) (*food, error) {
	if s == "" {
		s = "F"
	} else if s == " " {
		return &food{}, errors.New("warning! symbol cannot be space")
	}

	if len(s) > 1 {
		return &food{}, errors.New("warning! symbol of food should be more than 1")
	}

	var x, y = getRandomCoordinates(fr)
	return &food{
		foodSymbol: s,
		coordinate: &foodCoordinates{x: x, y: y},
		frame:      fr,
	}, nil
}

// Get x position and y position of food
func (f *food) GetFoodCoordinates() *foodCoordinates {
	return f.coordinate
}

func (f *food) GetFoodSymbol() string {
	return f.foodSymbol
}

// Set random cooedinates of food
func (f *food) setFoodRandomCoordinates() {
	f.coordinate.x, f.coordinate.y = getRandomCoordinates(f.frame)
}

// Create random number for spawn food
func getRandomCoordinates(mp *frame) (int, int) {
	rand.Seed(time.Now().UnixNano())
	numberX := 2 + rand.Intn(mp.size-2+1)
	numberY := 2 + rand.Intn(mp.size-2+1)
	return numberX, numberY
}
