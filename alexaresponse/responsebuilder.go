package alexaresponse

// Builder builds Alexa responses.
type Builder struct {
	response          Response
	sessionAttributes interface{}
	outputSpeech      *OutputSpeech // using pointers so unused structs don't get marshalled.
	card              *Card
	reprompt          *Reprompt
	directives        []interface{}
	shouldEndSession  bool
}

// NewBuilder returns a new builder object.
func NewBuilder() *Builder {
	return &Builder{
		sessionAttributes: make(map[string]interface{}),
		shouldEndSession:  true,
	}
}

// WithAttributes takes an attributes map and sets the attributes, overwriting any existing attributes.
func (b *Builder) WithAttributes(attributes interface{}) *Builder {
	b.sessionAttributes = attributes
	return b
}

// WithOutputSpeech sets the OutputSpeech object of the builder to the given OutputSpeech object.
func (b *Builder) WithOutputSpeech(outputSpeech OutputSpeech) *Builder {
	b.outputSpeech = &outputSpeech
	return b
}

// WithTextOutputSpeech sets the OutputSpeech type to "PlainText" and sets the text.
func (b *Builder) WithTextOutputSpeech(text string) *Builder {
	if b.outputSpeech == nil {
		b.outputSpeech = &OutputSpeech{}
	}
	b.outputSpeech.Type = "PlainText"
	b.outputSpeech.Text = text
	return b
}

// WithSsmlOutputSpeech sets the OutputSpeech type to "ssml" and sets the ssml encoded string.
func (b *Builder) WithSsmlOutputSpeech(ssml string) *Builder {
	if b.outputSpeech == nil {
		b.outputSpeech = &OutputSpeech{}
	}
	b.outputSpeech.Type = "ssml"
	b.outputSpeech.Ssml = ssml
	return b
}

// OutputSpeechPlayBehavior sets the play behavior of the output speech.
func (b *Builder) OutputSpeechPlayBehavior(behavior string) *Builder {
	if b.outputSpeech == nil {
		b.outputSpeech = &OutputSpeech{}
	}
	b.outputSpeech.PlayBehavior = behavior
	return b
}

// WithCard sets the Card object of the builder to the given Card object.
func (b *Builder) WithCard(card Card) *Builder {
	b.card = &card
	return b
}

// WithSimpleCard creates a simple card.
func (b *Builder) WithSimpleCard(title string, content string) *Builder {
	b.card = &Card{
		Type:    "Simple",
		Title:   title,
		Content: content,
	}
	return b
}

// WithStandardCard creates a standard card.
func (b *Builder) WithStandardCard(title string, text string, image Image) *Builder {
	b.card = &Card{
		Title: title,
		Text:  text,
		Image: &image,
	}
	return b
}

// WithReprompt sets the reprompt object of the builder to the given reprompt object.
func (b *Builder) WithReprompt(reprompt Reprompt) *Builder {
	b.reprompt = &reprompt
	return b
}

// WithTextReprompt creates the reprompt object with a "PlainText" type OutputSpeech.
func (b *Builder) WithTextReprompt(text string) *Builder {
	if b.reprompt == nil {
		b.reprompt = &Reprompt{
			OutputSpeech: &OutputSpeech{},
		}
	}
	b.reprompt.OutputSpeech.Type = "PlainText"
	b.reprompt.OutputSpeech.Text = text
	return b
}

// WithSsmlReprompt creates the reprompt object with a "ssml" type OutputSpeech.
func (b *Builder) WithSsmlReprompt(ssml string) *Builder {
	if b.reprompt == nil {
		b.reprompt = &Reprompt{
			OutputSpeech: &OutputSpeech{},
		}
	}
	b.reprompt.OutputSpeech.Type = "ssml"
	b.reprompt.OutputSpeech.Ssml = ssml
	return b
}

// RepromptPlayBehavior sets the play behavior of the reprompt.
func (b *Builder) RepromptPlayBehavior(behavior string) *Builder {
	if b.reprompt == nil {
		b.reprompt = &Reprompt{
			OutputSpeech: &OutputSpeech{},
		}
	}
	b.reprompt.OutputSpeech.PlayBehavior = behavior
	return b
}

// WithDirectives sets the directives for the response.
func (b *Builder) WithDirectives(directives []interface{}) *Builder {
	b.directives = directives
	return b
}

// ShouldEndSession determines whether the session should end after this request.
func (b *Builder) ShouldEndSession(boolean bool) *Builder {
	b.shouldEndSession = boolean
	return b
}

// Build builds and returns the Response object.
func (b *Builder) Build() Response {
	return Response{
		Version:           "1.0",
		SessionAttributes: b.sessionAttributes,
		Response: &Body{
			OutputSpeech:     b.outputSpeech,
			Card:             b.card,
			Reprompt:         b.reprompt,
			Directives:       b.directives,
			ShouldEndSession: b.shouldEndSession,
		},
	}
}
