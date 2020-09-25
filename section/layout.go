package section

import (
	"fmt"
	"log"

	"github.com/capsci/bulletin/git"
)

// Layout for client to use
type Layout struct {
	sections       []*Section
	defaultSection *Section
}

// AddSection adds new section
func (l *Layout) AddSection(section *Section) {
	l.sections = append(l.sections, section)
}

// AddSectionIfAbsent adds new section
func (l *Layout) AddSectionIfAbsent(name string) {
	for _, section := range l.sections {
		if section.Name == name {
			return
		}
	}
	l.sections = append(l.sections, &Section{Name: name})
}

// SetDefaultSection adds new section
func (l *Layout) SetDefaultSection(section *Section) {
	l.defaultSection = section
}

// PrintSectioNames prints all section names
func (l *Layout) PrintSectioNames() {
	fmt.Println("Section names are:")
	for _, section := range l.sections {
		fmt.Println(section.Name)
	}
}

// AddToSection buckets to appropriate section
func (l *Layout) AddToSection(sectionName string, commit *git.Commit) {
	for _, section := range l.sections {
		if section.Name == sectionName {
			section.AddCommit(commit)
			return
		}
	}
	log.Fatal("Count not find section: " + sectionName)
}

// BucketToSection buckets to appropriate section
func (l *Layout) BucketToSection(commit *git.Commit) {
	for _, section := range l.sections {
		if section.Belongs(commit.FullCommitMessage()) {
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
	if l.defaultSection != nil {
		doc += l.defaultSection.GenerateDocumentation()
	}
	return doc
}
