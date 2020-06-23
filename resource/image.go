package resource 

// Image resource type image
type Image struct {
	resName   string
	resType   string
	dataSize  int
	localPath string
	urlPath   string
}

const (
	_URLPrefix string ="/image/"
	_ResDir string ="../resdata/image/"
	_RecDir string ="../receive/image/"
)

func getURLPath(relativePath string) string {
	return _URLPrefix + relativePath
}

func getLocalPath(relativePath string) string {
	return _ResDir + relativePath 
}


// NewImageByURL new image by parse image
func NewImageByURL(urlComs []string) IResource {
	if comsSize := len(urlComs); comsSize == 2 {
		//https://localhost:8080/image/0.png
		// 0: image
		// 1: 0.png
		resName := urlComs[1]
		return NewImage(resName)
	}
	return nil
}

// NewImage new image type resource
func NewImage(resName string) *Image {
	return &Image{
		resName:   resName,
		resType: Img,
		localPath: getLocalPath(resName),
		urlPath:   getURLPath(resName),
	}
}



// GetLocalPath for avc
func (res *Image) GetLocalPath() string {
	return res.localPath
}

// GetURLPath for avc
func (res *Image) GetURLPath() string {
	return res.urlPath
}

// GetFileName for avc
func (res *Image) GetFileName() string {
	return res.resName
}

//GetSaveParentPath 获取保存父路径
func (res *Image) GetSaveParentPath() string {
	return _RecDir + "/" + res.resName
}

//GetSavePath 获取保持路径
func (res *Image) GetSavePath() string {
	return res.GetSaveParentPath() + "/" + res.resName
}

// SetDataSize for avc
func (res *Image) SetDataSize(size int) {
	res.dataSize = size
}

// GetDataSize for avc
func (res *Image) GetDataSize() int {
	return res.dataSize
}