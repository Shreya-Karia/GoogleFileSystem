package main

import (
	"github.com/Shreya-Karia/gfs/internal/master"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {
	masterInstance := &master.Master{
		FileTable:  make(map[string]*master.FileMetaData),
		ChunkTable: make(map[string]*master.ChunkMetadata),
		LeaseTime:  time.Minute,
	}

	rpc.Register(masterInstance)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

	go http.Serve(listener, nil)
	log.Println("Master server started on port 1234")
}
