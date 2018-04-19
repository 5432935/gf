// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

package ghttp

import (
    "io/ioutil"
    "net/http"
    "gitee.com/johng/gf/g/util/gconv"
    "gitee.com/johng/gf/g/encoding/gjson"
    "gitee.com/johng/gf/g/container/gtype"
)

// 请求对象
type Request struct {
    http.Request
    parsedGet  *gtype.Bool         // GET参数是否已经解析
    parsedPost *gtype.Bool         // POST参数是否已经解析
    values     map[string][]string // GET参数
    exit       *gtype.Bool         // 是否退出当前请求流程执行
    Id         int                 // 请求id(唯一)
    Server     *Server             // 请求关联的服务器对象
    Cookie     *Cookie             // 与当前请求绑定的Cookie对象(并发安全)
    Session    *Session            // 与当前请求绑定的Session对象(并发安全)
    Response   *Response           // 对应请求的返回数据操作对象
}

// 创建一个Request对象
func newRequest(s *Server, r *http.Request, w http.ResponseWriter) *Request {
    request := &Request{
        parsedGet  : gtype.NewBool(),
        parsedPost : gtype.NewBool(),
        values     : make(map[string][]string),
        exit       : gtype.NewBool(),
        Id         : s.servedCount.Add(1),
        Server     : s,
        Request    : *r,
        Response   : &Response {
            status         : http.StatusOK,
            ResponseWriter : w,
        },
    }
    // 会话处理
    request.Cookie           = GetCookie(request)
    request.Session          = GetSession(request)
    request.Response.request = request
    return request
}


// 初始化GET请求参数
func (r *Request) initGet() {
    if !r.parsedGet.Val() {
        if len(r.values) == 0 {
            r.values = r.URL.Query()
        } else {
            for k, v := range r.URL.Query() {
                r.values[k] = v
            }
        }
    }
}

// 初始化POST请求参数
func (r *Request) initPost() {
    if !r.parsedPost.Val() {
        // 快速保存，尽量避免并发问题
        r.parsedPost.Set(true)
        // MultiMedia表单请求解析允许最大使用内存：1GB
        r.ParseMultipartForm(1024*1024*1024)
    }
}

// 获得指定名称的参数字符串(GET/POST)，同 GetRequestString
// 这是常用方法的简化别名
func (r *Request) Get(k string) string {
    return r.GetRequestString(k)
}

// 获得指定名称的get参数列表
func (r *Request) GetQuery(k string) []string {
    r.initGet()
    if v, ok := r.values[k]; ok {
        return v
    }
    return nil
}

func (r *Request) GetQueryBool(k string) bool {
    return gconv.Bool(r.Get(k))
}

func (r *Request) GetQueryInt(k string) int {
    return gconv.Int(r.Get(k))
}

func (r *Request) GetQueryUint(k string) uint {
    return gconv.Uint(r.Get(k))
}

func (r *Request) GetQueryFloat32(k string) float32 {
    return gconv.Float32(r.Get(k))
}

func (r *Request) GetQueryFloat64(k string) float64 {
    return gconv.Float64(r.Get(k))
}

func (r *Request) GetQueryString(k string) string {
    v := r.GetQuery(k)
    if v == nil {
        return ""
    } else {
        return v[0]
    }
}

func (r *Request) GetQueryArray(k string) []string {
    return r.GetQuery(k)
}

// 获取指定键名的关联数组，并且给定当指定键名不存在时的默认值
func (r *Request) GetQueryMap(defaultMap...map[string]string) map[string]string {
    r.initGet()
    m := make(map[string]string)
    if len(defaultMap) == 0 {
        for k, v := range r.values {
            m[k] = v[0]
        }
    } else {
        for k, v := range defaultMap[0] {
            v2 := r.GetQueryArray(k)
            if v2 == nil {
                m[k] = v
            } else {
                m[k] = v2[0]
            }
        }
    }
    return m
}

// 获得post参数
func (r *Request) GetPost(k string) []string {
    r.initPost()
    if v, ok := r.PostForm[k]; ok {
        return v
    }
    return nil
}

func (r *Request) GetPostBool(k string) bool {
    return gconv.Bool(r.GetPostString(k))
}

func (r *Request) GetPostInt(k string) int {
    return gconv.Int(r.GetPostString(k))
}

func (r *Request) GetPostUint(k string) uint {
    return gconv.Uint(r.GetPostString(k))
}

func (r *Request) GetPostFloat32(k string) float32 {
    return gconv.Float32(r.GetPostString(k))
}

func (r *Request) GetPostFloat64(k string) float64 {
    return gconv.Float64(r.GetPostString(k))
}

func (r *Request) GetPostString(k string) string {
    v := r.GetPost(k)
    if v == nil {
        return ""
    } else {
        return v[0]
    }
}

func (r *Request) GetPostArray(k string) []string {
    return r.GetPost(k)
}

// 获取指定键名的关联数组，并且给定当指定键名不存在时的默认值
// 需要注意的是，如果其中一个字段为数组形式，那么只会返回第一个元素，如果需要获取全部的元素，请使用GetPostArray获取特定字段内容
func (r *Request) GetPostMap(defaultMap...map[string]string) map[string]string {
    r.initPost()
    m := make(map[string]string)
    if len(defaultMap) == 0 {
        for k, v := range r.PostForm {
            m[k] = v[0]
        }
    } else {
        for k, v := range defaultMap[0] {
            if v2, ok := r.PostForm[k]; ok {
                m[k] = v2[0]
            } else {
                m[k] = v
            }
        }
    }
    return m
}

// 获得post或者get提交的参数，如果有同名参数，那么按照get->post优先级进行覆盖
func (r *Request) GetRequest(k string) []string {
    v := r.GetQuery(k)
    if v == nil {
        return r.GetPost(k)
    }
    return v
}

func (r *Request) GetRequestString(k string) string {
    v := r.GetRequest(k)
    if v == nil {
        return ""
    } else {
        return v[0]
    }
}

func (r *Request) GetRequestBool(k string) bool {
    return gconv.Bool(r.GetRequestString(k))
}

func (r *Request) GetRequestInt(k string) int {
    return gconv.Int(r.GetRequestString(k))
}

func (r *Request) GetRequestUint(k string) uint {
    return gconv.Uint(r.GetRequestString(k))
}

func (r *Request) GetRequestFloat32(k string) float32 {
    return gconv.Float32(r.GetRequestString(k))
}

func (r *Request) GetRequestFloat64(k string) float64 {
    return gconv.Float64(r.GetRequestString(k))
}

func (r *Request) GetRequestArray(k string) []string {
    return r.GetRequest(k)
}

// 获取指定键名的关联数组，并且给定当指定键名不存在时的默认值
// 需要注意的是，如果其中一个字段为数组形式，那么只会返回第一个元素，如果需要获取全部的元素，请使用GetRequestArray获取特定字段内容
func (r *Request) GetRequestMap(defaultMap...map[string]string) map[string]string {
    m := r.GetQueryMap()
    if len(defaultMap) == 0 {
        for k, v := range r.GetPostMap() {
            if _, ok := m[k]; !ok {
                m[k] = v
            }
        }
    } else {
        for k, v := range defaultMap[0] {
            v2 := r.GetRequest(k)
            if v2 != nil {
                m[k] = v2[0]
            } else {
                m[k] = v
            }
        }
    }
    return m
}

// 获取原始请求输入字符串，注意：只能获取一次，读完就没了
func (r *Request) GetRaw() []byte {
    result, _ := ioutil.ReadAll(r.Body)
    return result
}

// 获取原始json请求输入字符串，并解析为json对象
func (r *Request) GetJson() *gjson.Json {
    data := r.GetRaw()
    if data != nil {
        if j, err := gjson.DecodeToJson(data); err == nil {
            return j
        }
    }
    return nil
}

// 退出当前请求执行，原理是在Request.exit做标记，由服务逻辑流程做判断，自行停止
func (r *Request) Exit() {
    r.exit.Set(true)
}

// 判断当前请求是否停止执行
func (r *Request) IsExited() bool {
    return r.exit.Val()
}



