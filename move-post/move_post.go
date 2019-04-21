package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const pthSep = string(os.PathSeparator)

// 获取指定目录下的 md 文档
func getSourceFileWithMd(sDirPath string) (files []string, err error) {
	sDir, err := ioutil.ReadDir(sDirPath)
	checkErr(err)
	for _, f := range sDir {
		// 当前文件|目录路径
		chFile := sDirPath + pthSep + f.Name()
		// 如果是目录，递归查找子目录文件
		if f.IsDir() {
			chFiles, err := getSourceFileWithMd(chFile)
			checkErr(err)
			files = append(files, chFiles...)
		} else { // 如果是文件，判断是否为所需文件
			fName := strings.ToLower(f.Name())
			// 判断文件名：不是 readme.md | readme.markdown 的 .md | .markdown 文档
			isMDFileWithoutReadme :=
				(!strings.EqualFold(fName, "readme.md") && !strings.EqualFold(fName, "readme.markdown")) &&
					(strings.HasSuffix(fName, ".md") || strings.HasSuffix(fName, ".markdown"))
			// 如果是 md 文档，加入返回结果中
			if isMDFileWithoutReadme {
				files = append(files, chFile)
			}
		}
	}
	return files, nil
}

// 判断目标目录是否存在，如果不存在则创建
func targetDirExists(tDirPath string) (err error) {
	_, err = os.Stat(tDirPath)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = os.Mkdir(tDirPath, os.ModePerm)
		checkErr(err)
	}

	return nil
}

// 复制文件
func copyFile(sDirPath, tDirPath string) (written int64, err error) {
	src, err := os.Open(sDirPath)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(tDirPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func move(sDirPath string, tDirPath string) (err error) {
	mdFiles, err := getSourceFileWithMd(sDirPath)
	checkErr(err)
	err = targetDirExists(tDirPath)
	checkErr(err)
	for _, mdFile := range mdFiles {
		targetFile := fmt.Sprintf("%s/%s", tDirPath, path.Base(mdFile))
		_, err = copyFile(mdFile, targetFile)
		checkErr(err)
	}
	return nil
}
