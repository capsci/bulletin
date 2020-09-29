package git

// SortBy holds which field should be used to sort
type SortBy int

const (
	// SortBySubject sort by this field
	SortBySubject SortBy = iota
	// SortByAuthor sort by this field
	SortByAuthor
	// SortByTime sort by this field
	SortByTime
)

type commitSorter struct {
	commits []Commit
	sortBy  SortBy
}

func (cs *commitSorter) Len() int {
	return len(cs.commits)
}

func (cs *commitSorter) Swap(i, j int) {
	cs.commits[i], cs.commits[j] = cs.commits[j], cs.commits[i]
}

func (cs *commitSorter) Less(i, j int) bool {
	switch cs.sortBy {
	case SortBySubject:
		return cs.commits[i].Subject < cs.commits[j].Subject
	case SortByAuthor:
		return cs.commits[i].Author.Name < cs.commits[j].Author.Name
	case SortByTime:
		return cs.commits[i].Author.Time.After(cs.commits[j].Author.Time)
	default:
		return cs.commits[i].Author.Name < cs.commits[j].Author.Name
	}
}
