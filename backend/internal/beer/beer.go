package beer

type Beer struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Style    string  `json:"style"`
	Brewery  string  `json:"brewery"`
	FullName string  `json:"fullname"`
	ABV      float32 `json:"abv"`
	MinIBU   int16   `json:"min_ibu"`
	MaxIBU   int16   `json:"max_ibu"`
}
