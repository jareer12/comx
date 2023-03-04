package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

type YamlStore struct {
	Compilers []string
}

func DecodeStore(bytes []byte) (YamlStore, error) {
	var ys YamlStore

	if err := yaml.Unmarshal(bytes, ys); err != nil {
		return YamlStore{}, err
	} else {
		return ys, nil
	}

}

func EncodeStore(ys YamlStore) (string, error) {
	bytes, err := yaml.Marshal(ys)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func SaveConfig(ys YamlStore) error {
	bytes, err := yaml.Marshal(ys)

	if err != nil {
		return err
	}

	SaveConfigBytes(bytes)

	return nil
}

func SaveConfigBytes(code []byte) {
	os.WriteFile("./config.yaml", code, os.ModePerm)
}
