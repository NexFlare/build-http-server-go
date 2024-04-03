# Build Your Own HTTP Server

Welcome to the HTTP server project! This challenge is an exciting journey into the backbone of the web, giving you the opportunity to build a server from the ground up.

## What is HTTP ?

[HTTP](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol) stands for Hypertext Transfer Protocol. It's the foundation of data communication on the World Wide Web. In this project, you'll construct an HTTP/1.1 server capable of serving multiple clients simultaneously.

## Project Overview

Dive into my implementation of an HTTP server crafted from scratch. This server adeptly manages simple GET/POST requests, serves files, and effortlessly handles multiple concurrent connections.

## Run the Server Locally

To get your HTTP server up and running on your local machine, follow these steps:

1. **Prerequisites**: Ensure `Go (version 1.19)` is installed on your system.
2. **Start the Server**: Execute `./your_server.sh` to launch your program, located in `app/server.go`. By default, the server listens on port 4421.

This project supports following end-points

1. `/echo/<message_to_echo>`: Echoes back the message sent to this endpoint.
2. `/user-agent`: Returns the User-Agent header value in plain text.
3. `/files/<path_to_file>`: Serves the specified file if it exists; returns 404 Not Found otherwise.
4. `/files/<file_name>` : Saves the content of the request body to the specified file.

**Note** : For using `files` path you need to supply a directory to ./your_server.sh by `./your_server.sh --directory <directory>`

## Extend and Experiment

This project is a sandbox for your creativity. Feel free to modify existing endpoints or add new ones to expand the server's capabilities. Dive in, explore HTTP, and make this server uniquely yours. Build you own X by visiting [CodeCrafters](https://app.codecrafters.io/r/good-cicada-186943)
