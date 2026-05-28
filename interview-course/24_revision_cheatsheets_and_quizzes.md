# Module 24: Revision Cheat Sheets and Quizzes

## 1) One-Page Rapid Revision Sheet
Memorize these anchors:
- CVE lifecycle: discover, validate, prioritize, remediate, verify, prevent.
- Hardening anchors: minimal base, non-root, multi-stage, capabilities drop, read-only FS.
- Supply chain anchors: SBOM, signing, provenance, trust policy.
- CI/CD anchors: scan gate, exception with expiry, canary rollout, rollback.

## 2) 20 Quick Quiz Questions
1. Why is digest pinning safer than tag pinning?
2. Name two reasons HIGH CVEs stay after patching.
3. What does runAsNonRoot protect against?
4. Why generate SBOM in pipeline?
5. What is the purpose of cosign?
6. Difference between update and upgrade in apt?
7. Why remove build tools from runtime stage?
8. What is a compensating control?
9. Why are network policies important for CVE risk?
10. What does canary rollout reduce?
11. What makes a good exception record?
12. Why can two scanners disagree?
13. What does readOnlyRootFilesystem prevent?
14. Why not bake secrets into images?
15. What is MTTR in remediation?
16. Why avoid privileged containers?
17. What should a post-incident review include?
18. Why does kernel security matter for containers?
19. What is a golden base image program?
20. How do you prove remediation success?

## 3) Answer Key (Short Form)
1. Tags are mutable, digests are immutable.
2. Stale base digest, package still present, DB lag, no upstream fix.
3. Reduces privileged runtime risk.
4. Component transparency and faster impact analysis.
5. Artifact signing and trust verification.
6. Update metadata vs upgrade packages.
7. Reduce attack surface and CVEs.
8. Temporary risk reduction when full fix unavailable.
9. Reduce exploit path and lateral movement.
10. Production risk by limited blast radius.
11. CVE, owner, reason, mitigation, expiry.
12. Different databases and matching logic.
13. Limits write-based persistence/tampering.
14. Prevent secret leakage in layers.
15. Mean time to remediate.
16. Avoid broad host-level power.
17. Root cause, impact, fixes, preventive actions.
18. Shared kernel means host vulnerability affects containers.
19. Centralized secure baseline for many services.
20. Rescan, verify package version, validate runtime tests.

## 4) 7-Day Final Revision Sprint
Day 1: Linux + Containers
Day 2: Dockerfile + Hardening
Day 3: CVE Remediation + Scanners
Day 4: SBOM + Signing + SLSA
Day 5: CI/CD + Kubernetes Controls
Day 6: Mock Interview Full Round
Day 7: Weak Area Repair + Final Drill
