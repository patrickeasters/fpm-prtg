// script to monitor php-fpm status via fastcgi
package main

import "./fcgiclient"
import "io/ioutil"
import "log"
import "os"
import "fmt"

func main() {
    var content []byte

    env := make(map[string]string)
    env["SCRIPT_FILENAME"] = "/ping"
    env["SCRIPT_NAME"] = "/ping"
    env["SERVER_SOFTWARE"] = "go/fcgiclient"

    var host = "127.0.0.1:9000"

    if len(os.Args) >= 2 {
        host = os.Args[1]
    }

    fcgi, err := fcgiclient.Dial("tcp", host)
    if err != nil {
            log.Println("err:", err)
            os.Exit(1)
    }

    resp, err := fcgi.Get(env)
    if err != nil {
            log.Println("err:", err)
            os.Exit(1)
    }

    content, err = ioutil.ReadAll(resp.Body)
    if err != nil {
            log.Println("err:", err)
            os.Exit(1)
    }
    if(string(content) == "pong"){
        fmt.Println("0")
    }
}
