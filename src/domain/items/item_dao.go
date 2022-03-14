package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/laithrafid/items-api/src/clients/elasticsearch"
	"github.com/laithrafid/items-api/src/domain/queries"
	"github.com/laithrafid/utils-go/errors_utils"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

func (i *Item) Save() errors_utils.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return errors_utils.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() errors_utils.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, typeItem, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return errors_utils.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
		}
		return errors_utils.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database error"))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return errors_utils.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}

	if err := json.Unmarshal(bytes, &i); err != nil {
		return errors_utils.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}
	i.Id = itemId
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, errors_utils.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, errors_utils.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, errors_utils.NewInternalServerError("error when trying to parse response", errors.New("database error"))
		}
		item.Id = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, errors_utils.NewNotFoundError("no items found matching given criteria")
	}
	return items, nil
}
