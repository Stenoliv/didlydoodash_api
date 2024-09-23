package middleware

import (
	"DidlyDoodash-api/src/data"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			output := data.APIErrors{}
			for _, e := range c.Errors {
				var err data.APIErrorMeta
				byteSlice, _ := json.Marshal(e.Meta) // Convert e.Meta to []byte
				json.Unmarshal(byteSlice, &err)      // Unmarshal into the struct
				output.NewAPIError(err.Code, err.Title, e.Error())
			}
			c.JSON(output.Status(), output)
		}
	}
}
