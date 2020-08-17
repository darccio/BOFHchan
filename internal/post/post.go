package post

type Post struct {
	Id          string `db:"id"`
	Title       string `db:"title"`
	Url         string `db:"url"`
	Description string `db:"description"`
	Creator     string `db:"creator_id"`
}
