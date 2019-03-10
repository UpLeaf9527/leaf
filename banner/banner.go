/*
  @Author leaf
  
  @DATE   3/9/19-23:02

  @Description
    
*/

package banner

import ("log"
		"org.leaf/doc-web/elog")

var banner string = "\n\n\n\n" +
	"//////////////////////////////////////////////////////////////////\n" +
	"// function: 	文档\n" +
	"// Author:  	leaf\n" +
	"// Blog:    	none\n" +
	"// E-Mail:  	none\n" +
	"// Version: 	1.0\n" +
	"//////////////////////////////////////////////////////////////////\n"
var logger *log.Logger

func init(){
	logger = elog.Logger
}

func Banner() {
	logger.Print(banner)
}