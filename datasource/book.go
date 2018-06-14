package datasource

import "github.com/vladmashtak/restapi/datamodel"

// ToDo implement DB connection
var Books []datamodel.Book = []datamodel.Book{
	{"1", "234234", "Book One", &datamodel.Author{"Vlad", "Mashtak"}},
	{"2", "234677", "Book Two", &datamodel.Author{"Dima", "Drav4enko"}},
}