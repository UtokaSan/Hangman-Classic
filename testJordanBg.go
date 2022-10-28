+package main

import (
"fmt"
"io/ioutil"
"os"
)

func main() {

mydata := []byte("All the data I wish to write to a file\n")

// the WriteFile method returns an error if unsuccessful
err := ioutil.WriteFile("myfile.data", mydata, 0777)
// handle this error
if err != nil {
// print it out
fmt.Println(err)
}

data, err := ioutil.ReadFile("myfile.data")
if err != nil {
fmt.Println(err)
}

fmt.Print(string(data))

f, err := os.OpenFile("myfile.data", os.O_WRONLY|os.O_APPEND, 0777)
if err != nil {
panic(err)
}
defer f.Close()

data, err = ioutil.ReadFile("test.txt")
if err != nil {
fmt.Println(err)
}
if _, err = f.WriteString("new data that wasn't there originally\n"); err != nil {
panic(err)
}

fmt.Print(string(data))

}