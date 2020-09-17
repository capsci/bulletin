package section

import (
	"strings"

	"github.com/capsci/bulletin/git"
	"github.com/capsci/bulletin/markdown"
)

// Section is a section in bulletin
type Section struct {
	Name            string
	EmojiShortcodes []string
	TextShortcodes  []string
	Commits         []*git.Commit
}

// AddEmojiShortcodes adds new emojis to existing list
func (s *Section) AddEmojiShortcodes(emojis []string) {
	s.EmojiShortcodes = append(s.EmojiShortcodes, emojis...)
}

// AddTextMarkers adds new emojis to existing list
func (s *Section) AddTextMarkers(markers []string) {
	s.TextShortcodes = append(s.TextShortcodes, markers...)
}

// AddCommit adds commit to section
func (s *Section) AddCommit(commit *git.Commit) {
	s.Commits = append(s.Commits, commit)
}

// Belongs check if the commit message belongs to a section based on commit message
func (s *Section) Belongs(commitMessage string, matchEmoji bool) bool {
	var markers []string
	if matchEmoji {
		markers = s.EmojiShortcodes
	} else {
		markers = s.TextShortcodes
	}
	for _, marker := range markers {
		if strings.Contains(commitMessage, marker) {
			return true
		}
	}
	return false
}

// GenerateDocumentation generates documentation
func (s *Section) GenerateDocumentation() string {
	doc := markdown.GetH3(s.Name)
	for _, commit := range s.Commits {
		doc += markdown.GetListItem(commit.Subject)
	}
	return doc
}
