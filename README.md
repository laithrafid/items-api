![ItemsAPI-docker-Image](https://github.com/laithrafid/items-api/actions/workflows/main.yml/badge.svg?branch=main)


#ITEMSAPI service written in go 
### development follows MVC Pattern


## Development Pattern
![alt text](https://github.com/laithrafid/infra-api/blob/main/Images/devpattern.png?raw=true)

## Archietecture 
![alt text](https://github.com/laithrafid/infra-api/blob/main/Images/items.png?raw=true)



## ITEMS API Endpoints and calls 
    GET /ping
	POST /items
	GET /items/{id}
	POST /items/search

## creating index using postman request or curl
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
