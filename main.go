package main

import (
	"errors"
	"fmt"
	"log"
	"net/textproto"
	"os"
	"strings"
	"time"
)

func RunCmd(conn *textproto.Conn, cmd string) error {

	_, err := conn.Cmd(cmd)

	if err != nil {
		return err
	}

	resp, err := conn.ReadLine()

	if err != nil {
		return err
	}

	if !strings.HasPrefix(resp, "+OK") {
		return errors.New(fmt.Sprintf("Bad reponse from server: %s\n", resp))
	}

	return nil
}

func main() {

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	forever := os.Getenv("FOREVER")

	for {
		fmt.Printf("Connecting to %s\n", host)

		conn, err := textproto.Dial("tcp", fmt.Sprintf("%s:110", host))
		if err != nil {
			log.Println(err)
			goto next
		}
		err = RunCmd(conn, fmt.Sprintf("USER %s", user))
		if err != nil {
			log.Println(err)
			goto next
		}
		RunCmd(conn, fmt.Sprintf("PASS %s", pass))
		if err != nil {
			log.Println(err)
			goto next
		}

	next:
		if forever == "" {
			break
		}
		sleepy := 5
		log.Printf("Sleeping for %d minutes...\n", sleepy)
		time.Sleep(time.Duration(sleepy) * time.Minute)
	}
}
