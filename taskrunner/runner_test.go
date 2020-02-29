package taskrunner

import (
	"log"
	"testing"
	"time"
)

func TsetRunner(t *testing.T){
	d:= func(dc dataChan)error {
		for i:=0;i<30 ;i++  {
			dc<-i
			log.Panicf("Dispatcher sent:%v",i)
		}
		return nil
	}
	e:= func(dc dataChan) error{
		forloop:
			for  {
				select {
				case d:=<-dc:
					log.Printf("Executor received:%v",d)
				default:
					break forloop

				}
			}
		return nil
	}
	renner:=NewRunner(30,false,d,e)
	go renner.StartAll()
	time.Sleep(3*time.Second)
}
