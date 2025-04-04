package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//go:embed templates/*.gitignore
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
		fmt.Println("Usage: gitignore <template> e.g 'gitignore python'. \nRun `gitignore list` to see available templates")
		return
	}

	template := strings.ToLower(flag.Args()[0])
	content, err := LoadTemplate(template)
	if err != nil {
		fmt.Printf("Error loading template '%s': %v\n", template, err)
		os.Exit(1)
	}

	fmt.Print(content)
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

func LoadTemplate(templateName string) (string, error) {
	filename := fmt.Sprintf("templates/%s.gitignore", templateName)
	data, err := templateFS.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(data), nil
}
