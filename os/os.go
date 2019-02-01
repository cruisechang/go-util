package os

import (
	"fmt"
	"os"
)

//CreateDirectory creates dir with target name
//If you want to full access dir,pass os.ModePerm as FileMode parameter
func CreateDirectory(name string, permissionBits os.FileMode) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return os.Mkdir(name, permissionBits)
	}
	return nil
}

//CreateFile try to open a target file,if error,create one
//If you want to full access file,pass os.ModePerm as FileMode parameter
func CreateFile(fullPath string, mode os.FileMode) (*os.File, error) {

	f, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, mode)
	defer f.Close()
	if err != nil {
		f, err = os.Create(fullPath)
		if err != nil {
			fmt.Printf("Create file error: %v \n", err)
			return nil, err
		}
	}
	return f, nil
}
