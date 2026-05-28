# Module 22: System Design Interview Scenarios for OS Security Engineers

## Learning Objectives
- Answer architecture and scale questions with security depth.
- Defend tradeoffs clearly under interview pressure.

## Scenario 1: Secure Image Factory for 300 Services
Prompt:
Design a centralized platform that builds, scans, signs, and publishes secure images for 300 microservices.

Expected components:
- Source trigger and build workers.
- Golden base image pipeline.
- Scanner stage with policy gate.
- SBOM generation and artifact store.
- Signature and verification policy.
- Dashboard for risk and SLA tracking.

Tradeoffs to explain:
- Strict blocking vs delivery speed.
- Central policy vs team autonomy.

## Scenario 2: Emergency CVE Response Platform
Prompt:
A critical CVE affects libc used across most images. Design response flow.

Expected design:
- Fleet inventory from SBOM.
- Prioritized remediation by exposure.
- Automated rebuild and canary rollout.
- Exception process for no-fix workloads.
- Executive and engineering reporting stream.

## Scenario 3: Multi-Cluster Kubernetes Security Baseline
Prompt:
Design security guardrails for dev, staging, and prod clusters.

Expected controls:
- Namespace and RBAC model.
- Admission policies and signed-image enforcement.
- Network segmentation and egress controls.
- Centralized audit and alerting.

## Scenario 4: Balancing Security and Developer Experience
Prompt:
Developers complain that security gates slow delivery.

Expected answer:
- Fast feedback in pull requests.
- Actionable remediation guidance.
- Time-bound exceptions with approval.
- Golden templates and platform support.

## Evaluation Rubric for Design Answers
1. Architecture completeness.
2. Risk-based prioritization.
3. Automation strategy.
4. Operational realism.
5. Clear tradeoff discussion.
