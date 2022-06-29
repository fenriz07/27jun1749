package entity

type Link struct {
	ID   string
	Code string
}

type entityLinkError string

func (e entityLinkError) Error() string { return string(e) }

const (
	ErrEntityNotFound = entityLinkError("error link entity not found")
	ErrCreatingEntity = entityLinkError("error generating the url")
)
