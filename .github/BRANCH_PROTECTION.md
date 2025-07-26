# Branch Protection Configuration

This document describes the branch protection settings for this repository using GitHub Rulesets.

## Overview

- **Protected Branch**: `main`
- **Enforcement Status**: Active
- **Bypass Permissions**: None (administrators included)

## Ruleset Configuration

### Pull Request Requirements

- ✅ **Require a pull request before merging**
  - Required approving reviews: **1**
  - ✅ Dismiss stale pull request approvals when new commits are pushed
  - ✅ Require conversation resolution before merging
  - ❌ Require approval of the most recent push
  - ❌ Require review from CODEOWNERS

### Status Checks

- ✅ **Require status checks to pass before merging**
  - ✅ Require branches to be up to date before merging (Strict mode)
  - Status checks to be added when CI is configured:
    - `build`
    - `test`
    - `lint`

### Branch Protection

- ✅ **Restrict deletions** - Prevent branch deletion
- ✅ **Block force pushes** - Prevent history rewriting

### Not Enabled

- ❌ Require linear history
- ❌ Require signed commits
- ❌ Restrict creations
- ❌ Restrict updates

## Setup Instructions

### Using GitHub Web UI

1. Go to **Settings** → **Rules** → **Rulesets**
2. Click **New ruleset** → **New branch ruleset**
3. Configure:
   - **Ruleset Name**: `Protect main branch`
   - **Enforcement status**: Active
   - **Bypass list**: Leave empty (no bypasses)
   - **Target branches**: Add `main`
4. Enable rules as specified above
5. Click **Create**

### Using GitHub CLI (API)

```bash
# Create ruleset using gh api
gh api repos/{owner}/{repo}/rulesets \
  --method POST \
  --field name='Protect main branch' \
  --field target='branch' \
  --field enforcement='active' \
  --field conditions='{"ref_name":{"include":["refs/heads/main"],"exclude":[]}}' \
  --field bypass_actors='[]' \
  --field rules='[
    {
      "type": "pull_request",
      "parameters": {
        "required_approving_review_count": 1,
        "dismiss_stale_reviews_on_push": true,
        "require_code_owner_review": false,
        "require_last_push_approval": false,
        "required_review_thread_resolution": true
      }
    },
    {
      "type": "required_status_checks",
      "parameters": {
        "strict_required_status_checks_policy": true,
        "required_status_checks": []
      }
    },
    {
      "type": "deletion"
    },
    {
      "type": "non_fast_forward"
    }
  ]'
```

## Rationale

This configuration provides:

1. **Code Review**: All changes must be reviewed before merging
2. **Up-to-date Checks**: Ensures branches are tested with latest main
3. **Conversation Tracking**: All feedback must be addressed
4. **History Protection**: No force pushes or branch deletion
5. **No Bypasses**: Even administrators follow the rules

## Adding CI Checks

When CI is configured, update the ruleset to require specific status checks:

1. Go to **Settings** → **Rules** → **Rulesets**
2. Click on `Protect main branch`
3. Under **Require status checks**, add:
   - Your CI workflow names (e.g., `build`, `test`, `lint`)
4. Save changes

## Emergency Procedures

If you need to bypass protection in an emergency:

1. Temporarily disable the ruleset (Settings → Rules → Rulesets → Edit → Disabled)
2. Make necessary changes
3. Re-enable the ruleset immediately after

⚠️ **Always document why protection was bypassed**