# MyRag

基于 Go 和 Milvus 的 RAG（检索增强生成）学习搜索助手。

## 功能特性

- **文档向量化**: 自动读取 `document/` 目录下的 Markdown 文件并进行分片处理
- **向量检索**: 使用 Milvus 向量数据库存储和检索相似文档
- **智能问答**: 结合检索结果与 LLM 生成回答
- **交互式对话**: 支持在控制台输入问题进行实时问答

## 技术栈

- **Go 1.25+**
- **CloudWeGo Eino**: LLM 应用开发框架
- **Milvus**: 向量数据库
- **Ark API**: 大语言模型与文本嵌入

## 项目结构

```
MyRag/
├── cmd/
│   └── main.go           # 程序入口
├── document/             # 文档存放目录
│   ├── test.md
│   └── test2.md
├── internal/
│   ├── doc.go            # 文档加载与分片
│   ├── entity.go         # 核心结构体定义
│   ├── initilize.go      # 组件初始化
│   ├── milvus.go         # Milvus 客户端
│   └── process.go        # RAG 查询处理
├── tool/
│   ├── count_time.go     # 耗时统计
│   └── hash.go           # 哈希工具
├── docker-compose.yml    # Milvus 服务配置
└── go.mod
```

## 快速开始

### 1. 启动 Milvus 服务

```bash
docker-compose up -d
```

服务启动后：
- Milvus: `localhost:19530`
- Attu (可视化面板): `localhost:8000`

### 2. 配置环境变量

创建 `.env` 文件：

```env
API_KEY=your_api_key
CHAT_MODEL=your_chat_model
EMBEDDING_MODEL=your_embedding_model
```

### 3. 添加文档

将 Markdown 文件放入 `document/` 目录。

### 4. 运行程序

```bash
go run cmd/main.go
```

### 5. 开始问答

在控制台输入问题，按回车获取回答。输入 `quit` 退出。

## RAG 工作流程

1. **初始化**: 加载 ChatModel、EmbeddingModel，连接 Milvus
2. **文档处理**: 读取本地 Markdown 文档，按标题分片，向量化后存入 Milvus
3. **检索**: 用户输入查询，从向量数据库检索 Top-K 相关文档
4. **生成**: 将检索结果与用户问题合并，发送给 LLM
5. **输出**: 返回 LLM 回答及 token 消耗统计

## 配置说明

### Milvus 表结构

| 字段 | 类型 | 说明 |
|------|------|------|
| id | VarChar | 文档唯一标识 |
| vector | BinaryVector | 嵌入向量 (65536维) |
| content | VarChar | 文档原文 |
| metadata | JSON | 文档元数据 |

### 分片策略

使用 Markdown 标题层级分片：
- `#` → h1
- `##` → h2
- `###` → h3

### 注意

本项目采用火山引擎下doubao-seed-2.0-pro和doubao-embedder-vision
请根据自己的模型修改ChatModel和Embedder的配置
(ps:好像火山引擎现在只提供多模态向量化模型,俩文本向量化模型不给用了)
