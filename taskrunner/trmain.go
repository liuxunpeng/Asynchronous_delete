package taskrunner

import "time"

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}
func NewWorker(interval time.Duration,r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker(){
	for  {
		select {
		//通过时间通道触发 定时响应携程
		   case<-w.ticker.C:
		   	go w.runner.StartAll()
		}
	}
}

func Start(){
	r:=NewRunner(3,true,VideoClearDispatcher,VideoClearExecutor)
	w:=NewWorker(60,r)
	go w.startWorker()
}
