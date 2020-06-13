package font

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	f, err := Get(URL("DUigFnRh"))
	require.NoError(t, err)
	assert.NotEmpty(t, f)
}
