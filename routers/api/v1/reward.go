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

// GetRewards 获取奖励列表（根据类型和项目名称查询）
// @Summary 获取奖励列表
// @Description 根据类型和项目名称查询
// @Tags 奖励
// @Accept  json
// @Produce  json
// @Param   project  query     string  false  "项目名称"
// @Param   type     query     string  false  "类型（竞赛/其他）"
// @Param page query int false "页码"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'success', data: {lists, total}"
// @Router /api/rewards [get]
func GetRewards(c *gin.Context) {
	name := c.Query("name")
	type_ := c.Query("type")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["`name` LIKE ?"] = "%" + name + "%"
	}

	if type_ != "" {
		maps["type"] = type_
	}

	code := e.SUCCESS

	data["lists"] = models.GetRewards(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetRewardsTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// AddReward 新增奖励
// @Summary 新增奖励
// @Description 新增奖励（竞赛名称，项目名称，年份，类型，获奖人）
// @Tags 奖励
// @Accept  json
// @Produce  json
// @Param   reward  body     models.Reward  true  "Add Reward Request"
// @Param token query string true "token"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'Success', data: nil"
// @Failure 200 {object} map[string]interface{} "code: 500, msg: 'INVALID_PARAMS', data: nil"
// @Router /api/admin/reward [post]
func AddReward(c *gin.Context) {
	req := &models.Reward{} // 初始化 req 为一个指向新的 models.Qa 的指针
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
	valid.Required(req.Name, "name").Message("竞赛名称不能为空")
	valid.Required(req.Project, "project").Message("竞赛项目不能为空")
	valid.Required(req.Year, "year").Message("年份不能为空")
	valid.Required(req.Type, "type").Message("类型不能为空")
	valid.Required(req.Winner, "winner").Message("获奖人不能为空")

	if !valid.HasErrors() {
		code = e.SUCCESS
		models.AddReward(req)
	} else {
		for _, err := range valid.Errors {
			log.Printf("Validation error: %s, %s", err.Key, err.Message)
			logging.Error(err)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// UpdateReward
// @Summary 编辑奖励
// @Description 根据id更新已存在的奖励
// @Tags 奖励
// @Accept  json
// @Produce  json
// @Param   id      path      int   true  "Reward ID"
// @Param   reward  body      models.Reward  true  "Update Reward Request"
// @Param token query string true "token"
// @Success 200     {object}  map[string]interface{} "code: 200, msg: 'Success', data: nil"
// @Failure 200     {object}  map[string]interface{} "code: 400, msg: 'INVALID_PARAMS', data: nil"
// @Failure 200     {object}  map[string]interface{} "code: 10002, msg: 'ERROR_NOT_EXIST_REWARD', data: nil"
// @Router /api/admin/reward/{id} [put]
func UpdateReward(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	req := &models.Reward{}
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
	valid.Required(req.Name, "name").Message("竞赛名称不能为空")
	valid.Required(req.Project, "project").Message("竞赛项目不能为空")
	valid.Required(req.Year, "year").Message("年份不能为空")
	valid.Required(req.Type, "type").Message("类型不能为空")
	valid.Required(req.Winner, "winner").Message("获奖人不能为空")

	if !valid.HasErrors() {
		if models.ExistRewardByID(id) {
			code = e.SUCCESS
			models.EditReward(id, req)
		} else {
			code = e.ERROR_NOT_EXIST_REWARD
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
		"data": make(map[string]string),
	})

}

// DeleteReward
// @Summary 删除奖励
// @Description 通过奖励ID删除一个奖励项
// @Tags 奖励
// @Accept json
// @Produce json
// @Param id path int true "奖励ID"
// @Param token query string true "token"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data: {}"
// @Failure 200 {object} map[string]interface{} "code: 10002, msg: 'ERROR_NOT_EXIST_REWARD', data: {}"
// @Router /api/admin/reward/{id} [delete]
func DeleteReward(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_PARAMS

	if models.ExistRewardByID(id) {
		code = e.SUCCESS
		models.DeleteReward(id)
	} else {
		code = e.ERROR_NOT_EXIST_REWARD
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}
