export interface KeyItem {
  name: string
  type: string
  ttl?: number
}

interface ScanResponse {
  cursor: number
  keys: KeyItem[]
}

export const useKeys = () => {
  // Use useState for shared state across components (e.g. KeyList and KeyDetail)
  const selectedDb = useState<number>('selectedDb', () => 0)
  const keys = useState<KeyItem[]>('keys', () => [])
  const cursor = useState<number>('cursor', () => 0)
  const searchPattern = useState<string>('searchPattern', () => '*')
  const loading = useState<boolean>('keysLoading', () => false)
  const hasMore = useState<boolean>('keysHasMore', () => true)
  const selectedKey = useState<string | undefined>('selectedKey', () => undefined)

  const fetchKeys = async (reset = false) => {
    if (loading.value && !reset) return
    loading.value = true

    if (reset) {
      cursor.value = 0
      keys.value = []
      hasMore.value = true
      selectedKey.value = undefined
    }

    try {
      const { data, error } = await useFetch<ScanResponse>('/api/keys', {
        params: {
          cursor: cursor.value,
          match: searchPattern.value,
          count: 100,
          db: selectedDb.value,
        },
      })

      if (error.value) {
        console.error('Error fetching keys:', error.value)
      } else if (data.value) {
        const newKeys = data.value.keys || []
        keys.value = reset ? newKeys : [...keys.value, ...newKeys]
        cursor.value = data.value.cursor
        hasMore.value = data.value.cursor !== 0
      }
    } finally {
      loading.value = false
    }
  }

  return {
    selectedDb,
    keys,
    loading,
    searchPattern,
    hasMore,
    selectedKey,
    fetchKeys,
  }
}