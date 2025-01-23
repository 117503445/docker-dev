package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/117503445/goutils"
)

func fastestDownloader(mirrors map[string]string) (string, error) {
	var wg sync.WaitGroup
	fastest := make(chan string, 1)
	errors := make(chan error, len(mirrors))

	for name, url := range mirrors {
		wg.Add(1)
		go func(n string, u string) {
			defer wg.Done()
			start := time.Now()
			resp, err := http.Get(u)
			if err != nil {
				errors <- fmt.Errorf("error downloading from %s: %v", n, err)
				return
			}
			defer resp.Body.Close()

			// Read the whole body to actually download the file.
			_, err = io.Copy(io.Discard, resp.Body)
			if err != nil {
				errors <- fmt.Errorf("error reading response body from %s: %v", n, err)
				return
			}

			duration := time.Since(start)
			select {
			case fastest <- n:
			default:
			}
			fmt.Printf("Downloaded from %s in %v\n", n, duration)
		}(name, url)
	}

	go func() {
		wg.Wait()
		close(fastest)
	}()

	select {
	case mirror := <-fastest:
		return mirror, nil
	case err := <-errors:
		return "", err
	}
}

func main() {
	goutils.InitZeroLog()

	mirrors := map[string]string{
		// "USTC":     "https://mirrors.ustc.edu.cn/archlinux/core/os/x86_64/gcc-go-14.2.1%2Br134%2Bgab884fffe3fc-2-x86_64.pkg.tar.zst",
		"PkgBuild": "https://geo.mirror.pkgbuild.com/core/os/x86_64/gcc-go-14.2.1%2Br134%2Bgab884fffe3fc-2-x86_64.pkg.tar.zst",
	}

	mirror, err := fastestDownloader(mirrors)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Fastest mirror is:", mirror)
	}
}
