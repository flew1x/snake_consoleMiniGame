package main

import (
	"fmt"
	"module/internal/logic"
	"testing"
)

func TestCreateNewFoodInstance(t *testing.T) {
	t.Run("empty food symbol", func(t *testing.T) {
		frameBuilder := logic.NewFrameBuilder()
		frame, _ := frameBuilder.SizeMap(20).BackgroundSymbol(" ").CreateMap()

		food, _ := logic.CreateFood("", frame)

		if food.GetFoodSymbol() != "F" {
			t.Fatalf("warning! default symbol %q != %v", "F", food.GetFoodSymbol())
		}
	})

	t.Run("a lot of food symbols", func(t *testing.T) {
		frameBuilder := logic.NewFrameBuilder()
		frame, _ := frameBuilder.SizeMap(20).BackgroundSymbol(" ").CreateMap()

		food, err := logic.CreateFood("12", frame)

		if err != nil {
			fmt.Println(len(food.GetFoodSymbol()))
			t.Fatalf("warning! default symbol more than 1")
		}
	})
}
