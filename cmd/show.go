package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type config struct {
	Dir string `json:"dir"`
}

var (
	homeDir     = getHomedir()
	commandPath = filepath.Join(homeDir, commandDir)
	configPath  = filepath.Join(homeDir, commandDir, commandCfg)
	defaultCfg  = config{Dir: commandPath}
	lineWidth   = 80
	lineLeftPad = 2
)

func NewShowCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show <command>",
		Short: "Show the specified command usage.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("[sorry]: the show command does not accept any arguments")
				return
			}
			force, _ := cmd.Flags().GetBool("force")
			showCmd(args[0], force)
		},
	}
	cmd.Flags().BoolP("force", "f", false, "force to refresh command usage from remote.")
	return cmd
}

func showCmd(cmd string, force bool) {
	cmd = strings.ToLower(cmd)

	if force {
		if err := retryDownloadCmd(cmd); err != nil {
			if errors.Is(err, ErrCommandNotFound) {
				fmt.Printf("[sorry]: could not found command <%s>\n", cmd)
				return
			}
			fmt.Printf("[sorry]: failed to download command <%s>\n", cmd)
		}
	}

	cfg, err := getConfigContent()
	if err != nil {
		fmt.Println("[sorry]: failed to get config content")
		return
	}

	p := path.Join(cfg.Dir, fmt.Sprintf("%s.md", cmd))
	if !isFileExist(p) {
		err := retryDownloadCmd(cmd)
		if err != nil {
			fmt.Printf("[sorry]: failed to retrieve command <%s>\n", cmd)
			return
		}
		if errors.Is(err, ErrCommandNotFound) {
			fmt.Printf("[sorry]: could not found command <%s>\n", cmd)
			return
		}
	}

	source, err := os.ReadFile(p)
	if err != nil {
		fmt.Printf("[sorry]: failed to open file <%s>\n", p)
		return
	}
	markdown.BlueBgItalic = color.New(color.FgBlue).SprintFunc()
	result := markdown.Render(string(source), lineWidth, lineLeftPad)
	fmt.Println(string(result))
}

// getHomedir get the home directory
func getHomedir() string {
	home, _ := os.UserHomeDir()
	return home
}

func makeCmdDir(dir string) error {
	if _, err := os.Stat(dir); err != nil && !os.IsExist(err) {
		return os.Mkdir(dir, 0755)
	}
	return nil
}

func isFileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func genConfigFile() error {
	if err := makeCmdDir(commandPath); err != nil {
		return err
	}

	if !isFileExist(configPath) {
		bs, _ := json.Marshal(defaultCfg)
		return os.WriteFile(configPath, bs, 0666)
	}

	return nil
}

func getConfigContent() (*config, error) {
	if err := genConfigFile(); err != nil {
		return nil, err
	}

	bs, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	cfgPath := config{}
	if err := json.Unmarshal(bs, &cfgPath); err != nil {
		return nil, err
	}

	return &cfgPath, nil
}

func retryDownloadCmd(cmd string) error {
	for j := 0; j < maxRetry; j++ {
		if err := downloadCmd(cmd); err != nil {
			fmt.Println("[Info]: Retrying to download at", j+1, "time.")
			continue
		}
		break
	}

	return nil
}

// downloadCmd download the command content to local disk.
func downloadCmd(cmd string) error {
	c, err := getConfigContent()
	if err != nil {
		return err
	}

	if err := makeCmdDir(c.Dir); err != nil {
		return err
	}

	resp, err := returnResp(fmt.Sprintf(commandUrl, cmd))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("[Error]: Download", cmd, "error", err.Error())
			return
		}
	}(resp.Body)

	content := make([]byte, 0)
	reader := bufio.NewReader(resp.Body)
	for {
		line, _, err := reader.ReadLine()
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			break
		}
		content = append(content, line...)
		content = append(content, []byte("\n")...)
	}

	fp := path.Join(c.Dir, fmt.Sprintf("%s.md", cmd))
	return os.WriteFile(fp, content, 0666)
}
