package beer

type Beer struct {
	ID       int64   `json:"id" db:"id"`
	Name     string  `json:"name" db:"name"`
	Style    string  `json:"style" db:"style"`
	Brewery  string  `json:"brewery" db:"brewery"`
	FullName string  `json:"fullname" db:"fullname"`
	ABV      float32 `json:"abv" db:"abv"`
	MinIBU   int16   `json:"minibu" db:"minibu"`
	MaxIBU   int16   `json:"maxibu" db:"maxibu"`
}
