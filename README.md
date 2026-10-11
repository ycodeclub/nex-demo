# nex-demo

工作区演示服务：登录 → JWT → 获取当前用户信息。业务 HTTP **强制 proto**（`nex.v1.http` + `protoc-gen-nex`）。

## 准备

```powershell
# 工作区根目录已 install 本地 CLI 时可跳过
cd ..\nex
go install ./cmd/nex ./cmd/protoc-gen-nex

cd ..\demo
go mod tidy
make proto
```

`go.mod` 已 `replace` 到同级 `../nex` 与 `../nex-contrib`。

## 运行

```powershell
go run ./cmd/server
# 或 make run
```

默认 `:8080`。演示账号：

| 用户名 | 密码 | 角色 |
|--------|------|------|
| alice | password123 | user, admin |
| bob | password123 | user |

## 试接口

```powershell
# 登录
curl -s http://127.0.0.1:8080/v1/auth/login -H "Content-Type: application/json" -d "{\"username\":\"alice\",\"password\":\"password123\"}"

# 个人资料（把 <token> 换成 access_token）
curl -s http://127.0.0.1:8080/v1/auth/me -H "Authorization: Bearer <token>"
```

## 契约

见 [`proto/v1/auth.proto`](proto/v1/auth.proto)：

| 传输 | 方法 | 说明 |
|------|------|------|
| HTTP | `POST /v1/auth/login` | public |
| HTTP | `GET /v1/auth/me` | private，JWT |
| gRPC | `auth.v1.Auth/Login` | 同端口 HTTP/2，无需 token |
| gRPC | `auth.v1.Auth/GetProfile` | metadata `authorization: Bearer <token>` |

同端口 mux：HTTP 与 gRPC 共用 `server.addr`（默认 `:8080`）。

## License

Apache License 2.0 — see [LICENSE](./LICENSE).
