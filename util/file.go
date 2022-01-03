package util

import (
    "bufio"
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "os"
)

const (
    MaxSQLId = "MaxSQLId"
)

type MemoData struct {
    Id   int
    Key  string
    Data string
}

func GetSQL(id int) MemoData {
    return getMemoData(id, sqlFilePath())
}

func GetAllSQL() map[int]MemoData {
    return getAllMemoData(sqlFilePath())
}

func PutSQL(memoData MemoData) {
    putMemoData(memoData, sqlFilePath())
}

func UpdateSQL(memoData MemoData, id int) {
    updateMemoData(memoData, id, sqlFilePath())
}

func DelSQL(id int) {
    ms := GetAllSQL()
    delete(ms, id)
    writeMemoData(ms, sqlFilePath())
}

func GetMemo(id int) MemoData {
    return getMemoData(id, memoFilePath())
}

func PutMemo(memoData MemoData) {
    putMemoData(memoData, memoFilePath())
}

func getMemoData(id int, filename string) MemoData {
    m := readMemoData(filename)
    return m[id]
}

func getAllMemoData(filename string) map[int]MemoData {
    return readMemoData(filename)
}

func updateMemoData(memoData MemoData, id int, filename string) {
    ms := readMemoData(filename)
    ms[id] = memoData
    writeMemoData(ms, filename)
}

func putMemoData(memoData MemoData, filename string) {
    ms := readMemoData(filename)
    if checkKeyExist(ms, memoData.Key) {
        err := fmt.Errorf("error: key=%v already existed", memoData.Key)
        fmt.Println(err)
        return
    }
    memoData.Id = getMaxSQLId()
    ms[memoData.Id] = memoData
    writeMemoData(ms, filename)
}

func checkKeyExist(ms map[int]MemoData, key string) bool {
    for _, m := range ms {
        if m.Key == key {
            return true
        }
    }
    return false
}

func getMaxSQLId() int {
    var id int
    res := GetConfig(MaxSQLId)
    if res == nil {
        id = 1
    } else {
        id = res.(int)
        id++
    }
    PutConfig(MaxSQLId, id)
    return id
}

func readMemoData(filename string) map[int]MemoData {
    ms := make(map[int]MemoData)
    data := read(filename)
    err := yaml.Unmarshal(data, &ms)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    return ms
}

func writeMemoData(ms map[int]MemoData, filename string) {
    d, err := yaml.Marshal(ms)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    write(filename, d)
}

func GetConfig(key string) interface{} {
    m := readConfig()
    return m[key]
}

func PutConfig(key string, value interface{}) {
    m := readConfig()
    m[key] = value
    writeConfig(m)
}

func readConfig() map[string]interface{} {
    m := make(map[string]interface{})
    data := read(configFilePath())
    err := yaml.Unmarshal(data, &m)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    return m
}

func writeConfig(in map[string]interface{}) {
    d, err := yaml.Marshal(in)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    write(configFilePath(), d)
}

func read(filename string) []byte {
    var (
        err     error
        content []byte
    )
    content, err = ioutil.ReadFile(filename)
    if err != nil {
        _ = os.MkdirAll(igoHomeDir(), os.ModePerm)
        return nil
    }
    return content
}

func write(filename string, data []byte) {
    // 拿到一个文件对象
    fileObj, err := os.Create(filename)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    writer := bufio.NewWriter(fileObj)
    defer writer.Flush()
    _, _ = writer.Write(data)
}

func sqlFilePath() string {
    return igoHomeDir() + string(os.PathSeparator) + ".sql.yaml"
}

func memoFilePath() string {
    return igoHomeDir() + string(os.PathSeparator) + ".memo.yaml"
}

func configFilePath() string {
    return igoHomeDir() + string(os.PathSeparator) + ".config.yaml"
}

func igoHomeDir() string {
    return UserHomeDir() + string(os.PathSeparator) + ".igo"
}

func UserHomeDir() string {
    dir, _ := os.UserHomeDir()
    return dir
}

func PathExists(path string) bool {
    _, err := os.Stat(path) //os.Stat获取文件信息
    if err != nil {
        if os.IsExist(err) {
            return true
        }
        return false
    }
    return true
}

func IsDir(path string) bool {
    s, err := os.Stat(path)
    if err != nil {
        return false
    }
    return s.IsDir()
}

func IsFile(path string) bool {
    return !IsDir(path)
    
}
