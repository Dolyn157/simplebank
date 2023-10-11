package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrInvalidSigningMethodToken = errors.New("token is sgined with an invalid algorithm")
var ErrInvalidToken = errors.New("token is invalid")
var ErrExpiredToken = errors.New("token is expired")

// Payload cotaines the payload data of a token.
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// Valid implements jwt.Claims.
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

// NewPayload creates a new payload with a specific username and duration.
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	TokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        TokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}
