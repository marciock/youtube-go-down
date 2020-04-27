package main

import (
	"flag"
	"fmt"
	"log"
	//"os/user"
	"path/filepath"

	. "github.com/kkdai/youtube"
)

func main() {

	customName :=flag.String("n","youtube","Name video")
	customDir :=flag.String("d","./","Directory")
	flag.Parse()
	log.Println(flag.Args())
	
	//usr, _ := user.Current()
	//currentDir := fmt.Sprintf("%v/Vídeos/", usr.HomeDir)
	currentDir := fmt.Sprintf("%v/./", "./")
	log.Println("download to dir=", *customDir)
	y := NewYoutube(true)
	arg := flag.Arg(0)
	if err := y.DecodeURL(arg); err != nil {
		fmt.Println("err:", err)
	}
	/*if err := y.StartDownload(filepath.Join(currentDir, "dl.mp4")); err != nil {
		fmt.Println("err:", err)
	}*/
	if err := y.StartDownload(filepath.Join(currentDir, *customName+".mp4")); err != nil {
		fmt.Println("err:", err)
	}
}

