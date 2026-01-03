<script setup lang="ts">
definePageMeta({
  layout: 'blank',
})

const { connection, connect } = useConnection()
const { isDark, toggleTheme } = useTheme()
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
  <div class="min-h-screen flex bg-gray-100 dark:bg-gray-900">
    <!-- Left Side: Hero / Info -->
    <div
      class="hidden lg:flex flex-1 flex-col justify-center items-center bg-gradient-to-br from-blue-600 to-indigo-800 text-white p-12 relative"
    >
      <div class="max-w-lg text-center">
        <h2 class="text-4xl font-bold mb-6">Valkey Inspector</h2>
        <p class="text-lg opacity-90 mb-8">
          A powerful, modern GUI for inspecting and managing your Valkey and Redis clusters. Visualize keys, monitor
          performance, and manage data with ease.
        </p>
        <!-- Placeholder for an illustration/image -->
        <div class="w-64 h-64 bg-white/10 rounded-full mx-auto flex items-center justify-center backdrop-blur-sm">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="w-32 h-32 opacity-80"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M20.25 6.375c0 2.278-3.694 4.125-8.25 4.125S3.75 8.653 3.75 6.375m16.5 0c0-2.278-3.694-4.125-8.25-4.125S3.75 4.097 3.75 6.375m16.5 0v11.25c0 2.278-3.694 4.125-8.25 4.125s-8.25-1.847-8.25-4.125V6.375m16.5 0v3.75m-16.5-3.75v3.75m16.5 0v3.75C20.25 16.153 16.556 18 12 18s-8.25-1.847-8.25-4.125v-3.75m16.5 0c0 2.278-3.694 4.125-8.25 4.125s-8.25-1.847-8.25-4.125"
            />
          </svg>
        </div>
      </div>
    </div>

    <!-- Right Side: Login Form -->
    <div class="w-full lg:w-[500px] flex flex-col justify-center p-8 bg-white dark:bg-gray-800 relative shadow-2xl">
      <!-- Dark Mode Toggle -->
      <button
        class="absolute top-6 right-6 text-gray-400 hover:text-gray-600 dark:hover:text-gray-200 p-2 rounded-md hover:bg-gray-100 dark:hover:bg-gray-700 transition-colors"
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
          class="w-6 h-6"
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
          class="w-6 h-6"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M21.752 15.002A9.718 9.718 0 0118 15.75c-5.385 0-9.75-4.365-9.75-9.75 0-1.33.266-2.597.748-3.752A9.753 9.753 0 003 11.25C3 16.635 7.365 21 12.75 21a9.753 9.753 0 009.002-5.998z"
          />
        </svg>
      </button>

      <div class="w-full max-w-sm mx-auto">
        <h1 class="text-2xl font-bold mb-6 text-center text-gray-800 dark:text-white">Connect to Valkey</h1>

        <form class="space-y-4" @submit.prevent="handleConnect">
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Host</label>
            <input
              v-model="connection.host"
              type="text"
              required
              class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md focus:ring-blue-500 focus:border-blue-500 bg-white dark:bg-gray-700 text-gray-900 dark:text-white"
              placeholder="localhost"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Port</label>
            <input
              v-model.number="connection.port"
              type="number"
              required
              class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md focus:ring-blue-500 focus:border-blue-500 bg-white dark:bg-gray-700 text-gray-900 dark:text-white"
              placeholder="6379"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Username (Optional)</label>
            <input
              v-model="connection.username"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md focus:ring-blue-500 focus:border-blue-500 bg-white dark:bg-gray-700 text-gray-900 dark:text-white"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Password (Optional)</label>
            <input
              v-model="connection.password"
              type="password"
              class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md focus:ring-blue-500 focus:border-blue-500 bg-white dark:bg-gray-700 text-gray-900 dark:text-white"
            />
          </div>

          <div v-if="error" class="text-red-500 text-sm text-center">
            {{ error }}
          </div>

          <button
            type="submit"
            :disabled="loading"
            class="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 dark:focus:ring-offset-gray-800 disabled:opacity-50 transition-colors"
          >
            {{ loading ? 'Connecting...' : 'Connect' }}
          </button>
        </form>
      </div>
    </div>
  </div>
</template>
