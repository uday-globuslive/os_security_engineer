# Module 09: Trivy and Grype (Scanning and Interpretation)

## Learning Objectives
- Run and interpret scanner results effectively.
- Understand why tool outputs can differ.

## 1) What Scanners Do
Scanners map known vulnerabilities to components in your image:
- OS packages.
- Language dependencies.
- Sometimes configuration issues.

## 2) Trivy Basics
```bash
trivy image myimage:tag
trivy image --severity CRITICAL,HIGH myimage:tag
```

## 3) Grype Basics
```bash
grype myimage:tag
grype myimage:tag -o table
```

## 4) Why Results Differ
- Different vulnerability databases.
- Different package matching logic.
- Different advisory timing and vendor status.

## 5) High CVEs After Patching: Root Causes
- Base image still vulnerable.
- Package still present in final layer.
- No upstream patch yet.
- Stale scanner database.
- Related images (sidecars/nodes) still vulnerable.

## 6) Scanning Best Practices
- Rescan with updated DB after each fix.
- Keep a baseline trend, not one-off snapshots.
- Correlate findings with actual exploitability context.
- Feed results into ticketing and exception workflow.

## 7) Interview Questions
- Why do scanners disagree sometimes?
- How do you separate false positives from real risk?
- How do you enforce scan policies without stopping all delivery?
