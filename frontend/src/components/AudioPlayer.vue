<script setup>
  import { ref } from 'vue'
  import { useAudioPlayer } from '@/composables/useAudioPlayer'

  const props = defineProps({
    src: { type: String, required: true },
    title: { type: String, default: '' },
  })

  const { isPlaying, currentTime, duration, error, play, pause } = useAudioPlayer(props.src)

  const formatTime = (s) => {
    if (!s || isNaN(s)) return '0:00'
    const mins = Math.floor(s / 60)
    const secs = Math.floor(s % 60)
      .toString()
      .padStart(2, '0')
    return `${mins}:${secs}`
  }
</script>

<template>
  <div class="player-card">
    <h3>{{ title }}</h3>
    <div class="controls">
      <button @click="isPlaying ? pause() : play()" class="play-btn">
        {{ isPlaying ? '⏸' : '▶' }}
      </button>
      <div class="time">{{ formatTime(currentTime) }} / {{ formatTime(duration) }}</div>
    </div>
    <p v-if="error" class="error">{{ error }}</p>
  </div>
</template>

<style scoped>
  .player-card {
    background: #fff;
    border: 1px solid rgba(88, 204, 2, 0.4);
    border-radius: 30px;
    padding: 1.5rem;
    text-align: center;
  }
  .controls {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 1rem;
  }
  .play-btn {
    font-size: 2rem;
    background: none;
    border: none;
    cursor: pointer;
  }
  .time {
    font-family: monospace;
  }
  .error {
    color: #e74c3c;
  }
</style>
