# Module 07: CVE Management End-to-End

## Learning Objectives
- Master real-world CVE triage, remediation, and validation.
- Build senior-level decision-making for vulnerability response.

## 1) What is a CVE?
- CVE is an identifier for a publicly disclosed vulnerability.
- Severity often includes CVSS score, but context determines practical risk.

## 2) CVE Lifecycle in Production
1. Discovery by scanner/advisory.
2. Validation and impact assessment.
3. Prioritization.
4. Remediation.
5. Verification.
6. Documentation and prevention.

## 3) Prioritization Model
Consider:
- Severity and exploitability.
- Internet exposure.
- Privilege context.
- Data sensitivity.
- Compensating controls.

## 4) Remediation Decision Tree
- Fix available now: patch immediately.
- No fix available: mitigate exposure and track exception.
- Not used package but present: remove from image.
- False positive suspected: verify with alternate sources and document.

## 5) Verification Checklist
- Rescan image with updated DB.
- Validate package versions in container.
- Run smoke/integration tests.
- Roll out with canary strategy.

## 6) Metrics
- MTTR by severity.
- Vulnerability backlog trend.
- Exception count and expiry health.
- Reopen rate for remediated CVEs.

## 7) Interview Questions
- Walk through a critical CVE remediation you owned.
- How do you justify exceptions for no-fix CVEs?
- How do you balance urgent patching vs stability?
