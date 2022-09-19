// file golang-sqlx-named-parameters/golang/repository/bookrepo/searchParams.go
package bookrepo

type SearchParams struct {
	Title *string `db:"title"`
}
