package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/drgo/mdson"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func run() error {
	files, err := filepath.Glob("./content/*.mdson")
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return fmt.Errorf("nothing to parse ")
	}
	t, err := template.ParseGlob("./layouts/html/*.html")
	if err != nil {
		return err
	}
	for _, f := range files {
		data, err := parse(f)
		if err != nil {
			return err
		}
		err = t.ExecuteTemplate(os.Stdout, "datasets", data)
		if err != nil {
			return err
		}
	}
	return nil
}

func parse(fileName string) (mdson.Node, error) {
	root, err := mdson.ParseFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file '%s': %v", fileName, err)
	}
	if root.Kind() != "Block" {
		return nil, fmt.Errorf("parser returned unexpected type: root is not a block")
	}
	return root, nil
}
