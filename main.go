package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

const (
	WallPaperPath string = "./wallpapers/"
	WallPaperUrl  string = "/wallpapers/"
	ListenPort    string = ":1234"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(WallPaperPath))))

	log.Printf("服务端即将启动\n")
	log.Printf("电脑测试地址：http://127.0.0.1%s\n", ListenPort)

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalln("Network interface Error:", err.Error())
		os.Exit(1)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() {
			if ipnet.IP.To4() != nil {
				log.Printf("PS4访问地址：http://%s%s", ipnet.IP.String(), ListenPort)
			}
		}
	}

	if err := http.ListenAndServe(ListenPort, nil); err != nil {
		log.Fatalln("WallPaper Server Start Error: ", err)
	}
}
