/*
  @Author leaf
  
  @DATE   3/9/19-10:35

  @Description
    
*/

package main

import (
	"log"
	"net/http"
	"org.leaf/doc-web/banner"
	"org.leaf/doc-web/config"
	"org.leaf/doc-web/elog"
)

var logger *log.Logger

func init() {
	logger = elog.Logger
}

func main() {
	banner.Banner()
	var docs = config.Docs()
	for _, val := range *docs {
		logger.Printf("begin %s doc ,visit /%s ...", val.Name, val.Uri)
		http.Handle(val.Uri, http.FileServer(http.Dir(val.Path)))
	}

	http.ListenAndServe(":10000", nil)
}
