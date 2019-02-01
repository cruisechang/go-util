package encode

import (
	"encoding/base64"
	"encoding/json"
)

//Base64Encode encode []byte to base64 string
func Base64Encode(strByte []byte) string {
	return base64.StdEncoding.EncodeToString(strByte)
}

//Base64Decode decode base64 string to utf8 string
func Base64Decode(base64Str string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {

		return nil, err
	}
	return data, nil
}

//JSONEncodeFromStruct encode struct to json string
func JSONEncodeFromStruct(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	// Convert bytes to string.
	return b, nil
}

//JSONDecodeToStruct decode json to struct
//v is a struct pointer = &Stru{}
//struct field name 必須大寫開頭且名稱跟json相同
func JSONDecodeToStruct(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}

//StructToMap decode struct to map[string]interface{}
func StructToMap(v interface{}) (map[string]interface{}, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}
//IsJosn check if a string is json string
func IsJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}