package main

import (
	"fmt"
	"os"

	"github.com/ycodeclub/Nex/doc"
)

// OpenAPI 导出入口。nex proto 生成 Register*HTTPDoc 后在 registerDocs 中挂上即可。
//
//	nex doc export
//	nex doc serve
func main() {
	doc.Configure(doc.Info{
		Title:       "demo API",
		Description: "demo HTTP API",
		Version:     "1.0.0",
		BasePath:    "/",
		ServerURL:   "http://localhost:8080",
	})
	m := doc.DefaultModule()
	registerDocs(m)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: openapi <output.yaml> | serve [addr]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "serve":
		addr := ":6633"
		if len(os.Args) > 2 {
			addr = os.Args[2]
		}
		if err := doc.ServeModule(m, addr); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		if err := doc.ExportFile(m, os.Args[1]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println("exported", os.Args[1])
	}
}

// registerDocs 注册 proto 生成的 HTTP 文档。
// 示例（生成后取消注释并按服务名替换）：
//
//	import v1 "github.com/ycodeclub/nex-demo/proto/v1"
//	v1.RegisterGreeterHTTPDoc(m)
func registerDocs(m *doc.Module) {
	_ = m
}
