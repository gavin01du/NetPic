package main

import (
	"bytes"
	"fmt"
	"strings"
	"io"
	"os"
	"bufio"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func getAllUrls(args string) []string {

	var urls []string

	fileName := "D:\\Temp\\urls_" + args + ".txt"

	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return urls
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		// fmt.Println(line)
		urls = append(urls, line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return urls
			}
		}
	}

	return urls
}

// 下载图片
func download(img_url string, _dir string) int {

	file_name := _dir + "\\" + img_url[strings.LastIndex(img_url,"/")+1:len(img_url)]
	// fmt.Println(file_name)
	fmt.Println("get file : " + img_url)
	t1 := time.Now()
	resp, err := http.Get(img_url)
	fmt.Println("get http response, time : ",time.Since(t1))
	if err != nil {
		fmt.Println("get file by http error")
		return 0
	}

	t2 := time.Now()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("get body, time : " ,time.Since(t2))
	if err != nil {
		fmt.Println("get body error")
		return 0
	}

	t3 := time.Now()
	out, _ := os.Create(file_name)
	io.Copy(out, bytes.NewReader(body))
	fmt.Println("create file, time : ", time.Since(t3))
	fmt.Println(file_name)
	fmt.Println("get file finished, time : ", time.Since(t1))

	return 0
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main()  {
	start := time.Now()

	for idx, args := range os.Args {
		fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
		if (len(args) > 10){
			continue
		}
		urls := getAllUrls(args)

		_dir := "D:\\Temp\\" + args
		exist, err := PathExists(_dir)
		if err != nil {
			fmt.Printf("get dir error![%v]\n", err)
			return
		}

		if exist {
			fmt.Printf("has dir![%v]\n", _dir)
		} else {
			fmt.Printf("no dir![%v]\n", _dir)
			// 创建文件夹
			err := os.Mkdir(_dir, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir failed![%v]\n", err)
			} else {
				fmt.Printf("mkdir success!\n")
			}
		}

		for _, url := range urls {
			download(url,_dir)
			//fmt.Println(url)
		}

		fmt.Println("url count : " + strconv.Itoa(len(urls)))
		fmt.Println("All task finished, time : ", time.Since(start))
	}

}