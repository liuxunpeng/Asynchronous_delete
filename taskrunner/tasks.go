package taskrunner

import (
	"errors"
	"log"
	"os"
	"sync"
	"video_server/Scheduler/dbops"
)
//删除文件
func deleteVideo(vid string)error{
	err:=os.Remove(VIDED_PATH+vid)

	if err!=nil && !os.IsNotExist(err){
		log.Printf("Deleting video error:%v",err)
		return err
	}
	return nil
}
func VideoClearDispatcher(dc dataChan)error{
  res,err:=dbops.ReadVideoDeletionRecord(3)
  if err!=nil{
  	log.Panicf("Video Clear dispatcher error:%v",err)
  	return err
  }
  if len(res)==0{
	return errors.New("All tasks finshed")
  }
	for _,id:=range res{
		dc<-id
	}
	return nil
}
func VideoClearExecutor(dc dataChan)error{
	errMap:=&sync.Map{}  //线程安全
	var err error
	forloop:
		for{
			select {
				case vid:=<-dc:
					go func(id interface{}) {
						if err:=deleteVideo(id.(string));err!=nil {
							errMap.Store(id, err)
							return
						}
						if err:=dbops.DelVideoDeletionRecord(id.(string));err!=nil{
							errMap.Store(id,err)
							return
						}
					}(vid)
			default:
				break forloop
			}
		}
	errMap.Range(func(k,v interface{})bool{
		err=v.(error)
		if err!=nil{
			return  false
		}
		return true
	})
	return err
}
