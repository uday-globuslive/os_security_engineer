# Interview Questions and Strong Answer Guide

## A) Core Technical Questions

### 1) Tell me about a critical CVE you remediated.
Strong answer structure:
- State workload and risk.
- Describe exact vulnerable component and verification method.
- Explain Dockerfile/package changes.
- Show validation and measurable result.

Sample answer:
"In an AKS-hosted Java service, Trivy flagged a critical OpenSSL CVE in our Ubuntu base image. I confirmed the vulnerable package version inside the container and mapped it to advisory details. I rebased to a patched image digest, applied package upgrades during build, removed unused utilities, and switched runtime to a non-root user. After rebuild, the critical finding was removed and high findings dropped significantly. We rolled out through canary, monitored error rates, and had no regression incidents."

### 2) How do you reduce image attack surface?
Sample answer points:
- Minimal or distroless base when compatible.
- Multi-stage builds to drop compilers/tools from runtime.
- Remove unnecessary packages and shells.
- Non-root user, least capabilities, read-only FS where possible.
- Expose only required ports.

### 3) Why do HIGH CVEs remain after patching?
Sample answer points:
- Vulnerable package still present in final image.
- Base image digest still old.
- Scanner DB lag or advisory mismatch.
- No upstream fix available yet.
- Other images (sidecars, node image) still vulnerable.

## B) Advanced Scenario Questions

### 4) How would you design a CVE remediation pipeline at scale?
Expected points:
- Automatic rebuild triggers from base image updates.
- Severity-based policy gates.
- Exception workflow with expiry.
- SBOM generation and artifact retention.
- Runtime validation to avoid breaking releases.

### 5) How do you handle no-fix-yet critical CVEs?
Expected points:
- Compensating controls (network restrictions, capability drops, WAF, feature flags).
- Temporary exception with review cadence.
- Track vendor patch availability and auto-retest.

### 6) How do you compare Trivy and Grype findings?
Expected points:
- Different vulnerability data sources and matching logic.
- Normalize by package and CVE.
- Investigate disagreements manually for high risk findings.

## C) Linux and Container Fundamentals

### 7) Difference between apt, yum/dnf, and apk?
- Ecosystem and distribution differences.
- Package metadata and update behavior.
- Alpine apk often smaller footprint; possible compatibility tradeoffs.

### 8) Explain container layers and why they matter for security.
- Vulnerable files can persist in lower layers.
- Layer cache can hide stale packages unless rebuild strategy is correct.

### 9) What is the security risk of running as root in a container?
- Higher impact if container escapes or mounts sensitive paths.
- Violates least privilege.

## D) Supply Chain and Governance

### 10) What is an SBOM and why is it valuable?
- Component inventory for vulnerability and compliance visibility.
- Faster blast-radius analysis.

### 11) Why sign images?
- Ensure authenticity and integrity.
- Enforce deployment of trusted artifacts only.

### 12) What is SLSA in practical terms?
- Better build provenance and tamper resistance.
- Structured maturity model for software supply chain controls.

## E) Behavioral Questions

### 13) Tell me about a time patching caused a regression.
Use STAR:
- Situation: patch urgency.
- Task: preserve uptime and security.
- Action: staged rollout, feature flag, rollback strategy.
- Result: restored stability while maintaining remediation plan.

### 14) How do you balance speed vs stability?
Expected points:
- Risk-based prioritization.
- Canary and phased rollout.
- Automated tests and observability gates.

### 15) How do you influence developers to adopt secure image patterns?
Expected points:
- Golden base images.
- Reusable pipeline templates.
- Developer-friendly policy feedback.

## F) Rapid-Fire Practice Questions
- Distroless vs slim: when would you choose each?
- How do you prove a CVE is not exploitable in your context?
- What would you log for audit evidence in remediation workflows?
- How do you detect stale base images across many services?
- What metrics would you show leadership monthly?

## G) Interview-Day Answer Quality Checklist
- Use specific tools, commands, and decisions.
- Include tradeoffs, not just one-sided claims.
- Quantify outcomes.
- Show ownership and collaboration.
- End with lessons learned and prevention improvements.
