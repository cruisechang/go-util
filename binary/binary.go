package binary

import (
	"encoding/binary"
	"errors"
)

//ByteToIntByLittleEndian is a byte[] to int converter.
//c# BitConverter.GetBytes is LittleEndian
//LittleEndian 從最小開始 (最低位元組在前)
//BigEndian  從最大開始 (最高位元組在前)
//Ex: long 0x12345678
//littleEndian 0x78 0x56 0x34 0x12
//BigEndian 78 56 43 12
func ByteToIntByLittleEndian(data []byte) (int, error) {
	//Check data length
	if len(data) < 4 {
		return 0, errors.New("byteToInt32 data length not enougth")
	}

	return int(binary.LittleEndian.Uint32(data)), nil
}

//Uint32ToByteByBigEndian is a uint32 to byte[] converter.
func Uint32ToByteByBigEndian(num uint32) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, num)
	return bs
}

//SingleByteToIntByBigEndian is single byte to uint16 converter.
func SingleByteToIntByBigEndian(data byte) (int, error) {

	//Check data type
	typeBytes := make([]byte, 2)
	//高位元補0
	typeBytes[0] = 0
	typeBytes[1] = data
	return int(binary.BigEndian.Uint16(typeBytes)), nil
}
