FROM python:3.12-bookworm AS builder

RUN apt-get update && apt-get install -y --no-install-recommends python3-venv curl

ENV PATH="/root/.local/bin:$PATH"

RUN curl -sSL https://astral.sh/uv/install.sh | sh && . /root/.profile

# 设置工作目录
WORKDIR /app

# 复制项目依赖声明文件（uv 需要 pyproject.toml 或 requirements.txt）
COPY pyproject.toml uv.lock ./

# 安装项目依赖（使用 uv sync 或 uv pip）
RUN uv sync --no-install-project

# 生产阶段
FROM python:3.12-bookworm
COPY --from=builder /app/.venv /app/.venv
ENV PATH="/app/.venv/bin:$PATH"
COPY . .
EXPOSE 8000
CMD ["python", "main.py"]