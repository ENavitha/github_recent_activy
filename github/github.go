package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchGitHubActivity(username string) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()
	// Handle API errors
	switch resp.StatusCode {
	case 404:
		fmt.Println(" User not found. Please check the username.")
		return
	case 403:
		fmt.Println(" Rate limit exceeded. Try again later.")
		return
	case 500:
		fmt.Println("SGitHub server error. Please try again later.")
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println("Failed to fetch data. Status:", resp.Status)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var events []map[string]interface{}
	if err := json.Unmarshal(body, &events); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	if len(events) == 0 {
		fmt.Println("No recent activity found.")
		return
	}
	// Displaying only the first few events
	fmt.Println("\nRecent GitHub Activity:")
	for i, event := range events {
		if i >= 5 { // Show only the first 5 events
			break
		}
		printEvent(event)
	}

}

// printEvent formats and prints an event based on its type
func printEvent(event map[string]interface{}) {
	eventType := event["type"].(string)
	repoName := event["repo"].(map[string]interface{})["name"].(string)

	switch eventType {
	case "PushEvent":
		commitCount := len(event["payload"].(map[string]interface{})["commits"].([]interface{}))
		fmt.Printf("- Pushed %d commits to %s\n", commitCount, repoName)
	case "IssuesEvent":
		action := event["payload"].(map[string]interface{})["action"].(string)
		fmt.Printf("- %s a new issue in %s\n", capitalize(action), repoName)
	case "WatchEvent":
		fmt.Printf("- Starred %s\n", repoName)
	case "ForkEvent":
		fmt.Printf("- Forked %s\n", repoName)
	case "CreateEvent":
		refType := event["payload"].(map[string]interface{})["ref_type"].(string)
		fmt.Printf("- Created a new %s in %s\n", refType, repoName)
	case "PullRequestEvent":
		action := event["payload"].(map[string]interface{})["action"].(string)
		fmt.Printf("- %s a pull request in %s\n", capitalize(action), repoName)
	case "ReleaseEvent":
		action := event["payload"].(map[string]interface{})["action"].(string)
		fmt.Printf("- %s a release in %s\n", capitalize(action), repoName)
	default:
		fmt.Printf("- %s in %s\n", eventType, repoName)
	}
}

// capitalize makes the first letter uppercase
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}
