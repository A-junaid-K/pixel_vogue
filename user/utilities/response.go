package utilities

import "github.com/gin-gonic/gin"

type Response struct {
	StatusCode int         `json:"statuscode"`
	Message    string      `json:"message"`
	Data       any         `json:"data"`
	Error      interface{} `json:"error"`
}

func ResponseResult(c *gin.Context, response Response) {
	c.JSON(response.StatusCode, response)
}
