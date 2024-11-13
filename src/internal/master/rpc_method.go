package master

import "errors"

// Method for client communications

func (master *Master) GetFileMetaData(fileName string, fileMetaData *FileMetaData) error {
	master.Mutex.Lock()
	defer master.Mutex.Unlock()

	fileMeta, exists := master.FileTable[fileName]

	if !exists {
		return errors.New("File not found")
	}

	*fileMetaData = *fileMeta
	return nil
}

// Method for creating a new file

func (master *Master) CreateFile(fileName string, fileMetaData *FileMetaData) error {
	master.Mutex.Lock()
	defer master.Mutex.Unlock()

	if _, exists := master.FileTable[fileName]; exists {
		return errors.New("File already exists")
	}

	newFileMetaData := &FileMetaData{
		FileName: fileName,
		ChunkIDs: []uint64{},
	}

	master.FileTable[fileName] = newFileMetaData
	*fileMetaData = *newFileMetaData

	return nil
}
