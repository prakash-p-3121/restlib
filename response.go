package restlib

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OkResponse(c *gin.Context, respContent interface{}) {
	c.JSON(http.StatusOK, respContent)
}

func BadReqResponse(c *gin.Context, respContent interface{}) {
	c.JSON(http.StatusBadRequest, respContent)
}

func NotFoundResponse(c *gin.Context, respContent interface{}) {
	c.JSON(http.StatusNotFound, respContent)
}

func ConflictResponse(c *gin.Context, respContent interface{}) {
	c.JSON(http.StatusConflict, respContent)
}

func InternalServerErrorResponse(c *gin.Context, respContent interface{}) {
	c.JSON(http.StatusInternalServerError, respContent)
}

func OkNoContentResponse(c *gin.Context, respContent interface{}) {
	c.JSON(http.StatusNoContent, respContent)
}
