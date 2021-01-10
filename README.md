# checkup
CLI to run simple health checks against endpoints 

---
## Example Usage :


``` bash
checkup listen https://google.com
checkup listen -c 200 https://google.com
checkup listen --code 302 https://potatoe.com/farm
checkup exam -f list.json https://potatoe.com
checkup exam --file list.yml https://tomatoe.io
```

## Exam File Example :

### Yaml
``` yaml
name: Test Name
endpoint: https://duckduckgo.com
test:
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
  "endpoint": "https://google.com",
  "test": [
    {
      "code": 200,
      "paths": [
          "/farm",
          "/something",
          "/else"
      ]
    },
    {
      "code": 404,
        "paths": [
          "/this",
          "/not",
          "/found"
      ]
    }
  ]
}
```