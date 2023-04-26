package author

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {
	author := &Author{ID: "4362771"}
	err := author.Fetch(context.Background())
	require.NoError(t, err)

	assert.Equal(t, "4362771", author.ID)
	assert.Equal(t, "7687417", author.UserID)
	assert.Equal(t, "忘语", author.Name)
	assert.Equal(t, "https://ccportrait.yuewen.com/apimg/349573/p_3626391704023601/100", author.AvatarURL)
	assert.Equal(t, "阅文集团白金作家，《凡人修仙传》一书在业内创造传奇，成为“凡人流”作品开山鼻祖。", author.Biography)
}
