package tutor

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"encoding/json"
	"io"
	"bytes"
)
// todo database, file?
var lessons = []string{"Present Continuous для действий в момент речи (Subject + be + V-ing)",
"Past Simple для завершенных действий (Subject + V2/ed)",
"Present Perfect для жизненного опыта (Subject + have/has + V3)",
"Future with 'going to' для планов и намерений (Subject + am/is/are + going to + V1)",
"Future with 'will' для спонтанных решений (Subject + will + V1)",
"Phrasal Verbs — фразовые глаголы (Verb + particle)",
"Modal verbs of ability — модальные глаголы способности (Subject + can/could + V1)",
"Modal verbs of advice — модальные глаголы совета (Subject + should + V1)",
"Question Tags — уточняющие вопросы (Statement, auxiliary + subject?)",
"Comparative adjectives — сравнительная степень (Adjective-er + than / more + adjective + than)",
"Superlative adjectives — превосходная степень (The + adjective-est / the most + adjective)",
"Zero Conditional — общеизвестные истины (If + Present Simple, Present Simple)",
"First Conditional — реальное будущее (If + Present Simple, will + V1)",
"Second Conditional — гипотетическое настоящее/будущее (If + Past Simple, would + V1)",
"Used to для прошлых привычек (Subject + used to + V1)",
"Emphatic 'Do' — эмфаза (Subject + do/does/did + Verb)",
"Present Perfect with 'just' — недавние события (Subject + have/has + just + V3)",
"Modal verbs of deduction in present — предположения в настоящем (Must/Might/Can't + V1)",
"Gerund after prepositions — герундий после предлогов (Preposition + V-ing)",
"Connectors of contrast — союзы противопоставления (But / However / Although)"}

// todo The prompt is overwelming, need to think how to make it better, so rule is not buried under additional text
// todo embeded file
// todo translation should be user settings
var prompt = `Role: Act as a warm, encouraging, and creative Elementary School English Teacher who makes learning fun and easy to understand.

Task: Your goal is to explain the following English grammar rule: %s.

Persona Guidelines:
Terminology: Use child-friendly metaphors. Treat parts of speech as "building blocks," "colors," or "puzzle pieces." Describe sentence structure as "putting toys in the right toy box" or "following a magic recipe."
Tone: Be enthusiastic, patient, and incredibly supportive. Use gentle corrections and lots of verbal "gold stars."
Simplicity: Explain the rule so clearly and simply that a 7-year-old could understand it.

Structure & Formatting: 
Format your response EXCLUSIVELY using Telegram-supported HTML tags (<b>, <i>, <u>, <s>, <code>, <pre>). 
- DO NOT use any Markdown formatting (no **, ##, or backticks).
- Telegram does not support standard HTML headers (like <h1>). Simulate headers by using bold text and double line breaks (e.g., <b>The Golden Rule:</b>\n\n). 
- Required sections: <b>The Golden Rule</b> (overview), <b>How We Build It</b> (syntax/structure), and <b>Oopsie Daisies!</b> (common mistakes).

Edge Cases: Always include a <b>Special Magic Exceptions</b> section where the rule acts a little silly or breaks the normal pattern.

Strict HTML Constraints:
- You MUST escape any literal '<', '>', and '&' symbols in your text or code blocks as '&lt;', '&gt;', and '&amp;' to prevent HTML parsing crashes.
- Keep the total response length well under 4000 characters.
- Do not be overly wordy; focus on bright, engaging, and clear explanations.
- Maintain the warm, elementary teacher persona throughout.
- Translate the entire final response to Russian.`

const connectionTtlInSeconds = 240

type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model string `json:"model"`
	Messages []Message `json:"messages"`
	Temperature float64 `json:"temperature"`
}

type ChatResponse struct {
 	Choices []struct{
  Message Message `json:"message"`
  FinishReason string `json:"finish_reason"`
  } `json:"choices"`
}

type EnglishTutor interface {
	GetLesson(ctx context.Context) (string, error)
}

type LlmEnglishTutor struct {
	BaseUrl string
	HTTPClient *http.Client
}

func NewLlmEnglishTutor(baseUrl string) *LlmEnglishTutor {
	return &LlmEnglishTutor{
		BaseUrl: baseUrl,
		HTTPClient: &http.Client{
			Timeout: time.Second*connectionTtlInSeconds,
		},
	}
}


func (t *LlmEnglishTutor) GetLesson(ctx context.Context) (string, error) {
	url := fmt.Sprintf("%s/v1/chat/completions", t.BaseUrl)
	randomIndex := rand.Intn(len(lessons))
	lesson := lessons[randomIndex]

	var msgs []Message
	msgs = append(msgs, Message{Role: "user", Content: fmt.Sprintf(prompt, lesson)})
	reqBody := ChatRequest{
			Model:       "local-model", // LM Studio uses whatever model is loaded
			Messages:    msgs,
			Temperature: 0.7,
		}

		jsonData, err := json.Marshal(reqBody)
		if err != nil {
			return "", err
		}

		// Create the HTTP Request
		req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return "", err
		}
		req.Header.Set("Content-Type", "application/json")

		// Execute the request
		resp, err := t.HTTPClient.Do(req)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("API Error (Status %d): %s", resp.StatusCode, string(body))
		}

		// Parse the JSON response
		var chatResp ChatResponse
		if err := json.Unmarshal(body, &chatResp); err != nil {
			return "", err
		}

		if len(chatResp.Choices) == 0 {
			return "", fmt.Errorf("no choices returned from model")
		}

		return chatResp.Choices[0].Message.Content, nil
}