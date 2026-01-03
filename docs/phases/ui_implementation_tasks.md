# UI Implementation Tasks

Based on the [UI/UX Design Concepts](./ui_ux_design.md), this document outlines the specific development tasks required to implement the interface using Nuxt and Tailwind CSS.

## 1. Layout & Structure
- [x] **Create Main Layout**: Implement a split-pane layout in `layouts/default.vue` or a new `layouts/inspector.vue`.
    - Left Pane: 30% width, fixed/scrollable (Key Browser).
    - Right Pane: 70% width, scrollable (Data Inspector).
    - Use Tailwind classes: `flex h-screen`, `w-1/3`, `w-2/3`, `border-r`.

## 2. Key Browser (Left Panel)
- [x] **Search & Filter Component**:
    - Add an input field for `MATCH` patterns.
    - Add a dropdown for Database selection (0-15).
    - Bind these to a reactive state to trigger list refreshes.
- [x] **Key List Item Component**:
    - Create `components/KeyListItem.vue`.
    - Implement Tailwind styling for the "compact list" look (hover effects, cursor pointer).
    - Add dynamic badges for types (String, Hash, List, Set, ZSet) using specific color classes.
    - Add a visual indicator for TTL (e.g., Green dot for persistent, Yellow for volatile).
- [ ] **Infinite Scroll / Pagination**:
    - Update `components/KeyList.vue` to handle `cursor`-based pagination.
    - Implement a "Load More" button or use an Intersection Observer for infinite scroll.
    - **Backend Update**: Ensure `GET /api/keys` accepts `cursor` and returns `next_cursor`.

## 3. Data Inspector (Right Panel)
- [x] **Inspector Container**:
    - Create `components/KeyDetail.vue`.
    - Implement the Header section:
        - Key Name (large font, monospace, truncate).
        - Copy to Clipboard button.
        - TTL Countdown timer (auto-refreshing).
        - Delete button (red, right-aligned) with confirmation modal.
- [x] **Type-Specific Viewers**:
    - **String**: `components/viewers/StringViewer.vue`
        - Detect JSON content and use a syntax highlighter or `vue-json-pretty`.
        - Fallback to a read-only text area for raw strings.
    - **Hash**: `components/viewers/HashViewer.vue`
        - Render a table with Key/Value columns.
        - Add a local search input for hash fields.
    - **List/Set/ZSet**: `components/viewers/CollectionViewer.vue`
        - Implement a virtualized list or internal pagination for large collections.

## 4. State Management & Polish
- [ ] **Composables**:
    - Create `composables/useKeys.ts` to manage the list state, cursor, and loading status.
- [ ] **Loading States**: Add skeleton loaders (pulsing gray bars) for the Key List and Detail view while fetching data.
- [x] **Empty States**: Design a "Select a key to view details" placeholder for the right pane.
