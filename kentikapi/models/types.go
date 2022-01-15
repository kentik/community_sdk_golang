package models

// ID is a common identifier type across all resources.
type ID = int

func IDPtr(i ID) *ID {
	return &i
}
