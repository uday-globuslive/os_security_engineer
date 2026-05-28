# Module 12: Kubernetes Runtime Security Basics

## Learning Objectives
- Understand key Kubernetes controls relevant to container security.
- Connect image hardening decisions to runtime enforcement.

## 1) Pod Security Basics
Critical settings:
- `runAsNonRoot`
- `allowPrivilegeEscalation: false`
- `readOnlyRootFilesystem: true`
- Drop Linux capabilities

## 2) Network Policies
- Restrict pod-to-pod and pod-to-external communication.
- Reduce lateral movement risk after compromise.

## 3) Secrets Management
- Avoid baking secrets in images.
- Use secret stores and least-privilege access.

## 4) Admission Controls
- Enforce signed image policy.
- Enforce no-root policy.
- Enforce approved registries only.

## 5) Observability and Response
- Centralized logs.
- Alerting on suspicious behavior.
- Runtime policy violations as incidents.

## 6) Interview Questions
- How do image and runtime security complement each other?
- What runtime controls reduce exploit impact?
- How do you enforce secure deployment standards across teams?
