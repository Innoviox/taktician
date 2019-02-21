package annotatedb

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/google/subcommands"
	"github.com/nelhage/taktician/ai"
	"github.com/nelhage/taktician/logs"
	"github.com/nelhage/taktician/ptn"
	"golang.org/x/sync/errgroup"
)

type Command struct {
	minPly, maxPly int
	workers        int
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
	flags.IntVar(&c.workers, "workers", 2, "parallel workers")
}

const (
	maxBatch       = 50
	commitInterval = 10000
)

type batch []annotation

func (c *Command) annotateRows(ctx context.Context, todo <-chan todoRow, out chan<- batch) error {
	for row := range todo {
		var batch batch
		game, err := ptn.ParsePTN(strings.NewReader(row.PTN))
		if err != nil {
			return err
		}
		it := game.Iterator()

		init, err := game.InitialPosition()
		if err != nil {
			return err
		}

		cfg := ai.MinimaxConfig{
			Size: init.Size(),
		}
		engine := ai.NewMinimax(cfg)
		for it.Next() {
			p := it.Position()
			if o, _ := p.GameOver(); o {
				break
			}
			for d := c.minPly; d <= c.maxPly; d++ {
				pv, v, _ := engine.AnalyzeDepth(ctx, d, p)
				batch = append(batch, annotation{
					Game:     row.Id,
					Ply:      p.MoveNumber(),
					Depth:    d,
					Analysis: v,
					Move:     ptn.FormatMove(pv[0]),
				})
			}
			if len(batch) > maxBatch {
				out <- batch
				batch = nil
			}
		}
		out <- batch
		log.Printf("game=%d moves=%d", row.Id, it.Position().MoveNumber())
	}
	return nil
}

func (c *Command) Execute(ctx context.Context, flag *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	repo, err := logs.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("open(%q): %s", flag.Arg(0), err.Error())
	}

	db := repo.DB()

	for _, sql := range createSchema {
		_, err := db.Exec(sql)
		if err != nil {
			log.Fatalf("create schema: %v", err)
		}
	}

	todo := make(chan todoRow)
	results := make(chan batch)

	go func() {
		defer close(todo)
		rows, err := db.Queryx(selectTODO)
		if err != nil {
			log.Fatalf("query: %s", err.Error())
		}
		for rows.Next() {
			var row todoRow
			err := rows.StructScan(&row)
			if err != nil {
				log.Fatalf("scan: %s", err.Error())
			}
			todo <- row
		}
	}()

	var grp errgroup.Group
	for j := 0; j < c.workers; j++ {
		grp.Go(func() error {
			return c.annotateRows(ctx, todo, results)
		})
	}
	go func() {
		if err := grp.Wait(); err != nil {
			log.Fatalf("parse: %v", err)
		}
		close(results)
	}()

	var tx *sql.Tx
	n := 0

	for batch := range results {
		if len(batch) == 0 {
			continue
		}
		if tx == nil {
			tx, err = db.Begin()
			if err != nil {
				log.Fatalf("begin: %v", err)
			}
		}

		vals := make([]interface{}, 5*len(batch))
		i := 0
		for _, a := range batch {
			vals[i] = a.Game
			i++
			vals[i] = a.Ply
			i++
			vals[i] = a.Depth
			i++
			vals[i] = a.Analysis
			i++
			vals[i] = a.Move
			i++
		}

		placeholders := strings.Repeat("(?, ?, ?, ?, ?), ", len(batch)-1)
		if _, err := db.Exec(
			fmt.Sprintf("INSERT OR REPLACE INTO annotations VALUES %s (?, ?, ?, ?, ?)", placeholders),
			vals...); err != nil {
			log.Fatalf("INSERT: %v", err)
		}
		n += len(batch)

		if n > commitInterval {
			if err := tx.Commit(); err != nil {
				log.Fatalf("commit: %v", err)
			}
			tx = nil
		}
	}

	return subcommands.ExitSuccess
}
