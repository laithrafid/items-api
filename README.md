
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