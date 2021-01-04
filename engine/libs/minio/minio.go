package minio

import (
	"context"
	_minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	lua "github.com/yuin/gopher-lua"
)

var Endpoint string = "127.0.0.1:9000"
var AccessKeyID string = "minioadmin"
var SecretAccessKey string = "minioadmin"
var UseSSL bool = false

func Preload(L *lua.LState) {
	L.PreloadModule("minio", Loader)
}

func Loader(L *lua.LState) int {
	minio_ud := L.NewTypeMetatable(`minio_ud`)
	L.SetGlobal(`minio_ud`, minio_ud)
	L.SetField(minio_ud, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"list_buckets": ListBuckets,
	}))
	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}

var api = map[string]lua.LGFunction{
	"open": Open,
}

func Open(L *lua.LState) int {
	ud := L.NewUserData()
	L.SetMetatable(ud, L.GetTypeMetatable("minio_ud"))
	L.Push(ud)
	return 1
}

func ListBuckets(L *lua.LState) int {
	client, err := _minio.New(Endpoint, &_minio.Options{
		Creds: credentials.NewStaticV4(AccessKeyID, SecretAccessKey, ""),
		Secure: UseSSL,
	})
	if err != nil {
		panic(err)
	}
	buckets, err := client.ListBuckets(context.Background())
	if err != nil {
		panic(err)
	}
	result := L.NewTable()
	for _, bucket := range buckets {
		result.Append(lua.LString(bucket.Name))
	}
	L.Push(result)
	return 1
}