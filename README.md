# ServerKit

## makefile help
1.Set You $GOROOT and $GOPATH correct 
2.make--make demo client&server 
3.install protoc(can find here -> ServerKit/src/github.com/)
4.install protoc-gen-go(make sure set correct bin $PATH)

```bash
cd ServerKit/src/github.com/golang/protobuf&&make
```
5.make p--make protocol

## Hotplugin Package

Hotplugin is base on plugin package.It can reload ''.so" automatically.

1.Code in $GOPATH/src/plugins

2.Make plug--It will build ".so" to $GOPATH/plugins

3.Use hotplugin.Call to Call function in your ".so"

