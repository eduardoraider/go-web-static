# Go Web Static Page

This project serves a basic HTML page using Go. It utilizes the `net/http` package to create a simple HTTP server.

## Usage

To run the server, use the following command:

```
go run main.go
```

You can then access the HTML page at:

```
http://localhost:8080
```

## Implementation Details

- The `http.FileServer` function creates a static file server that serves files from the specified directory (in this case, the "static" directory).
- The root path ("/") of the HTTP server is associated with the static file server using `http.Handle("/", fs)`.
- The server is started on port 8080 using `http.ListenAndServe(":8080", nil)`.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

by **Eduardo Raider** - Software Engineer