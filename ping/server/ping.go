// Ping runs an http server that takes a request containing the UPD address and
// public key and attempts to contact the overlay node. This is done over TCP
// rather than UDP to prevent accidentally hole-punching.
package main

import (
	"github.com/dist-ribut-us/crypto"
	"github.com/dist-ribut-us/overlay"
	"github.com/dist-ribut-us/rnet"
	"net/http"
	"strconv"
	"time"
)

func main() {
	srv, err := overlay.NewServer(":7667")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		keyStr := r.URL.Query().Get("key")
		portStr := r.URL.Query().Get("port")
		w.Write([]byte(r.RemoteAddr))

		pub, err := crypto.PubFromString(keyStr)
		if err != nil {
			return
		}

		addr, err := rnet.ResolveAddr(r.RemoteAddr)
		if err != nil {
			return
		}
		addr.Port, err = strconv.Atoi(portStr)
		if err != nil {
			return
		}

		n := &overlay.Node{
			Pub:      pub,
			ToAddr:   addr,
			FromAddr: addr,
		}
		srv.AddNode(n)
		srv.Handshake(n)
		time.Sleep(time.Millisecond * 10)
		srv.Send([]byte("pong"), n)
	})

	http.ListenAndServe(":11111", nil)
}
