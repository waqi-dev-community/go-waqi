package lib_test

import "encoding/json"

func payloadStringToStruct(str string, s interface{}) error {
	bb := []byte(str)
	err := json.Unmarshal(bb, s)
	if err != nil {
		return err
	}
	return nil
}
