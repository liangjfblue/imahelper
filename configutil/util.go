package configutil

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
)

func openFile(path string) (data []byte, err error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 666)
	if err != nil {
		return
	}
	defer f.Close()

	data, err = ioutil.ReadAll(f)
	if err != nil {
		return
	}

	return
}

func openAndUnmarshalFile(path string, obj any, unmarshalFn func(data []byte, v any) error) error {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Pointer {
		return errors.New("obj must pointer")
	}

	data, err := openFile(path)
	if err != nil {
		return err
	}

	err = unmarshalFn(data, obj)
	if err != nil {
		return err
	}

	return nil
}
