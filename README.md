# ServerKit
## How To Begin
1.Set You $GOROOT and $GOPATH correct

2.Install protoc(can find here -> ServerKit/src/github.com/)

3.install protoc-gen-go(make sure set correct bin $PATH)

```bash
cd ServerKit/src/github.com/golang/protobuf&&make
```
## makefile help
1.Make demo client&server -- make 

2.Make protocol -- make p

3.Make Server Only -- make s

4.Make Client Only -- make c

5.Make plugin -- make plug

## Hotplugin Package

Hotplugin is base on plugin package.It can reload ''.so" automatically.

1.Code in $GOPATH/src/plugins

2.Make plug--It will build ".so" to $GOPATH/plugins

3.Use hotplugin.Call to Call function in your ".so"

