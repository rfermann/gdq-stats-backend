// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

type CreateEvent struct {
	Name string `json:"name"`
	Year int64  `json:"year"`
}

type CreateEventTypeInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteEventTypeInput struct {
	ID string `json:"id"`
}

type MigrateEventDataInput struct {
	ID string `json:"id"`
}

type UpdateEventTypeInput struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}