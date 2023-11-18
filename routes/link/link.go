package routes

import (
	"read-it-later/api/utils"

	"github.com/gofiber/fiber/v2"
)

// Models
type Link struct {
	Url string `json:"url"`
}

// Link Router
func AddLinkRouter(app *fiber.App) {
	linkRouter := app.Group(utils.GetVersionedRouteString("/link"))

	linkRouter.Post("/", createLink)
}

// Handlers
func createLink(c *fiber.Ctx) error {
	url := new(Link)

	if err := c.BodyParser(url); err != nil {
		return err
	}

	utils.PrettyPrintJson("URL", url)

	return c.JSON(fiber.Map{
		"added": "sucess",
	})
}
