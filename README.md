<h1 align="center">Checkup</h1>
<p>CLI to run simple health checks against endpoints</p>
<p>
  <a href="https://opensource.org/licenses/MIT">
    <img alt="License: MIT" src="https://img.shields.io/github/license/JoseThen/checkup" target="_blank" />
  </a>
</p>

<p align="center">
  <img align="center" width="160px" src="./assets/gopher.png">
</p>

---
## Example Usage :


``` bash
checkup listen -e https://google.com # default code is 200
checkup listen -c 200 -e https://google.com
checkup listen --code 302 --endpoint https://potatoe.com/farm
checkup exam -f list.json
checkup exam --file list.yml
CU_USER=admin CU_PASS=pass go run main.go listen -e http://localhost:8080 -a # with basic auth
```

## Exam File Example :

### Yaml
``` yaml
name: Test Name
endpoint: https://duckduckgo.com
tests:
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
  "tests": [
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