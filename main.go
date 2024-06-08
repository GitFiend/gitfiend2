package main

import (
	"embed"
	"fmt"
	"gitfiend2/core/git"
	"gitfiend2/core/server"
	"os"
	"path"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:generate go run core/parser/genand/main.go
//go:generate gofmt -w core/parser

func main() {
	//runCommitsTest()
	//runWails()
	server.StartServer()
}

func runCommitsTest() {
	start := time.Now()
	home, err := os.UserHomeDir()

	if err != nil {
		fmt.Println(err)
		return
	}

	res := git.LoadCommits(git.RunOpts{RepoPath: path.Join(home, "Repos", "vscode")}, 1000)
	duration := time.Since(start)
	fmt.Println("********** vscode: ", duration.Milliseconds(), "ms")
	fmt.Printf("Loaded %d commits", len(res))
}

func runWails() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "gitfiend2",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []any{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
