package logic

import (
	"fmt"
	"module/pkg"
	"time"
)

type game struct {
	score      int
	isGameOver bool
}

func CreateGame(startScore int) *game {
	return &game{
		score: startScore,
	}
}

func (g *game) GameInit(s *snake, fr *frame, f *food) {
	go s.GoSnake(fr, f)
	for !g.isGameOver {
		t := time.NewTimer(400 * time.Millisecond)
		fr.DrawMap(s, f)
		fmt.Println(f.GetFoodCoordinates())
		fmt.Println("Score: ", s.score)
		<-t.C
		pkg.CallClear()

		if s.isDiy {
			CloseGame(g, s)
		} else if s.isMaxLength {
			CloseGame(g, s)
		}
	}
}

// Exit the game
func CloseGame(g *game, s *snake) {
	g.isGameOver = true
	pkg.CallClear()

	if s.isMaxLength {
		fmt.Println("It is max length of your snake :)")
	}

	fmt.Println("Game closed")
	fmt.Println("Your score: ", s.score)
}
