package post

import "database/sql"

type SQLRepository struct {
	DB *sql.DB
}

func (r *SQLRepository) Save(p *Post) error {
	return nil
}

func (r *SQLRepository) Find(id string) (*Post, error) {
	var p Post
	return &p, nil
}

func (r *SQLRepository) Delete(id string) error {
	return nil
}
