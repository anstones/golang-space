package tcppage

import (
	"bufio"
	"net"
)

func handleSession(conn net.Conn, callback func(net.Conn, []byte) bool)  {

	// c创建一个socket的读取器
	dataReader := bufio.NewReader(conn)


	for {
		// 从链接读取封包
		pkt,err := readPacket(dataReader)

		if err != nil || !callback(conn, pkt.Body){

			// 回调要求退出
			conn.Close()
			break
		}
	}

}