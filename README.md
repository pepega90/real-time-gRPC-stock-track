# real-time-gRPC-stock-track

is a real-time stock monitoring system that streams live stock prices using gRPC and displays them on a web-based dashboard. It utilizes Golang on the backend for handling gRPC connections and HTTP streaming, while the frontend provides a user-friendly interface for visualizing stock price updates.

<img src="https://github.com/pepega90/real-time-gRPC-stock-track/blob/main/preview.gif" />

## Features

- **Real-Time Data Streaming**: Live stock price updates with real-time price changes displayed in a seamless streaming format.
- **Backend with gRPC**: Efficient data handling using gRPC to maintain continuous streams and low-latency data transfer.
- **Interactive Dashboard**: Stock prices are displayed dynamically with a visual chart, built with Lightweight Charts and styled with TailwindCSS.
- **Event-Driven Interface**: Utilizes Server-Sent Events (SSE) to stream stock price data to the frontend without refreshing the page.

## Technologies Used

- **Backend**: Golang, gRPC, HTTP
- **Frontend**: HTML, JavaScript, Lightweight Charts, TailwindCSS
- **Protocol Buffers**: To define stock service data structures
- **Server-Sent Events (SSE)**: For real-time stock data streaming to the web interface

## Prerequisites

- **Go** (>=1.17)
- **Protocol Buffers** for Go
- **gRPC** package for Go

## Project Structure

```plaintext
.
├── grpc                     # Folder for gRPC server files
│   └── main.go              # Backend for gRPC server
├── http                     # Folder for HTTP server with SSE handler
│   └── main.go              # HTTP server with SSE handler
├── protos                   # Folder for Protocol Buffers files
│   ├── stock_grpc.pb.go     # Generated gRPC code
│   ├── stock.pb.go          # Generated protocol buffer code
│   └── stock.proto          # gRPC definitions for StockService
├── static                   # Folder for frontend assets
│   └── index.html           # Frontend dashboard with TailwindCSS and Lightweight Charts
├── go.mod                   # Go module file
├── go.sum                   # Go dependencies file
└── Makefile                 # Makefile for building and running the project
