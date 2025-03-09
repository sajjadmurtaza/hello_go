package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type ChatServer struct {
	APIKey string
}

func NewChatServer(apiKey string) *ChatServer {
	return &ChatServer{
		APIKey: apiKey,
	}
}

func (cs *ChatServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/":
		http.ServeFile(w, r, "index.html")
	case r.URL.Path == "/assets/style.css" && r.Method == http.MethodGet:
		http.ServeFile(w, r, "assets/style.css")
	case r.URL.Path == "/assets/app.js" && r.Method == http.MethodGet:
		http.ServeFile(w, r, "assets/app.js")
	case r.URL.Path == "/api/conversation":
		if r.Method == http.MethodGet {
			cs.handleConversationSSE(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	default:
		http.NotFound(w, r)
	}
}

func (cs *ChatServer) handleConversationSSE(w http.ResponseWriter, r *http.Request) {
	// Header settings for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Get "message" from URL parameters
	userMessage := r.URL.Query().Get("message")
	if userMessage == "" {
		http.Error(w, "message parameter is missing", http.StatusBadRequest)
		return
	}

	// Flusher is required for SSE transmission
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}
	// Pass the message to the Gemini API and receive the API result via channel
	chMsg, err := cs.requestGeminiAPI(r.Context(), userMessage)
	if err != nil {
		slog.ErrorContext(r.Context(), err.Error())
		http.Error(w, "Streaming error.", http.StatusInternalServerError)
		return
	}
	// Execute processing in goroutine
	done := make(chan struct{})
	// When handling a channel, if it is open and there is no data in the channel, it will be blocked, so
	// use goroutine to keep it in a state where it can be blocked.
	go func() {
		defer close(done)
		// Receive messages from the channel (will be blocked unless the channel is closed)
		for msg := range chMsg {
			// SSE format: "data: <message>\n\n"
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()
		}
		// End of SSE
		fmt.Fprintf(w, "data: %s\n\n", "[DONE]")
		flusher.Flush()
	}()

	<-done // Wait for the goroutine to complete
}

const Gemini2_0_Flash_Exp = "gemini-2.0-flash-001"

// requestGeminiAPI Requests the Gemini API and performs the response via channel
func (cs *ChatServer) requestGeminiAPI(ctx context.Context, message string) (chan string, error) {
	// Generate Gemini api client
	client, err := genai.NewClient(ctx, option.WithAPIKey(cs.APIKey))
	if err != nil {
		return nil, fmt.Errorf("Error creating client: %w", err)
	}
	model := client.GenerativeModel(Gemini2_0_Flash_Exp)
	session := model.StartChat()
	// Receive streaming
	respIter := session.SendMessageStream(ctx, genai.Text(message))
	// Start channel
	chMsg := make(chan string)
	// Use goroutine to advance the iterator and implement data acquisition and propagation.
	go func() {
		defer func() {
			// Client disconnection
			client.Close()
			// Close the channel to release the blocking
			close(chMsg)
		}()
		// Pass the iterator and channel for cooperation.
		responseStream(respIter, chMsg)
	}()
	return chMsg, nil
}

func responseStream(respIter *genai.GenerateContentResponseIterator, chMsg chan string) {
	// Loop until the iterator is finished
	for {
		resp, err := respIter.Next()
		if err != nil {
			// Ends the process when the iterator ends
			if err == iterator.Done {
				break
			}
			// Pass the error message to the channel and propagate the error message
			chMsg <- fmt.Sprintf("%v", err)
			return
		}
		for _, cand := range resp.Candidates {
			if cand.Content != nil {
				for _, part := range cand.Content.Parts {
					switch p := part.(type) {
					case genai.Text:
						// Propagate the message via channel
						chMsg <- string(p)
					}
				}
			}
		}
	}
}

// ======================================
// ================ main ================
// ======================================
// func main() {
// 	server := NewChatServer("AIzaSyCIthvZ-G5Mk3VZWq8vLPHPpt9O82ZS5HA")

// 	fmt.Println(server)
// 	fmt.Println(server.APIKey)

// 	resp, err := http.Get("http://google.com")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(resp)
// }

func main() {
	cs := NewChatServer("AIzaSyCIthvZ-G5Mk3VZWq8vLPHPpt9O82ZS5HA")
	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", cs)
}
