package token

import "time"

// Maker is the interface that wraps the basic Token operations.

type Maker interface {
	// Create creates a new token for a specific username and duration.
	CreateToken(username string, duration time.Duration) (string, error)
	// Verify verifies if a token is valid or not.
	VerifyToken(token string) (*Payload, error)
}
