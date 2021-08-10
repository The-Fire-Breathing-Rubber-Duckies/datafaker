package main

import (
	"github.com/the-fire-breathing-duckies/datafaker/cmd"
)

func main() {
	cmd.Execute()
	router := NewRouter()
	router.Logger.Fatal(router.Start(":1234"))
}
