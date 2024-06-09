package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

// Structs for API responses
type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type PoemResponse struct {
	Poem string `json:"poem"`
}

type RequestBody struct {
	Prompt string `json:"prompt"`
}

// Function to make a request to OpenAI API
func getPoemFromOpenAI(prompt string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()

	apiKey := os.Getenv("OPENAI_API_KEY")
	url := "https://api.openai.com/v1/chat/completions"
	systemPrompt := "Sen yaratıcı ve yetenekli bir Türk şairisin. Türkçe şiir kurallarına ve şiir temsillerine uyduğunuzdan emin olun. Cevapların TÜRKÇE olmalıdır. Şiirler için öncelikle bir adet TÜRKÇE yaratıcı başlık da oluştur. Verilen girdiyi temel alarak güzel ve ilham verici TÜRKÇE bir şiir yaz:\n\n"

	requestBody, _ := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "system", "content": systemPrompt},
			{"role": "user", "content": prompt},
		},
	})

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("OpenAI request creation error:", err)
		ch <- ""
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("OpenAI request error:", err)
		ch <- ""
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("OpenAI response read error:", err)
		ch <- ""
		return
	}

	var response OpenAIResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("OpenAI response parsing error:", err)
		ch <- ""
		return
	}

	if len(response.Choices) > 0 {
		ch <- response.Choices[0].Message.Content
	} else {
		ch <- ""
	}
}

// Function to handle poem generation requests
func generatePoemHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	prompt := reqBody.Prompt
	var wg sync.WaitGroup
	poemCh := make(chan string, 4) // Create a buffered channel with capacity 4

	wg.Add(1)
	go getPoemFromOpenAI(prompt, &wg, poemCh)

	// TODO: Add similar goroutines for Groq, HuggingFace, and Gemini

	wg.Wait()
	close(poemCh)

	// Aggregate results (for simplicity, we'll take the first non-empty result)
	var poem string
	for p := range poemCh {
		if p != "" {
			poem = p
			break
		}
	}

	response := PoemResponse{Poem: poem}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/generate_poem", generatePoemHandler)
	port := ":8000"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
