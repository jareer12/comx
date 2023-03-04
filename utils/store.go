package utils

import (
	"encoding/json"
	"os"
)

const (
	StorePath = "./config/config.json"
)

type StoreStruct struct {
	Compilers []string
}

func StoreContents() (StoreStruct, error) {
	if bytes, err := os.ReadFile(StorePath); err != nil {
		return StoreStruct{}, nil
	} else {
		if store, err := GetStore(bytes); err != nil {
			return StoreStruct{}, err
		} else {
			return store, nil
		}
	}
}

func GetStore(bytes []byte) (StoreStruct, error) {
	var ys StoreStruct

	if err := json.Unmarshal(bytes, &ys); err != nil {
		return StoreStruct{}, err
	}

	return ys, nil
}

func StoreToText(ys StoreStruct) ([]byte, error) {
	bytes, err := json.Marshal(ys)

	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}

func SaveConfig(ys StoreStruct) error {
	bytes, err := StoreToText(ys)

	if err != nil {
		return err
	}

	SaveConfigBytes(bytes)

	return nil
}

func SaveConfigBytes(code []byte) {
	os.WriteFile(StorePath, code, os.ModePerm)
}

func AddCompiler(compiler string, store *StoreStruct) {
	store.Compilers = append(store.Compilers, compiler)
}

func HasCompiler(comp string, store StoreStruct) bool {
	for i := 0; i < len(store.Compilers); i++ {
		if comp == store.Compilers[i] {
			return true
		}
	}

	return false
}
