package main

import (
	"github.com/JulyInSummer/cinematic/internal/app"
)

// @title           Cinematic API Specification
// @version         1.0
// @description     This is a simple Cinematic service which exposes CRUD APIs on movies
//
// @contact.name   JulyInSummer
// @contact.email  azimjanovbogdan@gmail.com
//
// @host      localhost:8009
// @BasePath  /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app.NewApp().Run()
}
