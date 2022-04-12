/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "随机讲个笑话",
	Long:  `随机讲个笑话.`,
	Run: func(cmd *cobra.Command, args []string) {
		getAJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Joke struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Joke string `json:"joke"`
}

func getAJoke() {
	ctx := context.Background()

	joke := Joke{}
	respBytes := getJokeFromWhan()
	if err := json.Unmarshal(respBytes, &joke); err != nil {
		g.Log().Error(ctx, "无法反序列化")
	}

	fmt.Printf("%s\n", joke.Joke)

}

func getJokeFromWhan() []byte {
	ctx := context.Background()
	url := "https://api.vvhan.com/api/xh?type=json"

	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil)
	if err != nil {
		g.Log().Error(ctx, "异常，无法发起连接")
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "dadjoke")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		g.Log().Error(ctx, "异常，无法发起连接")
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		g.Log().Error(ctx, "获取内容失败")
	}

	// fmt.Printf("%s", string(content))

	return content
}
