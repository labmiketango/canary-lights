# Harness Pipeline Status LED (Canary-Lights)

## Project Description

A GoLang application that drives a USB-connected LED to show the live status of multiple Harness.io pipelines at a glance.

The project is a demonstration of AI-driven software development and modern DevOps practices:

- **Spec-driven development:** the entire application was generated from a specification rather than written by hand.
- **Governed delivery:** compliance, governance, build, and deployment are all controlled end to end by Harness.io.

## Development Workflow

All work follows a spec-driven workflow, in this order:

**specify > plan > build > verify > review > release**

Do not skip or reorder phases.

---

## 1. Specify

- Requirements do **not** dictate every implementation detail.
- Requirements define the **behavior** that the implementation must demonstrate.
- Write specifications as observable, testable behavior, not as prescribed code structure.

## 2. Plan

Standing orders for planning:

1. Only modify code in this repository. Do not start to implement or fix anything in dependencies.
2. Record every change in the audit history.
3. Complete the integration and acceptance tests as part of the plan.

## 3. Build

1. Perform a local build based on the change.
2. Make create a feature branch for each spec
3. The full CI/CD procedure is executed by **Harness** and is triggered by a git event. Do not replicate or replace it locally.

## 4. Verify

All of the following must be performed:

1. Unit tests
2. Linting
3. Type checking
4. Dependency analysis

## 5. Review

1. Ensure that every verify step passes without issue.
2. Create documentation in `documentation/` containing:
   - A high-level release note
   - A summary of the implementation

## 6. Release

1. The release process is managed by **Harness** and is triggered by a git event.
2. Do not perform or simulate a release manually.
3. To release perform a commit & push. Add a summary of the changes in the commit message.

---

## General Standing Orders

1. **Do not assume.** If something needs clarification, ask questions before proceeding. This includes, but is not limited to, tooling choices (test framework, linter, type checker, dependency analysis tool), the audit history format and location, and any ambiguous requirement.
