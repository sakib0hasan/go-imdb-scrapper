package main

type APIResponse struct {
	V int    `json:"v"`
	Q string `json:"q"`
	D []struct {
		L  string        `json:"l"`
		ID string        `json:"id"`
		S  string        `json:"s"`
		Y  int           `json:"y"`
		Q  string        `json:"q"`
		I  []interface{} `json:"i"`
	} `json:"d"`
}

type IMDB struct {
	Name      string
	ImdbId    string
	Stars     string
	Year      int
	Type      string
	PosterURL string
}
