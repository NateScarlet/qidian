package book

import (
	"context"
	"testing"

	"github.com/NateScarlet/snapshot/pkg/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func snapshotBook(t *testing.T, book Book, opts ...snapshot.Option) {
	snapshot.MatchJSON(
		t,
		book,
		append(
			[]snapshot.Option{
				snapshot.OptionCleanRegex(
					snapshot.CleanAs(`"*count*"`),
					`(?m)^\s*".+Count": (\d+),?$`,
				),
			},
			opts...,
		)...,
	)
}

func TestBook_Fetch(t *testing.T) {
	var ctx = context.Background()

	var b = Book{ID: "1"}
	var err = b.Fetch(ctx)
	require.NoError(t, err)
	snapshotBook(t, b)
	assert.Equal(t, uint64(987300), b.WordCount)
	assert.LessOrEqual(t, uint64(94200), b.TotalRecommendCount)
}

func TestBook_Fetch_Free(t *testing.T) {
	var ctx = context.Background()

	var b = Book{ID: "8361"}
	var err = b.Fetch(ctx)
	require.NoError(t, err)
	snapshotBook(t, b)
	assert.Equal(t, uint64(731700), b.WordCount)
	assert.LessOrEqual(t, uint64(476300), b.TotalRecommendCount)
}
