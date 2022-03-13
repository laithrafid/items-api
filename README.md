![ItemsAPI-docker-Image](https://github.com/laithrafid/bookstore_items-api/actions/workflows/main.yml/badge.svg?branch=main)


1. creating index postman request
```
Put 127.0.0.1:9200/items
{
    "settings":{
        "index":{
        "number_of_shards": 4,
        "number_of_replicas": 2
        }
    }
}
```
