package v1

import (
	"UISwebsite_backend/pkg/logging"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"

	"UISwebsite_backend/models"
	"UISwebsite_backend/pkg/e"
	"UISwebsite_backend/pkg/setting"
	"UISwebsite_backend/pkg/util"
)

// GetWorks
// @Summary 获取成果列表
// @Description 根据名称和类型进行筛选，获取成果列表
// @Tags 成果
// @Accept json
// @Produce json
// @Param name query string false "成果名称（模糊匹配）"
// @Param type query string false "成果类型"
// @Param page query int false "页码"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data: {lists: 成果列表, total: 总数}"
// @Router /api/works [get]
func GetWorks(c *gin.Context) {
	desc := c.Query("desc")
	type_ := c.Query("type")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if desc != "" {
		// 使用 LIKE 进行模糊匹配
		maps["`desc` LIKE ?"] = "%" + desc + "%"
	}

	if type_ != "" {
		maps["type"] = type_
	}

	code := e.SUCCESS

	data["lists"] = models.GetWorks(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetWorksTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// AddWork
// @Summary 新增成果
// @Description 添加一个新的成果记录
// @Tags 成果
// @Accept json
// @Produce json
// @Param work body models.Work true "成果详情"
// @Param token query string true "token"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data:{}"
// @Failure 200 {object} map[string]interface{} "code: 400, msg: 'INVALID_PARAMS', data: {}"
// @Router /api/admin/work [post]
func AddWork(c *gin.Context) {
	req := &models.Work{}
	code := e.INVALID_PARAMS

	if err := c.BindJSON(&req); err != nil {
		logging.Error(fmt.Sprintf("Error binding JSON: %s", err))
		log.Printf("Error binding JSON: %s", err)
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]string),
		})
		return
	}

	valid := validation.Validation{}
	valid.Required(req.Desc, "desc").Message("名称不能为空")
	valid.Required(req.Year, "year").Message("年份不能为空")
	valid.Required(req.Type, "type").Message("类型不能为空")

	if !valid.HasErrors() {
		code = e.SUCCESS
		models.AddWork(req)
	} else {
		for _, err := range valid.Errors {
			log.Printf("Validation error: %s, %s", err.Key, err.Message)
			logging.Info(err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// UpdateWork 更新成果
// @Summary 更新成果
// @Description 根据ID更新成果的详细信息
// @Tags 成果
// @Accept json
// @Produce json
// @Param id path int true "成果ID"
// @Param work body models.Work true "更新成果详情"
// @Param token query string true "token"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data: {}"
// @Failure 200 {object} map[string]interface{} "code: 400, msg: 'INVALID_PARAMS', data: {}"
// @Failure 200 {object} map[string]interface{} "code: 10001, msg: 'ERROR_NOT_EXIST_WORK', data: {}"
// @Router /api/admin/work/{id} [put]
func UpdateWork(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	req := &models.Work{}
	code := e.INVALID_PARAMS

	if err := c.BindJSON(&req); err != nil {
		log.Printf("Error binding JSON: %s", err)
		logging.Error(fmt.Sprintf("Error binding JSON: %s", err))
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": make(map[string]string),
		})
		return
	}

	valid := validation.Validation{}
	valid.Required(req.Desc, "desc").Message("名称不能为空")
	valid.Required(req.Year, "year").Message("年份不能为空")
	valid.Required(req.Type, "type").Message("类型不能为空")

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistWorkByID(id) {
			models.EditWork(id, req)
		} else {
			code = e.ERROR_NOT_EXIST_WORK
		}
	} else {
		logging.Info("UpdateWork报错")
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			logging.Info(err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// DeleteWork 删除成果
// @Summary 删除成果
// @Description 通过成果ID删除一个成果记录
// @Tags 成果
// @Accept json
// @Produce json
// @Param id path int true "成果ID"
// @Param token query string true "token"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data: {}"
// @Failure 200 {object} map[string]interface{} "code: 10001, msg: 'ERROR_NOT_EXIST_WORK', data: {}"
// @Router /api/admin/work/{id} [delete]
func DeleteWork(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_PARAMS

	if models.ExistWorkByID(id) {
		code = e.SUCCESS
		models.DeleteWork(id)
	} else {
		code = e.ERROR_NOT_EXIST_WORK
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}
