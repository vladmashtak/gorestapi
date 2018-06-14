package datasource

import "github.com/vladmashtak/restapi/datamodels"

// ToDo implement DB connection
var Books []datamodels.Book = []datamodels.Book{
	{"1", "234234", "Book One", &datamodels.Author{"Vlad", "Mashtak"}},
	{"2", "234677", "Book Two", &datamodels.Author{"Dima", "Drav4enko"}},
}