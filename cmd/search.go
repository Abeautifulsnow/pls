package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// 结构使用 https://raw.githubusercontent.com/jaywcjlove/linux-command/master/dist/data.json
type commandIndex struct {
	Name        string `json:"n"`
	Path        string `json:"p"`
	Description string `json:"d"`
}

func NewSearchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search <command>",
		Short: "Search command by keywords",
		Run: func(_ *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("[sorry]: the search command does not accept any keywords")
				return
			}
			searchCmd(args[0])
		},
	}
	return cmd
}

// unmarshalIndex get all command descriptions list.
func unmarshalIndex() map[string]commandIndex {
	resp, err := returnResp(commandDataJson)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		if os.IsTimeout(err) {
			log.Fatalln(ErrTimeout)
		}
		log.Fatalln(err)
	}

	content := make([]byte, 0)
	reader := bufio.NewReader(resp.Body)
	for {
		line, _, err := reader.ReadLine()
		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}
		if err == io.EOF {
			break
		}

		content = append(content, line...)
	}

	ret := make(map[string]commandIndex)
	_ = json.Unmarshal(content, &ret)

	return ret
}

// searchCmd search the specified command.
func searchCmd(keywords string) {
	keywords = strings.ToLower(keywords)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"command", "description"})
	table.SetRowLine(true)

	for k, v := range unmarshalIndex() {
		desc := strings.ToLower(v.Description)
		if strings.Contains(desc, keywords) || strings.Contains(strings.ToLower(k), keywords) {
			table.Append([]string{k, v.Description})
		}
	}

	table.Render()
}
