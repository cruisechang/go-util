package debug

import (
	"fmt"
	"os"
	"testing"
)

func TestGetPackageFileName(t *testing.T) {
	t.Logf("package and function name : %s \n", GetFullPackageAndFunctionName(TestGetPackageFileName))
	t.Logf("package and function name : %s \n", GetShortPackageAndFunctionName(TestGetPackageFileName))
	t.Logf("package and function name : %s \n", GetFullPackageAndFunctionNameAndLine(TestGetPackageFileName))
	t.Logf("package and function name : %s \n", GetShortPackageAndFunctionNameAndLine(TestGetPackageFileName))

}

func TestGetFileLine(t *testing.T) {
	file, line := GetFileLine()
	fmt.Printf("File : %s \n", file)
	fmt.Printf("line : %d \n", line)
}

func TestLogWarning(t *testing.T) {
	//gameserver20160603.log
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf(GetShortPackageAndFunctionName(TestLogWarning)+" error opening file: %v \n", err)
	}
	defer f.Close()

	LogWarning(f, "LogWarning test")
}

func TestLogInfo(t *testing.T) {
	LogInfo(os.Stdout, "info ", "...")
}
