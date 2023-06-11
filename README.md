# OPENAI-GO
[![Go Reference](https://pkg.go.dev/badge/github.com/sashabaranov/go-openai.svg)](https://pkg.go.dev/github.com/sashabaranov/go-openai)
[![Go Report Card](https://goreportcard.com/badge/github.com/sashabaranov/go-openai)](https://goreportcard.com/report/github.com/sashabaranov/go-openai)
[![codecov](https://codecov.io/gh/sashabaranov/go-openai/branch/master/graph/badge.svg?token=bCbIfHLIsW)](https://codecov.io/gh/sashabaranov/go-openai)


# Example
```go
package main

import (
	"context"
	"fmt"
	"openai-go/openai"
)

func main() {
	// 一定要设置环境变量OPEN_API_KEY
	ai := openai.New(nil)
	ctx := context.Background()
	// 1.聊天
	input := "你是谁"
	count := 2 // 表示生成2个答案
	result, err := ai.ChatCompletions(ctx, input, 2)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
		// [我是一个人工智能语言模型，被称为OpenAI的GPT-3。 我是AI语言模型，一个能够自动回答问题和生成文本的人工智能程序。]
	}

	// 2.生成图片
	input = "韩国美女"
	result, err = ai.Images(ctx, input, count)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
		// 	[
		//	https://oaidalleapiprodscus.blob.core.windows.net/private/org-9MT5FUMC11U2zbyI5yIjpL1P/user-938N4hMdCq02FcSfB224yezE/img-YxKQf13QBxt9eFO31LAvmlKx.png?st=2023-03-26T13%3A02%3A42Z&se=2023-03-26T15%3A02%3A42Z&sp=r&sv=2021-08-06&sr=b&rscd=inline&rsct=image/png&skoid=6aaadede-4fb3-4698-a8f6-684d7786b067&sktid=a48cca56-e6da-484e-a814-9c849652bcb3&skt=2023-03-25T16%3A32%3A34Z&ske=2023-03-26T16%3A32%3A34Z&sks=b&skv=2021-08-06&sig=eVoEHWLcQ8pwjm7WVPkTMMgF8wv67XjnT5Ryq24kQiw%3D
		//	https://oaidalleapiprodscus.blob.core.windows.net/private/org-9MT5FUMC11U2zbyI5yIjpL1P/user-938N4hMdCq02FcSfB224yezE/img-O1n2SlAlljmLfnQPg1rbB9ZX.png?st=2023-03-26T13%3A02%3A42Z&se=2023-03-26T15%3A02%3A42Z&sp=r&sv=2021-08-06&sr=b&rscd=inline&rsct=image/png&skoid=6aaadede-4fb3-4698-a8f6-684d7786b067&sktid=a48cca56-e6da-484e-a814-9c849652bcb3&skt=2023-03-25T16%3A32%3A34Z&ske=2023-03-26T16%3A32%3A34Z&sks=b&skv=2021-08-06&sig=BDGjc4SzkLcEblytsLso0UMdL0s4u/xPxdCkmo/q2bk%3D
		//	]
	}
}

```

# Get Start
下载代码
```
git clone git@github.com:openfaas/faas.git
```

生成ai对象
```
// 一定要设置环境变量OPEN_API_KEY
ai := openai.New(nil)

// 或者通过config传进去
cfg := &config.Config{
		OpenAI: config.OpenAI{
			OpenAPIKey: "YOUR_OPEN_API_KEY",
		},
	}

ai := New(cfg)

// 或者使用openai包里面提供的默认方法，此时也需要设置环境变量OPEN_API_KEY
result, err := openai.ChatCompletions(ctx, "你是谁",1)
```

# Contributing
If you find any errors or suggestions for improvement, please submit a Pull Request or Issue.

# Licence
This project is open source under the MIT license.

# Author
@ZBIGBEAR