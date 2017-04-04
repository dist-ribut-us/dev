package main

import (
	"github.com/dist-ribut-us/crypto"
	"github.com/dist-ribut-us/log"
	"github.com/dist-ribut-us/packeter"
	"github.com/dist-ribut-us/rnet"
)

func main() {
	log.Go()

	pktr := packeter.New()
	srv, err := rnet.RunNew(rnet.RandomPort(), pktr)
	log.Panic(err)

	pubkey, err := crypto.XchgPubFromString("wyWIonPVgES739waNJESw8TIcy4PKnmHANx9LoNeWgA=")
	log.Panic(err)
	addr, err := rnet.ResolveAddr("adamcolton.net:58080")
	log.Panic(err)
	log.Info(log.Lbl("pinging"), addr, log.Lbl("from"), srv.Port())

	pkts, err := pktr.Make([]byte{}, 0.0001, .7)
	log.Panic(err)
	srv.SendAll(pkts, addr)

	msg := <-pktr.Chan()
	if ip, ok := pubkey.Verify(msg.Body); !ok {
		log.Info(log.Lbl("bad_sig"))
	} else {
		log.Info("my_ip", string(ip))
	}

}
