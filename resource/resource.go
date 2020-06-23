package resource

import ("strings")


// IResource resource interface
type IResource interface {
	//获取本地文件路径
	GetLocalPath() string
	//获取URL路径
	GetURLPath() string
	
	//获取保存父路径
	GetSaveParentPath() string
	//获取保存路径
	GetSavePath() string

	//获取资源文件名
	GetFileName() string

	//设置文件大小
	SetDataSize(int)
	//获取文件大小
	GetDataSize() int
}

const (
	// Img example type
	Img string = "image"
)

// IResParser resource request parser
type IResParser interface {
	ParseURL(string) IResource
}

var parserMap map[string]func([]string) IResource

func init() {
	parserMap = make(map[string]func([]string) IResource)
	parserMap[Img] = NewImageByURL
}
// ParseURL 从req从获取起源
func ParseURL(urlPath string) IResource {
	urlPath = strings.Trim(urlPath, "/")
	//以/拆分url
	components := strings.Split(urlPath, "/")
	if len(components) > 0 {
		//第一个代表请求类型
		resourceType := strings.ToLower(components[0])
		parser, ok := parserMap[resourceType]
		if ok {
			return parser(components)
		}
	}

	return nil
}

// ParsePromiseList parse promise list resource
func ParsePromiseList(promiseList string) []IResource{
	result := []IResource{}
	//以,拆分url
	components := strings.Split(promiseList, ",")
	for _, item:= range components {
		comps := strings.Split(item, ".")
		if length :=len(comps);length > 1{
			resFileType := comps[length-1]
			switch resFileType {
				case "png","jpeg","jpg":
					newImage := NewImage(item)
					if newImage !=nil {
						result = append(result, newImage)
					}
			}
		}
	}
	return result
}