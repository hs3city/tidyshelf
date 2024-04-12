package app

type App struct {
	Model interface{}
}

func NewApp() *App {
	return &App{}
}
