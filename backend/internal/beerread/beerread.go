package beerread

type BeerDTO struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	FullName string  `json:"fullname"`
	ABV      float32 `json:"abv"`
	MinIBU   int16   `json:"minibu"`
	MaxIBU   int16   `json:"maxibu"`
	Style    string  `json:"style"`
	Brewery  string  `json:"brewery"`
}
