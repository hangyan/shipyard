package main

import (
	"fmt"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/hangyan/shipyard/client"
)

var destroyCommand = cli.Command{
	Name:        "destroy",
	Usage:       "destroy a container",
	Description: "destroy <id> [<id>]",
	Action:      destroyAction,
}

func destroyAction(c *cli.Context) {
	cfg, err := loadConfig()
	if err != nil {
		logger.Fatal(err)
	}
	m := client.NewManager(cfg)
	containers, err := m.Containers(false)
	if err != nil {
		logger.Fatalf("error getting container info: %s", err)
	}
	ids := c.Args()
	if len(ids) == 0 {
		logger.Fatalf("you must specify at least one id")
	}
	for _, cnt := range containers {
		// this can probably be more efficient
		for _, i := range ids {
			if strings.HasPrefix(cnt.ID, i) {
				if err := m.Destroy(cnt); err != nil {
					logger.Fatalf("error destroying container: %s\n", err)
				}
				fmt.Printf("destroyed %s\n", cnt.ID[:12])
			}
		}
	}
}
