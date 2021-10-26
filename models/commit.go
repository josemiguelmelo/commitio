package models

import (
	"fmt"
)

type Commit struct {
	Subject string
	Body    *string
	Type    *ChangeType
	Issue   *Issue
}

func NewCommit(subject string, body *string, issue *Issue, commitType *ChangeType) Commit {
	return Commit{
		Subject: subject,
		Body:    body,
		Issue:   issue,
		Type:    commitType,
	}
}

func (c Commit) GetTypeOrDefault() ChangeType {
	if c.Type == nil {
		return DefaultType
	}
	return *c.Type
}

func (c Commit) GetSubject() string {
	baseSubject := fmt.Sprintf("%s: %s", c.GetTypeOrDefault(), c.Subject)
	if c.Issue != nil {
		return fmt.Sprintf("%s %s", baseSubject, c.Issue.Reference)
	}
	return baseSubject
}
