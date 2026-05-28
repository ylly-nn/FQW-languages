import { ref, onUnmounted } from 'vue'

export function useAudioPlayer(src) {
  const audio = new Audio(src)
  audio.preload = 'metadata'

  const isPlaying = ref(false)
  const currentTime = ref(0)
  const duration = ref(0)
  const error = ref(null)

  const play = () => {
    audio.play().catch(e => (error.value = e.message))
    isPlaying.value = true
  }
  const pause = () => {
    audio.pause()
    isPlaying.value = false
  }

  audio.addEventListener('loadedmetadata', () => {
    duration.value = audio.duration
  })
  audio.addEventListener('timeupdate', () => {
    currentTime.value = audio.currentTime
  })
  audio.addEventListener('ended', () => {
    isPlaying.value = false
  })

  onUnmounted(() => {
    audio.pause()
    audio.src = ''
    audio.load()
  })

  return { isPlaying, currentTime, duration, error, play, pause }
}
