package models

import "github.com/google/uuid"

type AuthTokenPayload struct {
	ID        uuid.UUID `json:"userid"`
	Username  string    `json:"username"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
}

type UserProfile struct {
	ID        uuid.UUID `json:"userid"`
	Username  string    `json:"username"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
}
