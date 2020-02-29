package main

import (
    "fmt"
    "os"
    "strings"
    "context"
    kafka "github.com/segmentio/kafka-go"
)


func main () {
    // guard: command line arguments
    if (len(os.Args) != 3) {
        fmt.Println("Syntax: "+os.Args[0]+" BROKERS TOPIC_NAME")
        fmt.Println("        "+os.Args[0]+" localhost:9092 topic.692628e9.3c82.43ee.9c7d.9baf284a97ca")
        os.Exit(1)
    }
    var brokers []string = strings.Split(os.Args[1], ",")
    var topic     string = os.Args[2]
    var groupid   string = "2fa2504c-7d96-40dd-8708-af022d657328"
    
    fmt.Println(fmt.Sprintf("Listening on topic %s", topic))
    
    var reader *kafka.Reader = kafka.NewReader(kafka.ReaderConfig{
        Brokers:  brokers,
        GroupID:  groupid,
        Topic:    topic,
        MinBytes: 0, // 10kB
        MaxBytes: 10e6, // 10MB
    })
    defer reader.Close()
    
    for {
        fmt.Println("============================================================================")
        m, err := reader.ReadMessage(context.Background())
        var message []byte = m.Value
        if err != nil {
            fmt.Println("Error reading lookup message:", err)
        } else {
            fmt.Println(string(message))
            fmt.Println("----------------------------------------------------------------------------")
            var b []byte = []byte(message)
            for char := range(b) {
                var s string
                if char==0xA {
                    s = "<NL>"
                } else if char==0xD {
                    s = "<CR>"
                } else {
                    s = string(char)
                }
                fmt.Println(fmt.Sprintf("%d : %s", char, s))
            }
        }
    }
    
    // enter service loop
    select{}
}

