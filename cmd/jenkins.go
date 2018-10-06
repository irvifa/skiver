// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/irvifa/skiver/builder"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// jenkinsCmd represents the jenkins command
var jenkinsCmd = &cobra.Command{
	Use:   "jenkins",
	Short: "Jenkins related utilities",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: buildPipeline,
}

func init() {
	jenkinsCmd.PersistentFlags().String("name", "", "Jenkin's Job")
	jenkinsCmd.PersistentFlags().String("pipeline", "", "Jenkin's Pipeline")
	rootCmd.AddCommand(jenkinsCmd)
}

func buildPipeline(cmd *cobra.Command, args []string) {
	filename, _ := filepath.Abs(".jenkins.yaml")
	jenkinsConfig, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config builder.JenkinsConfig

	err = yaml.Unmarshal(jenkinsConfig, &config)

	var url = fmt.Sprintf("%s/job/%s/job/%s/build?delay=0sec", config.JenkinsURL,
		cmd.Flag("pipeline").Value.String(),
		cmd.Flag("name").Value.String())
	println(url)
	sentRequest(config, url)
}

func sentRequest(config builder.JenkinsConfig, url string) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	req.SetBasicAuth(config.Username, config.Password)
	println("Set basic auth")
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	println(s)
}
