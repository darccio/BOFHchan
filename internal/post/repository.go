package post

type Repository interface {
	Init() error
	Save(p *Post) error
	Find(id string) (*Post, error)
	Delete(id string) error
}
