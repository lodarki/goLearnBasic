package main

import (
	"io"
	"os"
	"strings"
)

//一个常见模式是 io.Reader 包裹另一个 `io.Reader`，然后通过某种形式修改数据流。
//例如，gzip.NewReader 函数接受 `io.Reader`（压缩的数据流）并且返回同样实现了 io.Reader 的 `*gzip.Reader`（解压缩后的数据流）。
//编写一个实现了 io.Reader 的 `rot13Reader`， 并从一个 io.Reader 读取， 利用 rot13 代换密码对数据流进行修改。
//已经帮你构造了 rot13Reader 类型。 通过实现 Read 方法使其匹配 `io.Reader`。
type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	rot13 := Rot13{map[byte]byte{}}
	rot13.init()
	if err == nil {
		for index, element := range b[:n] {
			replace := rot13.rotMap[element]
			if replace > 0 {
				element = replace
			}
			b[index] = element
		}
	}
	return n, err
}

type Rot13 struct {
	rotMap map[byte]byte
}

func (rot13Map *Rot13) init() {
	input := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	output := []byte{'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm'}
	i := len(input)
	for j := 0; j < i; j++ {
		rot13Map.rotMap[input[j]] = output[j]
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
