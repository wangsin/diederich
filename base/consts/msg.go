package consts

const (
	Success   = 0
	ErrParams = -1000020
)

var CodeToName = map[int]string{
	Success:   "成功",
	ErrParams: "参数错误",
}
