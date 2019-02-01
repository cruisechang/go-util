/*

* https://github.com/satori/go.uuid
 */
package uuid

import (
	"github.com/satori/go.uuid"
)

func GetV1() (string, error) {

	u, err := uuid.NewV1()
	if err != nil {
		return "",err
	}
	return u.String(),nil
}




func GetV4() (string, error) {

	u, err := uuid.NewV4()
	if err != nil {
		return "",err
	}
	return u.String(),nil
}