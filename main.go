package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/http"
)

type TimeArgs struct {
	Format string `json:"format" jsonschema:"description=The time format to use"`
}

func main() {
	// Create a Gin transport
	transport := http.NewGinTransport()

	// Create a new server with the transport
	server := mcp_golang.NewServer(transport, mcp_golang.WithName("mcp-golang-gin-example"), mcp_golang.WithVersion("0.0.1"))

	// Register a simple tool
	err := server.RegisterTool("time", "Returns the current time in the specified format",
		func(ctx context.Context, args TimeArgs) (*mcp_golang.ToolResponse, error) {
			ginCtx, ok := ctx.Value("ginContext").(*gin.Context)
			if !ok {
				return nil, fmt.Errorf("ginContext not found in context")
			}
			userAgent := ginCtx.GetHeader("User-Agent")
			log.Printf("Request from User-Agent: %s", userAgent)

			format := args.Format
			if format == "" {
				format = time.RFC3339
			}
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(time.Now().Format(format))), nil
		})
	if err != nil {
		panic(err)
	}

	go server.Serve()

	// Create a Gin router
	r := gin.Default()

	// Add the MCP endpoint
	r.POST("/mcp", transport.Handler())

	// Start the server
	log.Println("Starting Gin server on :8081...")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
