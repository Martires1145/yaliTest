package util

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
)

var successMsg = viper.GetString("script.successMsg")

func RunCmd(name string, args []string, runState chan string) {
	cmd := exec.Command(name, args...)
	closed := make(chan struct{})
	defer close(closed)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stdoutPipe.Close()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stderr.Close()

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			data, err := simplifiedchinese.GB18030.NewDecoder().Bytes(scanner.Bytes())
			if err != nil {
				fmt.Println("transfer serror with bytes:", scanner.Bytes())
				runState <- err.Error()
				break
			}
			if string(data) == successMsg {
				runState <- successMsg
				break
			}
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			data, err := simplifiedchinese.GB18030.NewDecoder().Bytes(scanner.Bytes())
			if err != nil {
				runState <- err.Error()
				break
			}
			runState <- string(data)
			break
		}
	}()

	if err := cmd.Run(); err != nil {
		runState <- err.Error()
	}
}
