package author

type AuthorRow struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
