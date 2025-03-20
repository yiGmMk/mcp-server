# server.py
import os
from mcp.server.fastmcp import FastMCP
from pydantic import Field

# Create an MCP server
mcp = FastMCP("fetch-tools")

# Example: Accessing a secret from the client configuration
JINA_API_KEY = os.environ.get("JINA_API_KEY")
if not JINA_API_KEY:
    print("JINA_API_KEY not found in environment variables!")
else:
    print("JINA_API_KEY:", JINA_API_KEY)


# Add an addition tool
import requests


@mcp.tool(name="fetch", description="使用 r.jina.ai 读取 URL 并获取其内容")
def fetch(url: str = Field(description="需要抓取的网页url")) -> str:
    headers = {}
    if JINA_API_KEY:
        headers["Authorization"] = f"Bearer {JINA_API_KEY}"

    try:
        response = requests.get(
            f"https://r.jina.ai/{url}",
            headers=headers,
        )
        response.raise_for_status()  # Raise HTTPError for bad responses (4xx or 5xx)
        return response.text
    except requests.exceptions.RequestException as e:
        return f"Error fetching HTML: {e}"


@mcp.tool(
    name="search",
    description="使用 s.jina.ai 搜索网络并获取 SERP,Reader 就会搜索网络并返回前五个结果及其 URL 和内容，每个结果都以干净、LLM 友好的文本显示。这样，您就可以始终让您的 LLM 保持最新状态，提高其真实性，并减少幻觉。",
)
def search(q: str = Field(description="搜索关键词")) -> str:
    if not JINA_API_KEY:
        return "API_KEY is not configured."

    try:
        response = requests.get(
            f"https://s.jina.ai/?q={q}",
            headers={
                "Authorization": f"Bearer {JINA_API_KEY}",
                "X-Respond-With": "no-content",
            },
        )
        response.raise_for_status()  # Raise HTTPError for bad responses (4xx or 5xx)
        return response.text
    except requests.exceptions.RequestException as e:
        return f"Error searching: {e}"


def main():
    print("Hello from mcp-server!")
    mcp.run()


if __name__ == "__main__":
    main()
