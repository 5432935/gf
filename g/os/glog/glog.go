// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// 日志模块.
// 直接文件/输出操作，没有异步逻辑，没有使用缓存或者通道
package glog

import (
    "io"
    "sync"
    "gitee.com/johng/gf/g/container/gtype"
)

type Logger struct {
    mu    sync.RWMutex
    io    io.Writer           // 日志内容写入的IO接口
    path  *gtype.String       // 日志写入的目录路径
    debug *gtype.Bool         // 是否允许输出DEBUG信息
}

// 默认的日志对象
var logger = New()

// 新建自定义的日志操作对象
func New() *Logger {
    return &Logger {
        io    : nil,
        path  : gtype.NewString(),
        debug : gtype.NewBool(true),
    }
}

// 日志日志目录绝对路径
func SetPath(path string) {
    logger.SetPath(path)
}

// 设置是否允许输出DEBUG信息
func SetDebug(debug bool) {
    logger.SetDebug(debug)
}

// 获取日志目录绝对路径
func GetPath() string {
    return logger.path.Val()
}

func Print(v ...interface{}) {
    logger.Print(v ...)
}

func Printf(format string, v ...interface{}) {
    logger.Printf(format, v ...)
}

func Println(v ...interface{}) {
    logger.Println(v ...)
}

func Printfln(format string, v ...interface{}) {
    logger.Printfln(format, v ...)
}

func Fatal(v ...interface{}) {
    logger.Fatal(v ...)
}

func Fatalf(format string, v ...interface{}) {
    logger.Fatalf(format, v ...)
}

func Fatalln(v ...interface{}) {
    logger.Fatalln(v ...)
}

func Fatalfln(format string, v ...interface{}) {
    logger.Fatalfln(format, v ...)
}

func Panic(v ...interface{}) {
    logger.Panic(v ...)
}

func Panicf(format string, v ...interface{}) {
    logger.Panicf(format, v ...)
}

func Panicln(v ...interface{}) {
    logger.Panicln(v ...)
}

func Panicfln(format string, v ...interface{}) {
    logger.Panicfln(format, v ...)
}

func Info(v ...interface{}) {
    logger.Info(v...)
}

func Debug(v ...interface{}) {
    logger.Debug(v...)
}

func Notice(v ...interface{}) {
    logger.Notice(v...)
}

func Warning(v ...interface{}) {
    logger.Warning(v...)
}

func Error(v ...interface{}) {
    logger.Error(v...)
}

func Critical(v ...interface{}) {
    logger.Critical(v...)
}

func Infof(format string, v ...interface{}) {
    logger.Infof(format, v...)
}

func Debugf(format string, v ...interface{}) {
    logger.Debugf(format, v...)
}

func Noticef(format string, v ...interface{}) {
    logger.Noticef(format, v...)
}

func Warningf(format string, v ...interface{}) {
    logger.Warningf(format, v...)
}

func Errorf(format string, v ...interface{}) {
    logger.Errorf(format, v...)
}

func Criticalf(format string, v ...interface{}) {
    logger.Criticalf(format, v...)
}

func Infofln(format string, v ...interface{}) {
    logger.Infofln(format, v...)
}

func Debugfln(format string, v ...interface{}) {
    logger.Debugfln(format, v...)
}

func Noticefln(format string, v ...interface{}) {
    logger.Noticefln(format, v...)
}

func Warningfln(format string, v ...interface{}) {
    logger.Warningfln(format, v...)
}

func Errorfln(format string, v ...interface{}) {
    logger.Errorfln(format, v...)
}

func Criticalfln(format string, v ...interface{}) {
    logger.Criticalfln(format, v...)
}
