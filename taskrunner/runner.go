package taskrunner

import (

)

//runner  常驻 startDispatcher  长时间等待channel

//control channel
// data channels 数据信道
type Runner struct {
	Controller controlChan
	Error controlChan
	Data dataChan
	dataSize int //
	longlived bool //是否长期存活的channel
	Dispatcher fn
	Executer fn
}
/**
	size:数据长度
	longlived : 是否常驻
	d:生产者方法
	e:执行者方法
 */
func NewRunner(size int,longlived bool,d fn,e fn)*Runner{

	return &Runner{
		Controller: make(chan string,1),
		Error:      make(chan  string,1),
		Data:       make(chan interface{},size),
		dataSize:   size,
		longlived:  longlived,
		Dispatcher: d,
		Executer:   e,
	}

}
//常驻任务

func (r *Runner) startDispatch(){
	defer func() {
		if !r.longlived{
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()
	for {
		//异步不断接受channel发送请求 独立非阻塞
		select {
			case c:=<-r.Controller:
				if c==READY_TO_DISPATCH {
					err:=r.Dispatcher(r.Data)  //写入任务
					if err!=nil{
						r.Error<-CLOSE //进入colse 关闭程序分支
					}else{
						r.Controller<-READY_TO_EXECUTE //生产者执行完毕 进入消费者分支
					}
				}
				if c==READY_TO_EXECUTE{
					err:=r.Executer(r.Data)
					if err !=nil{
						r.Error<-CLOSE
					}else{
						r.Controller<-READY_TO_DISPATCH //消费者执行完毕 进入生产者分支
					}
				}
		    case e:=<-r.Error:
					if e==CLOSE{  //退出操作
						return
					}
		    default:
		}
	}
}
func(r *Runner) StartAll(){
	r.Controller<-READY_TO_DISPATCH
	r.startDispatch()
}