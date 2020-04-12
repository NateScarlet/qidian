package font

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/image/font/sfnt"
)

func TestDeobfuscate(t *testing.T) {
	ttf, err := ioutil.ReadFile("sample.ttf")
	require.NoError(t, err)
	font, err := sfnt.Parse(ttf)
	require.NoError(t, err)

	result, err := Deobfuscate("&#100187;&#100185;&#100188;&#100190;&#100190;", font)
	assert.NoError(t, err)
	assert.Equal(t, "68.55", result)
}
