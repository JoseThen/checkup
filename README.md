# checkup
Go Binary to run simple health checks with endpoints

---
## Example Usage :


``` bash
checkup listen -c 200 google.com
checkup listen --code 302 potatoe.com/farm
checkup exam -f list.json potatoe.com
checkup exam --file list.yml tomatoe.io
```

## Exam File Example :

### Yaml
``` yaml
name: Test Name
    - code: 200
      paths:
        - /farm
        - /something
        - /else
    - code: 404
      paths:
        - /this
        - /not
        - /found
```

``` json
{
    "name": "Exam Name",
    "test": [
        {
            "code": 200,
            "paths": [
                "/farm",
                "/something",
                "/else",
            ]
        },
        {
            "code": 404,
             "paths": [
                "/this",
                "/not",
                "/found",
            ]
        }
    ]
}
```