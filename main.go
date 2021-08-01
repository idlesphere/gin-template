package main

import (
	"gin-template/app"
	env "gin-template/config"
)

func main() {
	app.Start(env.Get("PORT"))
}
