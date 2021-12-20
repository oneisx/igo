package util

import (
    "bufio"
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "os"
)

func GetYaml(key string) interface{} {
    m := readYaml()
    return m[key]
}

func PutYaml(key string, value interface{}) {
    m := readYaml()
    m[key] = value
    writeYaml(m)
}

func readYaml() map[string]interface{} {
    m := make(map[string]interface{})
    data := read("D:\\igo\\test\\.config.yaml")
    err := yaml.Unmarshal(data, &m)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    return m
}

func writeYaml(in map[string]interface{}) {
    d, err := yaml.Marshal(in)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    write("D:\\igo\\test\\.config.yaml", d)
}

func read(filename string) []byte{
    var (
        err     error
        content []byte
    )
    content, err = ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    return content
}

func write(filename string, data []byte) {
    var (
        err error
    )
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
