package main

import (
	"context"
	"net"
	"fmt"
	"bufio"
	"strings"
	"log"
)
const(
	//Here I decalred port and tcp
	const_type = "tcp"
	const_port = ":1111"
)

func runServer(ctx context.Context) error {
	
	fmt.Println("Server started...")
	
	var lc net.ListenConfig

	ln, err := lc.Listen(ctx, const_type, const_port)
	if err != nil {
		return err
	}

	conn, _ := ln.Accept()
	for {
    	message, _ := bufio.NewReader(conn).ReadString('\n')
    	switch  {
    	case  (strings.TrimSpace(string(message)) == "") : //Here i also check for empty values
    		fmt.Print("Message Received: Text is Empty\n")
    	default:
    		fmt.Print("Message Received:", string(message))
	
	    	if(strings.TrimSpace(string(message)) == ""){
	    		str := "Emty "	
	    		conn.Write([]byte(str))
	    	}else {
		    	str := "Hello "+string(message)+"\n"
		    	conn.Write([]byte(str))
		    	newmessage := strings.ToUpper(message)
		    	conn.Write([]byte(newmessage + "\n"))
	    	}
    	}
  	}
  	return nil
}

func main(){
	ctx, _ := context.WithCancel(context.Background())
	err := runServer(ctx)
	switch{
	case err == context.Canceled:  
		log.Printf("Shutdown is done")
	case err != nil:
		log.Printf("ERROR: %v", err)
	}
}