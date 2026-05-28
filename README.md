# OS Security Engineer Interview Prep Workspace

This workspace contains a complete prep path for a Senior OS Security Engineer interview focused on CVE remediation, Linux-based container hardening, Kubernetes security, and supply chain controls.

## Table Of Contents
- [Start Here](#start-here)
- [What’s Inside](#whats-inside)
- [Recommended Order](#recommended-order)
- [Study Path](#study-path)
- [7-Day Start Path](#7-day-start-path)
- [Interview Day Checklist](#interview-day-checklist)
- [Main Goal](#main-goal)

## Start Here
1. Read [JD.txt](JD.txt) to understand the role expectations.
2. Read [secreening_questions.txt](secreening_questions.txt) to see the exact screening prompts.
3. Use [INDEX.md](INDEX.md) as the main workspace navigator.
4. Follow the short path in [interview-prep/00_study_plan.md](interview-prep/00_study_plan.md).
5. Use the full course in [interview-course/00_course_index.md](interview-course/00_course_index.md).

## What’s Inside
- Interview expectations and prep notes in [interview-prep](interview-prep)
- A full beginner-to-interview course in [interview-course](interview-course)
- Hands-on sample projects in [sample-projects](sample-projects)
- A top-level navigation file in [INDEX.md](INDEX.md)

## Recommended Order
- First: understand the role and questions.
- Second: study the course modules in order.
- Third: do the sample projects and labs.
- Fourth: practice the 100-question bank and mock interview simulator.

## Study Path
```text
JD + Screening Questions
	|
	v
Workspace INDEX.md
	|
	v
Quick Study Plan -> Full Course Index -> Sample Projects
	|
	v
Labs + Command Handbook + Mock Interview Simulator
	|
	v
Final Week Rapid Revision Pack
```

## 7-Day Start Path
### Day 1
- Read [JD.txt](JD.txt)
- Read [secreening_questions.txt](secreening_questions.txt)
- Open [INDEX.md](INDEX.md)

### Day 2
- Study [interview-course/01_learning_method.md](interview-course/01_learning_method.md)
- Study [interview-course/02_linux_fundamentals.md](interview-course/02_linux_fundamentals.md)

### Day 3
- Study [interview-course/03_networking_for_security.md](interview-course/03_networking_for_security.md)
- Study [interview-course/04_containers_from_zero.md](interview-course/04_containers_from_zero.md)

### Day 4
- Study [interview-course/05_dockerfile_mastery.md](interview-course/05_dockerfile_mastery.md)
- Study [interview-course/06_package_managers_and_os_layers.md](interview-course/06_package_managers_and_os_layers.md)

### Day 5
- Study [interview-course/07_cve_management_end_to_end.md](interview-course/07_cve_management_end_to_end.md)
- Study [interview-course/08_image_hardening_playbook.md](interview-course/08_image_hardening_playbook.md)

### Day 6
- Study [interview-course/09_scanning_tools_trivy_grype.md](interview-course/09_scanning_tools_trivy_grype.md)
- Study [interview-course/10_supply_chain_sbom_signing_slsa.md](interview-course/10_supply_chain_sbom_signing_slsa.md)
- Study [interview-course/11_cicd_security_automation.md](interview-course/11_cicd_security_automation.md)

### Day 7
- Study [interview-course/12_kubernetes_runtime_security_basics.md](interview-course/12_kubernetes_runtime_security_basics.md)
- Do one sample project from [sample-projects](sample-projects)
- Start [interview-course/14_interview_masterclass.md](interview-course/14_interview_masterclass.md)

## Interview Day Checklist
- Review your 5 strongest STAR stories.
- Re-read the CVE remediation and hardening modules.
- Practice 10 command recalls from the handbook.
- Review one system design scenario and one troubleshooting scenario.
- Be ready to explain why HIGH CVEs may remain after patching.
- Have one clear answer for how you reduce attack surface.
- Know your metrics: CVE reduction, image size reduction, MTTR, and rollout stability.

## Main Goal
By the end of this workspace, you should be able to:
- Explain Linux and container security from first principles.
- Remediate and validate CVEs in container images.
- Hardening images and Kubernetes runtime settings.
- Discuss SBOM, signing, provenance, CI/CD gates, and incident response confidently.
