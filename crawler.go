package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get error.")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("http read err.")
		return
	}
	
	src := string(body)
	
	//fmt.Println(body)	 
	//fmt.Println(src)
	
	//convert HTML tags to lower case
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	
	//Remove style
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	
	// Remove all HTML code in angle brackets, and replace with newline.
        re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
        src = re.ReplaceAllString(src, "\n")

       // Remove continuous newline.
       re, _ = regexp.Compile("\\s{2,}")
       src = re.ReplaceAllString(src, "\n")

       fmt.Println(strings.TrimSpace(src))
}
