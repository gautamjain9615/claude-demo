You are an autonomous developer agent. Implement the Jira ticket provided as the argument below, following the full development lifecycle from reading the ticket to creating a PR and monitoring CI.

**Ticket identifier**: $ARGUMENTS

---

## Step 1: Read the Jira Ticket

Use the Atlassian MCP to read the Jira ticket `$ARGUMENTS`:
- Extract: summary, description, acceptance criteria, story points
- Identify: what type of change this is (feature, bug, refactor)

## Step 2: Plan the Implementation

- Read CLAUDE.md for project conventions
- Explore the codebase to understand existing patterns (directory structure, handler style, store patterns, test style)
- Design the implementation approach
- List files you will create or modify

## Step 3: Implement the Changes

- Create a feature branch: `feature/$ARGUMENTS-<short-description>` (e.g., `feature/BOOK-1-add-reviews`)
- Write the code following existing patterns exactly:
  - Models in `internal/models/`
  - Store methods in `internal/store/memory.go`
  - Handlers in `internal/handlers/`
  - Wire new routes in `cmd/server/main.go`
- Add swaggo annotations for any new or modified endpoints
- Write table-driven tests for all new code
- Run tests locally: `go test ./...`
- Run linter locally: `golangci-lint run ./...`
- Fix any issues found by tests or linter before proceeding

## Step 4: Create the PR

- Stage and commit all changes with message: `feat(api): <description> [$ARGUMENTS]`
- Push the branch to remote
- Create a PR using `gh pr create` with:
  - Title referencing the Jira ticket (e.g., "feat(api): Add book reviews endpoint [BOOK-1]")
  - Body containing:
    - Summary of changes
    - Link to Jira ticket
    - List of files changed
    - Test plan
  - The PR should target `main`

## Step 5: Monitor CI

- Wait for GitHub Actions workflows to start, then monitor with `gh pr checks <pr-number> --watch`
- If any check fails:
  1. Read the workflow logs: `gh run view <run-id> --log-failed`
  2. Identify the failure (likely a lint issue like `nlreturn`, `godot`, or `wrapcheck`)
  3. Fix the issue in the code
  4. Run the linter locally to verify the fix: `golangci-lint run ./...`
  5. Commit the fix: `fix: resolve lint issues [$ARGUMENTS]`
  6. Push and wait for CI again
- Repeat until all checks pass

## Step 6: Report Completion

Summarize what was done:
- What the ticket asked for
- What was implemented
- Files created/modified
- Test results
- PR link
- CI status
- Note: "Ready for review and merge"
