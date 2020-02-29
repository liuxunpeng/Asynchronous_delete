package taskrunner
//消息类型
const (
	READY_TO_DISPATCH="d"
	READY_TO_EXECUTE="e"
	CLOSE="c"

	VIDED_PATH="./video/"
)
type controlChan chan string
type dataChan chan interface{}
type fn func(dc dataChan) error
