package jira

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func extractSprintIDFromLine(s string) string {
	id := strings.Split(s, "\t")[0] // extract SprintID
	return strings.Trim(id, " \n")
}

func GetSprintID(data SprintData) string {
	if data.SprintNum != 0 {
		// convert SprintNum to string
		target_string := strconv.FormatUint(uint64(data.SprintNum), 32)

		//get future splint list using jira Command
		sprint_list := GetFutureSprintList()

		for _, l := range sprint_list {
			fmt.Println(l)
			if strings.Contains(l, target_string) { // if is SprintNumber contained to the line
				return extractSprintIDFromLine(l)
			}
		}

		// if reaching this code, it means given SprintNum is not found
		fmt.Printf("no SprintNum found: %d\n", data.SprintNum)
		os.Exit(1)
	}

	if data.NextSprint {
		data.FutureSprint = 1
	}

	// decide current sprint or future sprint
	if data.FutureSprint == 0 {
		data.CurrentSprint = true
	}

	if data.CurrentSprint {
		sprint_line := GetCurrentSprintLine()
		return extractSprintIDFromLine(sprint_line)
	}

	//get future splint list using jira Command
	sprint_list := GetFutureSprintList()

	// check given future-sprint param is valid or not
	if len(sprint_list)-int(data.FutureSprint) < 0 {
		slog.Info("GetSprintID: given future-sprint value is greater than fetched the number of future sprint on the JIRA Server")
		slog.Info("GetSprintID", "future-sprint", data.FutureSprint)
		slog.Info("GetSprintID", "number of future sprint on JIRA", len(sprint_list))
		os.Exit(1)
	}

	// detect what line is target
	target_line := sprint_list[len(sprint_list)-int(data.FutureSprint)]
	return extractSprintIDFromLine(target_line)
}

func GetTicketID(result []string) string {
	// detect target line
	final_line := result[len(result)-1]
	divided := strings.Split(final_line, "/")

	// extract ticke id
	id := divided[len(divided)-1]
	id = strings.Trim(id, " \n\t") // delete space chars

	return id
}
