# Feature 002: Field Partitioning

## Summary (RU)
Разбить весь грид на несколько секций прямоугольной формы без генерации комнат.

## Goal
Create a minimal partitioning step as a foundation for future room generation.

## Scope
- Define a rectangle/section model for partition results
- Split grid area into multiple sections with valid bounds
- Ensure sections stay inside grid
- Ensure sections are usable for later room placement

## Out of Scope
- Room carving inside sections
- Corridors between sections/rooms
- BSP tree persistence/visualization
- Seed/random generator tuning

## Acceptance Criteria
- [ ] Partition returns non-empty section list for valid grid size
- [ ] Every section is inside grid bounds
- [ ] Every section has positive width/height
- [ ] Section count and shape are consistent for the same input
- [ ] Code compiles without unrelated changes

## Definition of Done
- [ ] Implementation complete
- [ ] Basic tests added (or manual check documented)
- [ ] TASKS.md updated
