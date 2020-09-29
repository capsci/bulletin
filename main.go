package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/capsci/bulletin/section"

	"github.com/capsci/bulletin/git"
)

// Arguments
var from, to string
var todo bool
var layout section.Layout

func init() {
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

	// Sections : features,fixes,enhancements,documentations and misc
	features := section.Section{Name: "Features"}
	features.AddTextQualifiers([]string{":sparkles:", ":racehorse:", ":zap:", ":ok_hand:", ":tada:", ":white_check_mark:"})
	features.AddTextQualifiers("feat:")

	fixes := section.Section{Name: "Fixes"}
	fixes.AddTextQualifiers([]string{":bug:", ":ambulance:", ":red:"})
	fixes.AddTextQualifiers("fix:")

	enhancements := section.Section{Name: "Enhancements"}
	enhancements.AddTextQualifiers([]string{":art:", ":hammer:", ":construction:", ":wrench:", ":ok_hand:", ":closed_lock_with_key:"})
	enhancements.AddTextQualifiers([]string{"refactor:", "style:", "chore:", "test:"})

	documentations := section.Section{Name: "Documentations"}
	documentations.AddTextQualifiers([]string{":books:", ":notebook:"})
	documentations.AddTextQualifiers("docs:")

	misc := section.Section{Name: "Misc"}

	layout.AddSection(&features)
	layout.AddSection(&fixes)
	layout.AddSection(&enhancements)
	layout.AddSection(&documentations)
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
