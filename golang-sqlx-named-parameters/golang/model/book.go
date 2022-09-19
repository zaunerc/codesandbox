// file golang-sqlx-named-parameters/golang/model/book.go
package model

type Book struct {
	ID    string `db:"id"`
	Title string `db:"title"`
}
