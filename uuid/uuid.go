/*

* https://github.com/satori/go.uuid
 */
package uuid

import (
	"github.com/satori/go.uuid"
	"strings"
)

func GetV1() (string, error) {

	u := uuid.NewV1()
	return u.String(), nil
}

func GetV4() (string, error) {

	u := uuid.NewV4()
	return u.String(), nil
}

//GetV1 去掉 "-"
func GetV1Trim() (string, error) {

	u  := uuid.NewV1()

	return strings.ReplaceAll(u.String(), "-", ""), nil
}

//GetV4 去掉 "-"
func GetV4Trim() (string, error) {

	u := uuid.NewV4()
	return strings.ReplaceAll(u.String(), "-", ""), nil
}
