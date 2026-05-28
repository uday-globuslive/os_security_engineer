# Module 10: Supply Chain Security, SBOM, Signing, and SLSA

## Learning Objectives
- Understand software supply chain risk in containerized systems.
- Implement SBOM and signing practices.

## 1) Supply Chain Risk Basics
Risk can enter from:
- Base images.
- Dependency registries.
- Build systems.
- Compromised credentials.

## 2) SBOM Fundamentals
SBOM is a manifest of components in your artifact.
Common formats:
- SPDX
- CycloneDX

Benefits:
- Faster vulnerability impact analysis.
- Better compliance and auditability.

## 3) Signing with Cosign
Goals:
- Verify image authenticity.
- Prevent untrusted image deployment.

Concept flow:
1. Build image.
2. Sign image.
3. Verify signature before deploy.

## 4) Provenance and SLSA
Provenance records:
- Source repo and commit.
- Builder identity.
- Build inputs and metadata.

SLSA offers maturity guidance for tamper-resistant builds.

## 5) Interview Questions
- Why is SBOM useful if you already scan images?
- What does image signing protect against?
- How would you enforce trust policies in CI/CD?
