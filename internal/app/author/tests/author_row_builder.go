//go:build unit
// +build unit

package tests

import (
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/author"
)

type AuthorBuilder struct {
	instance *author.AuthorRow
}

func Author() *AuthorBuilder {
	return &AuthorBuilder{instance: &author.AuthorRow{}}
}

func (b *AuthorBuilder) ID(v int64) *AuthorBuilder {
	b.instance.ID = v
	return b
}
func (b *AuthorBuilder) Name(v string) *AuthorBuilder {
	b.instance.Name = v
	return b
}

func (b *AuthorBuilder) P() *author.AuthorRow {
	return b.instance
}

func (b *AuthorBuilder) V() author.AuthorRow {
	return *b.instance
}

func (b *AuthorBuilder) Valid() *AuthorBuilder {
	return Author().ID(Author1ID).Name(AuthorName1)
}
