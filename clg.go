package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var stdout, stderr bytes.Buffer
	tag := "HEAD"

	if len(os.Args) > 1 {
		tag = os.Args[1]
	}

	cmd := exec.Command("git", "log", "--format=%h,%cI,%cn,%s,%b", tag)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("git log failed with %s:\n>> %s", err, string(stderr.Bytes()))
	}

	fmt.Println(ExtractChangelog(string(stdout.Bytes())))
}

// ExtractChangelog extract changelog from logs
func ExtractChangelog(s string) string {
	var changes []string
	var date string
	logs := strings.Split(s, "\n")

	for index, log := range logs {
		if log != "" {
			_, extDate, author, subject, body := splitLog(log)

			if index == 0 {
				date = extDate
				changes = append(changes, fmt.Sprintf("Changelog for %s:\n", date))
			}

			if strings.Contains(subject, "[CL]") {
				title := extractTitle(subject)

				if title == "" && body != "" {
					title = extractTitle(body)
				}

				if title != "" {
					changes = append(changes, fmt.Sprintf("- %s (by %s)", title, author))
				}
			}
		}
	}

	if len(changes) > 0 {
		return strings.Join(changes, "\n")
	}

	return ""
}

func splitLog(log string) (string, string, string, string, string) {
	info := strings.Split(log, ",")
	fmt.Println(info)
	return info[0], info[1], info[2], info[3], info[4]
}

func extractTitle(body string) string {
	info := strings.Split(body, "title: ")
	if len(info) == 2 {
		return info[1]
	}

	return ""
}
