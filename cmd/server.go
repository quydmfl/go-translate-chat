package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/quydmfl/go-translate-chat/pkg/websocket"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the WebSocket server",
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func startServer() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Server recovered from panic: %v", r)
			time.Sleep(2 * time.Second) // Prevent fast crash loops
			go startServer()            // Restart server
		}
	}()

	hub := websocket.NewHub()
	go hub.Run()

	server := &http.Server{Addr: ":8080"}

	// Channel to listen for OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start the server
	go func() {
		fmt.Println("Server started on ws://localhost:8080/ws")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
		}
	}()

	<-stop
	fmt.Println("Shutting down server...")

	// **Graceful Shutdown**
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	fmt.Println("Server exited cleanly")
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
