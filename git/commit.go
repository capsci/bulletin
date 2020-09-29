package git

import (
	"fmt"
	"sort"
)

// CommitIterator provides interator to run through available commits
type CommitIterator struct {
	commits []Commit
	offset  int
}

// GetCommitIterator gets a new log iterator
func GetCommitIterator() *CommitIterator {
	return &CommitIterator{
		offset: 0,
	}
}

// Reset resets the iterator to point to beginning
func (ci *CommitIterator) Reset() {
	ci.offset = 0
}

// Size returns the number of log items in the iterator
func (ci *CommitIterator) Size() int {
	return len(ci.commits)
}

// Push adds a new item to LogIterator
func (ci *CommitIterator) Push(commit Commit) {
	ci.commits = append(ci.commits, commit)
}

// PrintAll prints all commit messages
func (ci *CommitIterator) PrintAll() {
	for _, item := range ci.commits {
		fmt.Println(item)
	}
}

// Next returns item in the log iterator
func (ci *CommitIterator) Next() *Commit {
	if ci.offset >= ci.Size() {
		return nil
	}
	ci.offset++
	return &ci.commits[ci.offset-1]
}

// SortBySubject sorts commits based on Subject
func (ci *CommitIterator) SortBySubject() {
	ci.sortByField(SortBySubject)
}

// SortByAuthor sorts commits based on Subject
func (ci *CommitIterator) SortByAuthor() {
	ci.sortByField(SortByAuthor)
}

// SortByTime sorts commits based on Subject
func (ci *CommitIterator) SortByTime() {
	ci.sortByField(SortByTime)
}

func (ci *CommitIterator) sortByField(field SortBy) {
	cs := &commitSorter{
		commits: ci.commits,
		sortBy:  field,
	}
	sort.Sort(cs)
}

// Set sets commit value from records
func (commit *Commit) Set(record []string) {
	commit.CommitHash.Abbreviated = record[0]
	commit.CommitHash.Full = record[1]
	commit.Author.Name = record[2]
	commit.Author.Email = record[3]
	// TODO: Read Time
	commit.Subject = record[5]
	commit.Body = record[6]
}

// FullCommitMessage gets full commit message
func (commit *Commit) FullCommitMessage() string {
	return commit.Subject + "\n" + commit.Body
}
