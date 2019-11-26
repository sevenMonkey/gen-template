package e

const (
	SUCCESS               = 200
	ERROR                 = 500
	ERROR_INVALID_REQUEST = 100001
	ERROR_INVLIAD_PARA    = 100002
)

var MsgFlags = map[int]string{
	SUCCESS:               "请求完成",
	ERROR:                 "数据服务出小差了，稍后再试哦~",
	ERROR_INVALID_REQUEST: "无效接口",
	ERROR_INVLIAD_PARA:    "缺少必要的参数",
}

func GetMsg(code int) string {
	if msg, ok := MsgFlags[code]; ok {
		return msg
	}
	return MsgFlags[ERROR]
}
