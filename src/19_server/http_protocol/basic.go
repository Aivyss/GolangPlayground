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

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob(filepath.Join(util.GetGoPath(), "src", "18_template", "template_files", "*")))
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
			status := request(conn)
			response(conn, status)
		}(conn)
	}
}

func request(conn net.Conn) string {
	i := 0
	status := "200"
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			methodFields := strings.Fields(ln)
			md := methodFields[0]
			fmt.Println("[METHOD]", md)
			fmt.Println("[URI]", methodFields[1])
		}

		if ln == "" {
			// RFC 7230 HTTP standard condition when it was the end.
			break
		}

		i += 1
	}

	return status
}

func response(conn net.Conn, status string) {
	defer conn.Close()
	var msg []string

	if status == "200" {
		var tplOutput bytes.Buffer
		var body string
		err := tpl.ExecuteTemplate(&tplOutput, "tpl1.gohtml", nil)
		if err != nil {
			log.Println(err)
			status = "500"
			body = ""
		} else {
			body = tplOutput.String()
		}

		msg = createResponseMsg(body, status)
	} else {
		msg = createResponseMsg("", "500")
	}

	for _, value := range msg {
		fmt.Fprint(conn, value)
	}
}

func createResponseMsg(body string, status string) []string {
	return []string{
		"HTTP/1.1 " + status + "OK\r\n",
		"Content-Length: " + fmt.Sprint(len(body)) + "\r\n",
		"Content-Type: text/html\r\n",
		"\r\n",
		body,
	}
}
