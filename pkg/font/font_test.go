package font

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	f, err := Get(context.Background(), URL("DUigFnRh"))
	require.NoError(t, err)
	assert.NotEmpty(t, f)
}
