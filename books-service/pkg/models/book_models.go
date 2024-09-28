package models

import "database/sql"

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"published_year"`
	Category      string `json:"category"`
	Available     bool   `json:"available"`
	CreatedAt     string `json:"created_at"`
}

type BookService struct {
	DB *sql.DB
}

func (bs *BookService) GetBooks() ([]Book, error) {
	var books []Book
	var book Book

	rows, err := bs.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedYear, &book.Category, &book.Available, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

type AddBookParam struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"published_year"`
	Category      string `json:"category"`
	Available     bool   `json:"available"`
}

func (bs *BookService) AddBook(b AddBookParam) (Book, error) {
	var book Book
	query := "INSERT INTO books (b.title, b.author, b.published_year, b.category, b.available) VALUES ($1, $2, $3, $4, $5) RETURNING *"
	row := bs.DB.QueryRow(query, b.Title, b.Author, b.PublishedYear, b.Category, b.Available)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedYear, &book.Category, &book.Available, &book.CreatedAt)
	if err != nil {
		return Book{}, err
	}
	return book, nil
}

func (bs *BookService) UpdateBook(id int, b AddBookParam) (Book, error) {
	var book Book
	query := "UPDATE books SET title=$1, author=$2, published_year=$3, category=$4, available=$5 WHERE id=$6 RETURNING *"
	row := bs.DB.QueryRow(query, b.Title, b.Author, b.PublishedYear, b.Category, b.Available, id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedYear, &book.Category, &book.Available, &book.CreatedAt)
	if err != nil {
		return Book{}, err
	}
	return book, nil
}

func (bs *BookService) DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id=$1"
	_, err := bs.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
