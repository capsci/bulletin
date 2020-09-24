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
var emoji, text bool
var layout section.Layout

func init() {
	// Parse and validate arguments
	flag.StringVar(&from, "from", "", "from commit id/tag")
	flag.StringVar(&to, "to", "", "to commit id/tag")

	flag.BoolVar(&emoji, "emoji", false, "Check for emoji markers")
	flag.BoolVar(&text, "text", false, "Check for text markers")

	flag.Parse()
	if len(from) == 0 || len(to) == 0 {
		log.Fatal("Please provide from and to range")
	}
	if emoji == text {
		log.Fatal("Pass either -emoji or -text")
	}

	// Sections : features,fixes,enhancements,documentations and misc
	features := section.Section{Name: "Features"}
	features.AddEmojiShortcodes([]string{":sparkles:", ":racehorse:", ":zap:", ":ok_hand:", ":tada:", ":white_check_mark:"})
	features.AddTextMarkers([]string{"feat:"})

	fixes := section.Section{Name: "Fixes"}
	fixes.AddEmojiShortcodes([]string{":bug:", ":ambulance:", ":red:"})
	fixes.AddTextMarkers([]string{"fix:"})

	enhancements := section.Section{Name: "Enhancements"}
	enhancements.AddEmojiShortcodes([]string{":art:", ":hammer:", ":construction:", ":wrench:", ":ok_hand:", ":closed_lock_with_key:"})
	enhancements.AddTextMarkers([]string{"refactor:", "style:", "chore:", "test:"})

	documentations := section.Section{Name: "Documentations"}
	documentations.AddEmojiShortcodes([]string{":books:", ":notebook:"})
	documentations.AddTextMarkers([]string{"docs:"})

	misc := section.Section{Name: "Misc"}

	layout.AddSection(&features)
	layout.AddSection(&fixes)
	layout.AddSection(&enhancements)
	layout.AddSection(&documentations)
	layout.SetDefaultSection(&misc)
	if emoji {
		layout.BucketOnEmojiMarkers()
	} else {
		layout.BucketOnTextMarkers()
	}

}

func main() {
	iterator := git.GetLogs(from, to)

	commit := iterator.Next()
	for commit != nil {
		layout.BucketToSection(commit)
		commit = iterator.Next()
	}

	fmt.Println(layout.GenerateDocumentation())
}
