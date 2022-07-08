package client

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"text/template"
)

type JavaScriptEngine interface {
	// Run unsafeJS code, supports `eval` `escape`.
	Run(ctx context.Context, unsafeJS string) (output string, err error)
}

type nodeJSEngine struct {
	nodePath string
}

func NewNodeJSEngine(nodePath string) JavaScriptEngine {
	return nodeJSEngine{nodePath}

}

//go:embed javascript_engine.node.runner.cjs
var nodeRunnerJS string

var nodeJSRunTemplate = template.Must(template.New("").Funcs(template.FuncMap{
	"toJSON": func(s string) (_ string, err error) {
		data, err := json.Marshal(s)
		if err != nil {
			return
		}
		return string(data), nil
	},
	"__DEBUG__": func() bool {
		return isDebug
	},
	"runnerJS": func() string {
		return nodeRunnerJS
	},
}).Parse(`
"use strict";
const __DEBUG__ = {{ __DEBUG__ }};
const __CODE__ = {{ . | toJSON }};

{{ runnerJS }}
`))

// Run implements JavaScriptEngine
func (e nodeJSEngine) Run(ctx context.Context, unsafeJS string) (output string, err error) {
	var cmd = exec.CommandContext(ctx, e.nodePath, "-")
	var stdin bytes.Buffer
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	err = nodeJSRunTemplate.Execute(&stdin, unsafeJS)
	if err != nil {
		return
	}
	cmd.Stdin = &stdin
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		if stderr.Len() > 0 {
			print(stderr.String())
			err = fmt.Errorf("%w: %s", err, stderr.String())
		}
		return
	}
	return strings.TrimRight(stdout.String(), "\n"), nil
}

type errJSEngine struct {
	err error
}

func (e errJSEngine) Run(ctx context.Context, unsafeJS string) (output string, err error) {
	return "", e.err
}

var DefaultJavaScriptEngine JavaScriptEngine

func init() {
	var nodePath, err = exec.LookPath("node")
	if err != nil {
		DefaultJavaScriptEngine = errJSEngine{
			fmt.Errorf("qidian: client: nodejs executable not found, please configure DefaultJavaScriptEngine manually"),
		}
		return
	}

	DefaultJavaScriptEngine = NewNodeJSEngine(nodePath)
}

type contextKeyJavaScriptEngine struct{}

func WithJavaScriptEngine(ctx context.Context, v JavaScriptEngine) context.Context {
	return context.WithValue(ctx, contextKeyJavaScriptEngine{}, v)
}

func ContextJavaScriptEngine(ctx context.Context) JavaScriptEngine {
	if jsEngine, ok := ctx.Value(contextKeyJavaScriptEngine{}).(JavaScriptEngine); ok {
		return jsEngine
	}
	return DefaultJavaScriptEngine
}
