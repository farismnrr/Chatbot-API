package model

type OllamaRequest struct {
	Model    string `json:"model"`
	User     string `json:"user"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
		Name    string `json:"name"`
	} `json:"messages"`
}

type ChatCompletion struct {
	ID      string   `json:"id"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Index        int      `json:"index"`
	Message      Message  `json:"message"`
	Logprobs     *float64 `json:"logprobs,omitempty"`
	FinishReason string   `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int     `json:"prompt_tokens"`
	PromptTime       float64 `json:"prompt_time"`
	CompletionTokens int     `json:"completion_tokens"`
	CompletionTime   float64 `json:"completion_time"`
	TotalTokens      int     `json:"total_tokens"`
	TotalTime        float64 `json:"total_time"`
}
