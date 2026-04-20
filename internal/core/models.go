package site

type Post struct {
	Title       string
	Slug        string
	Date        string
	PrimaryTag  string
	Description string
	HTML        string
	Tags        []string
}

type Project struct {
	Title       string
	Slug        string
	Description string
	HTML        string
}
