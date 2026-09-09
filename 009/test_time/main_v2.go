package main

import (
	"fmt"
	"io"
	"math"
	"mime"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"strconv"
	"sync"
	"time"
)

var DEBUG bool = true
var mx sync.Mutex = sync.Mutex{}

func GetFileSize(url string) (int64, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}

	if DEBUG {
		reqDump, _ := httputil.DumpRequestOut(req, false)
		fmt.Println("REQUEST:")
		fmt.Println(string(reqDump))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if DEBUG {
		respDump, _ := httputil.DumpResponse(resp, false)
		fmt.Println("RESPONSE:")
		fmt.Println(string(respDump))
	}

	return resp.ContentLength, nil
}

func GetFilenameFromURL(urlStr string) string {
	resp, err := http.Head(urlStr)
	if err == nil {
		defer resp.Body.Close()

		if cd := resp.Header.Get("Content-Disposition"); cd != "" {
			_, params, err := mime.ParseMediaType(cd)
			if err == nil {
				if fn := params["filename"]; fn != "" {
					return fn
				}
			}
		}
	}

	u, err := url.Parse(urlStr)
	if err == nil {
		name := path.Base(u.Path)
		if name != "" && name != "/" {
			return name
		}
	}

	return "downloaded_file"
}

func PartialDownload(start, end int, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Download(start, end int, url string, out *os.File, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := PartialDownload(start, end, url)
	if err != nil {
		fmt.Printf("Fail Download: %d-%d\n", start, end)
		return
	}
	defer resp.Body.Close()

	mx.Lock()
	defer mx.Unlock()
	_, err = out.Seek(int64(start), 0)
	if err != nil {
		fmt.Println("Seek Error:", err)
		return
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Write Error:", err)
	}
}

func main() {
	NOW := time.Now()

	if len(os.Args) != 3 {
		fmt.Printf("USAGE : %s <url> <handler>\n", os.Args[0])
		return
	}
	url := os.Args[1]
	handler, err := strconv.Atoi(os.Args[2])

	fileSize, err := GetFileSize(url)
	if err != nil {
		fmt.Println("GetFileSize error:", err)
		return
	}

	chunk := int(math.Ceil(float64(fileSize) / float64(handler)))

	fmt.Println("url:", url)
	fmt.Println("chunk:", chunk)
	fmt.Println("file_size:", fileSize)
	fmt.Println("chunk_count:", handler)

	out, err := os.Create(GetFilenameFromURL(url))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	// pre-allocate file
	out.Seek(fileSize-1, 0)
	out.Write([]byte{0})

	var wg sync.WaitGroup

	for i := 0; i < handler; i++ {

		start := i * chunk
		end := (i+1)*chunk - 1

		if end >= int(fileSize) {
			end = int(fileSize) - 1
		}

		fmt.Printf("Range: %d-%d\n", start, end)

		wg.Add(1)
		go Download(start, end, url, out, &wg)
	}

	wg.Wait()

	fmt.Println("Download Completed!!!")

	fmt.Println("\n\nElapsed Time :", time.Since(NOW))

}
