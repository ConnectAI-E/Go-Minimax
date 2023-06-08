

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

1. 访问 [MiniMax](https://api.minimax.chat/document/guides/example) 管理后台。
2. 获取取 [groupID](https://api.minimax.chat/basic-information) 和 [API_token](https://api.minimax.chat/basic-information/interface-key)
3. 参考下面示例

### 示例

```go
package main

import (
	"context"
	"fmt"
	"os"
	textv1 "github.com/ConnectAI-E/go-minimax/gen/go/minimax/text/v1"
	"github.com/ConnectAI-E/go-minimax/minimax"
)



func main() {
	ctx := context.Background()
	
	//init client
	client, _ := minimax.New(
		minimax.WithApiToken(os.Getenv("TEST_MINIMAX_API_TOKEN")),
		minimax.WithGroupId(os.Getenv("TEST_MINIMAX_GROUP_ID")),
	)

	//chat
	req := &textv1.ChatCompletionsRequest{
		Messages: []*textv1.Message{
			{
				SenderType: "USER",
				Text:       "hi~",
			},
		},
		Model:       "abab5-chat",
		Temperature: 0.7,
	}
	res, _ := client.ChatCompletions(ctx, req)

	fmt.Println(res.Choices) // output: 你好！有什么我可以帮助你的吗？
}

```


### 快速上手:

<details>
<summary>MiniMax completion</summary>

```go
package main

import (
	"context"
	"fmt"
	"os"
	textv1 "github.com/ConnectAI-E/go-minimax/gen/go/minimax/text/v1"
	"github.com/ConnectAI-E/go-minimax/minimax"
)



func main() {
	ctx := context.Background()
	
	//init client
	client, _ := minimax.New(
		minimax.WithApiToken(os.Getenv("TEST_MINIMAX_API_TOKEN")),
		minimax.WithGroupId(os.Getenv("TEST_MINIMAX_GROUP_ID")),
	)

	//chat
	req := &textv1.ChatCompletionsRequest{
		Messages: []*textv1.Message{
			{
				SenderType: "USER",
				Text:       "hi~",
			},
		},
		Model:       "abab5-chat",
		Temperature: 0.7,
	}
	res, _ := client.ChatCompletions(ctx, req)

	fmt.Println(res.Choices) // output: 你好！有什么我可以帮助你的吗？
}
```
</details>


<details>
<summary>MiniMax stream completion</summary>

```go
package main

import (
	"context"
	"errors"
	"fmt"
	textv1 "github.com/ConnectAI-E/go-minimax/gen/go/minimax/text/v1"
	"github.com/ConnectAI-E/go-minimax/minimax"
	"io"
	"os"
)



func main() {
	ctx := context.Background()
	
	//init client
	client, _ := minimax.New(
		minimax.WithApiToken(os.Getenv("TEST_MINIMAX_API_TOKEN")),
		minimax.WithGroupId(os.Getenv("TEST_MINIMAX_GROUP_ID")),
	)

	//chat
	req := &textv1.ChatCompletionsRequest{
		Messages: []*textv1.Message{
			{
				SenderType: "USER",
				Text:       "hi~",
			},
		},
		Model:       "abab5-chat",
		Temperature: 0.7,
	}

	stream, _ := client.ChatCompletionStream(ctx, req)
	defer stream.CloseSend()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf(response.Choices[0].Delta + "\n") //嗨！有什么我可以帮助您的吗？
	}
}


```
</details>


<details>
<summary>MiniMax history stream completion</summary>

```go
package main

import (
	"context"
	"errors"
	"fmt"
	textv1 "github.com/ConnectAI-E/go-minimax/gen/go/minimax/text/v1"
	"github.com/ConnectAI-E/go-minimax/minimax"
	"io"
	"os"
)



func main() {
	ctx := context.Background()
	
	//init client
	client, _ := minimax.New(
		minimax.WithApiToken(os.Getenv("TEST_MINIMAX_API_TOKEN")),
		minimax.WithGroupId(os.Getenv("TEST_MINIMAX_GROUP_ID")),
	)
	
	//chat
	req := &textv1.ChatCompletionsRequest{
		Messages: []*textv1.Message{
			{
			    "sender_type": "USER",
			    "text": "路卡，今天在干什么呢？"
			},
			{
			    "sender_type": "BOT",
			    "text": "我今天在家里复习功课，准备期末考试呢！"
			},
			{
			    "sender_type": "USER",
			    "text": "期末考试怎么样，有把握吗？"
			}
		},
		Model:       "abab5-chat",
		Temperature: 0.7,
	}

	stream, _ := client.ChatCompletionStream(ctx, req)
	defer stream.CloseSend()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf(response.Choices[0].Delta + "\n") //放轻松，一切尽在掌握中
	}
}

```
</details>
