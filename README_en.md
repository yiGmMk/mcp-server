# mcp-server

Learning to write mcp server

## Installation (cline)

### Prerequisites

• Python 3.12 or newer
• uv package manager

If uv is not installed, please refer to the documentation [uv Installation Documentation](https://docs.astral.sh/uv/getting-started/installation/)

For macOS or Linux, use the following command to install:

```bash
curl -LsSf https://astral.sh/uv/install.sh | sh
```

For Windows, install in PowerShell:

```bash
powershell -ExecutionPolicy ByPass -c "irm https://astral.sh/uv/install.ps1 | iex"
```

### Configuration File

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
                "JINA_API_KEY": "jina_api_key, please obtain from https://jina.ai/reader",
                "PYTHONIOENCODING": "utf-8"
            },
            "disabled": false,
            "autoApprove": []
        }
    }
}
```

Or use venv (add environment variable configuration):

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
                "JINA_API_KEY": "jina_api_key, please obtain from https://jina.ai/reader",
                "PYTHONIOENCODING": "utf-8"
            },
            "disabled": false,
            "autoApprove": []
        }
    }
}
```

## jina.ai

Various tools provided by [jina.ai](https://jina.ai/), such as search, webpage reading, etc.
