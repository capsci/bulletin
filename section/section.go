package section

import (
	"log"
	"strings"

	"github.com/capsci/bulletin/git"
	"github.com/capsci/bulletin/markdown"
)

// Section is a section in bulletin
type Section struct {
	Name           string
	TextQualifiers []string
	Commits        []*git.Commit
}

// AddTextQualifiers adds new emojis to existing list
func (s *Section) AddTextQualifiers(qualifiers interface{}) {
	switch qualifiers.(type) {
	case []string:
		s.TextQualifiers = append(s.TextQualifiers, qualifiers.([]string)...)
	case string:
		s.TextQualifiers = append(s.TextQualifiers, qualifiers.(string))
	default:
		log.Fatal("ghj")
	}
}

// AddCommit adds commit to section
func (s *Section) AddCommit(commit *git.Commit) {
	s.Commits = append(s.Commits, commit)
}

// Belongs check if the commit message belongs to a section based on commit message
func (s *Section) Belongs(commitMessage string) bool {
	for _, marker := range s.TextQualifiers {
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
