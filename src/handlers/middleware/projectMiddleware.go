package middleware

import (
	"DidlyDoodash-api/src/db/daos"
	"DidlyDoodash-api/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProjectMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		orgID := c.Param("id")
		org, err := daos.GetOrg(orgID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.OrgNotFound)
			return
		}
		if org == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.OrgNotFound)
			return
		}

		c.Set("organisation", org)

		c.Next()
	}
}
