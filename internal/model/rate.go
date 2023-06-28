package model

import "github.com/google/uuid"

type Rate struct {
	Base
	Buy       float32    `json:"buy,omitempty"`
	Sell      float32    `json:"sell,omitempty"`
	ProjectID *uuid.UUID `json:"project_id,omitempty"`
}
