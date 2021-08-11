package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	user_name, addr, err_init := flag_init()
	if err_init != nil {
		println(fmt.Sprintf("Error: %s\n", err_init))
		return
	}
	conn, err_conn := net.Dial("tcp", addr)
	if err_conn != nil {
		println(fmt.Sprintf("Error: can't connect to server: %s: %s\n",
			addr, err_conn))
		return
	}
	if is_unique_name := check_name(user_name, conn); !is_unique_name {
		println(fmt.Sprintf("Error: name: \"%s\" is exist", user_name))
		return
	}
	session(conn)
}

func session(conn net.Conn) {
	go receive_msgs(conn)
	for {
		fmt.Print("-> ")
		msg, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		_, err := fmt.Fprint(conn, msg)
		if err != nil {
			println("Connection with server is lost")
			break
		}
	}
}

func receive_msgs(conn net.Conn) {
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			println("Connection with server is lost")
			os.Exit(1)
		}
		fmt.Print("\r\n" + msg)
	}
}

func check_name(name string, conn net.Conn) (is_unique_name bool) {
	fmt.Fprintln(conn, name)
	fmt.Fscanln(conn, &is_unique_name)
	return is_unique_name
}
