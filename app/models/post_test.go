package models

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllPosts(t *testing.T) {
	selectQuery := "SELECT * FROM posts"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "content"}).
		AddRow(1, "Title 1", "Content 1").
		AddRow(2, "Title 2", "Content 2").
		AddRow(3, "Title 3", "Content 3")

	mock.ExpectQuery(selectQuery).WillReturnRows(rows)
}

func TestGetPostByID(t *testing.T) {
	selectQuery := "SELECT * FROM posts WHERE id = ?"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "content"}).
		AddRow(1, "Title 1", "Content 1")

	mock.ExpectQuery(selectQuery).WithArgs(1).WillReturnRows(rows)
}

func TestCreatePost(t *testing.T) {
	insertQuery := "INSERT INTO posts (title, content) VALUES (?, ?)"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(insertQuery).WithArgs("Title 1", "Content 1").WillReturnResult(sqlmock.NewResult(1, 1))
}

func TestUpdatePost(t *testing.T) {
	updateQuery := "UPDATE posts SET title = ?, content = ? WHERE id = ?"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(updateQuery).WithArgs("Title 1", "Content 1", 1).WillReturnResult(sqlmock.NewResult(1, 1))
}

func TestDeletePost(t *testing.T) {
	deleteQuery := "DELETE FROM posts WHERE id = ?"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(deleteQuery).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
}
