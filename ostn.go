package main

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/ostensus/ostn/entropy"
	"os"
	"strconv"
	"time"
)

const attributeName = "ts_col"

func main() {

	vs, err := entropy.OpenStore("y.db")
	if err != nil {
		log.Errorf("Could not access version store: %s", err)
		os.Exit(1)
	}

	m := martini.Classic()

	type RepoForm struct {
		Name string `form:"name" binding:"required"`
	}

	type ChangeEventForm struct {
		Id      string `form:"id" binding:"required"`
		Version string `form:"version" binding:"required"`
	}

	m.Post("/repos", binding.Bind(RepoForm{}), func(repo RepoForm) (int, string) {

		parts := make(map[string]entropy.RangePartitionDescriptor)
		parts[attributeName] = entropy.RangePartitionDescriptor{}

		id, err := vs.NewRepository(repo.Name, parts)
		if err != nil {
			return 500, err.Error()
		}

		return 201, fmt.Sprintf("%d", id)
	})

	m.Post("/events/:repo", binding.Bind(ChangeEventForm{}), func(params martini.Params, form ChangeEventForm) (int, string) {

		repo, _ := strconv.ParseInt(params["repo"], 0, 64)
		ev := entropy.NewDatePartitionedEvent(form.Id, form.Version, attributeName, time.Now())

		if err := vs.Accept(repo, ev); err != nil {
			return 400, "Bad event"
		}

		return 202, "Good event"
	})

	m.Get("/digests/:repo", func(params martini.Params) (int, string) {

		repo, _ := strconv.ParseInt(params["repo"], 0, 64)

		digests, err := vs.Digest(repo)
		if err != nil {
			return 400, "No digest"
		}

		log.Infof("Digests: %+v", digests)

		// For now we'll just grab the first cab off the rank
		var digest string
		for _, digest = range digests {
			break
		}

		return 200, digest
	})

	m.Run()
}
