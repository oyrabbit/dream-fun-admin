package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/oyrabbit/dream-fun-admin/model"
	"github.com/oyrabbit/dream-fun-admin/utils/errmsg"
	"net/http"
)

// UpLoad 上传图片接口
func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	fileSize := fileHeader.Size

	url, code := model.UpLoadFile(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
