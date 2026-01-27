package beerstyle

type BeerStyle struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
