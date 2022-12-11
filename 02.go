package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func ToWork1(start, end int) {
	for i := start; i <= end; i++ {
		SpiderPageb(i)
	}
}

func save4File(Comment [][]string) {
	n := len(Comment)
	filePath := "demo02.txt"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE, 0775)
	if err != nil {
		fmt.Println("文件打开失败")
	}

	defer file.Close()
	write := bufio.NewWriter(file)
	for i := 0; i < n; i++ {
		write.WriteString(Comment[i][1] + "\n")
	}
	write.Flush()
}
func HttpGetb(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}

func SpiderPageb(i int) {
	url := "https://api.bilibili.com/x/v2/reply/main?=jQuery172043660583075908077_1670703333803&jsonp=jsonp&next=" + strconv.Itoa(i-1) + "&type=1&oid=21071819&mode=3&plat=1&_=1670703353259"
	result, err := HttpGetb(url)
	if err != nil {
		fmt.Println("HttpGetDb err:", err)
		return
	}
	ret1 := regexp.MustCompile(`:{"message":"(?s:(.*?))",`)
	if ret1 == nil {
		err = fmt.Errorf("%s", "MustCompile err")
		return
	}
	//fmt.Println(result)
	Comment := ret1.FindAllStringSubmatch(result, -1)

	save4File(Comment)

}

func main() {
	var start, end int
	start = 1
	end = 1451

	ToWork1(start, end)
}
