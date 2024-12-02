package main

import protocolbytes "github.com/mateusfdl/protocol-bytes"

func SendHelloMeassage() protocolbytes.Buffer {
	buf := protocolbytes.Buffer{}
	buf.WInt8(0)
	buf.WUTF("Matheus")
	buf.WUTF("Hello, World!")

	return buf
}

func SendThumbsUpMessage() protocolbytes.Buffer {
	buf := protocolbytes.Buffer{}
	buf.WInt8(1)
	buf.WUTF("Matheus")
	buf.WInt8(2)

	return buf
}
