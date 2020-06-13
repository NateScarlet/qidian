// +build generate

package book

// Run go get github.com/NateScarlet/gotmpl/cmd/gotmpl to get gotmpl command.
//go:generate go run ../../cmd/categories -o category.json
//go:generate gotmpl -i category.json -o category.go category.go.gotmpl
