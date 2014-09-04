package main

import (
	log "github.com/cihub/seelog"
	"github.com/go-martini/martini"
	"github.com/ostensus/ostn/entropy"
	"os"
	"time"
)

func main() {

	vs, err := entropy.OpenStore("y.db")
	if err != nil {
		log.Errorf("Could not access version store: %s", err)
		os.Exit(1)
	}

	m := martini.Classic()

	m.Get("/:version", func(params martini.Params) (int, string) {

		ev := entropy.NewDatePartitionedEvent("id", params["version"], "ts", time.Now())

		if err := vs.Accept(ev); err != nil {
			return 400, "Bad event"
		}

		digests, err := vs.Digest(200)
		if err != nil {
			return 400, "No digest"
		}

		return 200, digests["id"]
	})

	m.Run()
}
