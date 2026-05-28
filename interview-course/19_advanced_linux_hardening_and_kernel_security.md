# Module 19: Advanced Linux Hardening and Kernel Security

## Learning Objectives
- Understand kernel-level risks that affect container security.
- Apply advanced Linux hardening controls used in production.

## 1) Kernel Risk in Containerized Platforms
Containers share the host kernel, so kernel vulnerabilities can impact multiple workloads.

Key implication:
- Image hardening is necessary but not sufficient.
- Host OS and kernel patching strategy is equally critical.

## 2) Advanced Host Hardening Controls
- Keep host OS minimal and patched.
- Disable unused services.
- Enforce secure boot and trusted package sources.
- Use strict SSH controls and MFA where possible.

## 3) Syscall and Kernel Surface Reduction
- Use seccomp profiles to limit allowed syscalls.
- Use AppArmor or SELinux for process confinement.
- Avoid privileged containers and host namespace sharing.

## 4) File Integrity and Audit Controls
- Enable audit logging for sensitive events.
- Monitor key paths for unauthorized changes.
- Use immutable log pipelines for tamper resistance.

## 5) Resource Abuse Protection
- Set CPU and memory limits to reduce denial-of-service risk.
- Enforce process limits where applicable.
- Use cgroup policies as part of platform baseline.

## 6) Interview Questions
1. Why does kernel security matter if images are hardened?
2. What is seccomp and where does it help?
3. How do AppArmor and SELinux improve runtime security?

## 7) Practice Task
Write a one-page host hardening baseline for:
- Kubernetes worker nodes
- Container build servers
- CI runners
