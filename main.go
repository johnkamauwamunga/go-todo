package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("hello john, this is the innitial setup for a fiber app")

	a := 3
	b := 4

	fmt.Println("sum is:", a+b)
	app := fiber.New()

	log.Fatal(app.Listen(":4000"))
}
