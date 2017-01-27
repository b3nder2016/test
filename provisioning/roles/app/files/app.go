package main

import (
    "fmt"
    consulapi "github.com/hashicorp/consul/api"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
)

var consul_server string = "172.1.1.3:8500"

func getKVP(key string) string {

    config := consulapi.DefaultConfig()
    config.Address = consul_server
    consul, _ := consulapi.NewClient(config)

    kv := consul.KV()

    pair, _, err := kv.Get(key, nil)

    if err != nil {
        panic(err.Error())
    }

    return string(pair.Value)
}

func mysqlQuery(dsn string) string {
    db, err := sql.Open("mysql", dsn)

    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    rows, err := db.Query("SELECT Host FROM user WHERE User='root'")

    if err = rows.Err(); err != nil {
        panic(err.Error())
    } else {
        return "Hello world"
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s", mysqlQuery(getKVP("mysqluser") + ":" + getKVP("mysqlpass") + "@tcp(" + getKVP("mysqlhost") + ":" + getKVP("mysqlport") +")/mysql"))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":80", nil)
}
