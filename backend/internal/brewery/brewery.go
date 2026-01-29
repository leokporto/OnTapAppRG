package brewery

type Brewery struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
