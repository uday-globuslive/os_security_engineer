# Module 04: Containers from Zero

## Learning Objectives
- Understand what containers are and what they are not.
- Understand image layers, runtime, and isolation primitives.

## 1) Containers vs Virtual Machines
- VM: full guest OS per workload.
- Container: isolated process sharing host kernel.

Security implication:
- Kernel-level issues can affect all containers on a host.

## 2) Container Image Anatomy
- Image layers are immutable and stacked.
- Each Dockerfile instruction can create a layer.

Security implication:
- Vulnerable files in lower layers can still be detected.

## 3) Runtime Isolation
- Namespaces isolate process, network, mount, etc.
- Cgroups control resources.
- Capabilities split root privileges.

## 4) Registry Basics
- Images are stored and pulled from registries.
- Tags are mutable; digests are immutable.

Best practice:
- Prefer digest pinning for reproducibility.

## 5) Quick Command Walkthrough
```bash
docker images
docker history myimage:tag
docker run --rm -it myimage:tag sh
docker inspect myimage:tag
```

## 6) Common Misconceptions
- "Containers are secure by default" is false.
- "No CVEs means secure" is false.
- "Alpine always solves security" is false.

## Interview Questions
- Why do image layers matter in CVE remediation?
- What is the difference between image tag and digest?
- How can a container breakout increase risk?
