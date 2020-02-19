package hgraph

import (
	"strings"

	"github.com/google/uuid"
)

func UUID() string {
	return uuid.New().String()
}

func UUID32() string {
	return strings.ReplaceAll(UUID(), "-", "")
}
