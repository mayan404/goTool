package myos

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)


//一次性读取
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}


	return ioutil.ReadAll(f)
}

func processBlock(line []byte) {
	os.Stdout.Write(line)
}

//分块读取
//可在速度和内存占用之间取得很好的平衡。
func ReadBlock(filePth string, bufSize int, hookfn func([]byte)) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, bufSize) //一次读取多少个字节
	bfRd := bufio.NewReader(f)
	for {
		n, err := bfRd.Read(buf)
		hookfn(buf[:n]) // n 是成功读取字节数

		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}

	return nil
}

func ReadBlockTest()  {
	ReadBlock("test.txt", 10000, processBlock)
}





func processLine(line []byte) {
	os.Stdout.Write(line)
}
//逐行读取
//逐行读取有的时候真的很方便，性能可能慢一些，但是仅占用极少的内存空间。
func ReadLine(filePth string, hookfn func([]byte)) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		hookfn(line) //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func ReadLineTest()  {
	ReadLine("test.txt", processLine)
}
