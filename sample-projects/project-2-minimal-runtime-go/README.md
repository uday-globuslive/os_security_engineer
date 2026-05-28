# Project 2: Minimal Runtime Go Service

## Objective
Demonstrate attack-surface reduction using multi-stage builds and distroless runtime.

## Build
```bash
docker build -t go-minimal:latest .
```

## Scan
```bash
trivy image go-minimal:latest
```

## Why this project matters
- No package manager in runtime image.
- Very small footprint.
- Non-root execution in final stage.

## Interview Talking Point
Explain why minimal runtime images reduce CVE exposure and operational tradeoffs for debugging.
