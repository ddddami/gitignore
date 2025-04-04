package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//go:embed templates
var templateFS embed.FS

func main() {
	list := flag.Bool("list", false, "List available templates")
	flag.Parse()

	if *list {
		templates, err := listTemplates()
		if err != nil {
			fmt.Println("An error occured lisiting available templates")
			os.Exit(1)
		}
		fmt.Println("Available templates:")
		for _, t := range templates {
			fmt.Println(" -", t)
		}
		return
	}

	if len(flag.Args()) == 0 {
		fmt.Println("Please specify a template, e.g 'node', 'python'. Run `list` to list all available templates")
		fmt.Println("Available templates...")
		listTemplates()
		return
	}
}

func listTemplates() ([]string, error) {
	entries, err := templateFS.ReadDir("templates")
	if err != nil {
		return nil, err
	}
	var templates []string
	for _, entry := range entries {
		if !entry.IsDir() {

			name := entry.Name()
			templates = append(templates, strings.TrimSuffix(name, filepath.Ext(name)))
		}
	}
	return templates, nil
}
