package services

import (
	"errors"
	"github.com/pskclub/play-with-go/models"
)

type BookService struct {
}

func (BookService) FindAll() ([]models.Book) {
	books := []models.Book{
		models.Book{
			ID:   "1111",
			Name: "eiuei",
		},
		models.Book{
			ID:   "222",
			Name: "eiuei",
		},
	}

	return books
}
func (BookService) Find(id string) (*models.Book, error) {
	if id == "0" {
		return nil, errors.New("not found")
	}
	book := models.Book{
		ID:   id,
		Name: "เทพเจ้าสายฟ้า",
	}

	return &book, nil
}
