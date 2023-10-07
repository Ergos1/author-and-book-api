package models

type Book struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Rating   int64  `db:"rating"`
	AuthorID int64  `db:"author_id"`
}
