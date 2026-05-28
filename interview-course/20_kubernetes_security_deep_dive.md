# Module 20: Kubernetes Security Deep Dive

## Learning Objectives
- Move from basic pod security to platform-level controls.
- Explain Kubernetes security architecture in interview design rounds.

## 1) Cluster Trust Boundaries
Security layers:
- Control plane
- Worker nodes
- Namespaces
- Pods and containers
- Service-to-service communication

## 2) Identity and Access Management
- Use RBAC with least privilege.
- Separate roles for developers, operators, and security.
- Avoid broad cluster-admin usage.

## 3) Network Security at Scale
- Default deny network policies.
- Limit egress by namespace and workload type.
- Segment sensitive services and data planes.

## 4) Admission and Policy Controls
- Enforce runAsNonRoot, no privileged containers, approved registries.
- Enforce image signature verification.
- Block risky capabilities and hostPath mounts.

## 5) Secret and Key Management
- Use external secret managers where possible.
- Rotate secrets and short-lived credentials.
- Avoid static long-lived tokens in CI.

## 6) Runtime Detection and Response
- Collect audit logs, container runtime events, and network anomalies.
- Alert on policy violations and suspicious process behavior.
- Integrate with incident response workflow.

## 7) Interview Questions
1. How do you design secure multi-tenant Kubernetes clusters?
2. What controls stop untrusted images from running?
3. How do you reduce lateral movement in Kubernetes?

## 8) Design Exercise
Design a secure platform for 50 microservices with:
- Namespace isolation
- RBAC model
- Network policy model
- Admission policies
- Audit and alerting strategy
