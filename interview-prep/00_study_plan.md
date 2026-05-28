# Senior OS Security Engineer Interview Study Plan

## Goal
Become interview-ready for a role focused on CVE remediation, Linux container hardening, and secure image supply chain practices.

## 4-Week Plan (can be compressed to 2 weeks)

### Week 1: Core Foundations
- Linux security basics: users, permissions, processes, networking, package management.
- Container internals: layers, registries, runtime, namespaces, cgroups.
- CVE lifecycle basics: discovery, scoring, patching, verification.

Deliverables:
- Explain why the same image can show different scan results at different times.
- Build one container image and inspect layers.

### Week 2: Practical Remediation and Hardening
- Patch vulnerable packages using apt/yum/apk.
- Multi-stage builds and runtime minimization.
- Non-root containers, read-only filesystem, dropped capabilities.

Deliverables:
- Harden one vulnerable Dockerfile.
- Show before/after CVE count and image size.

### Week 3: Automation and Supply Chain
- CI/CD integration for Trivy/Grype scan gates.
- SBOM generation (SPDX/CycloneDX).
- Signing and provenance with cosign and sigstore.

Deliverables:
- Pipeline that fails on critical CVEs.
- Signed image and generated SBOM attached to build artifacts.

### Week 4: Interview Simulation and Deep Dives
- Practice behavioral + technical stories using STAR format.
- Debug tricky scanner findings and false positives.
- Tradeoffs: patch speed vs stability, slim vs distroless, blocking policy thresholds.

Deliverables:
- 10 strong story-based answers.
- 2 whiteboard-level architecture explanations.

## Daily 60-Minute Routine
1. 15 min: read one focused security concept.
2. 20 min: hands-on command practice.
3. 15 min: write one interview answer in STAR format.
4. 10 min: recap and document lessons.

## Interview Readiness Checklist
- I can explain CVE remediation end-to-end from triage to production rollout.
- I can justify base image selection and package strategy.
- I can explain scanner behavior differences and data lag.
- I can design CI/CD policy gates with exceptions and audit trails.
- I can discuss SBOM, signing, and provenance clearly.
- I can quantify outcomes: CVE reduction, image size reduction, deployment impact.
