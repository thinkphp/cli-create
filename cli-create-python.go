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

"""
Module Docstring
"""

__author__ = "Your Name"
__version__ = "0.1.0"
__license__ = "MIT"


def main():
    """ Main entry point of the app """
    print("hello world")


if __name__ == "__main__":
    """ This is executed when run from the command line """
    main()
`

func main() {

     if len(os.Args) < 2 {
     
        log.Fatal("Please tell me the file name")
     }

     fileName, path := os.Args[1], path.Dir(os.Args[1])

	   if !strings.HasSuffix(fileName, ".py") {
		             fileName += ".py"
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
