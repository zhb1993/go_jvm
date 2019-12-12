package classpath

import "errors"
import "strings"

type CompositeEntry []Entry

/**
	把参数按 ;分割成数组，然后遍历数组把每个路径都转成具Entry实例
 */
func newCompositeEntry(pathList string) CompositeEntry{
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator){
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry

}

//遍历每一个子路径的readclass 方法，如果读取到class数据，则返回数据
func (self CompositeEntry) readClass(className string)([]byte, Entry, error)  {
	for _, entry :=range self{
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, err
		}
	}
	return nil, nil,errors.New("class not found: "+ className)

}

//把每个子路径用分隔符拼接起来
func (self CompositeEntry) String() string  {
	strs := make([]string, len(self))
	for i, entry :=range self{
		strs[i] = entry.String()
	}
	return strings.Join(strs,pathListSeparator)
}
