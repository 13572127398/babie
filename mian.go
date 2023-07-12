package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./filedemo.go")
	if err != nil {
		fmt.Println("open file failed!,err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}

		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

/*
func main() {
	file, err := os.Open("./filedemo.go")
	if err != nil {
		fmt.Println("open file failed!,err:", err)
	}
	defer file.Close()
	var tmp = make([]byte, 128)
	var buffer []byte
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed!,err:", err)
			break
		}
		buffer = append(buffer, tmp[:n]...)
	}
	fmt.Println(string(buffer))

}
*/
