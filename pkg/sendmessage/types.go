package sendmessage

type WhatsApp struct {
	From     string     `json:"from"`
	To       string     `json:"to"`
	Contents []Contents `json:"contents"`
}

type Contents struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
