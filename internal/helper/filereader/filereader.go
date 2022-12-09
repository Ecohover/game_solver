package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

type Shape struct {
    cells [2][4]bool
}

// 定義一個 2*4 的方格圖形
var shape = Shape{
    cells: [2][4]bool{
        {true, true, true, true},
        {true, true, true, true},
    },
}

func main() {
    // 讀取 config.json 文件內容
    content, err := ioutil.ReadFile("config.json")
    if err != nil {
        panic(err)
    }

    // 定義一個 map[string]interface{} 用來存放解析後的 JSON 內容
    var data map[string]interface{}

    // 解析 JSON 字符串
    if err := json.Unmarshal(content, &data); err != nil {
        panic(err)
    }

    // 取出 "square" 鍵對應的值
    squares := data["square"].([]interface{})

    // 遍歷 "square" 鍵對應的值
    for _, square := range squares {
        // 將每一個 "square" 元素轉換成 map[string]interface{}
        squareMap := square.(map[string]interface{})
        // 取出 "value" 鍵對應的值
        value := squareMap["value"].([]interface{})
        // 定義一個 Shape 結構
        var s Shape
        // 遍歷 "value" 鍵對應的值，並且將每一個元素轉換成 bool 值存入 Shape 結
        for _, square := range squares {
            // 將每一個 "square" 元素轉換成 map[string]interface{}
            squareMap := square.(map[string]interface{})
            // 取出 "value" 鍵對應的值
            value := squareMap["value"].([]interface{})
            // 定義一個 Shape 結構
            var s Shape
            // 遍歷 "value" 鍵對應的值，並且將每一個元素轉換成 bool 值存入 Shape 結構的 cells 字段
            for i, row := range value {
                for j, cell := range row.([]interface{}) {
                    s.cells[i][j] = cell.(bool)
                }
            }

            // 打印 Shape 結構
            fmt.Println(s)
        }
    }
}
