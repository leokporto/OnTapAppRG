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

type BeerResponse struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	FullName string  `json:"fullname"`
	ABV      float32 `json:"abv"`
	MinIBU   int16   `json:"minibu"`
	MaxIBU   int16   `json:"maxibu"`
	Style    string  `json:"style"`
	Brewery  string  `json:"brewery"`
}

type BeerStyle struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type Brewery struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
