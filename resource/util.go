package resource

import (
	"os"
	"io/ioutil"
	"net/http"
	"log"

)

// MkDirAll create dir
func MkDirAll(dirPath string) error {
	_, err := os.Stat(dirPath)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0777)
	}
	return err
}


// SaveResByURLPath save receive file
func SaveResByURLPath(urlPath string, resp *http.Response) IResource {
	fileInfo := ParseURL(urlPath)
	if fileInfo != nil {
		return SaveRes(fileInfo, resp)
	}
	return nil
}

// SaveRes save file by FileInfo
func SaveRes(fileInfo IResource, resp *http.Response) IResource {
	MkDirAll(fileInfo.GetSaveParentPath())
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("push body read error is ", err)
		return nil
	}
	dataSize := len(body)
	err = ioutil.WriteFile(fileInfo.GetSavePath(), body, 0644)
	if err != nil {
		log.Println("push body write error is ", err)
		return nil
	}
	fileInfo.SetDataSize(dataSize)
	return fileInfo
}