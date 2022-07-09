package client

import (
	"testing"

	"github.com/NateScarlet/snapshot/pkg/snapshot"
	"github.com/stretchr/testify/require"
)

func TestParseSetCookie(t *testing.T) {
	res, err := parseSetCookie("Cc2838679FT=SOME_RANDOM_VALUE; path=/; expires=Fri, 15 Jul 2022 18:47:41 GMT")
	require.NoError(t, err)
	snapshot.MatchJSON(t, res)

}
