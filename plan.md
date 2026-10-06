# Standalone SMTP MCP release

## Context and discovery

Extract the reviewed SMTP submission adapter from smtpd PR 7 into the public
`github.com/sirerun/smtp-mcp` repository, which initially contains an Apache 2.0
license. Three supported use cases are preview, SMTP submission, and configuration
status. The adapter already has protocol and TLS tests. No running SMTP daemon,
mail campaign, credentials, or domain-specific settings belong in this release.

## Scope and acceptance

Ship a standalone Go module, `smtp-mcp` executable, useful README and client
configuration, CI and reproducible release packaging. Publish GitHub description
and v0.1.0 archives for Linux/macOS/Windows on amd64/arm64, with checksums.
Stdio is supported; authenticated HTTP remains future work.

## Work breakdown

- [x] T1 Freeze extraction contract. Owner: coordinator. Est: 10m. verifies: [infrastructure]. acc: [module, names, license, transport and release scope recorded].
- [ ] T2 Extract code and tests. Owner: implementation lane. Est: 20m. deps: [T1]. verifies: [preview, send, status]. acc: [standalone module builds and all adapter tests pass].
- [ ] T3 Write README, examples and release automation. Owner: coordinator. Est: 20m. deps: [T1]. verifies: [installation, configuration]. acc: [examples match flags and release archives contain binary and license].
- [ ] T4 Verify and independently review exact source. Owner: coordinator and independent reviewer. Est: 20m. deps: [T2, T3]. verifies: [infrastructure]. acc: [test, race, vet, lint, vulnerability and stdio smoke pass with no unresolved review findings].
- [ ] T5 Merge and verify landed source. Owner: coordinator. Est: 10m. deps: [T4]. verifies: [infrastructure]. acc: [remote main matches reviewed content and landed checks pass].
- [ ] T6 Publish v0.1.0 and verify assets. Owner: coordinator. Est: 15m. deps: [T5]. verifies: [installation]. acc: [public release tag targets landed source and downloaded checksums pass].
- [ ] T7 Record relocation in original smtpd PR. Owner: coordinator. Est: 5m. deps: [T6]. verifies: [infrastructure]. acc: [original draft closed as superseded and existing main unchanged].

## Parallel work and milestones

T2 and T3 proceed in isolated ownership lanes. Review starts on the combined exact
head. Verification uses an owned build lease and a one-minute load <=10 on the
shared Mac; existing qualified remote capacity may be used within its reservation.
No paid provisioning is included. Milestones are standalone source, verified and
reviewed merge, then published and downloaded release.

## Risks and operating procedure

- Shared machine load may defer builds; do not bypass the lease or load limit.
- Hosted CI may be billing-blocked. Use recorded local results without presenting
  them as hosted CI success or weakening branch protections.
- Publishing credentials is forbidden. Only synthetic example addresses belong
  in source and documentation; private operational configuration stays outside.
- SMTP acceptance is not inbox delivery. The package does not promise campaign
  scheduling, durable queues, deduplication, or sender reputation.
- Preserve the source daemon's main branch and existing Apache 2.0 license.

Definition of done: exact-head checks and independent review, guarded rebase
merge, landed verification, release publication and downloaded asset verification.

## Progress log

- 2026-10-05: Confirmed the destination is public and contains only its initial
  license commit. Source extraction head is smtpd `cb489aa`. Assigned isolated
  implementation and independent verification-capacity lanes. No external mail.

## Handoff and references

See README for the public interface and docs/release.md for release procedure.
The project-root ajent.social coordination file is ignored in this public repo.
Original source: https://github.com/sirerun/smtpd/pull/7.
