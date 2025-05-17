package goxmind

type NewJson []struct {
	ID        string `json:"id,omitempty"`
	Class     string `json:"class,omitempty"`
	RootTopic struct {
		ID             string `json:"id,omitempty"`
		Class          string `json:"class,omitempty"`
		Title          string `json:"title"`
		StructureClass string `json:"structureClass,omitempty"`
		Href           string `json:"href,omitempty"`
		Markers        []struct {
			MarkerId string `json:"markerId,omitempty"`
		} `json:"markers,omitempty"`
		Notes struct {
			Plain struct {
				Content string `json:"content,omitempty"`
			} `json:"plain,omitempty"`
		} `json:"notes,omitempty"`

		Childrennew struct {
			Attached []struct {
				ID      string `json:"id,omitempty"`
				Title   string `json:"title"`
				Href    string `json:"href,omitempty"`
				Markers []struct {
					MarkerId string `json:"markerId,omitempty"`
				} `json:"markers,omitempty"`
				Notes struct {
					Plain struct {
						Content string `json:"content,omitempty"`
					} `json:"plain,omitempty"`
				} `json:"notes,omitempty"`
				Childrennew struct{} `json:"children,omitempty"`
			} `json:"attached,omitempty"`
		} `json:"children,omitempty"`
	} `json:"rootTopic,omitempty"`
	Title string `json:"title,omitempty"`
}
