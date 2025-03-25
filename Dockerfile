# Generated by https://smithery.ai. See: https://smithery.ai/docs/config#dockerfile
FROM python:3.11-alpine

# Install build dependencies
RUN apk add --no-cache gcc musl-dev python3-dev

# Set work directory
WORKDIR /app

# Copy local code to the container image
COPY . /app

# Upgrade pip and install dependencies defined in pyproject.toml
RUN pip install --upgrade pip \
    && pip install .

# Expose the port if needed (optional)
# EXPOSE 8000

# Run the MCP server
CMD [ "python", "main.py" ]
