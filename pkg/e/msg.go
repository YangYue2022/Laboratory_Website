package e

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_NOT_EXIST_WORK:            "该成果不存在",
	ERROR_NOT_EXIST_REWARD:          "该奖励不存在",
	ERROR_NOT_EXIST_QA:              "该文章不存在",
	ERROR_NOT_EXIST_MEMBER:          "该成员不存在",
	ERROR_NOT_EXIST_RECORD:          "该记录不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
	ERROR_GET_MEMBER_FAIL:           "获取成员列表失败",
	ERROR_COUNT_MEMBER_FAIL:         "统计成员失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
