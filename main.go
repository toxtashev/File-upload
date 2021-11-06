package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {

    r.ParseMultipartForm(10 << 20)

    file, handler, err := r.FormFile("myFile")

    if err != nil {panic(err)}

    fmt.Printf("Name: %s\n", handler.Filename)
    fmt.Printf("Header: %s\n", handler.Header)
    fmt.Printf("Size: %d\n", handler.Size)

    defer file.Close()

    
    tempFile, err := ioutil.TempFile("photos", "upload-*.png")

    if err != nil {panic(err)}

    defer tempFile.Close()


    fileBytes, err := ioutil.ReadAll(file)

    if err != nil {panic(err)}

    tempFile.Write(fileBytes)

    fmt.Fprintf(w, "Successfully!\n")
}

func main() {

    http.HandleFunc("/upload", uploadFile)
    http.ListenAndServe(":9090", nil)    
}
