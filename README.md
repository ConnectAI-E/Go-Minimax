

<p align='center'>
    <img src='https://github.com/ConnectAI-E/go-minimax/assets/50035229/c87ca385-f9a6-4f81-a0a9-67828af63334' alt='' width='1300'/>
</p>

## MINIMAX-SDK
> 大语言模型 MiniMax API

[😎点击查看官方文档](https://www.yuque.com/minimax/api)

### 功能及特点

1. 全接口字段注释
2. Chatcompletion 文本对话接口
3. Embeddings 向量化接口
4. T2A 文本转语音接口
5. 无缝对接官方文档：单轮问答、历史记忆问答、流返回

### 使用方法

1. 访问 https://api.minimax.chat/document/guides/example 并创建应用。
2. 取 groupID 和 API_token
3. 使用 New 方法并根据传参说明生成 client。

### 示例

```go
package main

import (
	"context"
	"fmt"
	"github.com/ConnectAI-E/go-minimax"
    textv1 "github.com/ConnectAI-E/go-minimax/text/v1"
)

//init client

func main() {
	ctx := context.Background()
	client, _ := baidubce.New(
		WithApiToken(os.Getenv("TEST_MINIMAX_API_TOKEN")),
		WithGroupId(os.Getenv("TEST_MINIMAX_GROUP_ID")),
		)

	//chat
	req := &textv1.ChatCompletionsRequest{
		Messages: []*ai_customv1.Message{
			{
				SenderType: "USER",
				Text:       "hi~",
			},
		},
		Model:       "abab5-chat",
		Temperature: 0.7,
	}
	res, _ := client.ChatCompletions(ctx, req)

	fmt.Println(res.Result) // output: 嗨！有什么我可以帮助你的吗？
}

```
