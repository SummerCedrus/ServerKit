package misc

import (
	"bytes"
	"encoding/gob"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"strconv"
)

//序列化struct
func Serialize(v interface{}) ([]byte,error){
	buffer := new(bytes.Buffer)
	enc := gob.NewEncoder(buffer)
	err := enc.Encode(v)
	if nil != err{
		Errorf("Encode failed error[%s]", err.Error())
		return buffer.Bytes(), err
	}
	return buffer.Bytes(), nil
}
//反序列化struct, e为指针类型
func Deserialize(b []byte, e interface{}) error{
	reader := bytes.NewReader(b)
	dec := gob.NewDecoder(reader)
	err := dec.Decode(e)
	if nil != err{
		Errorf("Encode failed error[%s]", err.Error())
		return err
	}

	return nil
}
//struct转换成bson.M
func Struct2M(in interface{}, out bson.M) error{
	b, err := bson.Marshal(in)
	if nil != err{
		Errorf("bson.Marshal Failed Error[%s]", err.Error())
		return err
	}
	err = bson.Unmarshal(b, out)
	if nil != err{
		Errorf("bson.Unmarshal Failed Error[%s]", err.Error())
		return err
	}

	return nil
}
//bson.M 转换成struct
func M2Struct(in bson.M, out interface{}) error{
	b, err := bson.Marshal(in)
	if nil != err{
		Errorf("bson.Marshal Failed Error[%s]", err.Error())
		return err
	}
	err = bson.Unmarshal(b, out)
	if nil != err{
		Errorf("bson.UnMarshal Failed Error[%s]", err.Error())
		return err
	}

	return nil
}

func GetAllFileName(Dir string) ([]string, error){
	files , err := ioutil.ReadDir(Dir)
	filenames := make([]string, 0)
	if nil != err{
		Errorf("bson.Marshal Failed Error[%s]", err.Error())
		return filenames, err
	}
	for _, file := range files {
		if file.IsDir(){
			continue
		}
		filenames = append(filenames, file.Name())
	}

	return filenames, nil
}
func Stoi(s string) int{
	i, err := strconv.Atoi(s)
	if nil != err{
		Errorf("Sto32 Error %v", err)
		return 0
	}

	return i
}
func Sto32(s string) int32{
	return int32(Stoi(s))
}

func Sto64(s string)  int64{
	return int64(Stoi(s))
}

func ItoS(i int) string{
	return strconv.Itoa(i)
}