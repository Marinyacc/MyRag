package main

import (
	"MyRag/internal"
	"MyRag/tool"
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

//Rag基本流程:
//1.初始化ChatModel,EmbeddingModel,向量数据库Milvus连接,向量数据表建立
//2.本地文档切片（Transform）,存储到向量数据库（Embedding）
//3.用户输入查询字符串，retriver检索TopK个相关性最高的文档切片
//4.将 检索出的文档切片 和 用户输入 通过ChatTemplate合并 发送给ChatModel
//5.为用户返回ChatModel输出

func main() {

	ctx := context.Background()

	internal.Init(ctx)

	fmt.Println("启动成功!")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		var begin time.Time
		if scanner.Scan() {
			begin = time.Now()

			//从控制台读取输入
			query := scanner.Text()
			if query == "quit" {
				break
			}

			resp := internal.Process(ctx, query)

			fmt.Println(resp.Content)
			//记录对话处理时间
			fmt.Println("====================================")
			fmt.Println("本次对话耗时:", tool.CountTime(begin))
			fmt.Println("消耗:", resp.ResponseMeta.Usage.TotalTokens, "tokens")
			fmt.Println("====================================")
		}

	}

	fmt.Println("成功退出!")
}
