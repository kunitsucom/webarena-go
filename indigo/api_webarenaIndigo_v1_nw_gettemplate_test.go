package indigo

import (
	"context"
	"math"
	"testing"

	require "github.com/kunitsucom/util.go/testing/require"
)

func TestClient_GetWebArenaIndigoV1NwGetTemplate(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		client := NewTestClient(ctx, t)

		resp, err := client.GetWebArenaIndigoV1NwGetTemplate(ctx, math.MaxInt)
		require.ErrorIs(t, err, ErrUnexpectedStatusCode)
		require.Nil(t, resp)
	})
}