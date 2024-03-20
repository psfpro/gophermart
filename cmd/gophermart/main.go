package main

import "github.com/psfpro/gophermart/internal/gophermart"

func main() {
	container := gophermart.NewContainer()
	app := container.App()
	app.Run()
}
