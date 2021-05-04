package videos_domain

type VideoFromUser struct {
	Name      string
	Category  string
	EmbedID   string
	ID        string
	DateTime  string
	Thumbnail string
	Title     string
}

type VideoToDB struct {
	Name      string
	Category  string
	DateTime  string
	EmbedID   string
	Thumbnail string
	Title     string
}
