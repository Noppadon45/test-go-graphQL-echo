package handler

import "test-graphQL/graph/model"

var books = []*model.Book {
	{
		ID: "Book01",
		Name: "Apple",
		Type: &model.AllBookType[0],
		Publisher: []*model.Publisher{
			{
				ID: "001",
				Name: "CompanyA",
			},
			{
				ID: "002",
				Name: "StoreA",
			},
		},
	},
	{
		ID: "Book02",
		Name: "Solid",
		Type: &model.AllBookType[1],
		Publisher: []*model.Publisher{
			{
				ID: "003",
				Name: "CompanyB",
			},
			{
				ID: "004",
				Name: "StoreB",
			},
		},
	},
	{
		ID: "Book03",
		Name: "RealNumber",
		Type: &model.AllBookType[2],
		Publisher: []*model.Publisher{
			{
				ID: "001",
				Name: "CompanyA",
			},
			{
				ID: "004",
				Name: "StoreB",
			},
		},
	},
}

func GetBook() []*model.Book {
	return books
}

func CreateBook(req *model.NewBook) *model.Book {
	newBook := &model.Book{
		ID: "005",
		Name: req.Name,
		Type: req.Type,
		Publisher: []*model.Publisher{
			{
				ID: "002",
				Name: "StoreA",
			},
			{
				ID: "003",
				Name: "CompanyB",
			},
		},
	}
	books = append(books, newBook)
	return newBook
}