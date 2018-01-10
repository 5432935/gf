package main

import (
    "fmt"
    "time"
    "gitee.com/johng/gf/g/database/gdb"
)

// 本文件用于gf框架的mysql数据库操作示例，不作为单元测试使用

var db gdb.Link

// 初始化配置及创建数据库
func init () {
    gdb.AddDefaultConfigNode(gdb.ConfigNode {
        Host    : "127.0.0.1",
        Port    : "3306",
        User    : "root",
        Pass    : "123456",
        Name    : "test2",
        Type    : "mysql",
        Role    : "master",
        Charset : "utf8",
    })
    db, _ = gdb.Instance()

    //gdb.SetConfig(gdb.ConfigNode {
    //    Host : "127.0.0.1",
    //    Port : 3306,
    //    User : "root",
    //    Pass : "123456",
    //    Name : "test",
    //    Type : "mysql",
    //})
    //db, _ = gdb.Instance()

    //gdb.SetConfig(gdb.Config {
    //    "default" : gdb.ConfigGroup {
    //        gdb.ConfigNode {
    //            Host     : "127.0.0.1",
    //            Port     : "3306",
    //            User     : "root",
    //            Pass     : "123456",
    //            Name     : "test",
    //            Type     : "mysql",
    //            Role     : "master",
    //            Priority : 100,
    //        },
    //        gdb.ConfigNode {
    //            Host     : "127.0.0.2",
    //            Port     : "3306",
    //            User     : "root",
    //            Pass     : "123456",
    //            Name     : "test",
    //            Type     : "mysql",
    //            Role     : "master",
    //            Priority : 100,
    //        },
    //        gdb.ConfigNode {
    //            Host     : "127.0.0.3",
    //            Port     : "3306",
    //            User     : "root",
    //            Pass     : "123456",
    //            Name     : "test",
    //            Type     : "mysql",
    //            Role     : "master",
    //            Priority : 100,
    //        },
    //        gdb.ConfigNode {
    //            Host     : "127.0.0.4",
    //            Port     : "3306",
    //            User     : "root",
    //            Pass     : "123456",
    //            Name     : "test",
    //            Type     : "mysql",
    //            Role     : "master",
    //            Priority : 100,
    //        },
    //    },
    //})
    //db, _ = gdb.Instance()
}



// 创建测试数据库
func create() {
    fmt.Println("create:")
    _, err := db.Exec("CREATE DATABASE IF NOT EXISTS test")
    if (err != nil) {
        fmt.Println(err)
    }

    s := `
        CREATE TABLE IF NOT EXISTS user (
            uid  INT(10) UNSIGNED AUTO_INCREMENT,
            name VARCHAR(45),
            PRIMARY KEY (uid)
        )
        ENGINE = InnoDB
        DEFAULT CHARACTER SET = utf8
    `
    _, err = db.Exec(s)
    if (err != nil) {
        fmt.Println(err)
    }

    s = `
        CREATE TABLE IF NOT EXISTS user_detail (
            uid   INT(10) UNSIGNED AUTO_INCREMENT,
            site  VARCHAR(255),
            PRIMARY KEY (uid)
        )
        ENGINE = InnoDB
        DEFAULT CHARACTER SET = utf8
    `

    _, err = db.Exec(s)
    if (err != nil) {
        fmt.Println(err)
    }
    fmt.Println()
}

// 数据写入
func insert() {
    fmt.Println("insert:")
    r, err := db.Insert("user", gdb.Map {
        "name": "john",
    })
    if (err == nil) {
        uid, err2 := r.LastInsertId()
        if err2 == nil {
            r, err = db.Insert("user_detail", gdb.Map {
                "uid"  : uid,
                "site" : "http://johng.cn",
            })
            if err == nil {
                fmt.Printf("uid: %d\n", uid)
            } else {
                fmt.Println(err)
            }
        } else {
            fmt.Println(err2)
        }
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}


// 基本sql查询
func query() {
    fmt.Println("query:")
    list, err := db.GetAll("select * from user limit 2")
    if err == nil {
        fmt.Println(list)
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}

// replace into
func replace() {
    fmt.Println("replace:")
    r, err := db.Save("user", gdb.Map {
        "uid"  :  1,
        "name" : "john",
    })
    if (err == nil) {
        fmt.Println(r.LastInsertId())
        fmt.Println(r.RowsAffected())
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}

// 数据保存
func save() {
    fmt.Println("save:")
    r, err := db.Save("user", gdb.Map {
        "uid"  : 1,
        "name" : "john",
    })
    if (err == nil) {
        fmt.Println(r.LastInsertId())
        fmt.Println(r.RowsAffected())
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}

// 批量写入
func batchInsert() {
    fmt.Println("batchInsert:")
    _, err := db.BatchInsert("user", gdb.List {
        {"name": "john_1"},
        {"name": "john_2"},
        {"name": "john_3"},
        {"name": "john_4"},
    }, 10)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println()
}

// 数据更新
func update1() {
    fmt.Println("update1:")
    r, err := db.Update("user", gdb.Map {"name": "john1"}, "uid=?", 1)
    if (err == nil) {
        fmt.Println(r.LastInsertId())
        fmt.Println(r.RowsAffected())
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}

// 数据更新
func update2() {
    fmt.Println("update2:")
    r, err := db.Update("user", "name='john2'", "uid=1")
    if (err == nil) {
        fmt.Println(r.LastInsertId())
        fmt.Println(r.RowsAffected())
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}

// 数据更新
func update3() {
    fmt.Println("update3:")
    r, err := db.Update("user", "name=?", "uid=?", "john2", 1)
    if (err == nil) {
        fmt.Println(r.LastInsertId())
        fmt.Println(r.RowsAffected())
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}


// 链式查询操作1
func linkopSelect1() {
    fmt.Println("linkopSelect1:")
    r, err := db.Table("user u").LeftJoin("user_detail ud", "u.uid=ud.uid").Fields("u.*, ud.site").Where("u.uid > ?", 1).Limit(0, 2).Select()
    if (err == nil) {
        fmt.Println(r)
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}

// 链式查询操作2
func linkopSelect2() {
    fmt.Println("linkopSelect2:")
    r, err := db.Table("user u").LeftJoin("user_detail ud", "u.uid=ud.uid").Fields("u.*,ud.site").Where("u.uid=?", 1).One()
    if (err == nil) {
        fmt.Println(r)
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}

// 链式查询操作3
func linkopSelect3() {
    fmt.Println("linkopSelect3:")
    r, err := db.Table("user u").LeftJoin("user_detail ud", "u.uid=ud.uid").Fields("ud.site").Where("u.uid=?", 1).Value()
    if (err == nil) {
        fmt.Println(r.(string))
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}

// 错误操作
func linkopUpdate1() {
    fmt.Println("linkopUpdate1:")
    r, err := db.Table("henghe_setting").Update()
    if (err == nil) {
        fmt.Println(r.RowsAffected())
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}

// 通过Map指针方式传参方式
func linkopUpdate2() {
    fmt.Println("linkopUpdate2:")
    r, err := db.Table("user").Data(gdb.Map{"name" : "john2"}).Where("name=?", "john").Update()
    if (err == nil) {
        fmt.Println(r.RowsAffected())
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}

// 通过字符串方式传参
func linkopUpdate3() {
    fmt.Println("linkopUpdate3:")
    r, err := db.Table("user").Data("name='john3'").Where("name=?", "john2").Update()
    if (err == nil) {
        fmt.Println(r.RowsAffected())
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}


// 链式批量写入
func linkopBatchInsert1() {
    fmt.Println("linkopBatchInsert1:")
    r, err := db.Table("user").Data(gdb.List{
        {"name": "john_1"},
        {"name": "john_2"},
        {"name": "john_3"},
        {"name": "john_4"},
    }).Insert()
    if (err == nil) {
        fmt.Println(r.RowsAffected())
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}

// 链式批量写入，指定每批次写入的条数
func linkopBatchInsert2() {
    fmt.Println("linkopBatchInsert2:")
    r, err := db.Table("user").Data(gdb.List{
        {"name": "john_1"},
        {"name": "john_2"},
        {"name": "john_3"},
        {"name": "john_4"},
    }).Batch(2).Insert()
    if (err == nil) {
        fmt.Println(r.RowsAffected())
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}

// 链式批量保存
func linkopBatchSave() {
    fmt.Println("linkopBatchSave:")
    r, err := db.Table("user").Data(gdb.List{
        {"id":1, "name": "john_1"},
        {"id":2, "name": "john_2"},
        {"id":3, "name": "john_3"},
        {"id":4, "name": "john_4"},
    }).Save()
    if (err == nil) {
        fmt.Println(r.RowsAffected())
    } else {
        fmt.Println(err)
    }
    fmt.Println()
}

// 主从io复用测试，在mysql中使用 show full processlist 查看链接信息
func keepPing() {
    fmt.Println("keepPing:")
    for {
        fmt.Println("ping...")
        err := db.PingMaster()
        if err != nil {
            fmt.Println(err)
            return
        }
        err  = db.PingSlave()
        if err != nil {
            fmt.Println(err)
            return
        }
        time.Sleep(1*time.Second)
    }
}

// 数据库单例测试，在mysql中使用 show full processlist 查看链接信息
func instance() {
    fmt.Println("instance:")
    db1, _ := gdb.Instance()
    db2, _ := gdb.Instance()
    db3, _ := gdb.Instance()
    for {
        fmt.Println("ping...")
        db1.PingMaster()
        db1.PingSlave()
        db2.PingMaster()
        db2.PingSlave()
        db3.PingMaster()
        db3.PingSlave()
        time.Sleep(1*time.Second)
    }
}


func main() {

    //create()
    //create()
    //insert()
    //query()
    //replace()
    //save()
    //batchInsert()
    //update1()
    //update2()
    //update3()
    //linkopSelect1()
    //linkopSelect2()
    //linkopSelect3()
    //linkopUpdate1()
    //linkopUpdate2()
    //linkopUpdate3()
    keepPing()
}