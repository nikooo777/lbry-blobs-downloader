package example

import (
	"fmt"
	"os"

	"github.com/nikooo777/lbry-blobs-downloader/downloader"
	"github.com/nikooo777/lbry-blobs-downloader/shared"
)

// if you want to use the blobsdownloader as a library you can follow this example

func MySoftware() {
	sdHash := "c333e168b1adb5b8971af26ca2c882e60e7a908167fa9582b47a044f896484485df9f5a0ada7ef6dc976489301e8049d"
	downloadServer := "blobcache-eu.lbry.com"
	UDPPort := 5568
	TCPPort := 5567
	HTTPPort := 5569

	//static, it's ugly but it works for now
	shared.ReflectorPeerServer = fmt.Sprintf("%s:%d", downloadServer, TCPPort)
	shared.ReflectorQuicServer = fmt.Sprintf("%s:%d", downloadServer, UDPPort)
	shared.ReflectorHttpServer = fmt.Sprintf("%s:%d", downloadServer, HTTPPort)
	err := os.MkdirAll("./mypersonaldownloads/", os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = downloader.DownloadAndBuild(sdHash, false, downloader.HTTP, "jeremy.mp4", "./mypersonaldownloads/")
	if err != nil {
		panic(err)
	}
}
