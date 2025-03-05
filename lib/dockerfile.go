package lib

import (
	"os"
)

// EmitDockerfile 这个库函数主要的目的就是用来去生成一个dockerfile 模板
func EmitDockerfile(path string) error {
	// 首先构造出对应的一个dockerfile 模板
	file, err := os.OpenFile(path,
		os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString("FROM\n")
	file.WriteString("WORKDIR\n")
	file.WriteString("COPY\n")
	file.WriteString("RUN\n")
	file.WriteString("CMD\n")
	return nil
}
