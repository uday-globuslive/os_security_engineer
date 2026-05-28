# Module 21: Incident Response for Container Security

## Learning Objectives
- Handle vulnerability incidents and active container threats systematically.
- Communicate clearly during high-pressure production events.

## 1) Incident Types You Should Prepare For
- Critical CVE with active exploit intelligence.
- Malicious image pushed to registry.
- Suspicious runtime behavior in production pod.
- Credential or secret leakage from containerized app.

## 2) IR Lifecycle for Container Environments
1. Detect: alert from scanner, runtime detector, or monitoring.
2. Triage: confirm scope and business impact.
3. Contain: isolate workloads and block spread.
4. Eradicate: patch, replace artifact, rotate compromised secrets.
5. Recover: restore stable service and validate integrity.
6. Lessons learned: document and improve controls.

## 3) Fast Triage Checklist
- Which images, tags, and digests are impacted?
- Which namespaces/services are exposed?
- Is there known exploit activity in the wild?
- What immediate mitigation reduces risk fastest?

## 4) Containment Patterns
- Block deployment of affected images.
- Restrict network paths for impacted services.
- Scale down exposed workloads when needed.
- Revoke or rotate credentials tied to affected workloads.

## 5) Communication Plan
During incident updates include:
- What happened.
- Current impact.
- Immediate actions taken.
- Next actions and ETA.
- Risk status.

## 6) Post-Incident Improvement
- Add missing detection rules.
- Improve pipeline policy gates.
- Update runbooks and training.
- Add regression tests for incident class.

## 7) Interview Questions
1. How would you respond to a critical CVE with a 24-hour remediation SLA?
2. How do you balance containment with service availability?
3. What artifacts prove incident closure?
