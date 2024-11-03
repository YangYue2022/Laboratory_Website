package v1

import (
	"UISwebsite_backend/models"
	"UISwebsite_backend/pkg/e"
	"UISwebsite_backend/pkg/logging"
	"UISwebsite_backend/pkg/setting"
	"UISwebsite_backend/pkg/util"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

// GetQas 获取问答列表
// @Summary 获取问答列表
// @Description 根据问题关键字模糊查询问答列表
// @Tags 问答
// @Accept json
// @Produce json
// @Param que query string false "问题关键字（模糊查询）"
// @Param page query int false "页码"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data: {lists: 问答列表, total: 总数}"
// @Router /api/qas [get]
func GetQas(c *gin.Context) {
	que := c.Query("que")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if que != "" {
		maps["`que` LIKE ?"] = "%" + que + "%"
	}

	code := e.SUCCESS

	data["lists"] = models.GetQas(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetQasTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddQa 添加问答
// @Summary 添加问答
// @Description 添加一个新的问答记录，包括问题和答案
// @Tags 问答
// @Accept json
// @Produce json
// @Param qa body models.Qa true "问答详细信息"
// @Param token query string true "token"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data: {item:新增的qa}"
// @Failure 200 {object} map[string]interface{} "code: 400, msg: 'INVALID_PARAMS', data: {}"
// @Router /api/admin/qa [post]
func AddQa(c *gin.Context) {
	req := &models.Qa{} // 初始化 req 为一个指向新的 models.Qa 的指针

	code := e.INVALID_PARAMS

	if err := c.ShouldBindJSON(req); err != nil {
		logging.Error(fmt.Sprintf("Error binding JSON: %v", err))
		log.Println("Error binding JSON: %s", err)
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]interface{}),
		})
		return
	}

	valid := validation.Validation{}
	valid.Required(req.Que, "que").Message("问题不能为空")
	valid.Required(req.Ans, "ans").Message("回答不能为空")

	data := make(map[string]interface{})
	if !valid.HasErrors() {
		code = e.SUCCESS
		data["item"] = models.AddQa(req)
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			logging.Error(err)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// EditQa 编辑问答
// @Summary 编辑问答
// @Description 根据ID编辑问答记录，更新问题和答案
// @Tags 问答
// @Accept json
// @Produce json
// @Param id path int true "问答ID"
// @Param qa body models.Qa true "问答详细信息"
// @Param token query string true "token"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data: {item:想要编辑的问答条目}"
// @Failure 200 {object} map[string]interface{} "code: 400, msg: 'INVALID_PARAMS', data: {}"
// @Failure 200 {object} map[string]interface{} "code: 10003, msg: 'ERROR_NOT_EXIST_QA', data: {item:想要编辑的问答条目}"
// @Router /api/admin/qa/{id} [put]
func EditQa(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	req := &models.Qa{} // 初始化 req 为一个指向新的 models.Qa 的指针

	code := e.INVALID_PARAMS
	data := make(map[string]interface{})

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding JSON: %s", err)
		logging.Error(fmt.Sprintf("Error binding JSON: %v", err))
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]string),
		})
		return
	}

	valid := validation.Validation{}
	valid.Required(req.Que, "que").Message("问题不能为空")
	valid.Required(req.Ans, "ans").Message("回答不能为空")

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistQaByID(id) {
			data["item"] = models.EditQa(id, req)
		} else {
			code = e.ERROR_NOT_EXIST_QA
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			logging.Error(err)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// DeleteQa 删除问答
// @Summary 删除问答
// @Description 根据ID删除指定的问答记录
// @Tags 问答
// @Accept json
// @Produce json
// @Param id path int true "问答ID"
// @Param token query string true "token"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data: {}"
// @Failure 200 {object} map[string]interface{} "code: 10003, msg: 'ERROR_NOT_EXIST_QA ', data: {}"
// @Router /api/admin/qa/{id} [delete]
func DeleteQa(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_PARAMS

	if models.ExistQaByID(id) {
		code = e.SUCCESS
		models.DeleteQa(id)
	} else {
		code = e.ERROR_NOT_EXIST_QA
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
