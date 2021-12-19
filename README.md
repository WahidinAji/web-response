# api-response
this package is for api response ^^

* Created by Wahidin
* Open for public
* Enjoy your day ^-^

## how to install?
```bash
go get github.com/WahidinAji/web-response
```
## how to use?
```go
package <your-package-module>

import (
  alias "github.com/WahidinAji/web-response"
)

func Handle (context) {
  //do your code here
  return alias.WebResponse(200, "OK", data)
}
```

## note
* the response must JSON, example in fiber
```go
func (ctx *fiber.Ctx) error {
  //code
  return ctx.JSON(
     alias.WebResponse(fiber.StatusOk, "OK", data)
  )
}
```
* the response must JSON, example in echo
```go
func (ctx echo.Context) error {
  //code
  return ctx.JSON(http.StatusOK, response.WebResponse(http.StatusOK,"OK", data))
}
```
