# Linux and Container Security Basics (Beginner to Intermediate)

## 1) Linux Security Fundamentals

### Users, Groups, and Permissions
- Every process runs as a user identity (UID/GID).
- Permission model: read (r), write (w), execute (x) for user/group/others.
- Principle: give the least permissions needed.

Commands to know:
```bash
id
whoami
ls -l
chmod 640 file
chown appuser:appgroup file
```

### Processes and Services
- A process can be compromised if it runs with high privileges.
- Avoid running application processes as root.

Commands to know:
```bash
ps aux
top
ss -tulnp
```

### Package Management Basics
- Debian/Ubuntu: apt
- RHEL family: yum or dnf
- Alpine: apk

Security-relevant actions:
- Refresh package indexes.
- Upgrade vulnerable packages.
- Remove unneeded packages.

## 2) Container Fundamentals

### What Is a Container Image?
- Read-only layers stacked together.
- Built from Dockerfile instructions.
- A running container adds a writable layer on top.

Why this matters:
- Vulnerabilities in any layer can be detected.
- Old vulnerable layers persist until rebuilt with patched content.

### Key Runtime Isolation Concepts
- Namespaces: isolate process IDs, network, mount points, etc.
- Cgroups: resource limits.
- Capabilities: fine-grained Linux privileges.

## 3) Basic Hardening Controls

### Run as Non-Root
Example:
```dockerfile
RUN addgroup --system app && adduser --system --ingroup app app
USER app
```

### Use Minimal Base Images
- Prefer slim/minimal distributions.
- Use distroless if debugging needs are addressed elsewhere.

### Multi-Stage Builds
- Compile in builder stage.
- Copy only runtime artifacts into final image.

### Reduce Exposed Surface
- Remove package managers and build tools from runtime stage.
- Expose only required ports.
- Do not ship shells/editors/network tools unless required.

## 4) Basic Vulnerability Scanning Workflow
1. Build image.
2. Scan with Trivy or Grype.
3. Triage high and critical findings.
4. Patch or remove vulnerable components.
5. Rebuild and rescan.
6. Run functional tests.

## 5) Common Beginner Mistakes
- Using `latest` tags without digest pinning.
- Running as root by default.
- Installing tools in runtime image and forgetting to remove them.
- Ignoring scanner DB update timing.
- Assuming no CVEs means no risk.
