package router

import(
	"spewebinar/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App){
	api := app.Group("/api", logger.New())

	u := api.Group("/user")
	u.Get("/", user.UserGet)
	u.Post("/", user.UserPost)
	u.Patch("/:idx", user.UserUpdate)
	u.Delete("/:idx", user.UserDelete)
}
