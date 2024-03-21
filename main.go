package main

import "github.com/skip2/go-qrcode"

func main() {
	err := qrcode.WriteFile("hello this is qrcode writting via golang", qrcode.Medium, 256, "qr.png")
	if err != nil {
		panic(err)
	}
}
