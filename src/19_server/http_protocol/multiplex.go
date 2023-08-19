package main

import (
	"bufio"
	"bytes"
	"com.playground/18_template/util"
	"fmt"
	"html/template"
	"log"
	"net"
	"path/filepath"
	"strings"
)

type Method int
type methodUri struct {
	uri       string
	method    Method
	methodUri string
}

const (
	Get Method = iota
	Post
	Put
	Patch
	Delete
	Connect
	Options
	Trace
)

var templates *template.Template
var multiplexer = map[methodUri]func(status string) []string{}

func init() {
	templates = template.Must(template.ParseGlob(filepath.Join(util.GetGoPath(), "src", "18_template", "template_files", "*")))
	multiplexer[methodUri{
		uri:       "/hello/first",
		method:    Get,
		methodUri: "GET /hello/first",
	}] = func(status string) []string {
		var tplOutput bytes.Buffer
		var body string

		err := templates.ExecuteTemplate(&tplOutput, "tpl1.gohtml", nil)
		if err != nil {
			log.Println(err)
			status = "500"
			body = ""
		} else {
			body = tplOutput.String()
		}

		return resMsg(body, status)
	}

	multiplexer[methodUri{
		uri:       "/hello/second",
		method:    Get,
		methodUri: "GET /hello/second",
	}] = func(status string) []string {
		var tplOutput bytes.Buffer
		var body string

		err := templates.ExecuteTemplate(&tplOutput, "tpl2.gohtml", nil)
		if err != nil {
			log.Println(err)
			status = "500"
			body = ""
		} else {
			body = tplOutput.String()
		}

		return resMsg(body, status)
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
		}

		go func(conn net.Conn) {
			defer conn.Close()
			msg := req(conn)

			for _, value := range msg {
				fmt.Fprint(conn, value)
			}
		}(conn)
	}
}

func req(conn net.Conn) []string {
	i := 0
	status := "200"
	scanner := bufio.NewScanner(conn)
	var method string
	var uri string
	var msg []string

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			methodFields := strings.Fields(ln)
			method = methodFields[0]
			uri = methodFields[1]
			fmt.Println("[METHOD]", method)
			fmt.Println("[URI]", uri)
		}

		if ln == "" {
			// RFC 7230 HTTP standard condition when it was the end.
			break
		}

		i += 1
	}

	if status != "200" {
		msg = resMsg("", status)
	} else {
		response, ok := multiplexer[methodUri{
			uri:       uri,
			method:    getMethod(method),
			methodUri: method + " " + uri,
		}]

		if ok {
			msg = response(status)
		} else {
			msg = resMsg("", status)
		}
	}

	return msg
}

func resMsg(body string, status string) []string {
	return []string{
		"HTTP/1.1 " + status + "\r\n",
		"Content-Length: " + fmt.Sprint(len(body)) + "\r\n",
		"Content-Type: text/html\r\n",
		"\r\n",
		body,
	}
}

func getMethod(method string) Method {
	var result Method

	switch method {
	case "GET":
		result = Get
	case "POST":
		result = Post
	case "PUT":
		result = Put
	case "PATCH":
		result = Patch
	case "DELETE":
		result = Delete
	case "CONNECT":
		result = Connect
	case "TRACE":
		result = Trace
	case "OPTIONS":
		result = Options
	default:
		result = Get
	}

	return result
}
