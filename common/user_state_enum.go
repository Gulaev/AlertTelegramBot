package common

type UserState int

const (
	StateStart UserState = iota
	StateChoosingCity
	StateConfirmed
)
