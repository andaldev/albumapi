package data

type Album struct {
	ID           int    `json:"id"`
	TITLE        string `json:"title"`
	BAND         string `json:"band"`
	RELEASE_YEAR int    `json:"release_year"`
}

var albums []Album

func GetAlbum() []Album {
	albums = append(
		albums,
		Album{ID: 1, TITLE: "Sky Void of Stars", BAND: "Katatonia", RELEASE_YEAR: 2023},
		Album{ID: 2, TITLE: "Ghost Reveries", BAND: "Opeth", RELEASE_YEAR: 2005},
	)
	return albums
}
