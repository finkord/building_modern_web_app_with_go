package models

// TemplateData holds data sent from the handler to the template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float64
	Data      map[string]interface{} // any?
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
