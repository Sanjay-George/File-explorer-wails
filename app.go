package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
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
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, Let's do this now!", name)
}

// Get list of filenames from a directory
func (a *App) GetFiles() ([]string, error) {
	dir := "./" // Directory to list files from
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("Failed to read directory: %v", err)
		return nil, err
	}
	var filenames []string
	for _, file := range files {
		if !file.IsDir() {
			filenames = append(filenames, file.Name())
		}
	}
	return filenames, nil
}

func (a *App) GetFolders() ([]string, error) {
	dir := "./" // Directory to list folders from
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("Failed to read directory: %v", err)
		return nil, err
	}
	var folderNames []string
	for _, file := range files {
		if file.IsDir() {
			folderNames = append(folderNames, file.Name())
		}
	}
	return folderNames, nil
}
