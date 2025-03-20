# mcp-server

学习写 mcp server

## 安装 (cline)

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

