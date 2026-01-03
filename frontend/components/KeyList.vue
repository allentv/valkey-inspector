<script setup lang="ts">
import { ref } from 'vue'

const searchPattern = ref('*')
const selectedDb = ref(0)
const databases = Array.from({ length: 16 }, (_, i) => i)

// Placeholder data until API integration
const keys = ref([
  { name: 'user:1001:session', type: 'string', ttl: 300 },
  { name: 'user:1001:cart', type: 'hash', ttl: -1 },
  { name: 'system:logs', type: 'list', ttl: -1 },
  { name: 'leaderboard:2023', type: 'zset', ttl: 3600 },
])

const loadMore = () => {
  console.log('Load more keys...')
}

const refreshKeys = () => {
  console.log(`Refreshing keys for DB ${selectedDb.value} with pattern ${searchPattern.value}`)
}
</script>

<template>
  <div class="flex flex-col h-full bg-white">
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
      <div v-if="keys.length === 0" class="p-4 text-center text-gray-500 text-sm">No keys found.</div>

      <KeyListItem v-for="key in keys" :key="key.name" :name="key.name" :type="key.type" :ttl="key.ttl" />

      <div class="p-4 text-center">
        <button class="text-xs text-blue-600 hover:text-blue-800 font-medium" @click="loadMore">Load More...</button>
      </div>
    </div>
  </div>
</template>
