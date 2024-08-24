package routes

import (
	"github.com/davidlopez/sam-be/cmd/sam-be/routes/exercises"
	"github.com/davidlopez/sam-be/cmd/sam-be/routes/root"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	app.Get("/", root.Root())

	api := app.Group("/api")

	v1 := api.Group("/v1")

	exercisesRoutes := v1.Group("/exercises")
	exercisesRoutes.Get("/", exercises.AllExercises())
	exercisesRoutes.Post("/", exercises.CreateExercise())
	exercisesRoutes.Get("/:id", exercises.ExerciseById())
	exercisesRoutes.Put("/:id", exercises.UpdateExercise())
	exercisesRoutes.Delete("/:id", exercises.DeleteExercise())
}
