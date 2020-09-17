package section

import (
	"github.com/capsci/bulletin/git"
)

// Layout for client to use
type Layout struct {
	sections       []*Section
	defaultSection *Section
	emojiMarker    bool
}

// AddSection adds new section
func (l *Layout) AddSection(section *Section) {
	l.sections = append(l.sections, section)
}

// SetDefaultSection adds new section
func (l *Layout) SetDefaultSection(section *Section) {
	l.defaultSection = section
}

// BucketOnEmojiMarkers buckets commits based on emojis
func (l *Layout) BucketOnEmojiMarkers() {
	l.emojiMarker = true
}

// BucketOnTextMarkers buckets commits based on emojis
func (l *Layout) BucketOnTextMarkers() {
	l.emojiMarker = false
}

// BucketToSection buckets to appropriate section
func (l *Layout) BucketToSection(commit *git.Commit) {
	for _, section := range l.sections {
		if section.Belongs(commit.FullCommitMessage(), l.emojiMarker) {
			section.AddCommit(commit)
			return
		}
	}
	l.defaultSection.AddCommit(commit)
}

// GenerateDocumentation generates documentation
func (l *Layout) GenerateDocumentation() string {
	var doc string
	for _, section := range l.sections {
		doc += section.GenerateDocumentation()
	}
	doc += l.defaultSection.GenerateDocumentation()
	return doc
}
