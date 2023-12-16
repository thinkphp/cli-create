package main

import (
     "fmt"
     "io/ioutil"
     "log"
     "os"
     "path"
     "strings"
)

//will add more stuff we needed
var content = `
#include <iostream>
#include <chrono>

#define TIMER_START auto TIME_START = std::chrono::high_resolution_clock::now()
#define TIMER_END auto TIME_END = std::chrono::high_resolution_clock::now()
#define TIMECHECK std::cout << std::chrono::duration_cast<std::chrono::milliseconds>(TIME_END - TIME_START).count() << "ms" << std::endl

typedef double f64;
typedef long long i64;
typedef int i32;
typedef pair<i32, i32> pi32;
typedef unsigned long long u64;
typedef unsigned int u32;
typedef vector<i32> vi32;
typedef deque<i32> di32;

int main() {
    const int N = 1000000;

    TIMER_START;

    // Simulating some time-consuming operation, like a loop
    for (int i = 0; i < 1000000; ++i) {

        // Perform some computation
        int result = i * i;
    }

    TIMER_END;
    TIMECHECK;

    return 0;
}
`

func main() {

     if len(os.Args) < 2 {
       log.Fatal("Please tell me the file name")
     }

     fileName, path := os.Args[1], path.Dir(os.Args[1])

	   if !strings.HasSuffix(fileName, ".cpp") {
		             fileName += ".cpp"
	   }

     if path != "." {
	              	os.Mkdir(path, os.ModePerm)
	   }

     if len(os.Args) == 3 && os.Args[2] == "lc" {

	              	content = content[:strings.Index(content, "i32 main")]
	   }

     err := ioutil.WriteFile(fileName, []byte(content), os.ModePerm)

     if err != nil {
		    log.Fatal(err)
	   }
     
	   fmt.Printf("%s created \n", fileName)
}
