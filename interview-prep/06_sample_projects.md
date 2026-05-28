# Sample Projects for Hands-On Interview Prep

## Project 1: CVE Remediation Lab (Python)
Path: sample-projects/project-1-cve-remediation-lab

Purpose:
- Practice scanning and remediating OS and Python package vulnerabilities.
- Compare vulnerable vs hardened Dockerfiles.

What you practice:
- `apt` patching.
- Non-root runtime.
- Multi-stage and cleanup patterns.

## Project 2: Minimal Runtime Build (Go)
Path: sample-projects/project-2-minimal-runtime-go

Purpose:
- Build a tiny runtime image with reduced attack surface.

What you practice:
- Multi-stage build.
- Distroless runtime.
- Read-only and non-root runtime settings in deployment.

## Project 3: Secure Image CI/CD Pipeline
Path: sample-projects/project-3-secure-image-pipeline

Purpose:
- Create automated build + scan + SBOM + signing workflow.

What you practice:
- Trivy policy gate.
- SBOM generation.
- cosign signing pattern.

## Suggested Execution Order
1. Complete Project 1 and collect before/after metrics.
2. Complete Project 2 and explain image design choices.
3. Complete Project 3 and explain governance model (exceptions, approvals, audit).
