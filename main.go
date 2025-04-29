package main
import (
  "fmt"
  "net"
  "os"
)

func main() {
  // creates a new TCP server listening on port 8080
  listener, error := net.Listen("tcp", ":8080")
  
  // catch any error while trying to start server
  if error != nil {
    fmt.Println("Error starting server: ", error)
    // stop code execution securely
    os.Exit(1)
  }
  defer listener.Close() // defer means a command that will be executed by the final of the code
  fmt.Println("Server listening on port 8080")
  
  // infinite loop
  for {
    // accept new connections
    connection, error := listener.Accept()
    
    if error != nil {
      fmt.Println("Error accepting connections: ", error)
      continue
    }
    // it creates a new goroutine (soft thread) to avoid unexpected errors
    go handleConnection(connection)
  }
}

func handleConnection(connection net.Conn) {
  defer connection.Close()
  buffer := make([]byte, 1024) // buffer to keep received data
  
  for {
    // read received data and put it into buffer slice
    n, error := connection.Read(buffer)
    
    if error != nil {
      fmt.Println("No data received. Connection closed.")
      return
    }
    
    fmt.Println("Received message: " + string(buffer[:n]))
    
    // uses "n" length of bytes from buffer to write a message
    _, error = connection.Write(buffer[:n])
    
    if error != nil {
      fmt.Println("Error writing in server: ", error)
      return
    }
  }
}