package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	webSocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	clients ClientList
	sync.RWMutex
}

func newManager() *Manager {
	return &Manager{
		clients: make(ClientList),
	}
}

func (m *Manager) serveWS() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("new connection")
		// upgrade regular http connection into websocket
		conn, err := webSocketUpgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		client := NewClient(conn,m)
		m.addClient(client)

	}
}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	
	m.clients[client] = true
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	if _,ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients,client)
	}
}
