export const useConnection = () => {
  const connection = useState('connection', () => ({
    host: 'localhost',
    port: 6379,
    username: '',
    password: '',
  }))

  const isConnected = useState('isConnected', () => false)

  const connect = async () => {
    // TODO: Implement backend connection validation
    // const { error } = await useFetch('/api/connect', {
    //   method: 'POST',
    //   body: connection.value
    // })
    // if (error.value) throw error.value

    // For now, simulate success immediately so UI can be reviewed
    isConnected.value = true
    return true
  }

  const disconnect = () => {
    isConnected.value = false
    // State cleanup if necessary
  }

  return {
    connection,
    isConnected,
    connect,
    disconnect,
  }
}