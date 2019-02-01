package binary

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestBinaryToInt(t *testing.T) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, 175)

	fmt.Println(binary.BigEndian.Uint32(b))

	var myint int
	buf := bytes.NewBuffer(b)

	binary.Read(buf, binary.LittleEndian, &myint)

	fmt.Println(myint)

}

func TestSingleByteToInt(t *testing.T) {

	bytes := []byte{0}
	dataType, _ := SingleByteToIntByBigEndian(bytes[0])

	fmt.Println(dataType)
}
