package pokeapi

type LocationAreaResponse struct {
	Count			int			`json:"count"`
	Next			*string		`json:"next"`
	Previous		*string		`json:"previous"`
	Results			[]LocationAreaResponseResults	`json:"results"`

}

type LocationAreaResponseResults struct {
	Name		string		`json:"name"`
	Url			string		`json:"url"`
}
