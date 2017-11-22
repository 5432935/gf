package ghttp

import (
    "net/http"
    "time"
    "crypto/tls"
    "log"
    "net/url"
)

// http客户端
type Client struct {
    http.Client
}

// http server结构体
type Server struct {
    server     http.Server
    config     ServerConfig
    handlerMap HandlerMap
}

// 请求对象
type ClientRequest struct {
    http.Request
    getvals *url.Values
}

// 客户端请求结果对象
type ClientResponse struct {
    http.Response
}

// 服务端请求返回对象
type ServerResponse struct {
    http.ResponseWriter
}

// http回调函数
type HandlerFunc func(*ClientRequest, *ServerResponse)

// uri与回调函数的绑定记录表
type HandlerMap map[string]HandlerFunc

// HTTP Server 设置结构体
type ServerConfig struct {
    // HTTP Server基础字段
    Addr            string        // 监听IP和端口，监听本地所有IP使用":端口"
    Handler         http.Handler  // 默认的处理函数
    TLSConfig      *tls.Config    // TLS配置
    ReadTimeout     time.Duration
    WriteTimeout    time.Duration
    IdleTimeout     time.Duration
    MaxHeaderBytes  int           // 最大的header长度
    ErrorLog       *log.Logger    // 错误日志的处理接口
    // gf 扩展信息字段
    IndexFiles      []string      // 默认访问的文件列表
    IndexFolder     bool          // 如果访问目录是否显示目录列表
    ServerAgent     string        // server agent
    ServerRoot      string        // 服务器服务的本地目录根路径
}

// 默认HTTP Server
var defaultServerConfig = ServerConfig {
    Addr           : ":80",
    Handler        : nil,
    ReadTimeout    : 60 * time.Second,
    WriteTimeout   : 60 * time.Second,
    IdleTimeout    : 60 * time.Second,
    MaxHeaderBytes : 1024,
    IndexFiles     : []string{"index.html", "index.htm"},
    IndexFolder    : false,
    ServerAgent    : "gf",
    ServerRoot     : "",
}

// 修改默认的http server配置
func SetDefaultServerConfig (c ServerConfig) {
    defaultServerConfig = c
}

// 创建一个默认配置的HTTP Server(默认监听端口是80)
func NewServer() (*Server) {
    return NewServerByConfig(defaultServerConfig)
}

// 创建一个HTTP Server，返回指针
func NewServerByAddr(addr string) (*Server) {
    config     := defaultServerConfig
    config.Addr = addr
    return NewServerByConfig(config)
}

// 创建一个HTTP Server
func NewServerByAddrRoot(addr string, root string) (*Server) {
    config           := defaultServerConfig
    config.Addr       = addr
    config.ServerRoot = root
    return NewServerByConfig(config)
}

// 根据输入配置创建一个http server对象
func NewServerByConfig(s ServerConfig) (*Server) {
    var server Server
    server.SetConfig(s)
    return &server
}