package test

import (
    "igo/util"
    "testing"
)

func Test_PutYaml(t *testing.T)  {
    util.PutYaml("PowerBoot", true)
    t.Log("执行成功")
    res := util.GetYaml("PowerBoot")
    t.Log(res)
    util.PutYaml("Reminder", true)
}
