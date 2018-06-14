package repositories

import "github.com/vladmashtak/restapi/datamodels"

type Query func(book datamodels.Book)

type BookRepository interface {
	Select(query Query) (book datamodels.Book, ok bool)
	SelectAll() 		(result []datamodels.Book)
	Update(query Query) (updatedBook datamodels.Book, err error)
	Delete(query Query) (deleted bool)
}

var bookRepository *BookRepository = nil

/*func GetBookRepository(source []datamodels.Book) BookRepository {
	if bookRepository == nil {
		bookRepository = &BookRepository{}
	}
}*/
