# CI security incident — 2026-10-09

Repository: `ThreeAndTwo/comsclient`
Branch: `main`
Head inspected before this change: `806c4ffec76828430e0d3b385f601298e3e02494`

Unauthorized GitHub Actions workflows were identified during an owner-requested review.
The reviewed workflows attempted to send repository/history data or GitHub Actions secrets to an unapproved external endpoint.
Do not restore or execute these workflows. A successful Actions run alone does not establish which credentials were received or remained valid.

This commit removes the following verified malicious files:

- `.github/workflows/security-audit.yml` — Git blob `a5a0cbcf7f5aa30b5ac499d441e7ecc089fd1ea5`.

Review and containment notes

- Keep legitimate build/test/release workflows; suspicious names alone are not evidence.
- Revoke or rotate credentials that may have been exposed; deleting a workflow or a GitHub Secret does not revoke credentials at their issuing service.
- Historical commits are retained for evidence. This change does not rewrite Git history.
- The initial credential compromise remains under investigation; this record does not attribute it to a person or tool.
