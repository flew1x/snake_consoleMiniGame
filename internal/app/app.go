package app

import (
	"log"
	"module/internal/logic"
)

var stoppedGame bool

// Launch the game
func StartGame(sizeMap int, maxLength int) {

	//Create a new frame for map
	frameBuilder := logic.NewFrameBuilder()
	frame, err := frameBuilder.SizeMap(sizeMap).BackgroundSymbol(" ").CreateMap()

	CheckErrors(err)

	// Create a new snake
	snakeBuilder := logic.NewSnakeBuilder()
	snake, err := snakeBuilder.StartPos(10, 10).MaxLength(maxLength).SnakeSymbol("O").CreateSnake()

	CheckErrors(err)

	// Create a new food
	food, err := logic.CreateFood("", frame)

	CheckErrors(err)

	game := logic.CreateGame(0)
	game.GameInit(snake, frame, food)
}

func CheckErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
