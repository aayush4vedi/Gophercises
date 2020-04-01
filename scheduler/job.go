package scheduler

import (
	"encoding/json"
	"io"
)

type Priority uint8

const (
	_ Priority = iota
	P0
	P1
	P2
)

type UserType uint8

const (
	_ UserType = iota
	User
	Admin
	Root
)

type Job struct {
	Name     string `json:"name"`
	Duration uint16 `json:"duration"`
	Deadline uint16 `json:"deadline"`
	Priority `json:"priority"`
	UserType `json:"usertype"`
}

type JobList struct {
	Jobs []Job
}

func JsonParser(r io.Reader) ([]Job, error) {
	d := json.NewDecoder(r)
	jobsList := make([]Job, 0)
	if err := d.Decode(&jobsList); err != nil {
		return jobsList, err
	}
	return jobsList, nil
}
