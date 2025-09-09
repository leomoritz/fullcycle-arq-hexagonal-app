package handler_test

import (
	"testing"

	"github.com/leomoritz/fullcycle-arq-hexagonal-app/adapters/web/handler"
	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello Json"
	result := handler.JsonError(msg)
	require.Equal(t, []byte(`{"message":"Hello Json"}`), result)
}
