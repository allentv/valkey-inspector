<script setup lang="ts">
import StringViewer from './viewers/StringViewer.vue'
import HashViewer from './viewers/HashViewer.vue'
import CollectionViewer from './viewers/CollectionViewer.vue'
import { ref } from 'vue'

const props = defineProps<{
  keyName?: string
}>()

// Mock data state (to be replaced by API call in Task #4)
// This allows us to preview the UI without the backend ready yet
const keyData = ref({
  type: 'string',
  ttl: 300,
  value: '{"user_id": 1001, "name": "Alice", "roles": ["admin", "editor"], "last_login": "2023-10-27T10:00:00Z"}',
  // Uncomment to test other types:
  // type: 'hash', value: { "username": "alice", "email": "alice@example.com", "credits": "50" },
  // type: 'list', value: ["Log entry 1", "Log entry 2", "Log entry 3"],
})

const copyToClipboard = () => {
  if (props.keyName) {
    navigator.clipboard.writeText(props.keyName)
    // Ideally trigger a toast notification here
  }
}

const deleteKey = () => {
  if (confirm(`Are you sure you want to delete ${props.keyName}?`)) {
    console.log('Deleting key...')
  }
}
</script>

<template>
  <div v-if="!keyName" class="flex items-center justify-center h-full text-gray-400 bg-gray-50">
    <div class="text-center">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-16 h-16 mx-auto mb-4 text-gray-300"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M20.25 6.375c0 2.278-3.694 4.125-8.25 4.125S3.75 8.653 3.75 6.375m16.5 0c0-2.278-3.694-4.125-8.25-4.125S3.75 4.097 3.75 6.375m16.5 0v11.25c0 2.278-3.694 4.125-8.25 4.125s-8.25-1.847-8.25-4.125V6.375m16.5 0v3.75m-16.5-3.75v3.75m16.5 0v3.75C20.25 16.153 16.556 18 12 18s-8.25-1.847-8.25-4.125v-3.75m16.5 0c0 2.278-3.694 4.125-8.25 4.125s-8.25-1.847-8.25-4.125"
        />
      </svg>
      <p class="text-lg font-medium">Select a key to view details</p>
    </div>
  </div>

  <div v-else class="flex flex-col h-full bg-gray-50">
    <!-- Header -->
    <div class="p-6 border-b border-gray-200 bg-white flex justify-between items-start shadow-sm z-10">
      <div class="overflow-hidden mr-4 flex-1">
        <div class="flex items-center gap-2 mb-2">
          <h2 class="text-xl font-mono font-bold text-gray-900 truncate" :title="keyName">
            {{ keyName }}
          </h2>
          <button
            class="text-gray-400 hover:text-blue-600 transition-colors p-1 rounded hover:bg-blue-50"
            title="Copy Key Name"
            @click="copyToClipboard"
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
                d="M15.666 3.888A2.25 2.25 0 0013.5 2.25h-3c-1.03 0-1.9.693-2.166 1.638m7.332 0c.055.194.084.4.084.612v0a.75.75 0 01-.75.75H9a.75.75 0 01-.75-.75v0c0-.212.03-.418.084-.612m7.332 0c.646.049 1.288.11 1.927.184 1.1.128 1.907 1.077 1.907 2.185V19.5a2.25 2.25 0 01-2.25 2.25H6.75A2.25 2.25 0 014.5 19.5V6.257c0-1.108.806-2.057 1.907-2.185a48.208 48.208 0 011.927-.184"
              />
            </svg>
          </button>
        </div>

        <div class="flex items-center gap-4 text-sm text-gray-500">
          <span class="flex items-center gap-1 bg-gray-100 px-2 py-0.5 rounded">
            <span class="font-medium text-gray-700">Type:</span> {{ keyData.type.toUpperCase() }}
          </span>
          <span class="flex items-center gap-1 bg-gray-100 px-2 py-0.5 rounded">
            <span class="font-medium text-gray-700">TTL:</span> {{ keyData.ttl }}s
          </span>
        </div>
      </div>

      <button
        class="px-3 py-1.5 bg-white text-red-600 text-sm font-medium rounded hover:bg-red-50 transition-colors border border-red-200 shadow-sm"
        @click="deleteKey"
      >
        Delete
      </button>
    </div>

    <!-- Content -->
    <div class="flex-1 overflow-y-auto p-6">
      <StringViewer v-if="keyData.type === 'string'" :value="keyData.value as string" />
      <HashViewer v-else-if="keyData.type === 'hash'" :data="keyData.value as Record<string, string>" />
      <CollectionViewer
        v-else-if="['list', 'set', 'zset'].includes(keyData.type)"
        :data="keyData.value as string[]"
        :type="keyData.type as any"
      />
      <div v-else class="text-gray-500 italic">Unsupported type: {{ keyData.type }}</div>
    </div>
  </div>
</template>
