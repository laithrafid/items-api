package services

//2
import (
	"github.com/laithrafid/bookstore_items-api/src/domain/items"
	"github.com/laithrafid/bookstore_items-api/src/domain/queries"
	"github.com/laithrafid/bookstore_items-api/src/utils/errors_utils"
)

// 7
var (
	ItemsService itemsServiceInterface = &itemsService{}
)

// 3
type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, errors_utils.RestErr)
	Get(string) (*items.Item, errors_utils.RestErr)
	Search(queries.EsQuery) ([]items.Item, errors_utils.RestErr)
}

// 4
type itemsService struct{}

// 5
func (s *itemsService) Create(item items.Item) (*items.Item, errors_utils.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

// 6
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
