package access

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const localUserKey = "user"

func AccessWithAuthenticatedUser(c *fiber.Ctx) error {
	c.Locals(localUserKey, nil)

	// authentication

	// c.Locals(localUserKey,)
	fmt.Println("This is called")
	return c.Next()
}
