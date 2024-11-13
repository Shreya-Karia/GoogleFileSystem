package master

import "time"

// Definition for FileMetaData type

type FileMetaData struct {
	FileName string   // Name of the file
	ChunkIDs []uint64 // List of Chunk IDs associated with the file
}

// Definition for ChunkMetaData type

type ChunkMetadata struct {
	ChunkID     uint64    // Unique chunk identifier
	Replicas    []string  // Addresses of the Chunk Servers storing this chunk
	Version     uint64    // Version number for consistency
	LeaseExpiry time.Time // Lease expiration time for the primary replica
}
