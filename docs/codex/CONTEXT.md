# Hollowseed Context

## Project
- Name: Hollowseed
- Goal: learn Go + procedural generation
- Type: ASCII roguelike-style dungeon generator

## Current State
- Project scaffold exists
- No generation logic yet

## Architecture (current intent)
- cmd/hollowseed/main.go as entrypoint
- internal/dungeon for generation logic
- Keep package structure simple, avoid premature abstractions

## Technical Constraints
- Go 1.x
- ASCII output first

## Domain Terms
- Grid: 2D tile map
- Room: rectangular area of floor
- Corridor: connection path between rooms
- BSP node: partition rectangle in split tree

## Non-Goals (for now)
- No UI
- No enemies/loot systems
- No complex biome rules

## Working Agreement for Feature Chats
- One chat = one feature
- Agent must read this file + target feature SPEC
- Out-of-scope changes are not allowed

## Engineering Workflow
- TDD-first for new features: RED -> GREEN -> REFACTOR
- Start from tests, then minimal implementation
- Keep test scope aligned with current feature SPEC
