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

func HttpGetDb(url string) (result string, err error) {
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

func SpiderPageDb(idx int) {

	if idx <= 820 {
		url := "https://news.fzu.edu.cn/news/info/1002/" + strconv.Itoa(23578+idx) + ".htm"
		url2 := "https://news.fzu.edu.cn/system/resource/code/news/click/dynclicks.jsp?clickid=" + strconv.Itoa(23578+idx) + "&owner=1744991928&clicktype=wbnews"

		result1, err2 := HttpGetDb(url2)
		if err2 != nil {
			fmt.Println("HttpGetDb err:", err2)
			return
		}
		Ret := regexp.MustCompile(`.*`)
		if Ret == nil {
			err2 = fmt.Errorf("%s", "MustCompile err")
			return
		}
		Amount := Ret.FindAllStringSubmatch(result1, -1)

		result, err := HttpGetDb(url)
		if err != nil {
			fmt.Println("HttpGetDb err:", err)
			return
		}

		//fmt.Println("result=", result)
		ret1 := regexp.MustCompile(`28px;">(?s:(.*?))</p>`)
		if ret1 == nil {
			err = fmt.Errorf("%s", "MustCompile err")
			return
		}
		NewsName := ret1.FindAllStringSubmatch(result, -1)

		ret2 := regexp.MustCompile(`<span id="fbsj">(?s:(.*?))</span>`)
		if ret2 == nil {
			err = fmt.Errorf("%s", "MustCompile err")
			return
		}
		NewsTime := ret2.FindAllStringSubmatch(result, -1)

		ret3 := regexp.MustCompile(`<span id="author">(?s:(.*?))</span>`)
		if ret3 == nil {
			err = fmt.Errorf("%s", "MustCompile err")
			return
		}
		NewsAuthor := ret3.FindAllStringSubmatch(result, -1)

		ret4 := regexp.MustCompile(`<META Name="description" Content="(?s:(.*?))" />`)
		if ret4 == nil {
			err = fmt.Errorf("%s", "MustCompile err")
			return
		}
		NewsWork := ret4.FindAllStringSubmatch(result, -1)
		save3File(NewsName, NewsTime, NewsAuthor, NewsWork, Amount)
	} else {
		url := "https://news.fzu.edu.cn/info/1011/" + strconv.Itoa(22202+idx) + ".htm"
		url2 := "https://news.fzu.edu.cn/system/resource/code/news/click/dynclicks.jsp?clickid=" + strconv.Itoa(22202+idx) + "&owner=1779559075&clicktype=wbnews"

		result1, err2 := HttpGetDb(url2)
		if err2 != nil {
			fmt.Println("HttpGetDb err:", err2)
			return
		}
		Ret := regexp.MustCompile(`.*`)
		if Ret == nil {
			err2 = fmt.Errorf("%s", "MustCompile err")
			return
		}
		Amount := Ret.FindAllStringSubmatch(result1, -1)

		result, err := HttpGetDb(url)
		if err != nil {
			fmt.Println("HttpGetDb err:", err)
			return
		}

		ret1 := regexp.MustCompile(`<title>(?s:(.*?))</title>`)
		if ret1 == nil {
			err = fmt.Errorf("%s", "MustCompile err")
			return
		}
		NewsName := ret1.FindAllStringSubmatch(result, -1)

		ret2 := regexp.MustCompile(`<span>发布日期:  (?s:(.*?))</span>`)
		if ret2 == nil {
			err = fmt.Errorf("%s", "MustCompile err")
			return
		}
		NewsTime := ret2.FindAllStringSubmatch(result, -1)

		ret3 := regexp.MustCompile(`<span>作者： (?s:(.*?))</span>`)
		if ret3 == nil {
			err = fmt.Errorf("%s", "MustCompile err")
			return
		}
		NewsAuthor := ret3.FindAllStringSubmatch(result, -1)

		ret4 := regexp.MustCompile(`<META Name="description" Content="(?s:(.*?))" />`)
		if ret4 == nil {
			err = fmt.Errorf("%s", "MustCompile err")
			return
		}
		NewsWork := ret4.FindAllStringSubmatch(result, -1)
		save3File(NewsName, NewsTime, NewsAuthor, NewsWork, Amount)
	}

}

func save3File(NewsName, NewsTime, NewsAuthor, NewsWork, Amount [][]string) {
	n := len(NewsWork)
	filePath := "demo01.txt"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE, 0775)
	if err != nil {
		fmt.Println("文件打开失败")
	}

	defer file.Close()
	write := bufio.NewWriter(file)
	for i := 0; i < n; i++ {
		write.WriteString(NewsTime[i][1] + "\t" + NewsAuthor[i][1] + "\t\t" + Amount[i][0] + "\t" + NewsName[i][1] + "\n" + NewsWork[i][1] + "\n")
	}
	write.Flush()
}
func ToWork(start, end int) {

	for i := start; i <= end; i++ {
		SpiderPageDb(i)
	}

}

func main() {
	var start, end int

	start = 1
	end = 3500

	ToWork(start, end)

}
