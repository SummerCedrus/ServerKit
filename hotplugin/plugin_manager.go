package hotplugin

import (
	"errors"
	"fmt"
	. "github.com/SummerCedrus/ServerKit/misc"
	"path/filepath"
	"strings"
	"time"
)

const (
	CHECK_DUR = int64(10)
)

var lastCheckTime = int64(0)
var pluginMap PluginElems
func Run(){
	pluginMap.ElemMap = make(map[string]*PluginElem)
	lastCheckTime = time.Now().Unix()
	err := reloadAllPlugin()
	if nil != err{
		panic(err)
	}
	Log("Run..")
	loop()
}
func loop(){
	for {
		//1.10秒检查一次
		now := time.Now().Unix()
		if now < lastCheckTime+CHECK_DUR {
			continue
		}

		Logf("check %v", now)
		lastCheckTime = now
		//Test
		//result,err := Call("testplugin","Hello",1,4)
		//if nil != err{
		//	Logf("%v", err)
		//}else{
		//	Logf("result %v",result[0].(int))
		//}

		//2.取plugins下面所有文件名,文件名格式filename_datetime
		dir, err := filepath.Abs("plugins")
		if nil != err{
			panic(err)
		}
		fileNames, err := GetAllFileName(dir)
		if nil != err{
			panic(err)
		}
		//3.对比filename和datetime，如果有更新的就替换掉
		for _, fileName := range fileNames{
			if filepath.Ext(fileName) != ".so"{
				continue
			}

			pluginMap.Add(fileName)
		}
	}
}

func Call(modName,funName string, args ...interface{}) ([]interface{}, error){
	mod, ok := pluginMap.ElemMap[modName]
	if !ok{
		return nil, errors.New(fmt.Sprintf("can't find mod %s", modName))
	}

	result,err := mod.Call(funName, args...)

	if nil != err{
		return nil, err
	}

	return result, nil
}

func reloadAllPlugin() error{
	dir, err := filepath.Abs("plugins")
	if nil != err{
		return err
	}
	fileNames, err := GetAllFileName(dir)
	if nil != err{
		return err
	}
	loadMap := make(map[string]string)
	for _, fileName := range fileNames{
		if filepath.Ext(fileName) != ".so"{
			continue
		}
		fileNameWithoutExt := strings.TrimSuffix(fileName,".so")
		S := strings.Split(fileNameWithoutExt, "_")
		if len(S) != 2 {
			continue
		}
		if timeStampStr,ok := loadMap[S[0]];ok{
			if Sto64(timeStampStr) >= Sto64(S[1]){
				continue
			}
		}
		loadMap[S[0]] = S[1]
	}
	for modName, ts := range loadMap{
		pluginMap.Add(modName+"_"+ts+".so")
	}

	return nil
}