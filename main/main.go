package main

func main() {
	// editor := os.Getenv("EDITOR")
	// if editor == "" {
	// 	editor = "vim"
	// }
	// cmd := exec.Command(editor, "a.txt")
	// cmd.Stdin = os.Stdin
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// err := cmd.Run()
	// if err != nil {
	// 	panic(err)
	// }

	// githubissue.CreateIssue(githubissue.GITHUB_OWNER, githubissue.GITHUB_REPO, "new issue hello", "this is new world!!!")

	// issue, _ := githubissue.ReadIssue(githubissue.GITHUB_OWNER, githubissue.GITHUB_REPO, "2")
	// fmt.Printf("head: %s\nbody: %s\nstate: %s\n", issue.Title, issue.Body, issue.State)

	// issue, _ := githubissue.UpdateIssue(githubissue.GITHUB_OWNER, githubissue.GITHUB_REPO,
	// 	"2", "changed 2title!", "changed 2body!", false)
	// fmt.Printf("head: %s\nbody: %s\nstate: %s\n", issue.Title, issue.Body, issue.State)

	// _, err := githubissue.CloseIssue(githubissue.GITHUB_OWNER, githubissue.GITHUB_REPO, "2")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Issue closed successfully.")
	// githubissue.AddComment(githubissue.GITHUB_OWNER, githubissue.GITHUB_REPO, "2", "hello3!")
}
