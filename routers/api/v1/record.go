package v1

import (
	"UISwebsite_backend/models"
	"UISwebsite_backend/pkg/e"
	"UISwebsite_backend/pkg/setting"
	"UISwebsite_backend/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"strconv"
)

// GetRecords
// @Summary 获取记录列表
// @Description 根据活动名称模糊匹配，获取记录列表
// @Tags 记录
// @Produce json
// @Param activity_name query string false "活动名称"
// @Param visibility query bool false "可见"
// @Param page query int false "页码"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data: {lists: 记录列表, total: 总数}"
// @Router /api/records [get]
func GetRecords(c *gin.Context) {
	activityName := c.Query("activity_name")
	maps := make(map[string]interface{})
	if activityName != "" {
		activityIDs := models.GetActivityIDByName(activityName)

		maps["activity_ids"] = activityIDs

		//fmt.Println(activityID)
	}

	visibility := c.Query("visibility")
	// 解析 visibility 参数为 bool 类型
	if visibility != "" {
		visibility, _ := strconv.ParseBool(visibility)
		maps["visibility"] = visibility
	}
	data := make(map[string]interface{})

	code := e.SUCCESS

	data["lists"] = models.GetRecords(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetRecordsTotal(maps)

	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// Image 用于接收前端发送的图片数据
type Image struct {
	Name string `json:"name" binding:"required"`
	URL  string `json:"url" binding:"required"`
}

// AddRecordRequest 用于接收前端发送的JSON数据
type AddRecordRequest struct {
	Name   string  `json:"name" binding:"required"`
	Year   string  `json:"year" binding:"required"`
	Images []Image `json:"images" binding:"required"`
}

// AddRecord
// @Summary 新增记录
// @Description 新增一个活动及其图片记录
// @Tags 记录
// @Accept json
// @Produce json
// @Param data body AddRecordRequest true "活动信息和图片URL列表"
// @Param token query string true "token"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data:{}"
// @Failure 200 {object} map[string]interface{} "code: 400, msg: '请求参数错误', data: {}"
// @Router /api/admin/record [post]
func AddRecord(c *gin.Context) {
	var req AddRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  e.GetMsg(e.INVALID_PARAMS),
		})
		return
	}

	// 创建活动
	activity := models.Activity{Name: req.Name, Year: req.Year}
	models.AddActivity(&activity)

	// 创建记录
	for _, image := range req.Images {
		record := models.Record{
			ActivityID: activity.ID,
			Name:       image.Name,
			PhotoUrl:   image.URL,
			Visibility: true, // 默认为可见
		}
		models.AddRecord(&record)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
	})
}

type UpdateVisibilityRequest struct {
	ID         int   `json:"id" binding:"required"`
	Visibility *bool `json:"visibility" binding:"required"`
}

// UpdateRecordVisibility
// @Summary 更新记录可见性
// @Description 更新指定记录的可见性状态
// @Tags 记录
// @Accept json
// @Produce json
// @Param data body UpdateVisibilityRequest true "记录ID和可见性状态"
// @Param token query string true "token"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data:{}"
// @Failure 200 {object} map[string]interface{} "code: 400, msg: 'INVALID_PARAMS', error: '错误信息'"
// @Router /api/admin/record/visibility [put]
func UpdateRecordVisibility(c *gin.Context) {
	var req UpdateVisibilityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  e.INVALID_PARAMS,
			"msg":   e.GetMsg(e.INVALID_PARAMS),
			"error": err.Error(), // 返回错误信息以便调试
		})
		return
	}

	models.UpdateRecordVisibility(req.ID, *req.Visibility)

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
	})
}

type UpdateRecordInput struct {
	Name string `json:"name"`
}

func UpdateRecord(c *gin.Context) {
	var input UpdateRecordInput

	// 绑定 JSON 数据
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  e.INVALID_PARAMS,
			"msg":   e.GetMsg(e.INVALID_PARAMS),
			"error": err.Error(), // 返回错误信息以便调试
		})
		return
	}

	id := com.StrTo(c.Param("id")).MustInt()

	models.UpdateRecordName(id, input.Name)

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
	})
}

// DeleteRecord 删除记录
// @Summary 删除记录
// @Description 通过记录ID删除一个记录
// @Tags 记录
// @Produce json
// @Param id path int true "成果ID"
// @Param token query string true "token"
// @Success 200 {object} map[string]interface{} "code: 200, msg: 'SUCCESS', data: {}"
// @Failure 200 {object} map[string]interface{} "code: 10005, msg: '该记录不存在', data: {}"
// @Router /api/admin/record/{id} [delete]
func DeleteRecord(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_PARAMS

	if models.ExistRecordByID(id) {
		models.DeleteRecord(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_RECORD
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
