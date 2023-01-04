package main

import (
	"app/handlers"
	"github.com/cdanyl/celeritas"
	"log"
	"os"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cel := &celeritas.Celeritas{
		AppName: "app",
	}

	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	myHandlers := &handlers.Handlers{App: cel}

	app := &application{
		App:      cel,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	return app
}
