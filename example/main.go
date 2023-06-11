package main

import (
	"context"
	"fmt"
	"openai-go/common"
	"openai-go/openai"
	"os"
	"os/exec"
)

func main() {
	clear()
	// 一定要设置环境变量OPEN_API_KEY
	ai := openai.New(nil)

	args := os.Args[1:]
	input := ""
	if len(args) == 6 {
		// goland插件：通过goland extend获取选中的代码，发送到chatgpt
		if args[1] == args[3] && args[2] == args[4] {
			input = args[5]
		} else {
			fileArgs := args[:len(args)-1]
			inputArgs := args[len(args)-1]
			selectContent, err := common.ExtractFile(fileArgs[0], fileArgs[1], fileArgs[2], fileArgs[3], fileArgs[4])
			if err != nil {
				fmt.Println(err)
				return
			}
			input = string(selectContent) + " " + inputArgs
		}
	} else {
		input = args[0]
	}

	fmt.Printf("我：%s\n", input)

	if len(input) != 0 {
		result, err := common.Async(func(ctx context.Context) (any, error) {
			result, err := ai.ChatCompletions(ctx, input)
			if err != nil {
				return nil, err
			}

			return result[0], nil
		}, ".")
		fmt.Println()
		// 打印结果
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}

	//// 2.生成图片
	//input = "韩国美女"
	//result, err = ai.Images(ctx, input, count)
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println(result)
	//	// 	[
	//	//	https://oaidalleapiprodscus.blob.core.windows.net/private/org-9MT5FUMC11U2zbyI5yIjpL1P/user-938N4hMdCq02FcSfB224yezE/img-YxKQf13QBxt9eFO31LAvmlKx.png?st=2023-03-26T13%3A02%3A42Z&se=2023-03-26T15%3A02%3A42Z&sp=r&sv=2021-08-06&sr=b&rscd=inline&rsct=image/png&skoid=6aaadede-4fb3-4698-a8f6-684d7786b067&sktid=a48cca56-e6da-484e-a814-9c849652bcb3&skt=2023-03-25T16%3A32%3A34Z&ske=2023-03-26T16%3A32%3A34Z&sks=b&skv=2021-08-06&sig=eVoEHWLcQ8pwjm7WVPkTMMgF8wv67XjnT5Ryq24kQiw%3D
	//	//	https://oaidalleapiprodscus.blob.core.windows.net/private/org-9MT5FUMC11U2zbyI5yIjpL1P/user-938N4hMdCq02FcSfB224yezE/img-O1n2SlAlljmLfnQPg1rbB9ZX.png?st=2023-03-26T13%3A02%3A42Z&se=2023-03-26T15%3A02%3A42Z&sp=r&sv=2021-08-06&sr=b&rscd=inline&rsct=image/png&skoid=6aaadede-4fb3-4698-a8f6-684d7786b067&sktid=a48cca56-e6da-484e-a814-9c849652bcb3&skt=2023-03-25T16%3A32%3A34Z&ske=2023-03-26T16%3A32%3A34Z&sks=b&skv=2021-08-06&sig=BDGjc4SzkLcEblytsLso0UMdL0s4u/xPxdCkmo/q2bk%3D
	//	//	]
	//}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
