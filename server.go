package main

import (
    "log"
    "net/http"
    "io"
    "os"
    "sort"
    "quicksort"
    "time"
    "strconv"
    "fmt"
    "strings"
)

func main () {
    http.HandleFunc("/", form)
    http.HandleFunc("/result", result)
    http.ListenAndServe(":3000", nil)
}

func form (w http.ResponseWriter, r *http.Request) {
    f, err := os.Open("templates/form.html")
    if err != nil {
        log.Fatal(err)
    }
    io.Copy(w, f)
}

func result (w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    var errmsg string
    var numbers []int
    goroutines, err := strconv.Atoi(r.PostForm["goroutines"][0])
    if err != nil {
        errmsg = r.PostForm["goroutines"][0] + " is not a number.\n"
    }
    s := strings.Split(r.PostForm["numbers"][0], " ")
    for _, val := range s {
        num, err := strconv.Atoi(val)
        if err != nil {
            errmsg += val + " is not a number.\n"
        } else {
            numbers = append(numbers, num)
        }
    }
    if errmsg != "" {
        http.Error(w, errmsg, http.StatusBadRequest)
        return
    }
    data := sort.IntSlice(numbers)
    start := time.Now()
    quicksort.QuickSort(data, goroutines)
    end := time.Now()
    fmt.Fprint(w, "Result: ", data, "\n")
    fmt.Fprint(w, "Time  : ", end.Sub(start))
}