package main

import (
	"golang.org/x/sync/errgroup"
	"log"
	"myblog/router"
)

var g errgroup.Group

func main() {

	g.Go(func() error {
		return router.BackendServer().ListenAndServe()
	})

	g.Go(func() error {
		return router.BlogServer().ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}
