package models

type Issue struct {
	Reference string
}

func NewIssue(reference string) Issue {
	return Issue{Reference: reference}
}
