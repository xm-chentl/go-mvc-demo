package apisvc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
	"text/template"
)

const (
	metadataTpl = `package api

import (
    {{- range $i, $r := .packages }}
    {{ $r }}
    {{- end }}

	"github.com/xm-chentl/go-mvc"
	"github.com/xm-chentl/go-mvc/metadata"
)

// Register is 注册
func Register() {
	metadata.RegisterMap(map[string]mvc.IApi{
		{{- range $i, $r := .entrys }}
		{{ $r.Register }}{{ end }}
	})	
}`
)

var (
	formatQuote = `"github.com/xm-chentl/%s/api%s"`
	regexAPI    = regexp.MustCompile("[A-Za-z0-9]+API")
)

type apiInfo struct {
	Import   string
	Register string

	endpoint string // 终端名
	name     string //
	path     string // api文件路径
	quote    string // 引用名
	target   string // 目标
}

/*
	1. 读取目录
	2. 分析结构
	3. 读取API结构体文件
	4. 生成源文件
*/

func GenerateMatedata(apiPath string) (err error) {
	apiDir := apiPath
	projectName := ""
	if apiDir == "" {
		rootDir, _ := os.Getwd()
		rootDirArray := strings.Split(rootDir, "/")
		apiDir = path.Join(rootDir, "api")
		projectName = rootDirArray[len(rootDirArray)-1]
	} else {
		rootDirArray := strings.Split(apiDir, "/")
		projectName = rootDirArray[len(rootDirArray)-2]
	}
	fmt.Println(" >>>>>>>> ", projectName)

	apiDirFileArray, err := os.ReadDir(apiDir)
	if err != nil {
		return
	}

	apiInfoArray := make([]apiInfo, 0)
	for _, dirEntry := range apiDirFileArray {
		if dirEntry.IsDir() {
			getApiInfo(apiDir, dirEntry, &apiInfoArray)
		}
	}

	var wg sync.WaitGroup
	for index, apiInfo := range apiInfoArray {
		wg.Add(1)
		go getApiName(apiInfo.path, &apiInfoArray[index], &wg)
	}
	wg.Wait()

	// 整合数值
	packageMap := make(map[string]string)
	packageExistMap := make(map[string]string)
	deleteIndexArray := make([]int, 0)
	for index, info := range apiInfoArray {
		if info.name == "" {
			deleteIndexArray = append(deleteIndexArray, index)
			continue
		}

		apiInfoArray[index].target = strings.ReplaceAll(path.Dir(info.target), apiDir, "")
		importQuote := fmt.Sprintf(formatQuote, projectName, apiInfoArray[index].target)
		_, ok := packageExistMap[info.quote]
		if ok {
			if _, ok = packageMap[importQuote]; !ok {
				apiInfoArray[index].quote += "1"
				importQuote = apiInfoArray[index].quote + " " + importQuote
				packageMap[importQuote] = apiInfoArray[index].quote
				packageExistMap[apiInfoArray[index].quote] = importQuote
			}
		} else {
			packageMap[importQuote] = info.quote
			packageExistMap[info.quote] = importQuote
		}

		apiInfoArray[index].Import = importQuote
		apiInfoArray[index].Register = fmt.Sprintf(`"%s/%s": &%s.%s{},`,
			apiInfoArray[index].target,
			apiInfoArray[index].endpoint,
			apiInfoArray[index].quote,
			info.name,
		)
	}
	for _, index := range deleteIndexArray {
		apiInfoArray = append(apiInfoArray[:index], apiInfoArray[index+1:]...)
	}

	var packageArray []string
	for key := range packageMap {
		packageArray = append(packageArray, key)
	}

	metadataTemple, err := template.New("").Parse(metadataTpl)
	if err != nil {
		return
	}

	var contentBytes bytes.Buffer
	metadataTemple.Execute(&contentBytes, map[string]interface{}{
		"entrys":   apiInfoArray,
		"packages": packageArray,
	})

	filePath := path.Join(apiDir, "metadata.go")
	if _, err = os.Stat(filePath); err != nil {
		// 不存在
		_, _ = os.Create(filePath)
	}
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	if err = ioutil.WriteFile(filePath, contentBytes.Bytes(), 0644); err != nil {
		return
	}

	return
}

func getApiInfo(dirPath string, entry os.DirEntry, apiInfoArray *[]apiInfo) {
	if entry.IsDir() {
		dirPath = path.Join(dirPath, entry.Name())
		dirs, _ := os.ReadDir(dirPath)
		for _, dirEntry := range dirs {
			getApiInfo(dirPath, dirEntry, apiInfoArray)
		}
	} else {
		dirNameArray := strings.Split(dirPath, "/")
		*apiInfoArray = append(*apiInfoArray, apiInfo{
			endpoint: strings.Split(entry.Name(), ".")[0],
			path:     path.Join(dirPath, entry.Name()),
			quote:    dirNameArray[len(dirNameArray)-1],
			target:   path.Join(dirPath, entry.Name()),
		})
	}
}

func getApiName(apiPath string, apiInfo *apiInfo, wg *sync.WaitGroup) {
	var err error
	var goFile *os.File
	var name string
	defer func() {
		if goFile != nil {
			goFile.Close()
		}
		if err != nil {
			log.Fatal(err)
		}
		apiInfo.name = name
		wg.Done()
	}()

	goFile, err = os.Open(apiPath)
	if err != nil {
		return
	}

	content, err := ioutil.ReadAll(goFile)
	if err != nil {
		return
	}

	nameArray := regexAPI.FindAllString(string(content), 1)
	if len(nameArray) == 0 {
		return
	}
	name = nameArray[0]
}
