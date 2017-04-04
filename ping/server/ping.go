package main

import (
	"github.com/dist-ribut-us/crypto"
	"github.com/dist-ribut-us/log"
	"github.com/dist-ribut-us/packeter"
	"github.com/dist-ribut-us/rnet"
)

const PingPort uint32 = 58080

func main() {
	log.Go()

	_, priv, err := crypto.KeyPairFromString("wyWIonPVgES739waNJESw8TIcy4PKnmHANx9LoNeWgA=", "uEaBaST5LEbWEGWJSN7tNeGS-OLehY5D0K8RcYfZgAE=")
	log.Panic(err)

	pktr := packeter.New()
	srv, err := rnet.RunNew(rnet.Port(PingPort), pktr)
	log.Panic(err)

	for msg := range pktr.Chan() {
		if log.Error(msg.Err) {
			continue
		}
		go respond(msg, pktr, srv, priv)
	}
}

func respond(msg *packeter.Package, pktr *packeter.Packeter, srv *rnet.Server, priv *crypto.XchgPriv) {
	resp := priv.Sign([]byte(msg.Addr.IP.String()))
	pkts, err := pktr.Make(resp, 0.1, 0.99)
	if log.Error(err) {
		return
	}
	srv.SendAll(pkts, msg.Addr)
}
