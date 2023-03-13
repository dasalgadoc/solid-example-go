package main

import "solid-example-go/api/application"

func main() {
	var (
		engine = application.BuildEngine(application.BuildApplication())
		server = application.BuildServer(engine)
	)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
