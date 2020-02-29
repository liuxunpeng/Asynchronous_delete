package dbops

import (
	_"github.com/go-sql-driver/mysql"
	"log"
)

func AddUserCredential(vid string) error {
	stmIns,err:=dbConn.Prepare("INSERT INTO video_rec (video_id) VALUES(?)")
	if err!=nil{
		return err
	}
	_,err=stmIns.Exec(vid)
	if err !=nil{
		log.Printf("AddVideoDeletionRecord error: %v",err)
		return err
	}
	defer stmIns.Close()
	return nil
}
