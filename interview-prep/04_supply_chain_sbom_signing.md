# Supply Chain Security, SBOM, and Signing

## 1) Why Supply Chain Security Matters
Even if your app code is clean, compromised base images or dependencies can introduce risk.
Supply chain controls improve:
- Transparency (what is inside the image).
- Integrity (was the image tampered with).
- Traceability (who built it, from what source).

## 2) SBOM Basics

### What Is an SBOM?
A Software Bill of Materials is an inventory of components in your artifact.

Common formats:
- SPDX
- CycloneDX

Use cases:
- Vulnerability impact analysis.
- License compliance.
- Faster incident response.

### Practical Workflow
1. Build image.
2. Generate SBOM.
3. Store SBOM artifact with build metadata.
4. Use SBOM in ongoing vulnerability monitoring.

## 3) Signing and Provenance

### Image Signing
- Use cosign to sign container images.
- Store signatures in registry transparency logs where possible.

Why:
- Consumers can verify image authenticity.
- Helps prevent running untrusted images.

### Provenance
- Capture build details: source repo, commit SHA, build system identity, timestamps.
- Align with SLSA principles for stronger supply chain guarantees.

## 4) SLSA Simplified
SLSA is a framework for software supply chain security.
At a practical level, target:
- Controlled build process.
- Tamper-evident provenance.
- Policy checks before deployment.

## 5) Interview Talking Points
- How SBOM helps triage CVEs quickly.
- Why signed images reduce deployment risk.
- How provenance supports audit and incident response.
- Tradeoffs between strict policy enforcement and developer velocity.
