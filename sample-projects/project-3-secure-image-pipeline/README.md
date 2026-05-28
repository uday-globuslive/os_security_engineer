# Project 3: Secure Image Pipeline (GitHub Actions)

## Objective
Automate image build, vulnerability scanning, SBOM generation, and optional signing.

## What this workflow does
- Builds container image.
- Runs Trivy scan and fails on critical findings.
- Generates SBOM artifact.
- Optionally signs image with cosign if keyless/OIDC is configured.

## Setup Notes
- Ensure repository has Dockerfile.
- Configure registry credentials if pushing image.
- Enable GitHub OIDC for keyless signing if needed.

## Interview Talking Point
Describe how policy gates and exceptions are implemented to keep security strict but delivery practical.
