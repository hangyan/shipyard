package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/hangyan/shipyard/client"
)

var scaleCommand = cli.Command{
	Name:   "scale",
	Usage:  "scale a container",
	Action: scaleAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "id",
			Value: "",
			Usage: "container id",
		},
		cli.StringFlag{
			Name:  "count",
			Value: "",
			Usage: "total number of instances for container",
		},
	},
}

func scaleAction(c *cli.Context) {
	cfg, err := loadConfig()
	if err != nil {
		logger.Fatal(err)
	}
	m := client.NewManager(cfg)
	containerId := c.String("id")
	if containerId == "" {
		logger.Fatalf("you must specify a container id")
	}

	count := c.Int("count")
	container, err := m.Container(containerId, false)
	if err != nil {
		logger.Fatalf("error getting container info: %s", err)
	}

	if container == nil {
		logger.Fatalf("no container with id [%s] running", containerId)
	}

	if err := m.Scale(container, count); err != nil {
		logger.Fatalf("error scaling container: %s\n", err)
	}
	fmt.Printf("scaled %s to %d\n", container.ID[:12], count)
}
