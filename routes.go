package main

import "github.com/gofiber/fiber/v2"

func setupApp(app *fiber.App) {
	app.Post("/api/register", register)
	app.Post("/api/login", login)
	app.Get("/api/user", getUser)
	app.Post("/api/logout", logout)
}
