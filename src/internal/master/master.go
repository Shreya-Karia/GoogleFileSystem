package master

import (
	"sync"
	"time"
)

// Definition for Master type

type Master struct {
	FileTable  map[string]*FileMetaData  // Maps file names to their metadata
	ChunkTable map[string]*ChunkMetadata // Maps chunk IDs to their metadata
	LeaseTime  time.Duration             // Duration for which a lease is granted
	Mutex      sync.Mutex                // Mutex for thread-safe access to metadata
}
