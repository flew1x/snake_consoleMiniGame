package main

import (
	"module/internal/logic"
	"testing"
)

func TestCreateSnakeBuilderInstance(t *testing.T) {

	t.Run("positive snake's length", func(t *testing.T) {
		snakeBuilder := logic.NewSnakeBuilder()
		_, err := snakeBuilder.SnakeSymbol("0").MaxLength(10).StartPos(10, 10).CreateSnake()

		if err != nil {
			t.Fatal("something went wrong: ", err)
		}
	})

	t.Run("negative snake's length", func(t *testing.T) {
		snakeBuilder := logic.NewSnakeBuilder()
		_, err := snakeBuilder.SnakeSymbol("0").MaxLength(-1).StartPos(10, 10).CreateSnake()

		if err != nil {
			t.Fatal("something went wrong: ", err)
		}
	})

	t.Run("hard snake's length", func(t *testing.T) {
		snakeBuilder := logic.NewSnakeBuilder()
		_, err := snakeBuilder.SnakeSymbol("0").MaxLength(999999).StartPos(10, 10).CreateSnake()

		if err != nil {
			t.Fatal("something went wrong: ", err)
		}
	})

	t.Run("empty snake symbol", func(t *testing.T) {
		snakeBuilder := logic.NewSnakeBuilder()
		snake, err := snakeBuilder.SnakeSymbol("").MaxLength(2).StartPos(10, 10).CreateSnake()

		if err != nil {
			t.Fatal("something went wrong: ", err)
		}

		if snake.GetSymbol() != "O" {
			t.Fatalf("warning! default symbol %q != %v", "O", snake.GetSymbol())
		}
	})

	t.Run("space snake symbol", func(t *testing.T) {
		snakeBuilder := logic.NewSnakeBuilder()
		_, err := snakeBuilder.SnakeSymbol(" ").MaxLength(2).StartPos(10, 10).CreateSnake()

		if err != nil {
			t.Fatal("something went wrong: ", err)
		}
	})
}
