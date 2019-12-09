package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	absDir string
}

/**
类似构造函数
 */
func newDirEntry(path string) *DirEntry  {
	absDir , err :=filepath.Abs(path)
	if err != nil{
		panic(err)
	}
	//创建DirEntry实例并返回
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string)([]byte, Entry, error)  {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string  {
	return self.absDir
}