<template>
  <div class="flex h-screen w-full bg-white text-gray-900">
    <!-- Left Pane: Key Browser (30%) -->
    <aside class="w-1/3 border-r border-gray-200 h-full flex flex-col">
      <div class="flex-1 overflow-y-auto">
        <!-- This slot allows injecting the sidebar content, or defaults to a placeholder -->
        <slot name="sidebar">
          <KeyList />
        </slot>
      </div>
    </aside>

    <!-- Right Pane: Data Inspector (70%) -->
    <main class="w-2/3 h-full flex flex-col overflow-y-auto">
      <slot />
    </main>
  </div>
</template>
