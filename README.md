# mcp-server

[![smithery badge](https://smithery.ai/badge/@yiGmMk/mcp-server)](https://smithery.ai/server/@yiGmMk/mcp-server)

学习写 mcp server

## 安装 (cline)

### Installing via Smithery

To install mcp-server for Claude Desktop automatically via [Smithery](https://smithery.ai/server/@yiGmMk/mcp-server):

```bash
npx -y @smithery/cli install @yiGmMk/mcp-server --client claude
```

### 前置依赖

• Python 3.12 or newer
• uv package manager

没有安装uv的,请参考文档[uv 安装文档](https://docs.astral.sh/uv/getting-started/installation/)

macos或linux可使用如下命令安装

```bash
curl -LsSf https://astral.sh/uv/install.sh | sh
```

windwos,在powershell下安装

```bash
powershell -ExecutionPolicy ByPass -c "irm https://astral.sh/uv/install.ps1 | iex"
```

### 配置文件

优先使用下面这种配置方式

```json
{
    "mcpServers": {
        "yiGmMk/mcp-server": {
            "command": "uv",
            "args": [
                "--directory",
                "/path/to/your/mcp-server",
                "run",
                "main.py"
            ],
            "env": {
                "JINA_API_KEY": "jina_api_key,请从https://jina.ai/reader获取",
                "GOOGLE_AI_STUDIO_KEY": "Google AI Studio api key,请从https://aistudio.google.com/apikey获取",
                "PYTHONIOENCODING": "utf-8"
            },
            "disabled": false,
            "autoApprove": []
        }
    }
}
```

或使用venv(增加环境变量配置)

```json
{
    "mcpServers": {
        "yiGmMk/mcp-server": {
            "command": "uv",
            "args": [
                "run",
                "/path/to/your/mcp-server/main.py"
            ],
            "env": {
                "VIRTUAL_ENV": "/path/to/your/mcp-server/.venv",
                "JINA_API_KEY": "jina_api_key,请从https://jina.ai/reader获取",
                "GOOGLE_AI_STUDIO_KEY": "Google AI Studio api key,请从https://aistudio.google.com/apikey获取",
                "PYTHONIOENCODING": "utf-8"
            },
            "disabled": false,
            "autoApprove": []
        }
    }
}
```

## jina.ai

[jina.ai](https://jina.ai/)提供的各种工具,如搜索,读取网页...

## TODO

- [ ] 发布到 https://smithery.ai/server/@yiGmMk/mcp-server/tools
  dockerfile构建成功,但是mcp服务在docker中运行后立即退出,导致无法使用,相关issue:
  1.[MCP within a Docker Container exits, regardless of lifespan](https://github.com/modelcontextprotocol/python-sdk/issues/228)
  2. [ifespan shuts as soon as it opens](https://github.com/modelcontextprotocol/python-sdk/issues/223)