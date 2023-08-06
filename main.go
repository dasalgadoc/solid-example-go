package main

import (
	"solid-example-go/cmd/api"
)

func main() {
	var (
		engine = api.BuildEngine(api.BuildApplication())
		server = api.BuildServer(engine)
	)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
