package configutil

import "testing"

func TestLoadJsonConfig(t *testing.T) {
	path := "./config.json"

	type data struct {
		Port     int    `json:"port"`
		Env      string `json:"env"`
		LogLevel string `json:"logLevel"`
	}

	var obj data
	err := LoadJsonConfig(path, &obj)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(obj)

	var obj2 data
	err = LoadConfig(JsonConfig, path, &obj2)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(obj2)
}

func TestLoadXMLConfig(t *testing.T) {
	path := "./config.xml"

	type data struct {
		Port     int    `xml:"port"`
		Env      string `xml:"env"`
		LogLevel string `xml:"logLevel"`
	}

	var obj data
	err := LoadXMLConfig(path, &obj)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(obj)

	var obj2 data
	err = LoadConfig(XMLConfig, path, &obj2)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(obj2)
}

func TestLoadYAMLConfig(t *testing.T) {
	path := "./config.yaml"

	type data struct {
		Server struct {
			Port int `yaml:"port"`
		}
		Env string `yaml:"env"`
		Log struct {
			Level string `yaml:"level"`
		}
	}

	var obj data
	err := LoadYAMLConfig(path, &obj)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(obj)

	var obj2 data
	err = LoadConfig(YAMLConfig, path, &obj2)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(obj2)
}
