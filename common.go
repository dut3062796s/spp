package main

import (
	"context"
	"encoding/binary"
	"errors"
	"github.com/esrrhs/go-engine/src/common"
	"github.com/esrrhs/go-engine/src/loggo"
	"github.com/golang/protobuf/proto"
	"golang.org/x/sync/errgroup"
	"io"
	"net"
	"strconv"
)

const (
	MAX_MSG_SIZE       = 1024 * 1024 // 消息最大长度
	MAIN_BUFFER        = 1024 * 1024 // 主通道buffer最大长度
	CONN_BUFFER        = 1024        // 每个conn buffer最大长度
	LOGIN_TIMEOUT      = 10          // 主通道登录超时
	PING_INTER         = 1           // 主通道ping间隔
	PING_TIMEOUT_INTER = 5           // 主通道ping超时间隔
	CONN_TIMEOUT       = 300         // 每个conn的不活跃超时时间
	CONNNECT_TIMEOUT   = 10          // 每个conn的连接超时
)

func MarshalSrpFrame(f *SrpFrame, compress int) ([]byte, error) {

	if f.Type == SrpFrame_DATA && compress > 0 && len(f.DataFrame.Data) > compress {
		newb := common.CompressData(f.DataFrame.Data)
		if len(newb) < len(f.DataFrame.Data) {
			f.DataFrame.Data = newb
			f.DataFrame.Compress = true
		}
	}

	switch f.Type {
	case SrpFrame_LOGIN:
		if f.LoginFrame == nil {
			return nil, errors.New("LoginFrame nil")
		}
	case SrpFrame_LOGINRSP:
		if f.LoginRspFrame == nil {
			return nil, errors.New("LoginRspFrame nil")
		}
	case SrpFrame_DATA:
		if f.DataFrame == nil {
			return nil, errors.New("DataFrame nil")
		}
	case SrpFrame_PING:
		if f.PingFrame == nil {
			return nil, errors.New("PingFrame nil")
		}
	case SrpFrame_PONG:
		if f.PongFrame == nil {
			return nil, errors.New("PongFrame nil")
		}
	case SrpFrame_OPEN:
		if f.OpenFrame == nil {
			return nil, errors.New("OpenFrame nil")
		}
	case SrpFrame_CLOSE:
		if f.CloseFrame == nil {
			return nil, errors.New("CloseFrame nil")
		}
	default:
		return nil, errors.New("Type error")
	}

	mb, err := proto.Marshal(f)
	if err != nil {
		return nil, err
	}
	return mb, err
}

func recvFrom(ctx context.Context, wg *errgroup.Group, recvch chan<- *SrpFrame, conn *net.TCPConn) error {
	defer common.CrashLog()

	bs := make([]byte, 4)
	ds := make([]byte, MAX_MSG_SIZE)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			_, err := io.ReadFull(conn, bs)
			if err != nil {
				loggo.Error("recvFrom ReadFull fail: %s %s", conn.RemoteAddr().String(), err.Error())
				return err
			}

			len := binary.LittleEndian.Uint32(bs)
			if len > MAX_MSG_SIZE {
				loggo.Error("recvFrom len fail: %s %d", conn.RemoteAddr().String(), len)
				return errors.New("msg len fail " + strconv.Itoa(int(len)))
			}

			_, err = io.ReadFull(conn, ds[0:len])
			if err != nil {
				loggo.Error("recvFrom ReadFull fail: %s %s", conn.RemoteAddr().String(), err.Error())
				return err
			}

			f := &SrpFrame{}
			err = proto.Unmarshal(ds[0:len], f)
			if err != nil {
				loggo.Error("recvFrom Unmarshal fail: %s %s", conn.RemoteAddr().String(), err.Error())
				return err
			}

			recvch <- f
			//loggo.Info("recvFrom %s %s", conn.RemoteAddr().String(), f.Type.String())
		}
	}
}

func sendTo(ctx context.Context, wg *errgroup.Group, sendch <-chan *SrpFrame, conn *net.TCPConn, compress int) error {
	defer common.CrashLog()

	bs := make([]byte, 4)

	for {
		select {
		case <-ctx.Done():
			return nil
		case f := <-sendch:
			mb, err := MarshalSrpFrame(f, compress)
			if err != nil {
				loggo.Error("sendTo MarshalSrpFrame fail: %s %s", conn.RemoteAddr().String(), err.Error())
				return err
			}

			len := uint32(len(mb))
			binary.LittleEndian.PutUint32(bs, len)
			_, err = conn.Write(bs)
			if err != nil {
				loggo.Error("sendTo Write fail: %s %s", conn.RemoteAddr().String(), err.Error())
				return err
			}

			_, err = conn.Write(mb)
			if err != nil {
				loggo.Error("sendTo Write fail: %s %s", conn.RemoteAddr().String(), err.Error())
				return err
			}

			//loggo.Info("sendTo %s %s", conn.RemoteAddr().String(), f.Type.String())
		}
	}
}
