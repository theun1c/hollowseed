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
- [x] Grid initializes with walls
- [x] Get/Set work with bounds checks
- [x] Render returns consistent ASCII output
- [x] Code compiles without unrelated changes

## Definition of Done
- [x] Implementation complete
- [x] Basic tests added (or manual check documented)
- [x] TASKS.md updated

