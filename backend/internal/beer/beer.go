package beer

type Beer struct {
	ID        int64   `db:"id"`
	Name      string  `db:"name"`
	FullName  string  `db:"fullname"`
	ABV       float32 `db:"abv"`
	MinIBU    int16   `db:"minibu"`
	MaxIBU    int16   `db:"maxibu"`
	StyleID   int64   `db:"styleid"`
	BreweryID int64   `db:"breweryid"`
}
