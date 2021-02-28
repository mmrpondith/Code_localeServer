package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	nl, err := net.Listen("tcp", ":8888") //ip:prot
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) //1
	}
	conn, err := nl.Accept()
	if err != nil {
		fmt.Println(err.Error())
	}
	bs := make([]byte, 1024)
	n, e := conn.Read(bs)
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println(n)
	reqstr := string(bs)
	fmt.Println(reqstr)
	body := ` <!DOCTYPE html>
	<html>
		<title>HTML Tag tutorial in Webcoachbd</title>
		<body>
	
			<h1 style="color:yellow;">HTML Attribute tutorial</h1>
			<h2 title="Hello Webcoachbd">Bangladesh is a country of natural beauty</h2>
			<p align="right">Most browsers automatically left-justify a new paragraph.</p>
			<p align="center">To change this behavior, the align attribute for the tag and provide four kinds of content justification: left, right, center, and justify.</p>
			<p style="background-color:#f00;" title="This is tooltip tile">You may specify a paragraph only within a block, along with other paragraphs, lists, forms, and preformatted text</p>
			 <ul>
			  <li>Home</li>
			  <li>about</li>
			  <li>contect</li>
			</ul> 
			 <ol>
			  <li>Coffee</li>
			  <li>Tea</li>
			  <li>Milk</li>
			</ol> 
	
		</body>
	</html> `

	resp := "HTTP/1.1 200 OK\r\n" +
		"Content-Length: %d\r\n" +
		"Content-Type: text/html\r\n" +
		"\r\n%s"
	msg := fmt.Sprintf(resp, len(body), body)
	fmt.Println(msg)
	conn.Write([]byte(msg))
}
