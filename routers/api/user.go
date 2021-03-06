package api

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"gin-web/models"
)

// @Router /api/v1/user [get]
func GetUser(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	var user = models.User{}
	result := user.GetUser(id)
	c.JSON(200, result)
}
