package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	StorePath = fmt.Sprintf("%v/config/config.json", GetWD())
)

type StoreStruct struct {
	ProjectName      string
	ProjectLang      string
	SelectedCompiler string
	MainFile         string
	MainDir          string
	ModulesDir       string
	Compilers        []string
}

func GetWD() string {
	wd, err := os.Getwd()

	if err != nil {
		println(err)
		PrintError("Unable to get current working directory")
	}

	return wd
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

func SetCompiler(compiler string, store *StoreStruct) {
	store.SelectedCompiler = compiler
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
