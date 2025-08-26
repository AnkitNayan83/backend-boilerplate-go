package token

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID     uuid.UUID `json:"id"`
	UserID int32     `json:"user_id"`
	// add more fields if needed
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewPayload(userID int32, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenId,
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if payload.ExpiresAt.Before(time.Now()) {
		return ErrExpiredToken
	}

	return nil
}
