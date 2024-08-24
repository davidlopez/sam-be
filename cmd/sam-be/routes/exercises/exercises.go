package exercises

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func AllExercises() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, exercises!")
	}
}

func ExerciseById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello, exercises: %s!", c.Params("id")))
	}
}

func UpdateExercise() fiber.Handler {
	return ExerciseById()
}

func CreateExercise() fiber.Handler {
	return AllExercises()
}

func DeleteExercise() fiber.Handler {
	return ExerciseById()
}
