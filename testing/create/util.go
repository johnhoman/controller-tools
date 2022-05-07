package create

import (
	"github.com/google/uuid"
)

func uid() string { return uuid.New().String()[:8] }
