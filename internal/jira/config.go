package jira

type SprintData struct {
	SprintNum     int32
	CurrentSprint bool
	NextSprint    bool
	FutureSprint  int32
	BackLog       bool
}

type CreateParams struct {
	Summary      string
	Body         string
	Type         string
	Priority     string
	Labels       []string
	Epic         string
	StoryPoints  uint8
	SprintID     uint16
	SprintData   SprintData
	Assign       string
	TemplatePath string
}
