# Module 23: Troubleshooting Playbook and Common Failures

## Learning Objectives
- Diagnose frequent failures in remediation workflows quickly.
- Explain root cause and corrective action in interviews.

## 1) Problem: HIGH CVEs remain after patching
Likely causes:
- Base image digest not updated.
- Package still present in final runtime image.
- Scanner database lag.
- No upstream fix yet.

Fix workflow:
1. Verify package version inside final image.
2. Rebuild without stale cache when needed.
3. Update scanner DB and rescan.
4. Check related artifacts: sidecars and node images.

## 2) Problem: Hardened image breaks at runtime
Likely causes:
- Missing runtime dependency.
- Non-root user cannot access needed path.
- Read-only filesystem conflicts with write path.

Fix workflow:
1. Identify failing operation in logs.
2. Add minimum required permission or path.
3. Keep hardening controls and adjust safely.

## 3) Problem: Distroless image hard to debug
Likely causes:
- No shell/tools in runtime image by design.

Fix workflow:
1. Use temporary debug sidecar/ephemeral container.
2. Rely on logs/metrics/traces.
3. Avoid permanently bloating runtime image.

## 4) Problem: CI gate blocks too many builds
Likely causes:
- Unrealistic policy thresholds.
- Missing exception workflow.
- Poor remediation guidance.

Fix workflow:
1. Keep CRITICAL strict.
2. Add controlled exceptions for HIGH with expiry.
3. Improve developer feedback and templates.

## 5) Problem: Signed image verification fails
Likely causes:
- Signature not generated for pushed digest.
- Trust policy misconfiguration.
- Wrong registry reference.

Fix workflow:
1. Verify signed digest exactly matches deployed artifact.
2. Validate identity and trust policy bindings.
3. Retry signing and verification with deterministic references.

## 6) Interview Questions
1. Describe a hardening change that caused failure and how you fixed it.
2. How do you debug scanner disagreement under time pressure?
3. How do you recover developer trust after too many false gate failures?
