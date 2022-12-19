package main

import (
	"fmt"
	"log"
	"net/textproto"
	"os"
	"strings"
)

func RunCmd(conn *textproto.Conn, cmd string) {

	_, err := conn.Cmd(cmd)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := conn.ReadLine()

	if err != nil {
		log.Fatal(err)
	}

	if !strings.HasPrefix(resp, "+OK") {
		log.Fatal(fmt.Sprintf("Bad reponse from server: %s\n", resp))
	}

}

func main() {

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	pass := os.Getenv("PASS")

	fmt.Printf("Connecting to %s\n", host)

	conn, err := textproto.Dial("tcp", fmt.Sprintf("%s:110", host))
	if err != nil {
		log.Fatal(err)
	}

	RunCmd(conn, fmt.Sprintf("USER %s", user))
	RunCmd(conn, fmt.Sprintf("PASS %s", pass))

}
