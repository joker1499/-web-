package liner

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

func Extract(file File) {
	defer file.Close()

	r := bufio.NewReader(file)

	for {
		data2, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read err", err.Error())
			break
		}

		go Fire(data2, "突发公共事件分类、分级与分期")
		time.Sleep(time.Second)
	}

}
