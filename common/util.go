package common

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"os/exec"
	"time"
)

const (
	// 参数依次是：$startLine,$startColumn,$endLine,$endColumn,$file
	ExtractFileScript = `echo $(awk -v sline="%s" -v scol="%s" -v eline="%s" -v ecol="%s" 'NR>=sline && NR<=eline {if(NR==sline && NR==eline) {print substr($0,scol,ecol-scol+1)} else if(NR==sline) {print substr($0,scol)} else if(NR==eline) {print substr($0,1,ecol)} else {print}}' "%s")`
)

// Async 异步执行程序，并返回结果。在等待过程中打印动态加载的提示
func Async(f func(ctx context.Context) (any, error), prompt string) (result any, err error) {
	c := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer func() {
			c <- struct{}{}
		}()
		result, err = f(ctx)
	}()

	// 每隔1s打印提示语
	ticker := time.NewTicker(time.Second)
	// 设置1分钟超时
	timeout := time.NewTicker(time.Minute)
	for {
		select {
		case <-c:
			return
		case <-ticker.C:
			if len(prompt) > 0 {
				fmt.Print(prompt)
			}
		case <-timeout.C:
			cancel()
			return nil, TimeOutErr
		}
	}
}

func ExtractFile(fileName, beginLine, beginColumn, endLine, endColumn string) ([]byte, error) {
	cmdName := fmt.Sprintf(ExtractFileScript, beginLine, beginColumn, endLine, endColumn, fileName)
	fmt.Println("cmdName:", cmdName)
	cmd := exec.Command("sh", "-c", cmdName)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, errors.Wrapf(err, "fileName:%s, beginLine:%s, beginColumn:%s, endLine:%s, endColumn:%s", fileName, beginLine, beginColumn, endLine, endColumn)
	}

	return out.Bytes(), nil
}
