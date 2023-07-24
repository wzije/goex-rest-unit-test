package main

import "gitlab.com/jeelabs/learnings/goex-rest-unit-test/router"

func main() {
	r := router.RegisterRouter()
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
