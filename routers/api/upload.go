package api

import (
	"UISwebsite_backend/pkg/e"
	"UISwebsite_backend/pkg/logging"
	"UISwebsite_backend/pkg/upload"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadImage 上传图片
// @Summary 上传图片
// @Description 上传图片接口
// @Tags 图片
// @Accept multipart/form-data
// @Produce application/json
// @Param image formData file true "图片文件"
// @Success 200 {object} map[string]interface{}  "code": 200, "msg": “ok”, "data": {image_url:图片访问地址 image_save_url:图片保存地址}
// @Failure 200 {object} map[string]interface{} "code": 30001, "msg": "保存图片失败", "data": {} "保存图片失败"
// @Failure 200 {object} map[string]interface{} "code": 30002, "msg": "检查图片失败", "data": {} "检查图片失败"
//
//	@Failure 200 {object} map[string]interface{} "code": 30003, "msg": "校验图片错误，图片格式或大小有问题", "data": {}  "校验图片错误，图片格式或大小有问题"
//
// @Failure 200 {object} map[string]interface{} "code": 400, "msg": "请求参数错误", "data": {} "请求参数错误，图片为空"
// @Router /api/v1/upload [post]
func UploadImage(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}

	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
