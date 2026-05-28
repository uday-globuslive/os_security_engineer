# Module 26: Command Handbook (200 Practical Commands)

## How to use this handbook
- Practice 20 commands per day.
- For each command, know purpose, expected output, and one troubleshooting use-case.
- Use this as rapid recall before interviews.

## Section A: Linux Basics and System Discovery (1-20)
1. `uname -a` - Show kernel and system details.
2. `cat /etc/os-release` - Show Linux distribution info.
3. `hostnamectl` - Show host metadata.
4. `whoami` - Show current user.
5. `id` - Show UID/GID and groups.
6. `pwd` - Print working directory.
7. `ls -la` - List files with permissions and hidden entries.
8. `tree -L 2` - Show directory tree quickly.
9. `cd /path` - Change directory.
10. `stat file.txt` - Show file metadata.
11. `file binary` - Identify file type.
12. `du -sh *` - Show size summary by item.
13. `df -h` - Show disk usage.
14. `free -m` - Show memory usage.
15. `uptime` - Show load and uptime.
16. `date -u` - Show UTC date/time.
17. `timedatectl` - Show time sync status.
18. `env | sort` - Show environment variables.
19. `printenv PATH` - Show PATH.
20. `history | tail -20` - Show recent commands.

## Section B: Files, Permissions, and Search (21-40)
21. `touch demo.txt` - Create empty file.
22. `mkdir -p a/b/c` - Create nested directories.
23. `cp source.txt dest.txt` - Copy file.
24. `mv old.txt new.txt` - Move/rename file.
25. `rm -f temp.txt` - Remove file.
26. `rmdir emptydir` - Remove empty directory.
27. `chmod 640 secret.txt` - Set secure permissions.
28. `chmod -R 750 appdir` - Recursive permission change.
29. `chown app:app secret.txt` - Change owner/group.
30. `umask 027` - Set default file permission mask.
31. `find . -name "*.log"` - Find logs.
32. `find / -type f -perm -4000 2>/dev/null` - Find SUID files.
33. `locate docker.sock` - Fast file locate (if updatedb exists).
34. `grep -R "password" .` - Recursive text search.
35. `grep -n "ERROR" app.log` - Find line numbers for errors.
36. `awk '{print $1}' file.txt` - Column extraction.
37. `sed -n '1,20p' file.txt` - Print specific lines.
38. `sort file.txt | uniq -c` - Count unique lines.
39. `cut -d: -f1 /etc/passwd` - Parse usernames.
40. `head -20 file.txt && tail -20 file.txt` - Quick file preview.

## Section C: Processes, Services, and Logs (41-60)
41. `ps aux` - List processes.
42. `ps -ef | grep nginx` - Search process.
43. `top` - Live process view.
44. `htop` - Enhanced live process view.
45. `pgrep -a python` - Find process IDs and cmd lines.
46. `pkill -f process_name` - Kill by pattern.
47. `kill -15 <pid>` - Graceful terminate.
48. `kill -9 <pid>` - Force kill.
49. `systemctl status docker` - Service status.
50. `systemctl restart docker` - Restart service.
51. `systemctl list-units --type=service` - List services.
52. `journalctl -u docker -n 100` - Service logs.
53. `journalctl -xe` - Recent critical logs.
54. `dmesg | tail -50` - Kernel messages.
55. `last -n 10` - Recent logins.
56. `lastlog` - Last login per user.
57. `w` - Logged-in users and activity.
58. `lsof -i :8080` - Process using a port.
59. `strace -p <pid>` - Trace syscalls for a process.
60. `nohup longcmd &` - Run command detached from terminal.

## Section D: Networking and Troubleshooting (61-80)
61. `ip a` - Show interfaces and addresses.
62. `ip route` - Show routing table.
63. `ss -tulnp` - Show listening ports and processes.
64. `netstat -tulnp` - Alternative port view.
65. `ping -c 4 8.8.8.8` - Basic reachability test.
66. `curl -I https://example.com` - HTTP headers check.
67. `curl -v http://localhost:8080/health` - Verbose endpoint check.
68. `wget -S https://example.com -O /dev/null` - HTTP status check.
69. `dig example.com` - DNS lookup.
70. `nslookup example.com` - DNS query alternative.
71. `traceroute example.com` - Route path visibility.
72. `mtr -rw example.com` - Combined latency and route stats.
73. `tcpdump -i any port 443` - Packet capture by port.
74. `nc -zv localhost 5432` - TCP connectivity probe.
75. `telnet host 80` - Legacy connectivity test.
76. `arp -a` - ARP cache view.
77. `hostname -I` - Show host IPs.
78. `nmap -sV target` - Service discovery scan.
79. `openssl s_client -connect host:443` - TLS cert inspection.
80. `curl --resolve host:443:IP https://host` - Test host with custom DNS mapping.

## Section E: Package Managers and Patching (81-100)
81. `apt-get update` - Refresh apt metadata.
82. `apt-get upgrade -y` - Apply package upgrades.
83. `apt-get dist-upgrade -y` - Full upgrade with dependency changes.
84. `apt-cache policy openssl` - Package version policy.
85. `dpkg -l | grep openssl` - Debian package installed version.
86. `apt-get install --only-upgrade openssl -y` - Upgrade specific package.
87. `apt-get autoremove -y` - Remove unused dependencies.
88. `apt-get clean` - Clean package cache.
89. `yum check-update` - Check updates (RHEL).
90. `yum update -y` - Apply updates (RHEL).
91. `rpm -qa | grep openssl` - Installed RPM query.
92. `dnf update -y` - Apply updates with dnf.
93. `apk update` - Refresh Alpine indexes.
94. `apk upgrade` - Upgrade Alpine packages.
95. `apk add curl` - Install Alpine package.
96. `apk info -v openssl` - Alpine package info.
97. `zypper refresh` - Refresh SUSE repos.
98. `zypper update -y` - Update SUSE packages.
99. `needrestart -r a` - Restart services after updates (if installed).
100. `apt-mark hold package_name` - Hold package version.

## Section F: Docker Core Commands (101-120)
101. `docker version` - Show Docker client/server versions.
102. `docker info` - Docker daemon and config info.
103. `docker images` - List local images.
104. `docker ps` - List running containers.
105. `docker ps -a` - List all containers.
106. `docker pull nginx:1.27` - Pull image.
107. `docker run --rm -it alpine:3.20 sh` - Start interactive container.
108. `docker run -d -p 8080:80 nginx:1.27` - Run detached container with port mapping.
109. `docker exec -it <cid> sh` - Enter running container.
110. `docker logs -f <cid>` - Follow container logs.
111. `docker stop <cid>` - Stop container.
112. `docker rm <cid>` - Remove container.
113. `docker rmi image:tag` - Remove image.
114. `docker inspect <cid>` - Detailed container metadata.
115. `docker history image:tag` - Show image layers.
116. `docker diff <cid>` - Files changed in running container.
117. `docker cp <cid>:/file ./file` - Copy file from container.
118. `docker stats` - Resource usage live.
119. `docker top <cid>` - Processes inside container.
120. `docker system prune -f` - Remove unused Docker resources.

## Section G: Docker Build, Security, and Registry (121-140)
121. `docker build -t app:local .` - Build image.
122. `docker build --no-cache -t app:clean .` - Rebuild without cache.
123. `docker build --pull -t app:fresh .` - Refresh base image during build.
124. `docker tag app:local registry/app:1.0.0` - Tag image for registry.
125. `docker login registry.example.com` - Authenticate to registry.
126. `docker push registry/app:1.0.0` - Push image.
127. `docker manifest inspect image:tag` - Inspect manifest.
128. `docker save app:local -o app.tar` - Save image tarball.
129. `docker load -i app.tar` - Load image tarball.
130. `docker sbom app:local` - Generate SBOM (if plugin supported).
131. `docker scout quickview app:local` - Quick vulnerability insights (if available).
132. `docker scan app:local` - Scan image vulnerabilities (if available).
133. `docker run --read-only app:local` - Run with read-only filesystem.
134. `docker run --user 10001:10001 app:local` - Run as non-root UID/GID.
135. `docker run --cap-drop ALL app:local` - Drop capabilities.
136. `docker run --security-opt no-new-privileges app:local` - Prevent privilege escalation.
137. `docker network ls` - List Docker networks.
138. `docker volume ls` - List volumes.
139. `docker compose up -d` - Start stack from compose.
140. `docker compose logs -f` - Follow compose logs.

## Section H: Kubernetes Commands (141-160)
141. `kubectl version --short` - Client/server versions.
142. `kubectl cluster-info` - Cluster endpoints info.
143. `kubectl get nodes -o wide` - Node details.
144. `kubectl get ns` - Namespaces list.
145. `kubectl config get-contexts` - Available contexts.
146. `kubectl config use-context myctx` - Switch context.
147. `kubectl get pods -A` - All pods.
148. `kubectl get deployments -A` - All deployments.
149. `kubectl describe pod <pod> -n <ns>` - Pod detail and events.
150. `kubectl logs <pod> -n <ns>` - Pod logs.
151. `kubectl logs -f <pod> -n <ns>` - Follow pod logs.
152. `kubectl exec -it <pod> -n <ns> -- sh` - Exec into pod.
153. `kubectl get svc -A` - Services list.
154. `kubectl get ingress -A` - Ingress list.
155. `kubectl top nodes` - Node metrics.
156. `kubectl top pods -A` - Pod metrics.
157. `kubectl get events -A --sort-by=.lastTimestamp` - Cluster event timeline.
158. `kubectl rollout status deployment/app -n <ns>` - Rollout status.
159. `kubectl rollout undo deployment/app -n <ns>` - Rollback deployment.
160. `kubectl delete pod <pod> -n <ns>` - Delete pod (forces recreation if managed).

## Section I: Kubernetes Security and Policy (161-180)
161. `kubectl auth can-i create pods -n prod --as user@example.com` - RBAC permission check.
162. `kubectl get role,rolebinding -n prod` - Namespace RBAC bindings.
163. `kubectl get clusterrole,clusterrolebinding` - Cluster RBAC bindings.
164. `kubectl get networkpolicy -A` - Network policy inventory.
165. `kubectl describe networkpolicy np-name -n ns` - Policy details.
166. `kubectl get pod <pod> -n <ns> -o yaml` - Inspect securityContext and image.
167. `kubectl explain pod.spec.securityContext` - Security context schema reference.
168. `kubectl get psp` - PodSecurityPolicy list (legacy clusters).
169. `kubectl get pod -A -o jsonpath='{range .items[*]}{.metadata.namespace} {.metadata.name} {.spec.securityContext.runAsNonRoot}{"\n"}{end}'` - Check runAsNonRoot quickly.
170. `kubectl get pod -A -o jsonpath='{range .items[*]}{.metadata.namespace} {.metadata.name} {.spec.containers[*].securityContext.allowPrivilegeEscalation}{"\n"}{end}'` - Check privilege escalation settings.
171. `kubectl get pod -A -o jsonpath='{range .items[*]}{.metadata.namespace} {.metadata.name} {.spec.containers[*].securityContext.readOnlyRootFilesystem}{"\n"}{end}'` - Check read-only root FS settings.
172. `kubectl get secrets -A` - Secret inventory.
173. `kubectl describe secret mysecret -n ns` - Secret metadata inspection.
174. `kubectl create secret generic app-secret --from-literal=key=value -n ns` - Create secret.
175. `kubectl get serviceaccount -A` - Service accounts inventory.
176. `kubectl describe sa default -n ns` - Service account details.
177. `kubectl get mutatingwebhookconfigurations` - Mutation policies/webhooks.
178. `kubectl get validatingwebhookconfigurations` - Validation policies/webhooks.
179. `kubectl api-resources | grep -i policy` - Policy-related APIs.
180. `kubectl label ns prod pod-security.kubernetes.io/enforce=restricted --overwrite` - Enforce restricted Pod Security admission (cluster support required).

## Section J: Vulnerability Scanning, SBOM, Signing, and CI (181-200)
181. `trivy image app:local` - Scan image vulnerabilities.
182. `trivy image --severity CRITICAL,HIGH app:local` - Filter by severity.
183. `trivy image --ignore-unfixed app:local` - Show fixable vulnerabilities focus.
184. `trivy fs .` - Scan filesystem dependencies and configs.
185. `trivy config .` - Scan IaC misconfigurations.
186. `trivy sbom --format cyclonedx --output sbom.cdx.json app:local` - Generate SBOM with Trivy.
187. `grype app:local` - Scan with Grype.
188. `grype sbom:sbom.cdx.json` - Scan from SBOM file.
189. `syft app:local -o cyclonedx-json > sbom.cdx.json` - Generate CycloneDX SBOM.
190. `syft app:local -o spdx-json > sbom.spdx.json` - Generate SPDX SBOM.
191. `cosign generate-key-pair` - Create signing keys (if keyless not used).
192. `cosign sign --key cosign.key registry/app:1.0.0` - Sign container image.
193. `cosign verify --key cosign.pub registry/app:1.0.0` - Verify image signature.
194. `cosign attest --predicate sbom.cdx.json --key cosign.key registry/app:1.0.0` - Attach attestation.
195. `cosign verify-attestation --key cosign.pub registry/app:1.0.0` - Verify attestations.
196. `jq '.Results[]?.Vulnerabilities[]?.Severity' trivy-report.json | sort | uniq -c` - Severity summary from JSON report.
197. `git diff -- Dockerfile` - Review Dockerfile changes before merge.
198. `git log --oneline -n 20` - Recent change context.
199. `gh workflow run container-security.yml` - Trigger GitHub workflow manually.
200. `gh run list --limit 10` - Check recent CI runs.

## Interview drill practice
- Pick any 10 commands and explain:
1. Why to run it.
2. What output you expect.
3. What action follows if output is abnormal.
