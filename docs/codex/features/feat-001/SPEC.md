# Feature 001: Grid Foundation

## Summary (RU)
Сделать минимальную основу карты: размеры, хранение тайлов, базовый ASCII-рендер.

## Goal
Create a minimal grid model to support dungeon generation steps.

## Scope
- Define tile type (wall/floor)
- Define grid with width/height/cells
- Provide safe get/set by coordinates
- Provide ASCII render function

## Out of Scope
- Room placement
- Corridors
- BSP splitting
- Seed/random generator

## Acceptance Criteria
- [ ] Grid initializes with walls
- [ ] Get/Set work with bounds checks
- [ ] Render returns consistent ASCII output
- [ ] Code compiles without unrelated changes

## Definition of Done
- [ ] Implementation complete
- [ ] Basic tests added (or manual check documented)
- [ ] TASKS.md updated

