package services

import (
	"github.com/laithrafid/bookstore_items-api/src/domain/items"
	"github.com/laithrafid/bookstore_items-api/src/domain/queries"
	"github.com/laithrafid/bookstore_items-api/src/utils/errors_utils"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, errors_utils.RestErr)
	Get(string) (*items.Item, errors_utils.RestErr)
	Search(queries.EsQuery) ([]items.Item, errors_utils.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, errors_utils.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, errors_utils.RestErr) {
	item := items.Item{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, errors_utils.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}
