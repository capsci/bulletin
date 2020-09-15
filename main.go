package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/capsci/bulletin/markdown"

	"github.com/capsci/bulletin/git"
)

// Arguments
var from, to string

func init() {
	// Parse and validate arguments
	flag.StringVar(&from, "from", "", "from commit id/tag")
	flag.StringVar(&to, "to", "", "to commit id/tag")
	flag.Parse()
	if len(from) == 0 || len(to) == 0 {
		log.Fatal("Please provide from and to range")
	}
}

func main() {
	iterator := git.GetLogs(from, to)

	var features []*git.Commit
	var fixes []*git.Commit
	var enhancements []*git.Commit
	var documentations []*git.Commit
	var misc []*git.Commit

	commit := iterator.Next()
	for commit != nil {
		if sorter(commit.FullCommitMessage(), []string{":sparkles:", ":racehorse:", ":zap:", ":ok_hand:", ":tada:", ":white_check_mark:"}) {
			features = append(features, commit)
		} else if sorter(commit.FullCommitMessage(), []string{":bug:", ":ambulance:", ":red:"}) {
			fixes = append(fixes, commit)
		} else if sorter(commit.FullCommitMessage(), []string{":art:", ":hammer:", ":construction:", ":wrench:", ":ok_hand:", ":closed_lock_with_key:"}) {
			enhancements = append(enhancements, commit)
		} else if sorter(commit.FullCommitMessage(), []string{":books:", ":notebook:"}) {
			documentations = append(documentations, commit)
		} else {
			misc = append(misc, commit)
		}
		commit = iterator.Next()
	}

	doc := generateDoc("Features", features)
	doc += generateDoc("Fixes", fixes)
	doc += generateDoc("Enhancements", enhancements)
	doc += generateDoc("Documentation", documentations)
	doc += generateDoc("Misc", misc)
	fmt.Println(doc)
}

func sorter(commitMessage string, emojiShortcodes []string) bool {
	for _, emoji := range emojiShortcodes {
		if strings.Contains(commitMessage, emoji) {
			return true
		}
	}
	return false
}

func generateDoc(heading string, commits []*git.Commit) string {
	doc := markdown.GetH3(heading)
	for _, commit := range commits {
		doc += markdown.GetListItem(commit.Subject)
	}
	return doc
}
