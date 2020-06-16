package hotplugin

import (
	"errors"
	"fmt"
	."github.com/SummerCedrus/ServerKit/misc"
	"path/filepath"
	"plugin"
	"reflect"
	"strconv"
	"strings"
)
type PluginElems struct {
	ElemMap map[string]*PluginElem
}
type PluginElem struct {
	Timestamp int64          //时间戳
	ModName   string         //模块名,调用的名字
	FileName  string         //文件名Name_Timestamp，Open的名字
	P         *plugin.Plugin //插件的指针
	FuncCache map[string]func(... interface{}) []interface{} //函数指针缓存
}

func (pes *PluginElems) Add(fileName string) error {
	fileNameWithoutExt := strings.TrimSuffix(fileName, ".so")
	S := strings.Split(fileNameWithoutExt, "_")
	if len(S) != 2 {
		return errors.New(fmt.Sprintf("error fileName %v", fileName))
	}
	modName := S[0]
	timestamp, err := strconv.Atoi(S[1])
	if nil != err{
		return errors.New(fmt.Sprintf("error fileName %v", fileName))
	}
	if Elem,ok := pes.ElemMap[modName];ok{
		if Elem.Timestamp >= int64(timestamp) {
			return nil
		}
		Elem.Timestamp = int64(timestamp)
		Elem.FileName = fileName
		p, err := openPlugin(fileName)
		if err != nil{
			Logf("err %v",err)
			return err
		}
		Elem.P = p
		Elem.FuncCache = make(map[string]func(... interface{}) []interface{},0)
		Logf("reload %s",fileName)
	}else{
		p, err := openPlugin(fileName)
		if err != nil{
			return err
		}
		pes.ElemMap[modName] = &PluginElem{
			Timestamp:int64(timestamp),
			ModName:modName,
			FileName:fileName,
			P: p,
			FuncCache:make(map[string]func(... interface{}) []interface{},0),
		}
		Logf("load %s",fileName)
	}
	return nil
}

func openPlugin(filename string) (*plugin.Plugin, error){
	Path, err := filepath.Abs(fmt.Sprintf("plugins/%s", filename))
	if err != nil{
		return nil, err
	}
	p, err := plugin.Open(Path)
	if err != nil{
		return nil, err
	}
	return p, nil
}

func (pe *PluginElem) Call(funcName string, args ...interface{}) ([]interface{}, error){
	f, err := pe.getFunc(funcName)
	if nil != err {
		return nil, err
	}

	return f(args...), nil
}

func (pe *PluginElem) getFunc(funcName string) (func(... interface{}) []interface{}, error){
	//先从缓存中找
	if cache, ok := pe.FuncCache[funcName];ok{
		return cache, nil
	}
	f, err := pe.P.Lookup(funcName)
	if nil != err {
		return nil, err
	}
	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)
	fun := func(args ... interface{}) []interface{}{
		nIn := t.NumIn()
		if len(args) != nIn {
			err := errors.New(fmt.Sprintf("Mismatch Input Args Num Need %v But %v",nIn,len(args)))
			return []interface{}{err}
		}

		in := make([]reflect.Value, 0)

		for i, arg := range args{
			if t.In(i).Name() != reflect.TypeOf(arg).Name(){
				err := errors.New(fmt.Sprintf("Input Args Type Need %s But %s",t.In(i).Name(),reflect.TypeOf(arg).Name()))
				return []interface{}{err}
			}
			in = append(in, reflect.ValueOf(arg))
		}

		result := v.Call(in)

		out := make([]interface{}, 0)

		for _, r := range result{
			out = append(out, r.Interface())
		}

		return out
	}
	pe.FuncCache[funcName] = fun
	return fun, nil
}