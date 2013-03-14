package main

import (
        "flag"
	"fmt"
	"os"
)

var (
	Name = "anonymous"
        Encrypted = false
        Key = make([]byte, 0)
)

func init() {
        flUser := flag.String("user", Name, "set the client's user name")
        flKeyFile := flag.String("key", "", "set teh encryption key")
        flag.Parse()

        Name = *flUser
        if *flKeyFile != "" {
                file, err := os.Open(*flKeyFile)
                if err != nil {
                        fmt.Println("[!] couldn't open keyfile:", err.Error())
                        return
                }
                defer file.Close()

                Key = make([]byte, KeySize)
                n, err := file.Read(Key)
                if err != nil {
                        fmt.Println("[!] couldn't open keyfile:", err.Error())
                } else if n != KeySize {
                        fmt.Println("[!] invalid key")
                } else {
                        Encrypted = true
                }
        }
}

func main() {
}
