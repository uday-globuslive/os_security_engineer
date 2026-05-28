# Module 11: CI/CD Security Automation

## Learning Objectives
- Build automated remediation workflows.
- Implement practical policy gates and exception handling.

## 1) Pipeline Blueprint
1. Build image.
2. Generate SBOM.
3. Scan vulnerabilities.
4. Apply policy gate.
5. Sign image.
6. Publish artifact.
7. Deploy with runtime checks.

## 2) Policy Strategy
- Block all CRITICAL by default.
- Control HIGH with approved temporary exceptions.
- Keep complete audit logs for approvals and expiry.

## 3) Exception Governance
Each exception should include:
- CVE ID and package.
- Business justification.
- Mitigation controls.
- Expiry date.
- Owner.

## 4) Automation Triggers
- Base image update event.
- New CVE alert for included package.
- Scheduled patch rebuild windows.

## 5) Stability Controls
- Canary deployments.
- Smoke and integration tests.
- Fast rollback path.

## 6) Interview Questions
- How do you avoid blocking developers while maintaining security standards?
- What does a good CVE exception process look like?
- How do you monitor pipeline security KPIs?
