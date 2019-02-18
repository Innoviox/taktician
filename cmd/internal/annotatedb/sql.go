package annotatedb

var createSchema []string = []string{
	`
CREATE TABLE IF NOT EXISTS annotations (
  game integer,
  ply integer,
  depth integer,
  analysis integer,
  move string
)
`,
	`
CREATE UNIQUE INDEX IF NOT EXISTS annotations_pkey
ON annotations (game, ply, depth)
`}

type annotation struct {
	Game     int    `db:"game"`
	Ply      int    `db:"ply"`
	Depth    int    `db:"depth"`
	Analysis uint64 `db:"analysis"`
	Move     string `db:"move"`
}

/*
const selectTODO = `
WITH RECURSIVE iota AS (
 SELECT 1 AS n
 UNION
 SELECT n+1 FROM iota
   WHERE n < 100
)
 select * from iota;
`
*/

const selectTODO = `
SELECT id, ptn
FROM ptns
ORDER BY id ASC
`
