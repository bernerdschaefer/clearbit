package main

import (
	"github.com/codegangsta/cli"
	"github.com/thoughtbot/clearbit"
)

var prospectCommand = cli.Command{
	Name:      "prospect",
	Usage:     "fetch contacts for a company",
	ArgsUsage: "domain",
	Action:    prospect,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "name", Usage: "Filter by first or last name (case-insensitive)"},
		cli.StringSliceFlag{Name: "title", Usage: "Filter by job title"},
	},
}

func prospect(ctx *cli.Context) error {
	var (
		apiKey = apiKeyFromContext(ctx)
		client = clearbit.NewClient(apiKey, nil)

		domain = ctx.Args().First()
		name   = ctx.String("name")
		titles = ctx.StringSlice("title")
	)

	if domain == "" {
		return requiredArgError("Usage: clearbit prospect <domain>")
	}

	prospects, err := client.Prospect(clearbit.ProspectQuery{
		Domain: domain,
		Name:   name,
		Titles: titles,
	})
	if err != nil {
		return exitError(err)
	}

	display(prospects)
	return nil
}
