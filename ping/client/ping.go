package main

import (
	"bufio"
	"fmt"
	"github.com/dist-ribut-us/natt/igdp"
	"github.com/dist-ribut-us/overlay"
	"github.com/dist-ribut-us/rnet"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	// start overlay server
	port := fmt.Sprintf(":%d", rnet.RandomPort())
	overlayNode, err := overlay.NewServer(port)
	if err != nil {
		panic(err)
	}

	// natt
	err = igdp.Setup()
	if err != nil {
		panic(err)
	}
	_, err = igdp.AddPortMapping(overlayNode.Port(), overlayNode.Port())
	if err != nil {
		fmt.Println(err)
	}

	// delay to give user a chance to approve connection
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press Enter to send: ")
	reader.ReadString('\n')

	// send request
	url := fmt.Sprintf("http://dist.ribut.us:11111/?port=%d&key=%s", overlayNode.Port(), url.QueryEscape(overlayNode.PubStr()))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	r, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(r))

	// print response
	msg := <-overlayNode.Chan()
	fmt.Println(string(msg.Body))
}
