package dbops

import "log"

//api->videoid->mysql
//dispathcer->mysql->videoid->datachannel
//executor->dataChannel->videoid->delete videos

//取出数据
func ReadVideoDeletionRecord(count int)([]string,error){
	stmtOut,err:=dbConn.Prepare("SELECT video_id FROM video_del_rec LIMIT ?")
	var ids []string
	if err!=nil{
		return ids,err
	}
	rows,err:=stmtOut.Query(count)
	if err!=nil{
		log.Panicf("Query VideoDeletionRecond error:%v",err)
		return ids,err
	}
	for rows.Next(){
		var id string
		if err:=rows.Scan(&id);err!=nil{
			return ids,err
		}
		ids=append(ids,id)
	}
	defer stmtOut.Close()
	return ids ,nil
}
//已删除信息删除时候表删掉
func DelVideoDeletionRecord(vid string)error{
	stmtDel,err:= dbConn.Prepare("DELETE from video_del_rec where video_id=?")
	if err!=nil{
		return err
	}
	_,err=stmtDel.Exec(vid)
	if err!=nil{
		log.Printf("Deleting VideoDeletionRecord error:%v",err)
		return  err
	}
	defer stmtDel.Close()
	return nil
}