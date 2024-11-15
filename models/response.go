package models

import "github.com/google/uuid"

type AuthTokenPayload struct {
	ID        uuid.UUID `json:"userid"`
	Username  string    `json:"username"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Type      string    `json:"type"`
}

type UserProfile struct {
	ID        uuid.UUID `json:"userid"`
	Username  string    `json:"username"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Type      string    `json:"type"`
}
