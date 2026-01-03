<script setup lang="ts">
const databases = Array.from({ length: 16 }, (_, i) => i)

const { keys, loading, selectedDb, searchPattern, hasMore, selectedKey, fetchKeys } = useKeys()
const { connection, disconnect } = useConnection()
const { isDark, toggleTheme } = useTheme()
const router = useRouter()

const handleDisconnect = () => {
  disconnect()
  router.push('/connect')
}

const loadMore = () => {
  fetchKeys(false)
}

const refreshKeys = () => {
  fetchKeys(true)
}

// Initial load
onMounted(() => {
  if (keys.value.length === 0) {
    refreshKeys()
  }
})
</script>

<template>
  <div class="flex flex-col h-full bg-white dark:bg-gray-800">
    <!-- Connection Header -->
    <div
      class="p-4 border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 flex justify-between items-center"
    >
      <div
        class="flex items-center gap-2 text-sm font-medium text-gray-700 dark:text-gray-200 truncate"
        :title="`${connection.host}:${connection.port}`"
      >
        <div class="w-2 h-2 rounded-full bg-green-500 flex-shrink-0"></div>
        <span class="truncate">{{ connection.host }}:{{ connection.port }}</span>
      </div>
      <div class="flex items-center gap-1">
        <button
          class="text-gray-400 hover:text-gray-600 dark:hover:text-gray-200 p-2 rounded-md hover:bg-gray-100 dark:hover:bg-gray-700 transition-colors"
          :title="isDark ? 'Switch to Light Mode' : 'Switch to Dark Mode'"
          @click="toggleTheme"
        >
          <svg
            v-if="isDark"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="w-5 h-5"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M12 3v2.25m6.364.386l-1.591 1.591M21 12h-2.25m-.386 6.364l-1.591-1.591M12 18.75V21m-4.773-4.227l-1.591 1.591M5.25 12H3m4.227-4.773L5.636 5.636M15.75 12a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0z"
            />
          </svg>
          <svg
            v-else
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="w-5 h-5"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M21.752 15.002A9.718 9.718 0 0118 15.75c-5.385 0-9.75-4.365-9.75-9.75 0-1.33.266-2.597.748-3.752A9.753 9.753 0 003 11.25C3 16.635 7.365 21 12.75 21a9.753 9.753 0 009.002-5.998z"
            />
          </svg>
        </button>
        <button
          class="text-gray-400 hover:text-red-600 p-2 rounded-md hover:bg-gray-100 dark:hover:bg-gray-700 transition-colors"
          title="Disconnect"
          @click="handleDisconnect"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="w-5 h-5"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M15.75 9V5.25A2.25 2.25 0 0013.5 3h-6a2.25 2.25 0 00-2.25 2.25v13.5A2.25 2.25 0 007.5 21h6a2.25 2.25 0 002.25-2.25V15M12 9l-3 3m0 0l3 3m-3-3h12.75"
            />
          </svg>
        </button>
      </div>
    </div>

    <!-- Search & Filter Header -->
    <div class="p-4 border-b border-gray-200 dark:border-gray-700 space-y-3 bg-gray-50 dark:bg-gray-900">
      <div>
        <label class="block text-xs font-medium text-gray-500 dark:text-gray-400 mb-1">Search Pattern</label>
        <input
          v-model="searchPattern"
          type="text"
          class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white dark:bg-gray-800 text-gray-900 dark:text-white"
          placeholder="e.g. user:*"
          @keyup.enter="refreshKeys"
        />
      </div>

      <div>
        <label class="block text-xs font-medium text-gray-500 dark:text-gray-400 mb-1">Database</label>
        <select
          v-model="selectedDb"
          class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white dark:bg-gray-800 text-gray-900 dark:text-white"
          @change="refreshKeys"
        >
          <option v-for="db in databases" :key="db" :value="db">DB {{ db }}</option>
        </select>
      </div>
    </div>

    <!-- Key List -->
    <div class="flex-1 overflow-y-auto">
      <!-- Skeleton Loader -->
      <div v-if="loading && keys.length === 0" class="p-4 space-y-3">
        <div v-for="i in 5" :key="i" class="animate-pulse flex justify-between items-center dark:opacity-50">
          <div class="h-4 bg-gray-200 dark:bg-gray-700 rounded w-2/3"></div>
          <div class="h-4 bg-gray-200 dark:bg-gray-700 rounded w-8"></div>
        </div>
      </div>

      <div v-else-if="keys.length === 0" class="p-4 text-center text-gray-500 dark:text-gray-400 text-sm">
        No keys found.
      </div>

      <template v-else>
        <KeyListItem
          v-for="key in keys"
          :key="key.name"
          v-bind="key"
          :class="{ 'bg-blue-50 dark:bg-blue-900/30': selectedKey === key.name }"
          @click="selectedKey = key.name"
        />

        <div v-if="hasMore" class="p-4 text-center">
          <button
            class="text-xs text-blue-600 dark:text-blue-400 hover:text-blue-800 dark:hover:text-blue-300 font-medium disabled:opacity-50"
            :disabled="loading"
            @click="loadMore"
          >
            {{ loading ? 'Loading...' : 'Load More...' }}
          </button>
        </div>
      </template>
    </div>
  </div>
</template>
