package websocket

import (
	"errors"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"message/websocket/netpoll"
	"net"
	"net/url"
	"sync"
	"syscall"
)

var entityMap sync.Map
var unauthEntityMap sync.Map

func init() {
	entityMap = sync.Map{}
	unauthEntityMap = sync.Map{}
}

var (
	RejectAuthError =  ws.RejectConnectionError(ws.RejectionReason("授权失败"))
)

// 启动
func Setup(port int, authFilter func(token string) bool,onConnected func(iden string),  onMessage func(ident string, msg string,conn net.Conn) , onClose func(ident string)) {
	releaseLimitations()
	addr := net.TCPAddr{
		IP:   nil,
		Port: port,
		Zone: "",
	}
	ln, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		panic(err)
	} else {
		log.Printf("[message-websocket] Listening tcp on 0.0.0.0:%d\n", port)
	}
	poller, u, token := configPool(authFilter)
	dealConn(ln, u, poller, token, onConnected, onMessage, onClose)
}

// 设置最大文件数量 Mac可能会修改出错 linux上正常
func releaseLimitations() {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		//panic(err)
	}
}

// 配置epoll（linux） freeBSD（kqueue）
func configPool(authFilter func (t string) bool) (poller netpoll.Poller, u ws.Upgrader, token *string) {
	_t := ""
	token = &_t
	onError := func(err error) {
		fmt.Println(err)
	}
	config := netpoll.Config{OnWaitError: onError}
	poller, err := netpoll.New(&config)

	if err != nil {
		log.Fatal(err)
	}

	u = ws.Upgrader{
		OnHeader: func(key, value []byte) (err error) {
			return
		},
		OnRequest: func(uri []byte) (err error) {
			v, err := url.Parse(string(uri))
			if err != nil {
				fmt.Println(err)
			}
			m, err := url.ParseQuery(v.RawQuery)
			if err != nil || m == nil {
				fmt.Println(m)
			} else {
				t := m.Get("token")
				*token = t
				if t != "" && authFilter != nil && !authFilter(t) {
					return RejectAuthError
				}
			}
			return
		},
	}
	return
}

// 处理连接
func dealConn(ln net.Listener, u ws.Upgrader, poller netpoll.Poller, token *string,onConnected func(iden string),  onMessage func(ident string, msg string,conn net.Conn) , onClose func(ident string)) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		_, err = u.Upgrade(conn)

		if err != nil {
			if err != RejectAuthError {
				panic(err)
			} else {
				continue
			}
		}

		desc := netpoll.Must(netpoll.HandleRead(conn))

		if err = poller.Start(desc, func(e netpoll.Event) {
			if e&netpoll.EventReadHup != 0 {
				_ = poller.Stop(desc)
				_ = conn.Close()
				if *token == "" {
					unauthEntityMap.Delete(conn)
				} else {
					entityMap.Delete(*token)
				}
				return
			}

			msg, wsOpsCode, err := wsutil.ReadClientData(conn)
			if err != nil {
				fmt.Println("发生错误：", err)
				_ = poller.Stop(desc)
				_ = conn.Close()
				onClose(*token)
				if *token == "" {
					unauthEntityMap.Delete(conn)
				} else {
					entityMap.Delete(*token)
				}
				return
			}
			if wsOpsCode == ws.OpText { // 文本消息
				onMessage(*token, string(msg), conn)
			} else if wsOpsCode == ws.OpBinary { // 二进制的消息
				log.Println("发送了二进制数据")
			}
		}); err != nil {
			panic(err)
		}

		// 添加到map中
		if *token == "" {
			unauthEntityMap.Store(conn, true)
			onConnected(*token)
		} else {
			_, ok := entityMap.Load(*token)
			if ok {
				if e := wsutil.WriteServerText(conn, []byte(fmt.Sprintf("token: %s 已经存在", *token))); e != nil {
					fmt.Println(e)
				}
			} else {
				entityMap.Store(*token, conn)
			}
			onConnected(*token)
		}
	}
}

// 给固定id发送消息
func SendMsgToId(ident string, msg string) error {
	v, ok := entityMap.Load(ident)
	if !ok {
		return  errors.New(fmt.Sprintf("identifier: %s不存在", ident))
	}
	conn, ok := v.(net.Conn)
	if !ok {
		return  errors.New(fmt.Sprintf("类型错误 %T", v))
	}
	if err := wsutil.WriteServerText(conn, []byte(msg)); err != nil {
		return  err
	}
	return nil
}
