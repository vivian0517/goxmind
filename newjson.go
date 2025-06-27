package goxmind

// XMindSheet represents a single sheet in the XMind file
type XMindSheet struct {
	ID               string           `json:"id,omitempty"`
	RevisionID       string           `json:"revisionId,omitempty"`
	Class            string           `json:"class,omitempty"`
	RootTopic        *XMindTopic      `json:"rootTopic,omitempty"`
	Title            string           `json:"title,omitempty"`
	Style            *XMindStyle      `json:"style,omitempty"`
	TopicOverlapping string           `json:"topicOverlapping,omitempty"`
	TopicPositioning string           `json:"topicPositioning,omitempty"`
	Extensions       []XMindExtension `json:"extensions,omitempty"`
	Theme            *XMindTheme      `json:"theme,omitempty"`
}

// XMindTopic represents a topic in the XMind file
type XMindTopic struct {
	ID              string                 `json:"id,omitempty"`
	Class           string                 `json:"class,omitempty"`
	Title           string                 `json:"title,omitempty"`
	AttributedTitle []XMindAttributedTitle `json:"attributedTitle,omitempty"`
	CustomWidth     float64                `json:"customWidth,omitempty"`
	Position        *XMindPosition         `json:"position,omitempty"`
	StructureClass  string                 `json:"structureClass,omitempty"`
	Children        *XMindChildren         `json:"children,omitempty"`
	Notes           *XMindNotes            `json:"notes,omitempty"`
	Markers         []XMindMarker          `json:"markers,omitempty"`
	TitleUnedited   bool                   `json:"titleUnedited,omitempty"`
}

// XMindAttributedTitle represents attributed title text
type XMindAttributedTitle struct {
	Text string `json:"text,omitempty"`
}

// XMindPosition represents position coordinates
type XMindPosition struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

// XMindChildren represents children of a topic
type XMindChildren struct {
	Attached []*XMindTopic `json:"attached,omitempty"`
}

// XMindNotes represents notes content
type XMindNotes struct {
	Plain    *XMindPlainContent `json:"plain,omitempty"`
	RealHTML *XMindHTMLContent  `json:"realHTML,omitempty"`
}

// XMindPlainContent represents plain text content
type XMindPlainContent struct {
	Content string `json:"content,omitempty"`
}

// XMindHTMLContent represents HTML content
type XMindHTMLContent struct {
	Content string `json:"content,omitempty"`
}

// XMindMarker represents a marker/icon
type XMindMarker struct {
	MarkerID string `json:"markerId,omitempty"`
}

// XMindStyle represents style information
type XMindStyle struct {
	ID         string                 `json:"id,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// XMindExtension represents extension information
type XMindExtension struct {
	Provider string                 `json:"provider,omitempty"`
	Content  map[string]interface{} `json:"content,omitempty"`
}

// XMindTheme represents theme information
type XMindTheme struct {
	Map             *XMindThemeMap `json:"map,omitempty"`
	SkeletonThemeID string         `json:"skeletonThemeId,omitempty"`
	ColorThemeID    string         `json:"colorThemeId,omitempty"`
}

// XMindThemeMap represents theme map information
type XMindThemeMap struct {
	ID         string                 `json:"id,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// NewJson represents the new XMind JSON format as an array of sheets
type NewJson []*XMindSheet
