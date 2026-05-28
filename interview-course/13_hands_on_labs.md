# Module 13: Hands-On Labs

## Lab 1: Vulnerable vs Hardened Image
Use project:
- `sample-projects/project-1-cve-remediation-lab`

Tasks:
1. Build vulnerable image.
2. Scan with Trivy.
3. Build hardened image.
4. Rescan and compare findings.
5. Record before/after metrics.

Expected output:
- Documented CVE reduction and image size change.

## Lab 2: Minimal Runtime Strategy
Use project:
- `sample-projects/project-2-minimal-runtime-go`

Tasks:
1. Build and scan distroless runtime image.
2. Compare with a non-minimal variant (optional).
3. Explain tradeoffs and debugging plan.

Expected output:
- Clear justification for minimal runtime design.

## Lab 3: Security Pipeline Dry Run
Use project:
- `sample-projects/project-3-secure-image-pipeline`

Tasks:
1. Review workflow stages.
2. Explain where and why gate fails on CRITICAL.
3. Explain SBOM artifact and signing integration.

Expected output:
- End-to-end pipeline explanation in under 5 minutes.

## Lab 4: CVE Triage Simulation
- Pick 5 CVEs from scan output.
- Classify by severity, exploitability, and exposure.
- Decide patch, mitigate, or exception for each.

Expected output:
- Triage sheet with rationale and owner.

## Lab 5: Interview Story Builder
Write 5 STAR stories:
- Critical CVE remediation.
- Production-safe rollout.
- False-positive resolution.
- Security automation improvement.
- Cross-team collaboration outcome.
