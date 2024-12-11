package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandlerJsonEr(t *testing.T) {
	msg := "test error"

	result := jsonError(msg)

	require.Equal(t, []byte(`{"message":"test error"}`), result)
}
