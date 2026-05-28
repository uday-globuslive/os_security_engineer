# Module 28: Debugging Labs (50 Real-World Scenarios)

## How to use
For each lab:
1. Read symptom.
2. State likely root cause.
3. Run validation command(s).
4. Apply fix.
5. Explain interview-ready lesson.

## Labs 1-10: Linux and Package Issues
1. Symptom: CVE remains after patching Ubuntu image. Cause: base digest stale. Validate: check image digest and package version. Fix: rebase to patched digest and rebuild with pull.
2. Symptom: `apt-get install` fails in CI. Cause: outdated package indexes. Validate: inspect apt logs. Fix: run `apt-get update` in same build step.
3. Symptom: App cannot read config file. Cause: wrong file ownership after non-root switch. Validate: `ls -l`. Fix: `chown` config path to runtime user.
4. Symptom: High CPU in container after patch. Cause: changed library behavior or busy loop exposed. Validate: `top`, app metrics. Fix: rollback/canary comparison and patch-level pinning.
5. Symptom: Service restart loop after hardening. Cause: app writes to read-only root. Validate: runtime logs. Fix: mount writable temp path.
6. Symptom: Scanner flags package not used by app. Cause: extra package installed. Validate: package list in image. Fix: remove package from runtime stage.
7. Symptom: `apk add` fails on Alpine. Cause: mirror/index mismatch. Validate: repo config. Fix: refresh indexes and pin correct repo.
8. Symptom: Kernel warning spikes on nodes. Cause: host kernel regression. Validate: `dmesg`. Fix: node patch strategy and controlled roll.
9. Symptom: Missing library in distroless runtime. Cause: copied binary needs shared libs. Validate: run in debug image. Fix: static compile or include required libs.
10. Symptom: SSH key leaked in image layer. Cause: copied home directory in build context. Validate: image history and file scan. Fix: `.dockerignore` and secret hygiene.

## Labs 11-20: Dockerfile and Image Build
11. Symptom: Huge image size unexpectedly. Cause: build tools left in runtime. Validate: `docker history`. Fix: multi-stage build.
12. Symptom: Build reproducibility poor. Cause: mutable tags. Validate: compare digest over time. Fix: pin digest.
13. Symptom: CVE count increases after simple app change. Cause: base image drift. Validate: compare base digest. Fix: controlled base image update policy.
14. Symptom: Container runs as root despite user creation. Cause: missing `USER` instruction. Validate: runtime `id`. Fix: set `USER` explicitly.
15. Symptom: Build cache keeps vulnerable package. Cause: stale cached layer. Validate: run no-cache build. Fix: rebuild with `--no-cache`.
16. Symptom: Build fails on `pip install`. Cause: system dependency missing. Validate: pip error logs. Fix: install build dependency in builder stage only.
17. Symptom: Runtime command fails in distroless. Cause: shell-dependent startup script. Validate: entrypoint inspection. Fix: direct binary execution.
18. Symptom: Duplicate packages across layers. Cause: repeated install steps. Validate: history/layer diff. Fix: consolidate RUN instructions.
19. Symptom: App starts slowly after patch. Cause: package set changed startup path. Validate: startup profiling/logs. Fix: tune and trim dependencies.
20. Symptom: Unauthorized binary in image. Cause: inherited from base. Validate: SBOM diff. Fix: switch base or strip package.

## Labs 21-30: Kubernetes Runtime and Policy
21. Symptom: Pod blocked by admission policy. Cause: privileged setting denied. Validate: `kubectl describe pod` events. Fix: align securityContext.
22. Symptom: Pod cannot bind low port as non-root. Cause: privileged port requirement. Validate: app logs. Fix: move to high port or capability policy.
23. Symptom: Service unavailable after network policy. Cause: denied east-west traffic. Validate: policy and flow tests. Fix: add explicit allow rule.
24. Symptom: Secret not mounted in pod. Cause: wrong namespace/secret name. Validate: pod spec and secret list. Fix: correct reference.
25. Symptom: CrashLoopBackOff post-hardening. Cause: permission denied on temp path. Validate: logs and mount config. Fix: writable `emptyDir`.
26. Symptom: Rollout stuck. Cause: readiness probe failing. Validate: probe events and endpoint check. Fix: adjust probe path/timing.
27. Symptom: Image pull denied. Cause: missing imagePullSecret/registry policy. Validate: events. Fix: attach secret and allowed registry config.
28. Symptom: Signed image rejected. Cause: identity mismatch in verify policy. Validate: signature cert identity. Fix: update trust policy.
29. Symptom: Pod evicted frequently. Cause: resource limits too low. Validate: events and metrics. Fix: right-size requests/limits.
30. Symptom: Lateral movement concern after compromise. Cause: no namespace egress control. Validate: network policy inventory. Fix: default deny + minimal allow.

## Labs 31-40: Scanning and Supply Chain
31. Symptom: Trivy and Grype severity mismatch. Cause: advisory source difference. Validate: compare CVE records. Fix: triage with policy normalization.
32. Symptom: False closure of CVE ticket. Cause: no rescan evidence. Validate: ticket artifacts. Fix: mandatory before/after report.
33. Symptom: SBOM missing transitive component. Cause: generation at wrong stage. Validate: compare artifact stage. Fix: generate from final image.
34. Symptom: Signature exists but deploy blocked. Cause: signature attached to different digest. Validate: digest match. Fix: sign exact pushed digest.
35. Symptom: Pipeline scan too slow. Cause: scanning full monorepo each run. Validate: timing metrics. Fix: incremental scan strategy.
36. Symptom: Critical CVE no fix upstream. Cause: vendor status pending. Validate: advisory timeline. Fix: compensating controls + exception expiry.
37. Symptom: High count spikes overnight. Cause: scanner DB update. Validate: advisory update logs. Fix: communicate risk delta and triage.
38. Symptom: SBOM artifact lost. Cause: CI retention misconfigured. Validate: artifact policy. Fix: retention and external store.
39. Symptom: Unsigned image deployed in dev. Cause: relaxed environment policy. Validate: admission logs. Fix: unify trust policy by environment tier.
40. Symptom: Supply chain alert for dependency confusion. Cause: mixed registry resolution. Validate: lockfile and registry config. Fix: private registry precedence and pinning.

## Labs 41-50: CI/CD and Incident Response
41. Symptom: Too many failed builds from HIGH CVEs. Cause: over-strict gate with no exception process. Validate: pipeline policy. Fix: critical block + time-bound HIGH exceptions.
42. Symptom: Emergency patch causes incident. Cause: no canary rollout. Validate: deploy history. Fix: phased rollout and rollback checkpoints.
43. Symptom: Vulnerable sidecar missed in scan. Cause: app image only scanned. Validate: deployment image inventory. Fix: scan all images in manifest.
44. Symptom: MTTR remains high. Cause: manual triage bottleneck. Validate: process mapping. Fix: auto-ticketing and severity routing.
45. Symptom: CVE reopened repeatedly. Cause: base image lifecycle unmanaged. Validate: recurrence trend. Fix: golden base program.
46. Symptom: Incident response confusion. Cause: unclear ownership. Validate: postmortem. Fix: defined RACI and runbook.
47. Symptom: Rollback fails. Cause: artifact immutability gaps. Validate: release artifact trace. Fix: immutable versioning and retention.
48. Symptom: Production image differs from tested one. Cause: mutable tag reuse. Validate: digest mismatch. Fix: deploy by digest only.
49. Symptom: Audit asks for proof of fix and signing. Cause: incomplete evidence chain. Validate: artifact and logs. Fix: store scan, SBOM, signature, deploy verification.
50. Symptom: Team distrusts security gate output. Cause: noisy findings and poor guidance. Validate: feedback and false-positive stats. Fix: tuned policy and actionable remediation docs.

## Scoring your lab performance
- 1 point: identified likely cause.
- 1 point: selected correct validation command.
- 1 point: proposed safe fix.
- 1 point: explained prevention action.

Target: 160/200 across all 50 labs.
