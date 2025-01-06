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
	Components   []string
	Epic         string
	StoryPoints  uint8
	Project      string
	SprintID     uint16
	SprintData   SprintData
	Assignee     string
	Reporter     string
	TemplatePath string
}
