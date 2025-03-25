// mcp-golang 暂不支持cline https://github.com/metoro-io/mcp-golang/issues/88
// package main
// import (
// 	"fmt"

// 	mcp_golang "github.com/metoro-io/mcp-golang"
// 	"github.com/metoro-io/mcp-golang/transport/stdio"
// )

// // Tool arguments are just structs, annotated with jsonschema tags
// // More at https://mcpgolang.com/tools#schema-generation
// type Content struct {
// 	Title       string  `json:"title" jsonschema:"required,description=The title to submit"`
// 	Description *string `json:"description" jsonschema:"description=The description to submit"`
// }
// type MyFunctionsArguments struct {
// 	Submitter string  `json:"submitter" jsonschema:"required,description=The name of the thing calling this tool (openai, google, claude, etc)"`
// 	Content   Content `json:"content" jsonschema:"required,description=The content of the message"`
// }

// func main() {
// 	done := make(chan struct{})

// 	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())
// 	err := server.RegisterTool("hello", "Say hello to a person", func(arguments MyFunctionsArguments) (*mcp_golang.ToolResponse, error) {
// 		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Hello, %server!", arguments.Submitter))), nil
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = server.RegisterResource("test://resource", "resource_test", "This is a test resource", "application/json", func() (*mcp_golang.ResourceResponse, error) {
// 		return mcp_golang.NewResourceResponse(mcp_golang.NewTextEmbeddedResource("test://resource", "This is a test resource", "application/json")), nil
// 	})
// 	err = server.Serve()
// 	if err != nil {
// 		panic(err)
// 	}

// 	<-done
// }

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// FetchArgs defines the arguments for the fetch tool.
type FetchArgs struct {
	URL string `json:"url"`
}

// SearchArgs defines the arguments for the search tool.
type SearchArgs struct {
	Query string `json:"q"`
}

func main() {
	// Create MCP server
	s := server.NewMCPServer(
		"Demo 🚀",
		"1.0.0",
	)

	// Add tool
	tool := mcp.NewTool("hello_world",
		mcp.WithDescription("Say hello to someone"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the person to greet"),
		),
	)

	// Add tool handler
	s.AddTool(tool, helloHandler)

	// Add fetch tool
	fetchTool := mcp.NewTool(
		"fetch",
		mcp.WithDescription("使用 r.jina.ai 读取 URL 并获取其内容"),
		mcp.WithString("url",
			mcp.Required(),
			mcp.Description("需要抓取的网页url"),
		),
	)
	s.AddTool(fetchTool, fetchHandler)

	// Add search tool
	searchTool := mcp.NewTool(
		"search",
		mcp.WithDescription("使用 s.jina.ai 搜索网络并获取 SERP"),
		mcp.WithString("q",
			mcp.Required(),
			mcp.Description("搜索关键词"),
		),
	)
	s.AddTool(searchTool, searchHandler)

	// Start the stdio server
	fmt.Println("Server started")
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func helloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, ok := request.Params.Arguments["name"].(string)
	if !ok {
		return nil, errors.New("name must be a string")
	}

	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}

func fetchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var args FetchArgs
	if err := mapToStruct(request.Params.Arguments, &args); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	apiKey := os.Getenv("JINA_API_KEY")
	if apiKey == "" {
		return nil, errors.New("JINA_API_KEY is not set")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://r.jina.ai/%s", args.URL), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	return mcp.NewToolResultText(string(body)), nil
}

func searchHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var args SearchArgs
	if err := mapToStruct(request.Params.Arguments, &args); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	apiKey := os.Getenv("JINA_API_KEY")
	if apiKey == "" {
		return nil, errors.New("JINA_API_KEY is not set")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://s.jina.ai/?q=%s", args.Query), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("X-Respond-With", "no-content")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	return mcp.NewToolResultText(string(body)), nil
}

// mapToStruct converts a map[string]interface{} to a struct using JSON marshaling.
func mapToStruct(m map[string]interface{}, v interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}
