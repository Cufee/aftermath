# Documentation

This documentation is optimized for AI context when working on features.

## Structure

```
docs/
├── features/     # Context about existing features
├── adr/          # Architecture Decision Records
├── plans/        # Work plans (active and completed)
└── work/         # Work item tracking by feature
```

## Quick Links

### Features
Feature documentation describes how existing functionality works, key files, patterns, and gotchas.

- [Template](features/_template.md)

### ADRs
Architecture Decision Records track significant technical decisions with context and rationale.

- [Template](adr/_template.md)

### Plans
Work plans document goals, approach, and progress for features being implemented.

- [Template](plans/_template.md)
- [Active Plans](plans/active/)
- [Completed Plans](plans/completed/)

### Work Items
Track individual work items grouped by feature area.

- [Template](work/_template.md)

## Naming Conventions

| Type | Format | Example |
|------|--------|---------|
| Features | `kebab-case.md` | `discord-commands.md` |
| ADRs | `XXXX-kebab-case.md` | `0001-templ-for-frontend.md` |
| Plans | `kebab-case.md` | `frontend-v3.md` |
| Work Items | `feature/item.md` | `discord/add-slash-command.md` |

## Adding Documentation

### New Feature Doc
1. Copy `features/_template.md`
2. Fill in sections with key files, patterns, gotchas
3. Add link to this README

### New ADR
1. Copy `adr/_template.md`
2. Use next sequential number (check existing ADRs)
3. Fill in context, decision, consequences

### New Plan
1. Copy `plans/_template.md` to `plans/active/`
2. Update Progress Log as work proceeds
3. Move to `plans/completed/` when done

### New Work Item
1. Create feature subdirectory in `work/` if needed
2. Copy `work/_template.md`
3. Link to related feature doc and plan

## Writing Guidelines

For AI context optimization:
- Include exact file paths (e.g., `cmd/discord/commands/handler.go:42`)
- Use code blocks for function names and patterns
- Keep sections scannable with bullet points
- Highlight non-obvious behavior prominently
- Cross-reference related documents
