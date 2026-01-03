<template>
  <div class="flex h-screen w-full bg-white text-gray-900 dark:bg-gray-900 dark:text-gray-100">
    <!-- Left Pane: Key Browser (30%) -->
    <aside class="w-1/3 border-r border-gray-200 dark:border-gray-700 h-full flex flex-col bg-white dark:bg-gray-800">
      <div class="flex-1 overflow-y-auto">
        <!-- This slot allows injecting the sidebar content, or defaults to a placeholder -->
        <slot name="sidebar">
          <KeyList />
        </slot>
      </div>
    </aside>

    <!-- Right Pane: Data Inspector (70%) -->
    <main class="w-2/3 h-full flex flex-col overflow-y-auto bg-gray-50 dark:bg-gray-900">
      <slot />
    </main>
  </div>
</template>
