package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/studio-b12/gowebdav"
	"os"
	"strings"
)

var (
	url       string
	user      string
	pwd       string
	files     string
	remoteDir string
)

var uploadCmd = &cobra.Command{
	Use:   "Upload CLI",
	Short: "uc",
	Long:  `Upload CLI WebDav`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if required flags are provided
		if url == "" || files == "" || remoteDir == "" {
			fmt.Println("Error: Required flags are missing")
			cmd.Usage()
			os.Exit(1)
		}
		c := gowebdav.NewClient(url, user, pwd)
		c.Connect()
		parts := strings.Split(files, ",")

		for _, filePath := range parts {
			file, _ := os.Open(filePath)
			err := c.WriteStream(remoteDir+"/"+filePath, file, 0644)
			if err != nil {
				fmt.Println(err)
			}
			file.Close()
		}
		// Implement file upload logic here
		fmt.Printf("Uploading file %s to %s%s\n", files, url, remoteDir)
	},
}

func init() {
	uploadCmd.Flags().StringVarP(&url, "url", "u", "", "URL")
	uploadCmd.Flags().StringVarP(&user, "user", "a", "", "User")
	uploadCmd.Flags().StringVarP(&pwd, "pwd", "p", "", "Password")
	uploadCmd.Flags().StringVarP(&files, "files", "f", "", "Files:file1.txt,file2.txt")
	uploadCmd.Flags().StringVarP(&remoteDir, "remote-dir", "r", "", "/data/../auto+file")
}

func main() {
	if err := uploadCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
