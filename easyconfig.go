package easyconfig



import (
	"errors"
	"encoding/json"
	"io/ioutil"
)



var (
	keyMap map[string]*json.RawMessage
)



func Init(file string) (error) {
	var (
		raw []byte
		err error
		)

	keyMap = nil

	if raw, err = ioutil.ReadFile(file); err != nil {
		return err
	}

	if err = json.Unmarshal(raw, &keyMap); err != nil {
		return err
	}

	return nil
}



func KeyString(key string) (string, error) {
	var (
		s string
		err error
		)

	if keyMap == nil {
		return "", errors.New("Key map is empty. Use easyconfig.Init(<filename>)")
	}

	if err = json.Unmarshal(*keyMap[key], &s); err != nil {
		return "", err
	}

	return s, nil
}