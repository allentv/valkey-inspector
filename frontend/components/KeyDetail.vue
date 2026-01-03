<script setup lang="ts">
import StringViewer from './viewers/StringViewer.vue'
import HashViewer from './viewers/HashViewer.vue'
import CollectionViewer from './viewers/CollectionViewer.vue'

const props = defineProps<{
  keyName?: string
}>()

const { selectedDb } = useKeys()

const { data: keyData, pending, refresh } = useFetch(() => `/api/keys/${encodeURIComponent(props.keyName || '')}`, {
  params: {
    db: selectedDb
  },
  immediate: false,
  watch: [() => props.keyName]
})

// Trigger fetch when keyName changes and is present
watch(() => props.keyName, (newVal) => {
  if (newVal) {
    refresh()
  }
}, { immediate: true })

const copyToClipboard = () => {
  if (props.keyName) {
    navigator.clipboard.writeText(props.keyName)
    // Ideally trigger a toast notification here
  }
}

const deleteKey = () => {
  if (confirm(`Are you sure you want to delete ${props.keyName}?`)) {
    // TODO: Backend Integration - Delete Key
    // Need to implement DELETE /api/keys/{key} endpoint
    // await useFetch(`/api/keys/${props.keyName}`, { method: 'DELETE', params: { db: selectedDb.value } })
    console.log('TODO: Delete key', props.keyName)
  }
}
</script>

<template>
  <div v-if="!keyName" class="flex items-center justify-center h-full text-gray-400 dark:text-gray-500 bg-gray-50 dark:bg-gray-900">
    <div class="text-center">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-16 h-16 mx-auto mb-4 text-gray-300 dark:text-gray-600"
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

  <div v-else-if="pending" class="flex items-center justify-center h-full text-gray-500 dark:text-gray-400">
    <div class="animate-pulse flex flex-col items-center dark:opacity-50">
      <div class="h-8 w-32 bg-gray-200 dark:bg-gray-700 rounded mb-4"></div>
      <div class="h-64 w-full max-w-lg bg-gray-200 dark:bg-gray-700 rounded"></div>
    </div>
  </div>

  <div v-else-if="keyData" class="flex flex-col h-full bg-gray-50 dark:bg-gray-900">
    <!-- Header -->
    <div class="p-6 border-b border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 flex justify-between items-start shadow-sm z-10">
      <div class="overflow-hidden mr-4 flex-1">
        <div class="flex items-center gap-2 mb-2">
          <h2 class="text-xl font-mono font-bold text-gray-900 dark:text-white truncate" :title="keyName">
            {{ keyName }}
          </h2>
          <button
            class="text-gray-400 hover:text-blue-600 dark:hover:text-blue-400 transition-colors p-1 rounded hover:bg-blue-50 dark:hover:bg-blue-900/30"
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

        <div class="flex items-center gap-4 text-sm text-gray-500 dark:text-gray-400">
          <span class="flex items-center gap-1 bg-gray-100 dark:bg-gray-700 px-2 py-0.5 rounded">
            <span class="font-medium text-gray-700 dark:text-gray-300">Type:</span> {{ keyData?.type?.toUpperCase() }}
          </span>
          <span class="flex items-center gap-1 bg-gray-100 dark:bg-gray-700 px-2 py-0.5 rounded">
            <span class="font-medium text-gray-700 dark:text-gray-300">TTL:</span> {{ keyData?.ttl }}s
          </span>
        </div>
      </div>

      <button
        class="px-3 py-1.5 bg-white dark:bg-gray-800 text-red-600 dark:text-red-400 text-sm font-medium rounded hover:bg-red-50 dark:hover:bg-red-900/20 transition-colors border border-red-200 dark:border-red-800 shadow-sm"
        @click="deleteKey"
      >
        Delete
      </button>
    </div>

    <!-- Content -->
    <div class="flex-1 overflow-y-auto p-6">
      <StringViewer v-if="keyData.type === 'string'" :value="keyData.value" />
      <HashViewer v-else-if="keyData.type === 'hash'" :data="keyData.value" />
      <CollectionViewer
        v-else-if="['list', 'set', 'zset'].includes(keyData.type)"
        :data="keyData.value"
        :type="keyData.type"
      />
      <div v-else class="text-gray-500 dark:text-gray-400 italic">Unsupported type: {{ keyData.type }}</div>
    </div>
  </div>
</template>
