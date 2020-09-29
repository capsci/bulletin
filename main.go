package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/capsci/bulletin/config"

	"github.com/capsci/bulletin/section"

	"github.com/capsci/bulletin/git"
)

// Arguments
var from, to string
var todo bool
var layout section.Layout

func init() {

	cfg, err := config.GenerateFromYML("./config.yml")
	if err != nil {
		log.Fatal(err)
	}
	// Parse and validate arguments
	flag.StringVar(&from, "from", "", "from commit id/tag")
	flag.StringVar(&to, "to", "", "to commit id/tag")

	flag.BoolVar(&todo, "todo", false, "Generate a todo list for releasenotes")

	flag.Parse()
	if len(from) == 0 || len(to) == 0 {
		log.Fatal("Please provide from and to range")
	}

	if todo {
		return
	}

	for _, sect := range cfg.Sections {
		layout.AddSection(&section.Section{
			Name:           sect.Name,
			TextQualifiers: sect.Qualifiers,
		})
	}

	misc := section.Section{Name: "Misc"}
	layout.SetDefaultSection(&misc)
}

func main() {
	iterator := git.GetLogs(from, to)

	if todo {
		iterator.SortBySubject()
	}

	commit := iterator.Next()
	for commit != nil {
		if todo {
			layout.AddSectionIfAbsent(commit.Author.Name)
			layout.AddToSection(commit.Author.Name, commit)
		} else {
			layout.BucketToSection(commit)
		}
		commit = iterator.Next()
	}

	fmt.Println(layout.GenerateDocumentation())
}
