package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var (
	users map[string]net.Conn
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	users = make(map[string]net.Conn)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error: %s\n", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var (
		user_name string
	)
	fmt.Fscanln(conn, &user_name)
	log.Printf("User: \"%s\" try to connect...\n", user_name)
	if _, ok := users[user_name]; ok {
		fmt.Fprintln(conn, "false")
		log.Printf("User: \"%s\" connection refuse\n", user_name)
		return
	}
	users[user_name] = conn
	defer delete(users, user_name)
	fmt.Fprintln(conn, "true")
	log.Printf("User: \"%s\" connection accept\n", user_name)

	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("User: \"%s\" connetion is lost\n", user_name)
			break
		}
		fmt.Printf("%s: %s", user_name, msg)
	}
	log.Printf("Delete user \"%s\"\n", user_name)
}
