package tcppage

// 封包

import (
	"bytes"
	"encoding/binary"
	"io"
)

// 二进制封包格式
type Packet struct {
	Size uint16
	Body []byte
}


//将数据写入dataWriter
func WritePacket(dataWriter io.Writer, data []byte) error  {
	// 准备字节缓冲数组
	var buffer bytes.Buffer

	// 将size写入缓存
	if err := binary.Write(&buffer, binary.LittleEndian, uint16(len(data))); err != nil{
		return err
	}

	//写入body
	if _,err := buffer.Write(data);err != nil{
		return err
	}

	// 获得写入的完整数据
	out := buffer.Bytes()

	// 写入dataWriter
	if _, err := dataWriter.Write(out); err != nil{
		return nil
	}

	return nil
}