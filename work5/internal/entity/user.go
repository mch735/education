package entity

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        string    `json:"id"`
}

func (u *User) MarshalBinary() ([]byte, error) {
	buf, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("user marshal binary error: %w", err)
	}

	return buf, nil
}

func (u *User) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, u)
	if err != nil {
		return fmt.Errorf("user unmarshal binary error: %w", err)
	}

	return nil
}
