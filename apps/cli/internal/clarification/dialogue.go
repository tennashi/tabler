package clarification

import (
	"context"
)

// Exchange represents a question-answer pair in dialogue
type Exchange struct {
	Question string
	Answer   string
}

// DialogueSession maintains the state of a clarification dialogue
type DialogueSession struct {
	OriginalInput   string
	CurrentQuestion string
	History         []Exchange
	ExtractedInfo   map[string]string
	IsComplete      bool
	SkipRequested   bool
}

// QuestionGenerator creates contextual questions
type QuestionGenerator interface {
	GenerateQuestion(ctx context.Context, session *DialogueSession) (question string, isComplete bool, err error)
}

// ResponseProcessor extracts information from user responses
type ResponseProcessor interface {
	ProcessResponse(session *DialogueSession, response string) error
	ExtractInfo(session *DialogueSession) map[string]string
	BuildFinalTask(session *DialogueSession) string
	DetectsSkip(response string) bool
}

// DialogueManager orchestrates the clarification conversation
type DialogueManager struct {
	detector      *VaguenessDetector
	questionGen   QuestionGenerator
	processor     ResponseProcessor
	maxExchanges  int
}

// NewDialogueManager creates a new dialogue manager
func NewDialogueManager(
	detector *VaguenessDetector,
	questionGen QuestionGenerator,
	processor ResponseProcessor,
) *DialogueManager {
	return &DialogueManager{
		detector:     detector,
		questionGen:  questionGen,
		processor:    processor,
		maxExchanges: 3,
	}
}

// StartDialogue begins a clarification dialogue if needed
func (m *DialogueManager) StartDialogue(ctx context.Context, input string) (*DialogueSession, error) {
	// Check if input needs clarification
	isVague, _ := m.detector.DetectVagueness(input)
	if !isVague {
		return nil, nil // No dialogue needed
	}
	
	// Create new session
	session := &DialogueSession{
		OriginalInput: input,
		History:       []Exchange{},
		ExtractedInfo: make(map[string]string),
	}
	
	// Generate first question
	question, isComplete, err := m.questionGen.GenerateQuestion(ctx, session)
	if err != nil {
		return nil, err
	}
	
	if isComplete {
		session.IsComplete = true
		return session, nil
	}
	
	session.CurrentQuestion = question
	session.History = append(session.History, Exchange{Question: question, Answer: ""})
	
	return session, nil
}

// ProcessResponse handles user's answer and continues dialogue
func (m *DialogueManager) ProcessResponse(ctx context.Context, session *DialogueSession, response string) error {
	// Check for skip intent
	if m.processor.DetectsSkip(response) {
		session.SkipRequested = true
		session.IsComplete = true
		return nil
	}
	
	// Record the answer
	if len(session.History) > 0 {
		session.History[len(session.History)-1].Answer = response
	}
	
	// Process the response
	if err := m.processor.ProcessResponse(session, response); err != nil {
		return err
	}
	
	// Extract information
	session.ExtractedInfo = m.processor.ExtractInfo(session)
	
	// Check if we've reached exchange limit
	if len(session.History) >= m.maxExchanges {
		session.IsComplete = true
		return nil
	}
	
	// Generate next question
	question, isComplete, err := m.questionGen.GenerateQuestion(ctx, session)
	if err != nil {
		return err
	}
	
	if isComplete {
		session.IsComplete = true
		return nil
	}
	
	// Continue dialogue
	session.CurrentQuestion = question
	session.History = append(session.History, Exchange{Question: question, Answer: ""})
	
	return nil
}

// GetFinalTask constructs the final task from dialogue results
func (m *DialogueManager) GetFinalTask(session *DialogueSession) string {
	if session.SkipRequested || len(session.History) == 0 {
		return session.OriginalInput
	}
	
	return m.processor.BuildFinalTask(session)
}