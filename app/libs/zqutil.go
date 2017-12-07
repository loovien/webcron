package libs

import (
	"encoding/binary"
	"net"
	"errors"
	"fmt"
)
type Packet interface {
	Serialize() []byte
}

type ZqProtocol struct {
	ToType int32
	FromType int32
	ToIp int32
	FromIp int32
	AskId int32
	AskId2 int32
}

type ZqPacket struct {
	Header ZqProtocol
	Body string
}

func (z *ZqPacket) Serialize() []byte {
	bodyBytes := []byte(z.Body)
	bodyBytesLength := len(bodyBytes)
	buf := make([]byte, PACKET_HEADER_LENGTH + bodyBytesLength)

	binary.LittleEndian.PutUint32(buf[:4], uint32(bodyBytesLength))
	binary.LittleEndian.PutUint16(buf[4:], uint16(z.Header.ToType))
	binary.LittleEndian.PutUint32(buf[6:], uint32(z.Header.ToIp))
	binary.LittleEndian.PutUint16(buf[10:], uint16(z.Header.FromType))
	binary.LittleEndian.PutUint16(buf[12:], uint16(z.Header.FromIp))
	binary.LittleEndian.PutUint32(buf[16:], uint32(z.Header.AskId))
	binary.LittleEndian.PutUint32(buf[20:], uint32(z.Header.AskId2))

	copy(buf[24:], bodyBytes)

	return buf
}

type ZqUtil struct {
}

func NewZqUtil() *ZqUtil {
	return &ZqUtil{}
}

// SendNotifyEmail
func (z *ZqUtil) SendNotifyEmail(params string) error {
	p := ZqProtocol{
		ToIp: 0,
		ToType: NOTIFY_SERVER_PORT,
		FromType: 0,
		FromIp: 0,
		AskId: 0,
		AskId2: 0,
	}
	packet := &ZqPacket{
		Header:p,
		Body:params,
	}
	return z.SendTcpPacket(packet)
}

// SendTcpPacket
func (z *ZqUtil) SendTcpPacket(packet *ZqPacket) error {
	raddr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", LOCAL_PROXY_HOST, LOCAL_PROXY_PORT))
	conn, err := net.DialTCP("tcp", nil, raddr)
	defer conn.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("连接本地代理失败:%s", err.Error()))
	}
	conn.Write(packet.Serialize())
	return nil
}
