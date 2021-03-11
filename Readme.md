# Ahmad Rifai - Backend Developer

## 2. Build Microservice

### get movie detail

```
curl --location --request GET 'localhost:8080/api/imdb/detail/tt4853102'
```

### get movies list

## 4. Anagram

```
curl --location --request POST 'localhost:8080/api/anagram' \
--header 'Content-Type: application/json' \
--data-raw '{
    "data": [
        "kita",
        "atik",
        "tika",
        "aku",
        "kia",
        "makan",
        "kua"
    ]
}'
```

and response

```
{
    "status_code": 200,
    "message": null,
    "data": {
        "request": [
            "kita",
            "atik",
            "tika",
            "aku",
            "kia",
            "makan",
            "kua"
        ],
        "result": [
            [
                "kita",
                "atik",
                "tika"
            ],
            [
                "aku",
                "kua"
            ],
            [
                "kia"
            ],
            [
                "makan"
            ]
        ]
    }
}
```
