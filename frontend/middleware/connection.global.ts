export default defineNuxtRouteMiddleware((to) => {
  const { isConnected } = useConnection()

  if (!isConnected.value && to.path !== '/connect') {
    return navigateTo('/connect')
  }

  if (isConnected.value && to.path === '/connect') {
    return navigateTo('/')
  }
})
