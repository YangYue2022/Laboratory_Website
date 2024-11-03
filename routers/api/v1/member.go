package v1

import (
	"UISwebsite_backend/models"
	"UISwebsite_backend/pkg/app"
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

// GetMembers
// @Summary 获取成员列表
// @Description 根据身份和姓名进行筛选，获取成员列表
// @Tags 成员
// @Accept json
// @Produce json
// @Param name query string false "成员姓名"
// @Param type query string false "成员身份"
// @Param page query int false "页码"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data: {lists: 成果列表, total: 总数}"
// @Router /api/members [get]
func GetMembers(c *gin.Context) {
	appG := app.Gin{C: c}
	members, _ := models.GetMembers(util.GetPage(c), setting.AppSetting.PageSize)
	count, _ := models.GetMembersTotal()

	data := make(map[string]interface{})
	data["lists"] = members
	data["total"] = count

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// AddMember
// @Summary 新增成员
// @Description 添加一个新的成员记录
// @Tags 成员
// @Accept json
// @Produce json
// @Param member body models.Member true "成员详情"
// @Param token query string true "token"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data:{}"
// @Failure 200 {object} map[string]interface{} "code: 400, msg: 'INVALID_PARAMS', data: {}"
// @Router /api/admin/member [post]
func AddMember(c *gin.Context) {
	req := &models.Member{}
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
	valid.Required(req.Name, "name").Message("姓名不能为空")
	valid.Required(req.Identity, "identity").Message("身份不能为空")
	valid.Required(req.Title, "title").Message("职称/学历不能为空")

	if !valid.HasErrors() {
		code = e.SUCCESS
		models.AddMember(req)
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

// UpdateMember 更新成员
// @Summary 更新成员
// @Description 更新成员的详细信息
// @Tags 成员
// @Accept json
// @Produce json
// @Param member body models.Member true "成员详情"
// @Param token query string true "token"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data: {}"
// @Failure 200 {object} map[string]interface{} "code: 400, msg: 'INVALID_PARAMS', data: {}"
// @Failure 200 {object} map[string]interface{} "code: 1004, msg: 'ERROR_NOT_EXIST_MEMBER', data: {}"
// @Router /api/admin/member/{id} [put]
func UpdateMember(c *gin.Context) {
	req := &models.Member{}
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
	valid.Required(req.Name, "name").Message("姓名不能为空")
	valid.Required(req.Identity, "identity").Message("身份不能为空")
	valid.Required(req.Title, "title").Message("职称/学历不能为空")

	if !valid.HasErrors() {
		code = e.SUCCESS
		if ok, _ := models.ExistMemberByID(req.ID); ok {
			models.EditMember(req)
		} else {
			code = e.ERROR_NOT_EXIST_MEMBER
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

// DeleteMember 删除成员
// @Summary 删除成员
// @Description 通过成员ID删除一个成员记录
// @Tags 成员
// @Produce json
// @Param id path int true "成果ID"
// @Param token query string true "token"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data: {}"
// @Failure 200 {object} map[string]interface{} "code: 10001, msg: 'ERROR_NOT_EXIST_MEMBER', data: {}"
// @Router /api/admin/member/{id} [delete]
func DeleteMember(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_PARAMS

	if ok, _ := models.ExistMemberByID(id); ok {
		models.DeleteMember(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_MEMBER
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
