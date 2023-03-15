package configutil

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"gopkg.in/yaml.v3"
	"sync"
)

// ConfigType 文件配置类型
type ConfigType string

const (
	// JsonConfig json配置
	JsonConfig ConfigType = "json"
	// XMLConfig xml配置
	XMLConfig ConfigType = "xml"
	// YAMLConfig yaml配置
	YAMLConfig ConfigType = "yaml"
)

var (
	onceDo    sync.Once
	configMap map[ConfigType]func(path string, obj any) error
)

func init() {
	initConfig()
}

func initConfig() {
	onceDo.Do(func() {
		configMap = make(map[ConfigType]func(path string, obj any) error)

		configMap[JsonConfig] = LoadJsonConfig
		configMap[XMLConfig] = LoadXMLConfig
		configMap[YAMLConfig] = LoadYAMLConfig
	})
}

// LoadConfig 加载配置
func LoadConfig(configType ConfigType, path string, obj any) error {
	initConfig()

	c, ok := configMap[configType]
	if !ok {
		return errors.New("not fount type unmarshal fn")
	}

	return c(path, obj)
}

// LoadJsonConfig 加载json配置
func LoadJsonConfig(path string, obj any) error {
	return openAndUnmarshalFile(path, obj, json.Unmarshal)
}

// LoadXMLConfig 加载xml配置
func LoadXMLConfig(path string, obj any) error {
	return openAndUnmarshalFile(path, obj, xml.Unmarshal)
}

// LoadYAMLConfig 加载yaml配置
func LoadYAMLConfig(path string, obj any) error {
	return openAndUnmarshalFile(path, obj, yaml.Unmarshal)
}
