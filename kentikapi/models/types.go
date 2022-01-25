package models

// ID is a common identifier type across all resources.
type ID = string

func IDPtr(i ID) *ID {
	return &i
}
