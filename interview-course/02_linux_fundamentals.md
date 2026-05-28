# Module 02: Linux Fundamentals for Security Engineers

## Learning Objectives
- Understand Linux users, permissions, processes, filesystems, and package basics.
- Understand how Linux knowledge maps directly to container security.

## 1) Linux Architecture Basics
Linux is composed of:
- Kernel: manages hardware, memory, process scheduling.
- User space: shells, tools, libraries, applications.
- Filesystem hierarchy: organized directories and mount points.

Why interviewers care:
- Container vulnerabilities often come from user-space libraries and packages.

## 2) Users, Groups, and Permission Model
- User ID (UID) and Group ID (GID) identify process identity.
- Permission bits apply to user, group, others.

Command examples:
```bash
id
whoami
ls -l /etc/passwd
chmod 640 file.txt
chown appuser:appgroup file.txt
```

Security meaning:
- Least privilege limits blast radius.
- Root processes increase risk in compromise scenarios.

## 3) Process and Service Visibility
Core commands:
```bash
ps aux
ss -tulnp
top
```

Security meaning:
- Detect unexpected listeners and high-privilege processes.

## 4) Filesystem and Sensitive Paths
Important paths:
- `/etc` for configs.
- `/var/log` for logs.
- `/tmp` for temporary files.
- `/usr/lib` and `/lib` for libraries.

Security meaning:
- Vulnerable libraries often reside in common library paths.

## 5) Package Management Intro
- Debian/Ubuntu: `apt`
- RHEL: `yum` or `dnf`
- Alpine: `apk`

Basic flow:
```bash
apt-get update
apt-get upgrade -y
apt-cache policy openssl
```

## 6) Linux Security Basics You Must Explain
- Principle of least privilege.
- Defense in depth.
- Patch management lifecycle.
- Secure defaults.

## Common Interview Questions
- Why is running as root dangerous in containers?
- Difference between `apt-get update` and `apt-get upgrade`?
- How do you verify a vulnerable package version?

## Module Exercise
- Create a small Linux command cheat sheet from this module.
- Explain to yourself how each command supports security operations.
