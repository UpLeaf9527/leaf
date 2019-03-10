/*
  @Author leaf
  
  @DATE   3/10/19-08:40

  @Description
    
*/

package elog

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

var log_file string = "./log.log"

func init(){
	logFile,err :=os.Create(log_file)
	defer logFile.Close()
	if err != nil{
		log.Fatalf("open elog file(%s) failed, err: %s.",log_file,err)
	}
	writer := io.MultiWriter(logFile, os.Stdout)
	Logger = log.New(writer, "[info]", log.Llongfile)
}
