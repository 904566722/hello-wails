package main

import (
	"context"
	"fmt"
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
func (a *App) Greet(per Person) string {
	return fmt.Sprintf("Hello %s, It's show time!", per)
}

type Person struct {
	Name     string `json:"name"`
	NickName string `json:"nickName"`
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%s)", p.NickName, p.Name)
}
