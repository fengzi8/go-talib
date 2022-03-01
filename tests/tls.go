package main

import (
    // "log"
    // "net/http"
	"crypto/tls"
	"fmt"
	"time"
)

func main() {
	_, err := tls.Dial("tcp", "example.com:80", nil)
	if err != nil {
		panic("Server doesn't support SSL certificate err: " + err.Error())
	}
	_, err1 := tls.Dial("tcp", "https://golang.org/", nil)
	if err1 != nil {
		panic("Server doesn't support SSL certificate err: " + err.Error())
	}
    // _, err := http.Get("https://golang.org/")
    // if err != nil {
    //     log.Fatal(err)
    // }
	// else
	// {
	// 	log.Info(r)
	// }

	conn, err := tls.Dial("tcp", "blog.umesh.wtf:443", nil)
	if err != nil {
		panic("Server doesn't support SSL certificate err: " + err.Error())
	}

	err = conn.VerifyHostname("blog.umesh.wtf")
	if err != nil {
		panic("Hostname doesn't match with certificate: " + err.Error())
	}
	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	fmt.Printf("Issuer: %s\nExpiry: %v\n", conn.ConnectionState().PeerCertificates[0].Issuer, expiry.Format(time.RFC850))

}

