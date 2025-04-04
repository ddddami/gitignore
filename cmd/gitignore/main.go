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
			fmt.Println("An error occured listing available templates")
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

	templateName := strings.ToLower(flag.Args()[0])
	if err := generateGitignore(templateName); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
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

func loadTemplate(templateName string) (string, error) {
	filename := fmt.Sprintf("templates/%s.gitignore", templateName)
	data, err := templateFS.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(data), nil
}

func generateGitignore(templateName string) error {
	if _, err := os.Stat(".gitignore"); err == nil {
		return fmt.Errorf(".gitignore already exists in current directory")
	}

	content, err := loadTemplate(templateName)
	if err != nil {
		return fmt.Errorf("error loading template '%s': %v", templateName, err)
	}

	err = os.WriteFile(".gitignore", []byte(content), 0o644)
	if err != nil {
		return fmt.Errorf("failed to write .gitignore file: %v", err)
	}
	return nil
}
