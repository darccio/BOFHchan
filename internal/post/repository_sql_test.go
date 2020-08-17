package post

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func TestSQLRepository(t *testing.T) {
	id := "hello-world"
	r, err := initRepository()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	// Test 1: don't find the not-yet-created post.
	result, err := r.Find(id)
	if err != nil {
		t.Fatal(err)
	}
	if result != nil {
		t.Fatalf("expected nil, got %+v", result)
	}
	// Test 2: save and find the post.
	p := &Post{
		Id:          id,
		Title:       "Hello World",
		Description: "Ef mi primera puflifafión",
		Creator:     "tifu",
	}
	if err = r.Save(p); err != nil {
		t.Fatal(err)
	}
	result, err = r.Find(id)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatalf("expected %+v, got nil", result)
	}
	// Test 3: update the post.
	p.Title = "Hola mundo"
	if err = r.Save(p); err != nil {
		t.Fatal(err)
	}
	result, err = r.Find(id)
	if err != nil {
		t.Fatal(err)
	}
	if p.Title != result.Title {
		t.Fatalf("expected %q, got %q", p.Title, result.Title)
	}
	// Test 4: delete the post and check it doesn't exist.
	if err = r.Delete(id); err != nil {
		t.Fatal(err)
	}
	result, err = r.Find(id)
	if err != nil {
		t.Fatal(err)
	}
	if result != nil {
		t.Fatalf("expected nil, got %+v", result)
	}
}

func initRepository() (*SQLRepository, error) {
	db, err := sqlx.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return nil, err
	}
	r := SQLRepository{
		DB: db,
	}
	if err = r.Init(); err != nil {
		return nil, err
	}
	return &r, nil
}
