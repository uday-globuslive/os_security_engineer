# Module 05: Dockerfile Mastery for Security

## Learning Objectives
- Build efficient and secure Dockerfiles.
- Avoid anti-patterns that increase CVE risk and image size.

## 1) Dockerfile Fundamentals
Key instructions:
- `FROM`, `RUN`, `COPY`, `WORKDIR`, `USER`, `CMD`, `ENTRYPOINT`.

## 2) Security-First Dockerfile Principles
- Use minimal trusted base images.
- Use fixed versions or digests.
- Use multi-stage builds.
- Remove build-only dependencies.
- Run as non-root user.

## 3) Vulnerable Pattern Example
```dockerfile
FROM ubuntu:20.04
RUN apt-get update && apt-get install -y python3 python3-pip curl vim
COPY . /app
CMD ["python3", "/app/app.py"]
```

Issues:
- Large base, unnecessary tools.
- Likely stale packages if not maintained.
- Runs as root by default.

## 4) Hardened Pattern Example
```dockerfile
FROM python:3.12-slim
ENV PYTHONDONTWRITEBYTECODE=1 PYTHONUNBUFFERED=1
RUN apt-get update && apt-get dist-upgrade -y && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY app.py .
RUN addgroup --system app && adduser --system --ingroup app app
USER app
CMD ["python3", "app.py"]
```

## 5) Size and Security Optimization
- Combine apt commands and clean cache.
- Use `.dockerignore` to avoid copying unnecessary files.
- Keep runtime stage minimal.

## 6) Interview Questions
- Why multi-stage builds improve security?
- What is the purpose of `USER` in Dockerfile?
- Why clear `/var/lib/apt/lists`?
