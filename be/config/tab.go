package config

type Tab struct {
	ID       string `json:"id"`
	Icon     string `json:"icon"`
	Title    string `json:"title"`
	Endpoint string `json:"endpoint"`
}

var DefaultTabs = []Tab{
	{ID: "timeline", Title: "Timeline", Icon: "fa-solid fa-timeline", Endpoint: "/api/ai/timeline"},
	{ID: "characterdetails", Title: "Character Details", Icon: "fa-solid fa-users", Endpoint: "/api/ai/charactersdetails"},
	{ID: "conversations", Title: "Conversations", Icon: "fa-solid fa-comments", Endpoint: "/api/ai/conversations"},
}
