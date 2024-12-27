package jira

import (
	"bufio"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetCurrentSprintLine() string {
	sprint_list := make([]string, 0, 4)

	cmd := exec.Command("jira", "sprint", "list", "--state=active", "--table", "--plain")
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		slog.Error("GetCurrentSprintLine", "error", err)
		os.Exit(1)
	}

	cmd.Start()

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() { // get stdout text for each lines
		s := scanner.Text()
		slog.Debug("GetCurrentSprintLine", "scan content", s)
		sprint_list = append(sprint_list, s)
	}

	cmd.Wait()

	return sprint_list[1]
}

func GetFutureSprintList() []string {
	sprint_list := make([]string, 0, 16)

	cmd := exec.Command("jira", "sprint", "list", "--state=future", "--table", "--plain")
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		slog.Error("GetFutureSprintList", "error", err)
		os.Exit(1)
	}

	cmd.Start()

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() { // get stdout text for each lines
		s := scanner.Text()
		slog.Debug("GetFutureSprintList", "scan content", s)
		sprint_list = append(sprint_list, s)
	}

	cmd.Wait()

	return sprint_list[1:] // cut header line
}

func CreateTicket(param CreateParams) string {
	// prepare option strings
	options := make([]string, 0, 16)
	options = append(options, "-c")
	options = append(options, "jira")
	options = append(options, "issue")
	options = append(options, "create")
	options = append(options, "--no-input")
	options = append(options, "-t"+param.Type)
	options = append(options, "-s'"+param.Summary+"'")

	// decide assigner
	if param.Assign != "" {
		options = append(options, "-a'"+param.Assign+"'")
	} else {
		me := GetMe()
		options = append(options, "-a'"+me+"'")
	}

	if param.Body != "" {
		options = append(options, "-b'"+param.Body+"'")
	}

	if param.Priority != "" {
		options = append(options, "-y"+param.Priority)
	}

	if len(param.Labels) != 0 {
		for _, v := range param.Labels {
			options = append(options, "-l'"+v+"'")
		}
	}

	if param.Epic != "" {
		options = append(options, "-P'"+param.Epic+"'")
	}

	if param.TemplatePath != "" {
		options = append(options, "--template")
		options = append(options, "'"+param.TemplatePath+"'")
	}

	if param.StoryPoints != 0 {
		i := strconv.Itoa(int(param.StoryPoints))
		options = append(options, "--custom")
		options = append(options, "story-points="+i)
	}

	slog.Debug("CreateTicket", "options", strings.Join(options, " "))

	// publish Command
	result_lines := make([]string, 0, 4)

	cmd := exec.Command("sh", options...)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		slog.Error("CreateTicket", "error", err)
		os.Exit(1)
	}

	cmd.Start()

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() { // get stdout text for each lines
		s := scanner.Text()
		slog.Debug("CreateTicket", "scan content", s)
		result_lines = append(result_lines, s)
	}

	cmd.Wait()

	if len(result_lines) == 0 {
		slog.Error("CreateTicket", "error", "no strings received from jira issue create")
	}

	// return ticket id
	return GetTicketID(result_lines)
}

func GetMe() string {
	out, err := exec.Command("jira", "me").Output()
	if err != nil {
		slog.Error("GetMe", "error", err)
		os.Exit(1)
	}
	return strings.Trim(string(out), " \n\t")
}

func MoveTicketToSprint(ticket_id string, sprint_id string) error {
	result_lines := make([]string, 0, 4)

	cmd := exec.Command("jira", "sprint", "add", sprint_id, ticket_id)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		slog.Error("MoveTicketToSprint", "error", err)
		os.Exit(1)
	}

	cmd.Start()

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() { // get stdout text for each lines
		s := scanner.Text()
		slog.Debug("MoveTicketToSprint", "scan content", s)
		result_lines = append(result_lines, s)
	}

	cmd.Wait()

	for _, v := range result_lines {
		if strings.Contains(v, "Issues added to") {
			return nil
		}
	}

	fmt.Printf("ticket_id: %s, sprint_id: %s\n", ticket_id, sprint_id)

	return errors.New("moving tickets to specified sprint is failed")
}
