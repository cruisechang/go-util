package runtime

import (
	"fmt"
	"os"
	"path/filepath"
)

//MakeFolder make target dir
//will check if folder exists
func MakeFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}

//MakeFile maks a file with our without folder.
//will check if file exists.
func MakeFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {

		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
		defer f.Close()

		if err != nil {
			f, err = os.Create(path)
			if err != nil {
				fmt.Printf("error opening file: %v \n", err)
				return
			}
		}
	}
}

//Exist check if folder or file exist.
func Exist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

//GetNowPath return path of the current running file
func GetNowPath() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
