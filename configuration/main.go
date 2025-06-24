package main

import (
    "fmt"
    "github.com/hashicorp/consul/api"
)

func main() {
    // 1. קונפיגורציית ברירת מחדל (מתחברת ל-localhost:8500)
    config := api.DefaultConfig()

    // 2. יצירת לקוח Consul
    client, err := api.NewClient(config)
    if err != nil {
        panic(err)
    }

    // 3. קריאה ל-Key/Value
    kv := client.KV()
    pair, _, err := kv.Get("loadtool/config/speed_factor", nil)
    if err != nil {
        panic(err)
    }

    // 4. בדיקה אם הערך קיים
    if pair == nil {
        fmt.Println("Dont found value in consul")
    } else {
        fmt.Printf("The value from consul: %s\n", string(pair.Value))
    }
}