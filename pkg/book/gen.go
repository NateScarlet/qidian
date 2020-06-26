// +build generate

package book

// Run go get github.com/NateScarlet/gotmpl/cmd/gotmpl to get gotmpl command.

//go:generate gotmpl -i ../../internal/data/category.json -o category.go category.go.gotmpl
