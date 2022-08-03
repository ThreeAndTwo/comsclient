package rpc

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type Websocket struct {
	rawurl string
}

func NewWebsocket(rawurl string) *Websocket {
	return &Websocket{rawurl: rawurl}
}

func (ws *Websocket) Dial() (*websocket.Conn, error) {
	return ws.DialContext(context.Background())
}

// DialContext set default context
func (ws *Websocket) DialContext(ctx context.Context) (*websocket.Conn, error) {
	_ws, _, err := websocket.DefaultDialer.Dial(ws.rawurl, nil)
	if err != nil {
		return nil, err
	}
	return _ws, nil
}

func readMessage(_ws *websocket.Conn) {
	go func() {
		for {
			wErr := _ws.WriteMessage(websocket.BinaryMessage, []byte("ping"))
			if wErr != nil {
				fmt.Printf("Write message error: %s \n", wErr)
				continue
			}
			time.Sleep(time.Second * 3)
		}
	}()

	for {
		_, data, err := _ws.ReadMessage()
		if err != nil {
			fmt.Printf("Read message error: %s \n", err)
			continue
		}
		fmt.Println("receive:", string(data))
	}
}
