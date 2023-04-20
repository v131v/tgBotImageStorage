package loader

import "os"

type LoaderService struct {
	storagePath string
}

func New(storagePath string) *LoaderService {
	return &LoaderService{storagePath}
}

func (ls *LoaderService) Get(dirName string) (filePaths []string, err error) {
	dirPath := ls.storagePath + dirName + "\\"
	dirFile, err := os.Open(dirPath)
	if err != nil {
		return
	}

	defer dirFile.Close()

	filesInfo, err := dirFile.ReadDir(-1)
	if err != nil {
		return
	}

	for _, fileInfo := range filesInfo {
		filePath := dirPath + fileInfo.Name()
		filePaths = append(filePaths, filePath)
	}
	return
}
