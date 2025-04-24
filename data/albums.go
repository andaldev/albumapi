package data

type Album struct {
	ID           int
	TITLE        string
	BAND         string
	RELEASE_YEAR int
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
