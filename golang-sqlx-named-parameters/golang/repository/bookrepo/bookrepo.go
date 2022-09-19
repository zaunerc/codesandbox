// file golang-sqlx-named-parameters/golang/repository/bookrepo/bookrepo.go
package bookrepo

import (
	"github.com/jmoiron/sqlx"
	"main/model"
	"strings"
)

func buildCreateQuery() string {

	const createQuery = `
INSERT INTO books
(
 	title
)
VALUES
(
	:title
)
RETURNING *
`
	return createQuery
}

func Create(db *sqlx.DB, book *model.Book) (*model.Book, error) {

	rows, err := db.NamedQuery(buildCreateQuery(), book)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, rows.Err()
	}

	createdBook := &model.Book{}
	err = rows.StructScan(createdBook)
	return createdBook, err

}

func buildSearchQuery(searchParams *SearchParams) string {
	sb := strings.Builder{}
	sb.WriteString(`SELECT * FROM books WHERE (1 = 1) `)

	if searchParams.Title != nil {
		sb.WriteString(`AND (title = :title)`)
	}

	return sb.String()
}

func Search(db *sqlx.DB, searchParams *SearchParams) ([]model.Book, error) {
	var books []model.Book
	searchQuery := buildSearchQuery(searchParams)
	rows, err := db.NamedQuery(searchQuery, *searchParams)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book
		err = rows.StructScan(&book)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}

	return books, rows.Err()
}
