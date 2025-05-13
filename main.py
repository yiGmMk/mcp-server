# server.py
import os
from mcp.server.fastmcp import FastMCP
from pydantic import Field
from gpt import call_openai
from prompt import translate as translatePrompt
import requests
import json

# Create an MCP server
mcp = FastMCP(
    name="ai-tools",
    description="AI 工具集合",
    version="0.1.0",
    host="0.0.0.0",  # 添加 host 参数
    port=8000,  # 添加 port 参数
)

JINA_API_KEY = os.environ.get("JINA_API_KEY")
if not JINA_API_KEY:
    print("JINA_API_KEY not found in environment variables!")
else:
    print("JINA_API_KEY:", JINA_API_KEY)

GOOGLE_AI_STUDIO_KEY = os.environ.get("GOOGLE_AI_STUDIO_KEY")

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
    description="使用 s.jina.ai 搜索网络并获取 SERP,Reader 就会搜索网络并返回前五个结果及其 URL 和内容，每个结果都以干净、LLM 友好的文本显示。这样，您就可以始终让您的 LLM 保持最新状态，提高其真实性，并减少幻觉。不支持用作翻译",
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


@mcp.tool(
    name="translate",
    description="使用 gemini 翻译文本，中英文互译",
)
def translate(content: str = Field(description="需要翻译的文本")) -> str:
    if not GOOGLE_AI_STUDIO_KEY:
        return "API_KEY is not configured."
    try:
        res = call_openai(translatePrompt, content)
        return res
    except Exception as e:
        return f"Error translating: {e}"


@mcp.tool(
    name="translate",
    description="使用 deepl 翻译,支持多种语言互译",
)
def translate_deepl(
    content: str = Field(description="需要翻译的文本"),
    source_lang: str = Field(
        description="当前语言,source language,支持: AR,BG,CS,DA,DE,EL,EN-GB,EN-US,ES,ET,FI,FR,HU,ID,IT,JA,KO,LT,LV,NB,NL,PL,PT-BR,PT-PT,RO,RU,SK,SL,SV,TR,UK,ZH,ZH-HANS,ZH-HANT"
    ),
    target_lang: str = Field(
        description="目标语言,target language,支持:AR,BG,CS,DA,DE,EL,EN-GB,EN-US,ES,ET,FI,FR,HU,ID,IT,JA,KO,LT,LV,NB,NL,PL,PT-BR,PT-PT,RO,RU,SK,SL,SV,TR,UK,ZH,ZH-HANS,ZH-HANT"
    ),
) -> str:
    url = "https://nav.programnotes.cn/translate"
    headers = {"Content-Type": "application/json"}
    data = {"text": content, "source_lang": source_lang, "target_lang": target_lang}
    response = requests.post(url, headers=headers, data=json.dumps(data))
    return response.text


from prompt import common, prompt1, prompt2, prompt3


@mcp.prompt(name="prompt optimization", description="优化提示词")
def optimize_prompt(content: str, usage: str = "通用") -> str:
    msg = ""
    if usage == "通用":
        msg = common
    elif usage == "带格式输出":
        msg = prompt1
    elif "带建议的优化提示词":
        msg = prompt2
    elif usage == "指令型提示词的优化，优化的同时遵循原指令":
        msg = prompt3
    else:
        msg = common
    return f"{msg}{content}"


def main():
    try:
        # 服务启动成功时，代码会阻塞在此处，不会执行后续 print
        mcp.run()
    except Exception as e:
        # 捕获服务启动异常
        print(f"Error starting server: {e}", flush=True)
        raise  # 抛出异常确保容器退出并显示错误


if __name__ == "__main__":
    print("Hello from mcp-server!", flush=True)
    # print(translate("hello world"))
    main()
