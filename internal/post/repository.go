package post

type Repository interface {
	Save(p *Post) error
	Find(id string) (*Post, error)
	Delete(id string) error
}
