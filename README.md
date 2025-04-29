# This is a simple echo server builded in Go

### How does it work? 🤔

In this code, we use the `net` package to create server and handle connections.

The `Listen` function is used to create a server:
```go
listener, error := net.Listen("tcp", ":8080") // creates a TCP server listening on port 8080
```

The `Accept` function is responsible for accept new connections to the server:
```go
connection, error := listener.Accept()
```

---

In `handleConnection` function, the `buffer` slice is used to keep data while the current connection:
```go
buffer := make([]byte, 1024)
```

The server will read the input message and keep it into `buffer`:
```go
n, error := connection.Read(buffer)
```

Then, it will print back the exactly same message to user:
```go
_, error = connection.Write(buffer[:n]) // `n` refers to the message bytes length
```

---

### How to run this? 🧑‍💻

- Firstly, clone this repo:
```bash
git clone https://github.com/Gustanol/GoEchoServer.git && cd GoEchoServer
```

- Run the server:
```bash
go run main.go
```

- Then, open a new terminal window and execute:
```bash
telnet localhost 8080
```