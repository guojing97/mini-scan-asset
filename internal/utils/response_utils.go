package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequestResponse(c *gin.Context) {
	c.String(http.StatusBadRequest, "BAD REQUEST")
	c.Abort()
}
