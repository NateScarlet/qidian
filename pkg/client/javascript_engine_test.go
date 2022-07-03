package client

import (
	"context"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNodeJSEngine(t *testing.T) {

	nodePath, err := exec.LookPath("node")
	if err != nil {
		t.Skip("nodejs executable not found")
		return
	}
	var e = NewNodeJSEngine(nodePath)

	var ctx = context.Background()
	for _, tt := range []struct {
		name         string
		giveUnsafeJS string
		wantOutput   string
		wantError    string
	}{
		{name: "should return value", giveUnsafeJS: "'aa'", wantOutput: "aa"},
		{name: "should allow variable", giveUnsafeJS: "a = 1; a", wantOutput: "1"},
		{name: "should now allow require", giveUnsafeJS: "require('fs')", wantError: "ReferenceError: require is not defined"},
		{name: "should allow throw error", giveUnsafeJS: "throw 'my error'", wantError: "exit status 1: my error"},
		{name: "should wait promise", giveUnsafeJS: "Promise.resolve(3)", wantOutput: "3"},
		{name: "should error on invalid syntax", giveUnsafeJS: "(", wantError: "SyntaxError"},
		{name: "should allow eval", giveUnsafeJS: "eval('a = 1'); a", wantOutput: "1"},
	} {
		t.Run(tt.name, func(t *testing.T) {

			output, err := e.Run(ctx, tt.giveUnsafeJS)
			if tt.wantError != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.wantError)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantOutput, output)
		})

	}
}
