package main

import (
    "crypto/sha256"
    "encoding/hex"
    "io"
    "net/http"
    "google.golang.org/protobuf/proto"
    "log"
    pb "Homework/genproto/example"
)

func hashMessage(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var req pb.Request
    err = proto.Unmarshal(body, &req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    hash := sha256.New()
    hash.Write([]byte(req.Message))
    hashedMessage := hex.EncodeToString(hash.Sum(nil))

    res := &pb.Response{Message: hashedMessage}
    resData, err := proto.Marshal(res)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/octet-stream")
    w.Write(resData)
}

func main() {
    http.HandleFunc("/hash", hashMessage)
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
