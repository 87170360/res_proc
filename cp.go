package main

import (
	"path/filepath"
	"strings"
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"encoding/json"
	"crypto/md5"
)
var (
	conf *configInfo
)

type configInfo struct {
	Input1 string `json:"ui_path"`
	Input2 string `json:"tradition_res"`
	Input3 string `json:"img_path"`
	Output string `json:"output_path"`
	DropFile []string `json:"drop_file"`
	RepeatFile []string `json:"repeat_file"`
	NotWarn []string `json:"not_warn"`
}

func loadConf() (*configInfo, error) {
	data, err := ioutil.ReadFile("conf.json")
    if err != nil {
        return nil, err
    }

	ci := &configInfo{}
    err = json.Unmarshal(data, &ci)
    if err != nil {
        return nil, err
    }
    return ci, nil
}

func getFileList(path string) []string {
	fileList := []string{}
    err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
        if strings.HasSuffix(path, ".png"){
            fileList = append(fileList, path)
        }
        return nil
    })

    if err != nil {
        fmt.Println(err)
		panic("err")
    }

	return fileList
}

func getTarget(fp string) string {

	base := filepath.Base(fp)
	var srcdir string
	if base[:3] == "ui_" {
		srcdir = conf.Input1
	} else {
		srcdir = conf.Input3
	}
	srcdir = filepath.Dir(srcdir)

	return filepath.Join(srcdir, fp[len(conf.Output):])
}

func getHash(fp string) (string, error) {
    f, err := os.Open(fp)
    if err != nil {
        fmt.Println(err)
        return "", err
    }
    defer f.Close()

    h := md5.New()
    if _, err := io.Copy(h, f); err != nil {
        fmt.Println(err)
        return "", err
    }

    return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func main() {
	var err error
	conf, err = loadConf()
	if err != nil {
		fmt.Println(err)
		return
	}

	num := 0
	files := getFileList(conf.Output)
	for _, v := range files {
		h1, err := getHash(v)
		if err != nil {
			return
		}
		tar := getTarget(v)
		if _, err := os.Stat(tar); os.IsNotExist(err) {
			fmt.Printf("target noexist %s\n", tar)
			continue
		}
		h2, err := getHash(tar)
		if err != nil {
			return
		}
		if h1 != h2 {
			fmt.Printf("%s\ncopy to\n%s", v, tar)
			num++
			copy(v, tar)
		}
	}
	fmt.Printf("copy num:%d\n", num)
}
