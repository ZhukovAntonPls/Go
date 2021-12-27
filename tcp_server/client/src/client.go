package tcp_server_client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//open connection
	connection, err := net.Dial("tcp", ":3001")
	if err != nil {
		//No connection could be made because the target machine actively refused it
		fmt.Println("Error dialing", err.Error())
		return //terminate program
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")

	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err := connection.Write([]byte(trimmedClient + " says: " + trimmedInput + "\n"))
		if err != nil {
			return
		}
	}
}
