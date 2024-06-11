package main

import (
    "bytes"
    "io"
    "log"
    "net/http"
    "google.golang.org/protobuf/proto"
    pb "Homework/genproto/example"
)

func main() {
    req := &pb.Request{Message: "Abduazim!"}
    reqData, err := proto.Marshal(req)
    if err != nil {
        log.Fatalf("Failed to marshal request: %v", err)
    }

    resp, err := http.Post("http://localhost:8080/hash", "application/octet-stream", bytes.NewBuffer(reqData))
    if err != nil {
        log.Fatalf("Failed to make request: %v", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Failed to read response: %v", err)
    }

    var res pb.Response
    err = proto.Unmarshal(body, &res)
    if err != nil {
        log.Fatalf("Failed to unmarshal response: %v", err)
    }

    log.Printf("Hashed message: %s", res.Message)
}
