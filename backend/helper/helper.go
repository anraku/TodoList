package helper

import "github.com/labstack/echo"

type ResponseMap map[string]interface{}
type Responsible interface {
	ToResponseMap() ResponseMap
}

func JSONResponseObject(c echo.Context, statusCode int, object Responsible) error {
	body := object.ToResponseMap()
	return JSONResponse(c, statusCode, body)
}

func JSONResponse(c echo.Context, statusCode int, body ResponseMap) error {
	if statusCode != 200 {
		body["success"] = false
	} else {
		body["success"] = true
	}
	return c.JSON(statusCode, body)
}

func JSONResponseArray(c echo.Context, statusCode int, collection []ResponseMap) error {
	body := ResponseMap{}
	body["items"] = collection
	body["count"] = len(collection)

	return JSONResponse(c, statusCode, body)
}
