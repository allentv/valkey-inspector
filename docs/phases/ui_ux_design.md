# UI/UX Design Concepts for Valkey Inspector

## 1. The "Master-Detail" Layout
Since we are inspecting a database, a **Split-Pane** layout (like an email client or VS Code) works best.

*   **Sidebar (Left - 30% width):** The "Key Browser".
*   **Main Pane (Right - 70% width):** The "Data Inspector".

## 2. The Key Browser (Left Panel)
This is where the pagination logic lives. Since Valkey uses `SCAN` (cursor-based iteration) rather than offset-based pagination (like SQL `LIMIT/OFFSET`), standard numbered pages (1, 2, 3...) don't work well.

*   **Search Bar:**
    *   Input for `MATCH` patterns (e.g., `user:*`, `session:*`).
    *   Dropdown for **Database Selection** (DB 0-15).
*   **The List:**
    *   Display keys in a compact list.
    *   **Type Badges:** Color-coded pills next to the key name (e.g., <span style="color:blue">STRING</span>, <span style="color:green">HASH</span>, <span style="color:orange">LIST</span>).
    *   **TTL Indicator:** A small dot or icon indicating if the key is volatile (will expire) or persistent.
*   **Pagination (The "Load More" Approach):**
    *   Instead of "Page 1, 2, 3", use an **"Infinite Scroll"** or a **"Load More"** button at the bottom of the list.
    *   **Backend Logic:** The frontend sends the current `cursor` to the backend. The backend runs `SCAN cursor MATCH pattern COUNT 100`. It returns the keys *and* the `next_cursor`.
    *   **Frontend Logic:** Append the new keys to the existing array. If `next_cursor` is `0`, you've reached the end.

## 3. The Data Inspector (Right Panel)
This view changes dynamically based on the **Type** of the selected key.

*   **Header:**
    *   Key Name (Copy to clipboard button).
    *   TTL Countdown (e.g., "Expires in 300s").
    *   Memory Usage (using `MEMORY USAGE key`).
    *   **Delete Button:** With a confirmation modal.

*   **Type-Specific Views:**
    *   **String:**
        *   If the value looks like JSON, auto-format and syntax highlight it (using a library like `vue-json-pretty`).
        *   If binary, show a Hex dump or a "Binary Data" placeholder.
    *   **List / Set / ZSet:**
        *   Don't load all 1 million items of a list!
        *   Use internal pagination (`LRANGE`, `SSCAN`, `ZSCAN`).
        *   Display in a virtualized table (using `vue-virtual-scroller` if the list is huge) to keep the DOM light.
    *   **Hash:**
        *   A clean Key-Value table.
        *   Search bar specifically for Hash fields (`HSCAN`).

## 4. Cluster vs. Single Node
If you are connecting to a Cluster:
*   **Node Switcher:** A dropdown in the top header showing which specific node (Master/Replica) you are querying, or an "Auto" mode that routes requests intelligently.
*   **Slot Info:** When viewing a key, display which **Hash Slot** it belongs to and which Node is the owner.

## 5. Tailwind Implementation Tips
Structure the Key List item for readability:

```html
<!-- Example Key List Item -->
<div class="flex items-center justify-between p-3 border-b border-gray-200 hover:bg-gray-50 cursor-pointer transition-colors">
  <div class="flex items-center gap-2 overflow-hidden">
    <!-- Type Badge -->
    <span class="px-2 py-0.5 text-xs font-bold rounded bg-blue-100 text-blue-800">STR</span>
    <!-- Key Name -->
    <span class="truncate font-mono text-sm text-gray-700" title="user:1001:session">user:1001:session</span>
  </div>
  <!-- TTL Indicator (Green = Safe, Red = Expiring soon) -->
  <div class="w-2 h-2 rounded-full bg-green-500"></div>
</div>
```

## 6. Performance "Gotchas" to Avoid
*   **Never use `KEYS *`:** Ensure your backend strictly uses `SCAN`.
*   **Large Values:** If a String value is 50MB, don't try to render it immediately. Fetch the length first (`STRLEN`), warn the user, and offer a "Download" button or "Preview first 1KB" option.