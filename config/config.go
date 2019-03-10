/*
  @Author leaf
  
  @DATE   3/10/19-08:58

  @Description
    
*/

package config

import (
	"io/ioutil"
	"log"
	"org.leaf/doc-web/elog"
	"os"
	"path/filepath"
	"strings"
)

type document struct {
	Name		string
	Path        string
	Uri         string
	Description string
}

var docDir = "./doc"
var readme = "readme"
var logger *log.Logger

func init() {
	logger = elog.Logger
}

func Docs() *[]document {
	var result []document
	file, e := os.Open(docDir)
	fatalf(e, docDir)
	infos, err := file.Readdir(0)
	fatalf(err, docDir)
	for _, val := range infos {
		if val.IsDir() {
			filename := val.Name()
			uFile := filepath.Join(docDir, filename)
			docum := docu(uFile)
			if docum != nil {
				result = append(result, *docum)
			}
		}
	}
	return &result
}

func docu(pathAndName string) *document {
	file, e := os.Open(pathAndName)
	defer file.Close()
	if e != nil {
		logger.Printf("read %s directory failed,error: %s.", pathAndName, e)
		return nil
	}
	infos, e := file.Readdir(0)
	if e != nil {
		logger.Printf("read %s directory failed,error: %s.", pathAndName, e)
		return nil
	}
	var desc string
	for _, val := range infos {
		if !val.IsDir() {
			readmeFile := strings.ToLower(val.Name())
			if strings.Contains(readmeFile,readme) {
				readmeFile = filepath.Join(pathAndName,val.Name())
				readme, e := os.Open(readmeFile)
				if e != nil {
					logger.Printf("read %s file failed,error: %s.", readme.Name(), e)
				}
				content, err := ioutil.ReadAll(readme)
				if err != nil {
					logger.Printf("read %s file failed,error: %s.", readme, e)
				}
				desc = string(content)
				break
			}
		}
	}
	_, name := filepath.Split(pathAndName)
	return &document{
		Name:        name,
		Path:        pathAndName,
		Uri:         name,
		Description: desc,
	}
}

func fatalf(e error, name string) {
	if e != nil {
		logger.Fatalf("read %s directory failed,error: %s.", name, e)
	}
}
