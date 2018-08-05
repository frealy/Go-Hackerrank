package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	stdout, err := os.Create(dir + "/test.txt")

	str1 := `157.55.39.171 - - [03/Aug/2018:09:35:50 +0700] "GET /script/kosania.global.js HTTP/1.1" 200 4406 "-" "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"
207.46.13.226 - - [03/Aug/2018:09:35:53 +0700] "GET /script/jquery-ui.min.js HTTP/1.1" 200 253764 "-" "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"
207.46.13.226 - - [03/Aug/2018:09:35:53 +0700] "GET /script/bootstrap.min.js HTTP/1.1" 200 37097 "-" "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"
157.55.39.170 - - [03/Aug/2018:09:35:57 +0700] "GET /style/host-benefits.css HTTP/1.1" 200 116 "-" "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"
40.77.167.182 - - [03/Aug/2018:09:36:03 +0700] "GET /robots.txt HTTP/1.1" 200 79 "-" "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"
40.77.167.182 - - [03/Aug/2018:09:36:04 +0700] "GET /style/bootstrap.min.css HTTP/1.1" 200 121285 "-" "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"
40.77.167.182 - - [03/Aug/2018:09:36:04 +0700] "GET /style/button.css HTTP/1.1" 200 3111 "-" "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"
40.77.167.182 - - [03/Aug/2018:09:36:05 +0700] "GET /style/animate.min.css HTTP/1.1" 200 70897 "-" "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"
40.77.167.182 - - [03/Aug/2018:09:36:05 +0700] "GET /script/jquery-3.2.1.min.js HTTP/1.1" 200 86766 "-" "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"
40.77.167.182 - - [03/Aug/2018:09:36:09 +0700] "GET /script/signin.js HTTP/1.1" 200 0 "-" "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"
66.249.69.149 - - [03/Aug/2018:10:50:47 +0700] "GET /robots.txt HTTP/1.1" 200 79 "-" "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
66.249.69.151 - - [03/Aug/2018:10:50:49 +0700] "GET / HTTP/1.1" 200 29249 "-" "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"`

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Println(re.MatchString(str1))        // true

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Fprintf(writer, "%s\n", element)
		fmt.Println(element)
	}

	writer.Flush()
}
