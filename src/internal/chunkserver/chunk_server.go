package chunkserver

import "sync"

// Definition for ChunkServer type

type ChunkServer struct {
	Address   string            // Address of this Chunk Server
	ChunkData map[uint64][]byte // Stores chunk data, mapped by chunk ID
	Mutex     sync.Mutex
}
