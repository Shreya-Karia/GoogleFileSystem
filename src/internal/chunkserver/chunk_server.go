package chunkserver

import "sync"

// Definition for ChunkServer type

type ChunkServer struct {
	Address   string            // Address of this Chunk Server
	ChunkData map[uint64][]byte // Stores chunk data, mapped by chunk ID
	Mutex     sync.Mutex
}

// Method to store new chunk

func (chunkServer *ChunkServer) storeChunkServer(chunkID uint64, chunkData []byte, isChunkCreated *bool) error {
	chunkServer.Mutex.Lock()
	defer chunkServer.Mutex.Unlock()

	chunkServer.ChunkData[chunkID] = chunkData
	*isChunkCreated = true

	return nil
}
