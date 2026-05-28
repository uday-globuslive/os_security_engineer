# Module 16: 100 Interview Questions with Model Answers

## How to use this module
- Read the question.
- Say the model answer out loud in your own words.
- Check the common mistake and avoid it.
- Keep each answer in 60 to 120 seconds unless asked for deep dive.

## Section A: Linux and OS Security Basics (1-20)

1. What is the Linux kernel and why does it matter for container security?
Model answer: The kernel controls process isolation, memory, filesystems, and networking. Containers share the host kernel, so kernel weaknesses can affect all containers on that node.
Common mistake: Saying containers have their own full OS kernel.

2. What does least privilege mean in practice?
Model answer: Give only required permissions to users, processes, and services. In containers, run as non-root, restrict capabilities, and allow only needed network paths.
Common mistake: Defining it theoretically without implementation details.

3. Why is running as root risky in a container?
Model answer: Root in container increases impact of compromise, especially with misconfigurations like writable mounts or privileged mode. Non-root reduces blast radius.
Common mistake: Claiming root in container is always safe because of isolation.

4. Explain chmod vs chown.
Model answer: chmod changes file permission bits; chown changes file owner and group. Security teams use both to enforce access boundaries.
Common mistake: Mixing ownership and permission concepts.

5. Difference between apt update and apt upgrade?
Model answer: update refreshes package metadata; upgrade installs newer package versions according to metadata.
Common mistake: Saying both do the same thing.

6. Why do security engineers care about /etc and /var/log?
Model answer: /etc holds config that affects security posture; /var/log holds evidence for monitoring and incident response.
Common mistake: Treating logs as only ops concern.

7. What is a package manager?
Model answer: Tool to install, update, and remove software packages with dependency resolution, such as apt, yum/dnf, and apk.
Common mistake: Ignoring dependency handling aspect.

8. Why patching alone is not enough?
Model answer: You also need hardening, runtime controls, network restrictions, and monitoring because not all CVEs have immediate fixes.
Common mistake: Framing security as patch-only activity.

9. What is process visibility in Linux?
Model answer: Observing running processes, privileges, and open sockets using tools like ps and ss to detect suspicious behavior.
Common mistake: Listing tools without explaining why.

10. What is defense in depth?
Model answer: Multiple layers of controls across build, image, runtime, network, and policy so one failure does not expose everything.
Common mistake: Giving only one control example.

11. Why avoid unnecessary packages?
Model answer: Each package can introduce vulnerabilities and operational complexity, so removing unused packages reduces attack surface.
Common mistake: Only mentioning image size, not security.

12. What are Linux capabilities?
Model answer: Fine-grained privileges split from root. Dropping unnecessary capabilities reduces exploit power.
Common mistake: Confusing with Kubernetes RBAC.

13. What is a secure default?
Model answer: System starts in safest practical configuration, such as no root login, minimal open ports, and deny-by-default policies.
Common mistake: Treating secure defaults as optional.

14. Why do logs matter in CVE remediation?
Model answer: Logs help validate exploit attempts, detect regressions after patching, and provide audit trace for response.
Common mistake: Not connecting logs to remediation lifecycle.

15. How do you validate package version in container?
Model answer: Run package-specific query inside container, e.g., dpkg -l, rpm -qa, or apk info, then compare with advisory fixed version.
Common mistake: Trusting scanner output without verification.

16. What is UID/GID significance in containers?
Model answer: They define runtime identity and filesystem access. Correct UID/GID setup prevents privilege and permission issues.
Common mistake: Ignoring file ownership behavior with mounted volumes.

17. Why pin versions?
Model answer: Pinning improves reproducibility and reduces surprise changes; digest pinning is stronger than mutable tags.
Common mistake: Using latest everywhere.

18. What is exploitability context?
Model answer: Real risk depends on exposure, privileges, network path, and whether vulnerable code path is reachable.
Common mistake: Prioritizing purely by CVSS score.

19. What does hardening mean?
Model answer: Systematic reduction of attack surface and abuse paths while preserving required functionality.
Common mistake: Equating hardening with one-time patching.

20. How do you explain Linux security to non-security stakeholders?
Model answer: Use business impact language: lower breach risk, fewer emergency incidents, stronger compliance, predictable delivery.
Common mistake: Overusing jargon.

## Section B: Containers and Dockerfile (21-40)

21. What is a container image layer?
Model answer: Immutable filesystem snapshot created by build instructions. Final image is stacked layers.
Common mistake: Thinking layers are only performance feature.

22. Why do layers matter for vulnerabilities?
Model answer: Vulnerable files can persist across layers; scanning inspects components in the final artifact composition.
Common mistake: Assuming overwriting always removes risk context.

23. Tag vs digest?
Model answer: Tag is mutable label; digest is immutable content hash. Security workflows prefer digest pinning.
Common mistake: Treating tags as immutable.

24. Why use multi-stage builds?
Model answer: Keep build tools out of runtime image, reducing size and vulnerability surface.
Common mistake: Using multi-stage only for speed narrative.

25. Why remove apt cache in Dockerfile?
Model answer: It reduces final image size and avoids shipping unnecessary package metadata.
Common mistake: Leaving cache in runtime image.

26. What does .dockerignore do?
Model answer: Prevents unnecessary files from entering build context, reducing leakage risk and build time.
Common mistake: Not excluding secrets and local artifacts.

27. Distroless vs slim?
Model answer: Distroless minimizes runtime components and attack surface; slim is easier for debugging and compatibility.
Common mistake: Declaring one universally better.

28. Why set USER in Dockerfile?
Model answer: Enforces non-root runtime identity by default.
Common mistake: Creating user but forgetting USER directive.

29. What is a secure base image strategy?
Model answer: Trusted source, minimal footprint, update cadence, digest pinning, and periodic refresh.
Common mistake: Picking smallest image without compatibility and patch cadence review.

30. How do you verify image content?
Model answer: Use image history, package listings, and scanner outputs to inspect components and layers.
Common mistake: Relying on Dockerfile intent only.

31. Why avoid SSH in containers?
Model answer: Increases attack surface and violates immutable infrastructure principles. Use orchestration-native debugging workflows instead.
Common mistake: Treating container as traditional VM.

32. What is ENTRYPOINT vs CMD?
Model answer: ENTRYPOINT sets executable; CMD sets default args or command. Combined usage controls runtime behavior.
Common mistake: Not understanding override behavior.

33. Why separate build and runtime dependencies?
Model answer: Build deps are often high-risk and unnecessary at runtime; excluding them reduces CVEs.
Common mistake: Installing all tools in final stage.

34. How can image size reduction improve security?
Model answer: Fewer binaries and libraries means fewer vulnerable components and reduced attack surface.
Common mistake: Discussing cost only.

35. Why avoid privileged containers?
Model answer: Privileged mode grants broad host access, defeating isolation boundaries.
Common mistake: Using privileged for convenience in production.

36. What is reproducible build value?
Model answer: Predictable artifacts improve trust, auditing, and incident rollback analysis.
Common mistake: Ignoring build determinism.

37. How do you handle secret injection in build?
Model answer: Avoid baking secrets in image layers; use secure runtime secret management.
Common mistake: Passing secrets via ARG and committing layers.

38. What is immutable infrastructure principle in containers?
Model answer: Build new image for changes rather than modifying running containers.
Common mistake: Hot-fixing inside running container.

39. Why expose only required ports?
Model answer: Limits attack paths and accidental service exposure.
Common mistake: Exposing broad port ranges by default.

40. What is common Dockerfile anti-pattern?
Model answer: Large base + root user + extra debug tools + no patching + no cleanup.
Common mistake: Missing one of these while claiming image is hardened.

## Section C: CVE Management and Remediation (41-60)

41. What is your CVE remediation lifecycle?
Model answer: Discover, validate, prioritize, patch/mitigate, rescan, test, deploy safely, document, and automate prevention.
Common mistake: Skipping validation or deployment safety.

42. How do you prioritize CVEs?
Model answer: Severity + exploitability + exposure + privilege context + business criticality.
Common mistake: Prioritizing by CVSS only.

43. Why HIGH CVEs after patching?
Model answer: Old base digest, stale scanner DB, package still present, no upstream fix, or findings from sidecars/node images.
Common mistake: Assuming scanner is wrong without investigation.

44. How do you validate scanner finding?
Model answer: Confirm package and version in container, cross-check advisory, verify fixed version availability.
Common mistake: Blindly suppressing findings.

45. What if fix is unavailable?
Model answer: Apply compensating controls, document exception with expiry, monitor upstream fix, retest regularly.
Common mistake: Permanent exception without mitigation.

46. Patch vs remove package decision?
Model answer: If package is unnecessary, remove it. If required, patch to fixed version and validate functionality.
Common mistake: Keeping vulnerable unused packages.

47. How to handle false positive suspicion?
Model answer: Validate with secondary source, verify package mapping, review vendor advisory status, and document decision.
Common mistake: Ignoring finding without audit trail.

48. How do you avoid patch regressions?
Model answer: Canary rollout, smoke/integration tests, observability gates, and quick rollback path.
Common mistake: Patching directly to full production.

49. What remediation evidence do you capture?
Model answer: CVE ID, affected artifact, change set, scan before/after, test results, approval, deployment record.
Common mistake: No traceable audit evidence.

50. What is MTTR in vulnerability management?
Model answer: Mean time to remediate; key KPI for response efficiency.
Common mistake: Reporting counts without response time.

51. How to reduce recurring CVEs?
Model answer: Golden base images, automated rebuilds, dependency hygiene, and policy gates.
Common mistake: Fixing case-by-case only.

52. What is vulnerability backlog health?
Model answer: Trend of unresolved CVEs over time by severity and age.
Common mistake: Showing only current snapshot.

53. What is CVE exception expiry?
Model answer: Time-bound approval forcing periodic review, preventing permanent unmanaged risk.
Common mistake: Exceptions with no owner or expiry.

54. How do you communicate risk tradeoff?
Model answer: Explain exploit path, business impact, mitigation strength, and deadline for full fix.
Common mistake: Saying only "security best practice".

55. How do you handle transitive dependency CVE?
Model answer: Identify introducing layer/component, upgrade parent dependency or rebase image, then verify.
Common mistake: Assuming app does not import it so it is irrelevant.

56. How often should images be rebuilt?
Model answer: Event-driven plus scheduled cadence, e.g., weekly refresh and urgent rebuild on critical advisories.
Common mistake: Rebuilding only when app code changes.

57. What is canary in remediation rollout?
Model answer: Release patched image to small traffic slice first, monitor error and security signals, then expand.
Common mistake: Calling canary without monitoring criteria.

58. How to compare vulnerability tools?
Model answer: Normalize by package and CVE, check data-source differences, investigate disagreements for critical findings.
Common mistake: Selecting one tool as absolute truth.

59. How to prove CVE closed?
Model answer: Rescan updated artifact, verify fixed package version, confirm no regression in runtime tests.
Common mistake: Closing ticket after code change only.

60. What metrics to present leadership?
Model answer: Critical/high trend, MTTR, exception aging, remediation SLA compliance, and incident correlation.
Common mistake: Too technical dashboard without business framing.

## Section D: Hardening and Runtime Security (61-75)

61. What is attack surface reduction?
Model answer: Removing or constraining components and paths that attackers can abuse.
Common mistake: Equating only with smaller image.

62. Why read-only root filesystem?
Model answer: Limits write-based persistence and tampering in compromised containers.
Common mistake: Enabling without checking app write paths.

63. Why drop Linux capabilities?
Model answer: Reduces privileged operations available to process, minimizing impact of compromise.
Common mistake: Dropping none by default.

64. What is seccomp purpose?
Model answer: Restricts allowed system calls to reduce kernel attack surface.
Common mistake: Mentioning without syscall context.

65. What is AppArmor/SELinux role?
Model answer: Mandatory access control to constrain process behavior beyond basic permissions.
Common mistake: Confusing with RBAC.

66. What runtime settings in Kubernetes are high value?
Model answer: runAsNonRoot, allowPrivilegeEscalation false, readOnlyRootFilesystem true, dropped capabilities.
Common mistake: Only discussing image hardening.

67. Why network policy matters for CVEs?
Model answer: It reduces reachable exploit path and lateral movement potential.
Common mistake: Ignoring east-west exposure.

68. How to secure secrets with containers?
Model answer: Use secret stores and mount at runtime; never bake secrets into image layers.
Common mistake: ENV secrets in Dockerfile.

69. What is Pod Security admission purpose?
Model answer: Enforce baseline/restricted controls at admission so unsafe pods are blocked.
Common mistake: Treating policies as optional docs.

70. Why not install debug tools in production image?
Model answer: They add unnecessary binaries and potential vulnerabilities.
Common mistake: Keeping curl/bash/vim by default.

71. How do you debug distroless images safely?
Model answer: Use sidecar/ephemeral debug containers and observability tooling rather than bloating runtime image.
Common mistake: Replacing distroless with full distro permanently.

72. What is immutable deployment benefit?
Model answer: Predictable rollback and traceability because artifacts are fixed and versioned.
Common mistake: Ad-hoc shell changes in running pods.

73. How do you handle legacy apps needing root?
Model answer: Contain risk with compensating controls, roadmap refactor to non-root, and strict runtime policy boundaries.
Common mistake: Accepting root indefinitely without plan.

74. How does hardening impact reliability?
Model answer: Can improve resilience but must be tested because strict controls may break assumptions.
Common mistake: Deploying hardening without compatibility validation.

75. What is secure-by-default platform pattern?
Model answer: Golden images, policy-as-code, reusable secure templates, and automated guardrails.
Common mistake: Depending on manual team discipline only.

## Section E: SBOM, Signing, Supply Chain (76-88)

76. What is SBOM?
Model answer: Inventory of components in software artifact, used for vulnerability and compliance management.
Common mistake: Calling it just a dependency file.

77. Why SBOM if scanners exist?
Model answer: SBOM gives structured inventory and faster impact analysis across advisories and assets.
Common mistake: Treating SBOM as duplicate scan report.

78. SPDX vs CycloneDX?
Model answer: Both are SBOM formats; choice depends on ecosystem tooling and integration requirements.
Common mistake: Claiming one is always required.

79. What does cosign do?
Model answer: Signs container artifacts and enables signature verification for trusted deployment.
Common mistake: Confusing signing with encryption.

80. What is provenance?
Model answer: Metadata proving what source and process produced artifact, by whom and when.
Common mistake: Ignoring build identity and source linkage.

81. What risk does signing reduce?
Model answer: Deployment of tampered or untrusted images.
Common mistake: Saying it removes vulnerabilities.

82. How do you enforce signed images?
Model answer: Admission policy verifies signatures and trusted identities before workload is admitted.
Common mistake: Signing without verification policy.

83. What is SLSA practical value?
Model answer: Maturity framework for build integrity, provenance, and tamper resistance.
Common mistake: Describing only theory with no controls.

84. How to secure CI credentials?
Model answer: Short-lived identities, OIDC federation, secret minimization, scoped permissions.
Common mistake: Long-lived static tokens.

85. What is artifact trust chain?
Model answer: Link source commit, build pipeline identity, signed artifact, and deployment policy verification.
Common mistake: Missing one stage of trust.

86. How to respond to supply chain incident?
Model answer: Identify impacted artifacts via SBOM/provenance, block deployments, rotate trust material, rebuild from trusted base.
Common mistake: Focusing only on app code patch.

87. What is dependency confusion risk?
Model answer: Pulling malicious package from public source due to namespace/version misconfiguration.
Common mistake: Assuming private registry usage automatically prevents it.

88. How do you prove artifact integrity to auditors?
Model answer: Provide signed artifacts, verification logs, provenance records, and policy enforcement evidence.
Common mistake: Providing screenshots without verifiable records.

## Section F: CI/CD, Scale, and Leadership (89-100)

89. What should fail a security gate?
Model answer: Policy-defined thresholds, typically critical CVEs, failing signature checks, or prohibited base images.
Common mistake: Binary pass/fail without exception governance.

90. How do you avoid developer friction?
Model answer: Shift-left templates, clear remediation guidance, fast feedback, and controlled exceptions.
Common mistake: Strict policies without support paths.

91. How do you design remediation at scale for 300+ images?
Model answer: Golden base images, centralized rebuild orchestration, risk-based prioritization, and fleet-level dashboards.
Common mistake: Manual per-service operations.

92. What is a golden image program?
Model answer: Curated, hardened base images maintained centrally and consumed by teams.
Common mistake: Creating golden images without update SLA.

93. How do you measure program success?
Model answer: Reduction in critical exposure, faster MTTR, lower exception age, stable release quality.
Common mistake: Measuring only scan count.

94. How do you present to executives?
Model answer: Translate technical work into business risk reduction, uptime impact, and compliance outcomes.
Common mistake: Tool-level details only.

95. Describe conflict with delivery team and how you handled it.
Model answer: Align on risk, propose phased controls, add compensating safeguards, and agree measurable deadlines.
Common mistake: Framing security as veto-only function.

96. How do you coach juniors on remediation quality?
Model answer: Provide checklists, pair on triage, review evidence quality, and enforce repeatable templates.
Common mistake: Delegating without standards.

97. What is your first 30-day plan in this role?
Model answer: Baseline risk, identify top exposures, establish quick-win hardening, define policy and exception workflow, launch metrics.
Common mistake: Overpromising full transformation immediately.

98. What would you automate first?
Model answer: Build-scan-policy loop with critical blocking and auto-ticketing for high findings.
Common mistake: Automating everything at once.

99. How do you ensure security changes survive team turnover?
Model answer: Codify policies, templates, runbooks, and ownership; avoid tribal knowledge.
Common mistake: People-dependent processes.

100. Why should we hire you for this role?
Model answer: I combine hands-on remediation, platform automation, and clear risk communication. I can reduce CVE exposure quickly while keeping delivery stable and measurable.
Common mistake: Generic answer without role-specific outcomes.

## Practice Drill
- Pick 10 random questions daily.
- Answer aloud in 90 seconds each.
- Self-score on clarity, depth, and outcome evidence.
