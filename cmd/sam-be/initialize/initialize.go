package initialize

import (
	"github.com/davidlopez/sam-be/cmd/sam-be/routes"
	"github.com/davidlopez/sam-be/cmd/sam-be/utils"
	"github.com/gofiber/fiber/v2"
)

func Initialize() {
	app := fiber.New()

	routes.InitializeRoutes(app)

	_ = app.Listen(utils.GetPort())
}
