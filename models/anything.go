package models

// AnythingResponse defines the structure of the /anything response
type AnythingResponse struct {
	Method  string                `json:"method"`
	URL     string                `json:"url"`
	Headers map[string][]string   `json:"headers"`
	Args    map[string]string     `json:"args,omitempty"`
	Data    string                `json:"data,omitempty"`
	JSON    interface{}           `json:"json,omitempty"`
	Form    map[string][]string   `json:"form,omitempty"`
	Files   map[string][]FileInfo `json:"files,omitempty"`
}

// FileInfo holds information about uploaded files
type FileInfo struct {
	Filename string              `json:"filename"`
	Size     int64               `json:"size"`
	Header   map[string][]string `json:"header"`
}
