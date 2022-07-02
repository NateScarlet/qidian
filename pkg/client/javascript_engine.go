package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type JavaScriptEngine interface {
	Run(ctx context.Context, unsafeJS string) (output string, err error)
}

type nodeJSEngine struct {
	nodePath string
}

// Run implements JavaScriptEngine
func (e nodeJSEngine) Run(ctx context.Context, unsafeJS string) (output string, err error) {
	var cmd = exec.CommandContext(ctx, e.nodePath, "-")
	var stdin bytes.Buffer
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	_, err = stdin.WriteString(`
const vm = require('node:vm');

const code = `)
	if err != nil {
		return
	}
	var encoder = json.NewEncoder(&stdin)
	err = encoder.Encode(unsafeJS)
	if err != nil {
		return
	}
	_, err = stdin.WriteString(`;
(async () => {
	console.log(await vm.runInNewContext(code));
})().catch((err) => {
	console.error(err);
	process.exit(1);
})
`)
	if err != nil {
		return
	}
	cmd.Stdin = &stdin
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		if stderr.Len() > 0 {
			err = fmt.Errorf("%w: %s", err, stderr.String())
		}
		return
	}
	return strings.TrimRight(stdout.String(), "\n"), nil
}

type errorJSEngine struct {
	err error
}

func (e errorJSEngine) Run(ctx context.Context, unsafeJS string) (output string, err error) {
	return "", e.err
}

func NewNodeJSEngine(nodePath string) JavaScriptEngine {
	return nodeJSEngine{nodePath}
}

var DefaultJSEngine JavaScriptEngine

func init() {
	var nodePath, err = exec.LookPath("node")
	if err != nil {
		DefaultJSEngine = errorJSEngine{
			fmt.Errorf("qidian: client: nodejs executable not found, please configure javascript engine"),
		}
		return
	}

	DefaultJSEngine = NewNodeJSEngine(nodePath)
}

type contextKeyJavaScriptEngine struct{}

func WithJavaScriptEngine(ctx context.Context, v JavaScriptEngine) context.Context {
	return context.WithValue(ctx, contextKeyJavaScriptEngine{}, v)
}

func ContextJavaScriptEngine(ctx context.Context) JavaScriptEngine {
	if jsEngine, ok := ctx.Value(contextKeyJavaScriptEngine{}).(JavaScriptEngine); ok {
		return jsEngine
	}
	return DefaultJSEngine
}
