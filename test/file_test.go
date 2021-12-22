package test

import (
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
    util.PutSQL(util.MemoData{Key: "实例详情", Data: "select * from dbs_instance"})
    util.PutSQL(util.MemoData{Key: "备份详情", Data: "select * from dbs_backup_info"})
}
