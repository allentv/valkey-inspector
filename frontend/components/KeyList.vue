<script setup lang="ts">
const databases = Array.from({ length: 16 }, (_, i) => i)

const { keys, loading, selectedDb, searchPattern, hasMore, selectedKey, fetchKeys } = useKeys()
const { connection, disconnect } = useConnection()
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
  <div class="flex flex-col h-full bg-white">
    <!-- Connection Header -->
    <div class="p-4 border-b border-gray-200 bg-white flex justify-between items-center">
      <div class="flex items-center gap-2 text-sm font-medium text-gray-700 truncate" :title="`${connection.host}:${connection.port}`">
        <div class="w-2 h-2 rounded-full bg-green-500 flex-shrink-0"></div>
        <span class="truncate">{{ connection.host }}:{{ connection.port }}</span>
      </div>
      <button class="text-gray-400 hover:text-red-600 p-1" title="Disconnect" @click="handleDisconnect">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-5 h-5">
          <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 9V5.25A2.25 2.25 0 0013.5 3h-6a2.25 2.25 0 00-2.25 2.25v13.5A2.25 2.25 0 007.5 21h6a2.25 2.25 0 002.25-2.25V15M12 9l-3 3m0 0l3 3m-3-3h12.75" />
        </svg>
      </button>
    </div>

    <!-- Search & Filter Header -->
    <div class="p-4 border-b border-gray-200 space-y-3 bg-gray-50">
      <div>
        <label class="block text-xs font-medium text-gray-500 mb-1">Search Pattern</label>
        <input
          v-model="searchPattern"
          type="text"
          class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          placeholder="e.g. user:*"
          @keyup.enter="refreshKeys"
        />
      </div>

      <div>
        <label class="block text-xs font-medium text-gray-500 mb-1">Database</label>
        <select
          v-model="selectedDb"
          class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white"
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
        <div v-for="i in 5" :key="i" class="animate-pulse flex justify-between items-center">
          <div class="h-4 bg-gray-200 rounded w-2/3"></div>
          <div class="h-4 bg-gray-200 rounded w-8"></div>
        </div>
      </div>

      <div v-else-if="keys.length === 0" class="p-4 text-center text-gray-500 text-sm">No keys found.</div>

      <template v-else>
        <KeyListItem 
          v-for="key in keys" 
          :key="key.name" 
          v-bind="key"
          :class="{ 'bg-blue-50': selectedKey === key.name }"
          @click="selectedKey = key.name"
        />

        <div v-if="hasMore" class="p-4 text-center">
          <button class="text-xs text-blue-600 hover:text-blue-800 font-medium disabled:opacity-50" :disabled="loading" @click="loadMore">
            {{ loading ? 'Loading...' : 'Load More...' }}
          </button>
        </div>
      </template>
    </div>
  </div>
</template>
