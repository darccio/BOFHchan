package post

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type SQLRepository struct {
	DB *sqlx.DB
}

func (r *SQLRepository) Init() error {
	if _, err := r.DB.Exec("CREATE TABLE posts (id TEXT PRIMARY KEY, title TEXT, url TEXT NULL, description TEXT NULL, creator_id TEXT)"); err != nil {
		return err
	}
	return nil
}

func (r *SQLRepository) Save(p *Post) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	_, err = tx.NamedExec("INSERT INTO posts (id, title, url, description, creator_id) VALUES (:id, :title, :url, :description, :creator_id) ON CONFLICT(id) DO UPDATE SET title=:title, url=:url, description=:description, creator_id=:creator_id", p)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *SQLRepository) Find(id string) (*Post, error) {
	var p Post
	if err := r.DB.Get(&p, "SELECT * FROM posts WHERE id = ?", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *SQLRepository) Delete(id string) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (r *SQLRepository) Close() error {
	return r.DB.Close()
}
