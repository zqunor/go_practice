package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// Feed 包含需要处理的数据源的信息
type Feed struct {
	Name string `json:"site"`
	URI string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds 读取并反序列化源数据文件
func RetrieveFeeds() ([]*Feed, error) {
	// 打开文件
	file, err := os.Open(dataFile) 
	if err != nil {
		return nil, err
	}

	// 当函数返回时关闭文件
	defer file.Close()

	// 将文件解码到一个切片里
	// 这个切片的每一项是一个指向Feed类型值的指针
	var feeds []*Feed
	// 之后再调用这个指针的Decode方法，传入切片的地址。
	// 之后 Decode 方法会解码数据文件，并将解码后的值以Feed类型值的形式存入切片里
	err = json.NewDecoder(file).Decode(&feeds)

	// 这个函数不需要检查错误，调用者会做这件事
	return feeds, err
}