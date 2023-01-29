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

func Extract(file File, qikan string) {
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

		go Fire(data2, qikan)
		time.Sleep(time.Second)
	}

}
