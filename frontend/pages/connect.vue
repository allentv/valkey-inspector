<script setup lang="ts">
definePageMeta({
  layout: 'blank',
})

const { connection, connect } = useConnection()
const loading = ref(false)
const error = ref('')
const router = useRouter()

const handleConnect = async () => {
  loading.value = true
  error.value = ''
  try {
    await connect()
    router.push('/')
  } catch (e) {
    error.value = 'Failed to connect. Please check your settings.'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
      <h1 class="text-2xl font-bold mb-6 text-center text-gray-800">Connect to Valkey</h1>

      <form class="space-y-4" @submit.prevent="handleConnect">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Host</label>
          <input
            v-model="connection.host"
            type="text"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
            placeholder="localhost"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Port</label>
          <input
            v-model.number="connection.port"
            type="number"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
            placeholder="6379"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Username (Optional)</label>
          <input
            v-model="connection.username"
            type="text"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Password (Optional)</label>
          <input
            v-model="connection.password"
            type="password"
            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
          />
        </div>

        <div v-if="error" class="text-red-500 text-sm text-center">
          {{ error }}
        </div>

        <button
          type="submit"
          :disabled="loading"
          class="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 transition-colors"
        >
          {{ loading ? 'Connecting...' : 'Connect' }}
        </button>
      </form>
    </div>
  </div>
</template>