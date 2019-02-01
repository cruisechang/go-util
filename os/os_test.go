package os

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestCreateDir(t *testing.T) {

	name := "logs"

	err := CreateDir(name, os.ModePerm)

	if err != nil {
		fmt.Printf("Create dir error: %v \n", err)
		return
	}
	fmt.Printf("Create dir success, dir name:%s\n", name)
}

func TestCreateFile(t *testing.T) {

	ti := time.Now()
	path := "logs/serverLog" + ti.Format("20060102") + ".json"

	f, err := CreateFile(path, os.ModePerm)

	if err != nil {
		fmt.Printf("make file error: %v\n", err)
		return
	}
	fmt.Printf("make file success: %s\n", f.Name())

}
