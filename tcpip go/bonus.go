package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)
const(
	//Here i declared tcp and ip
	const_type = "tcp"
	const_ip_port = "127.0.0.1:1111"
)

func main() {
	conn,_ := net.Dial(const_type, const_ip_port)
	
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Send : ")
		
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		//Here i cheked the text  if the text is empty
		
		if strings.TrimSpace(string(text)) == "" {
        	fmt.Print("From server : Text is Empty \n")
        
        }else if(strings.TrimSpace(string(text)) == "EXIT") { // When I will write EXIT my code eill stopped
            fmt.Println("User exiting...")
            break;

        }else{ // Here just return server answers
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("From server : " + message)
		}

	}
}