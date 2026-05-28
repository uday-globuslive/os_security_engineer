# CI/CD Automation Playbook for CVE Remediation

## 1) Reference Pipeline Stages
1. Source checkout.
2. Build image.
3. Generate SBOM.
4. Scan with Trivy/Grype.
5. Enforce policy gate.
6. Sign image.
7. Push artifact.
8. Deploy to environment.
9. Post-deploy verification.

## 2) Policy Gate Example
- Fail build on any CRITICAL CVE.
- Allow HIGH CVEs only with approved exception and expiration date.
- Always log vulnerabilities and decisions.

## 3) Recommended Exception Model
For unavoidable CVEs:
- Document CVE ID, affected package, reason, mitigation.
- Define temporary expiry date.
- Require security approval.
- Auto-reopen if expiry is reached.

## 4) Auto-Rebuild Strategy
Trigger image rebuild when:
- Base image digest changes.
- New CVE affects current image contents.
- Scheduled weekly/monthly patch windows.

## 5) Quality and Reliability Guards
Before promotion:
- Rescan after patching.
- Run smoke/integration tests.
- Compare runtime behavior to previous release.

## 6) Operational Metrics
Track:
- Vulnerability density per image.
- MTTR by severity.
- Build failure rates due to policy gates.
- Exception count and aging.

## 7) Common Pipeline Pitfalls
- Scanning only application dependencies and not OS packages.
- Treating scan output as static truth without context.
- Blocking all builds without exception workflow.
- No alerting for stale exceptions.
