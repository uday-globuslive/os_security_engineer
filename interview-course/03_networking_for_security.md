# Module 03: Networking for Security Engineers

## Learning Objectives
- Understand networking concepts required for container and Kubernetes security.
- Explain common exposure and lateral movement risks.

## 1) Networking Basics
- IP address: host identifier.
- Port: process communication endpoint.
- Protocol: communication rules (TCP/UDP).

Command examples:
```bash
ip a
ip route
ss -tulnp
curl -v http://localhost:8080/health
```

## 2) Security-Relevant Networking Concepts
- Ingress traffic: external to workload.
- Egress traffic: workload to external.
- East-west traffic: service-to-service internal traffic.

Why this matters:
- CVE exploitability depends on reachable attack path.

## 3) DNS and Service Discovery
In container platforms, services discover each other by DNS names.

Risk:
- Misconfigured DNS or open egress can leak data or enable command-and-control traffic.

## 4) Network Hardening Basics
- Expose only required ports.
- Restrict egress where possible.
- Use network policies in Kubernetes.
- Separate environments and trust zones.

## 5) Interview Questions
- How does network exposure influence vulnerability priority?
- Why is egress control important for compromised containers?
- What is the difference between host networking and bridged networking?

## Module Exercise
- Map one app architecture and identify all ingress/egress paths.
- Mark where a high-risk CVE would be most exploitable.
