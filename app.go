package main

import (
	"context"

	"github.com/NuruProgramming/NuruIDLE/object"
	"github.com/NuruProgramming/NuruIDLE/repl"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Start(name string) (string, []string) {
	env := object.NewEnvironment()
	return repl.Start(name, env)
}
