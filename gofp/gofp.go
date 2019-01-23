package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	template2 "github.com/logic-building/functional-go/internal/template"
)

var (
	destination = flag.String("destination", "", "functional code for user defined data type")
	pkgName     = flag.String("pkg", "", "package name for generated files")
	types       = flag.String("type", "", "user defined type")
)

func main() {
	fmt.Println("Welcome to functional-go")

	flag.Parse()

	if *destination == "" || *pkgName == "" || *types == "" {
		fmt.Println("go:generate -destination fp.go -source employee.go -pkg Employee")
		os.Exit(1)
	}

	if len(*destination) > 0 {
		if err := os.MkdirAll(filepath.Dir(*destination), os.ModePerm); err != nil {
			log.Fatalf("Unable to create destination directory: %v", err)
		}
		f, err := os.Create(*destination)
		if err != nil {
			log.Fatalf("Failed opening destination file: %v", err)
		}
		generatedCode, err := generateFPCode(*pkgName, *types)
		if err != nil {
			log.Fatalf("Failed code generation: %v", err)
		}
		f.Write([]byte(generatedCode))
		defer f.Close()
	}

	log.Print("Code is generated successfully")
}

func generateFPCode(pkg, dataTypes string) (string, error) {
	conditionalType := ""
	types := strings.Split(dataTypes, ",")

	template := "// Code generated by 'gofp'. DO NOT EDIT.\n"
	template += "package <PACKAGE> \n"

	for _, t := range types {

		if strings.TrimSpace(strings.ToLower(t)) != strings.ToLower(pkg) {
			conditionalType = strings.TrimSpace(t)
		}
		t = strings.TrimSpace(t)
		r := strings.NewReplacer("<PACKAGE>", pkg, "<TYPE>", t, "<CONDITIONAL_TYPE>", conditionalType)

		template = r.Replace(template)

		template += template2.Map()
		template = r.Replace(template)

		template += template2.Filter()
		template = r.Replace(template)

		template += template2.Remove()
		template = r.Replace(template)

		template += template2.Some()
		template = r.Replace(template)
	}

	return template, nil
}
