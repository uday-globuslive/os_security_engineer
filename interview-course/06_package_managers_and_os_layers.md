# Module 06: Package Managers and OS Layers

## Learning Objectives
- Understand apt, yum/dnf, and apk behavior.
- Explain how OS packages map to scanner findings.

## 1) apt (Debian/Ubuntu)
Core commands:
```bash
apt-get update
apt-get install -y <pkg>
apt-get upgrade -y
apt-cache policy <pkg>
```

Important:
- `update` refreshes metadata.
- `upgrade` applies newer versions.

## 2) yum/dnf (RHEL Family)
Core commands:
```bash
yum check-update
yum update -y
yum info <pkg>
```

## 3) apk (Alpine)
Core commands:
```bash
apk update
apk add <pkg>
apk upgrade
apk info -v <pkg>
```

## 4) Layer-Aware Package Management
If a vulnerable package is added in one layer and removed incorrectly in another, traces can remain and scanners may still report risk depending on artifact state.

## 5) Patch Strategy
- Rebase on patched base image digest.
- Update packages in a deterministic build step.
- Remove unnecessary packages from runtime stage.
- Rebuild cleanly and rescan.

## 6) Interview Questions
- Why can scanners still detect vulnerabilities after package updates?
- How do you verify package version accurately in a built image?
- When would you choose Alpine vs Debian-slim?
