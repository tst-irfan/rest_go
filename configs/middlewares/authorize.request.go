package middlewares

import (
	. "auth_go/app/utils/token"

	"github.com/gin-gonic/gin"
)

func AuthorizeRequest(c *gin.Context) {
	err := TokenValid(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	id, err := ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.Set("UserID", id)
}
