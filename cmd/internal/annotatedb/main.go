package annotatedb

import (
	"context"
	"flag"
	"log"

	"github.com/google/subcommands"
	"github.com/nelhage/taktician/logs"
)

type Command struct {
	minPly, maxPly int
}

func (*Command) Name() string     { return "annotatedb" }
func (*Command) Synopsis() string { return "Annotate the playtak DB with analysis at every move" }
func (*Command) Usage() string {
	return `annotatedb [options] game.db
`
}

func (c *Command) SetFlags(flags *flag.FlagSet) {
	flags.IntVar(&c.minPly, "min-ply", 3, "minimum ply to analyze")
	flags.IntVar(&c.maxPly, "max-ply", 5, "maximum ply to analyze")
}

func (c *Command) Execute(ctx context.Context, flag *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	repo, err := logs.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("open(%q): %s", flag.Arg(0), err.Error())
	}

	db := repo.DB()

	rows, err := db.Queryx(selectTODO)
	if err != nil {
		log.Fatalf("query: %s", err.Error())
	}

	type ptn struct {
		Id  int    `db:"id"`
		Ptn string `db:"ptn"`
	}
	for rows.Next() {
		var ptn ptn
		err := rows.StructScan(&ptn)
		if err != nil {
			log.Fatalf("scan: %s", err.Error())
		}
	}

	return subcommands.ExitSuccess
}
