# Module 08: Image Hardening Playbook

## Learning Objectives
- Apply practical hardening controls and explain tradeoffs.
- Build hardened images that remain operable.

## 1) Hardening Goals
- Reduce attack surface.
- Enforce least privilege.
- Minimize vulnerable components.
- Improve deterministic builds.

## 2) Core Controls
- Minimal base image.
- Multi-stage builds.
- Non-root user.
- Remove package manager and shell where possible in runtime.
- Read-only root filesystem (at runtime policy level).
- Drop Linux capabilities.

## 3) Runtime Security Settings (Kubernetes aligned)
- `runAsNonRoot: true`
- `allowPrivilegeEscalation: false`
- `readOnlyRootFilesystem: true`
- Drop all capabilities and add only required ones.

## 4) Operational Tradeoffs
- Distroless improves security posture but complicates debugging.
- Very strict policies can break legacy apps.
- Hardening must be validated with real runtime tests.

## 5) Hardened Image Review Checklist
- Is base image pinned?
- Any unnecessary binaries left?
- Does process run as non-root?
- Are build tools absent in runtime?
- Are ports and permissions minimal?

## 6) Interview Questions
- Explain your image hardening process step-by-step.
- What are tradeoffs of distroless images?
- How do you prove hardening did not break production behavior?
