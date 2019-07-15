package tcppage

// 封包

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// 二进制封包格式
type Packet struct {
	Size uint16   //占2个字节
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

func readPacket(dataReader io.Writer) (pkt Packet, err error) {

	// Size 为unit16类型 占2个字节
	var sizeBuffer = make([]byte,2)

	//持续读取size直到读到为止
	_, err = io.ReadFull(dataReader, sizeBuffer)

	if err != nil{
		fmt.Println(err)
		return
	}

	// 使用bytes.Reader 读取sizeBuffer 中的数据
	sizeBuffer := bytes.NewReader(sizeBuffer)

	// 读取小端的unit16作为size
	err = binary.Read(sizeBuffer, binary.LittleEndian, &pkt.Size)

	if err != nil{
		fmt.Println(err)
		return
	}

	pkt.Body = make([]byte, pkt.Size)

	_, err = io.ReadFull(dataReader, pkt.Body)

	return
}