package main

import (
	"flag"
	"log"
)

func main(){
	var fatalErr error
	defer func() {
		if fatalErr != nil{
			log.Fatalln(fatalErr)
		}
	}()
	var(
		interval = flag.Int("interval", 10, "チェックの間隔（秒単位）")
		archive = flag.String("archive", "archive", "アーカイブの保存先")
		dbpath = flag.String("db", "./db", "filedbデータベースへのパス")
	)
	flag.Parse()
}
