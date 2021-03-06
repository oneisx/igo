package test

import (
    "fmt"
    "igo/util"
    "testing"
)

func Test_ConfigFile(t *testing.T) {
    util.PutConfig("PowerBoot", true)
    t.Log("执行成功")
    res := util.GetConfig("PowerBoot")
    t.Log(res)
    util.PutConfig("Reminder", true)
}

func Test_SQLFile(t *testing.T) {
    util.PutSQL(util.MemoData{Key: "实例详情1", Data: "select * from dbs_instance"})
    util.PutSQL(util.MemoData{Key: "备份详情1", Data: "select * from dbs_backup_info"})
    util.PutSQL(util.MemoData{Key: "实例详情2", Data: "select * from dbs_instance"})
    util.PutSQL(util.MemoData{Key: "备份详情2", Data: "select * from dbs_backup_info"})
    util.PutSQL(util.MemoData{Key: "实例详情3", Data: "select * from dbs_instance"})
    util.PutSQL(util.MemoData{Key: "备份详情3", Data: "select * from dbs_backup_info"})
    util.PutSQL(util.MemoData{Key: "实例详情4", Data: "select * from dbs_instance"})
    util.PutSQL(util.MemoData{Key: "备份详情4", Data: "select * from dbs_backup_info"})
    util.PutSQL(util.MemoData{Key: "实例详情5", Data: "select * from dbs_instance"})
    util.PutSQL(util.MemoData{Key: "备份详情5", Data: "select * from dbs_backup_info"})
    util.PutSQL(util.MemoData{Key: "实例详情6", Data: "select * from dbs_instance"})
    util.PutSQL(util.MemoData{Key: "备份详情6", Data: "select * from dbs_backup_info"})
    util.PutSQL(util.MemoData{Key: "实例详情7", Data: "select * from dbs_instance"})
    util.PutSQL(util.MemoData{Key: "备份详情7", Data: "select * from dbs_backup_info"})
    util.PutSQL(util.MemoData{Key: "实例详情8", Data: "select * from dbs_instance"})
    util.PutSQL(util.MemoData{Key: "备份详情8", Data: "select * from dbs_backup_info"})
    util.PutSQL(util.MemoData{Key: "实例详情9", Data: "select * from dbs_instance"})
    util.PutSQL(util.MemoData{Key: "备份详情9", Data: "select * from dbs_backup_info"})
    util.PutSQL(util.MemoData{Key: "实例详情10", Data: "select * from dbs_instance"})
    util.PutSQL(util.MemoData{Key: "备份详情10", Data: "select * from dbs_backup_info"})
    util.PutSQL(util.MemoData{Key: "实例详情11", Data: "select * from dbs_instance"})
    util.PutSQL(util.MemoData{Key: "备份详情11", Data: "select * from dbs_backup_info"})
    util.PutSQL(util.MemoData{Key: "实例详情12", Data: "select * from dbs_instance"})
    util.PutSQL(util.MemoData{Key: "备份详情12", Data: "select * from dbs_backup_info"})
}

func Test_UtilFunc(t *testing.T) {
    fmt.Println(util.IsNum("1"))
}
