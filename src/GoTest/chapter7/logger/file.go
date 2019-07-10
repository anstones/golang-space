package logger

import (
	"errors"
	"fmt"
	"os"
)

type fileWriter struct {
	flle *os.File
}

func (f *fileWriter) Setfile(filename string) (err error) {
	if f.flle != nil {
		f.flle.Close()
	}

	// 创建文件并保存句柄
	f.flle, err = os.Create(filename)
	// 如果出错，返回错误
	return err

}

func (f *fileWriter) Write(data interface{}) error {
	// 日志文件没有创建成功
	if f.flle == nil {
		return errors.New("file not create")
	}

	str := fmt.Sprintf("%v\n", data)

	_, err := f.flle.Write([]byte(str))

	return err
}

func newFileWriter() *fileWriter {
	return &fileWriter{}
}
