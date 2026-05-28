<script setup>
  import BackButton from '@/components/BackButton.vue'

  // Можно добавить логику выбора языка позже
  const selectLanguage = (lang) => {
    console.log('Выбран язык:', lang)
    // здесь будет сохранение в localStorage/store и, возможно, переход на главную
  }
</script>

<template>
  <div class="mosaic-wrap">
    <div id="i8z5xo3zc_0" class="root root--u-i8z5xo3zc root--s1-i8z5xo3zc root--s2-i8z5xo3zc">
      <div
        id="ikc0yythz_0"
        class="container container--u-ikc0yythz container--s1-ikc0yythz container--s2-ikc0yythz">
        <!-- Верхняя панель с кнопкой возврата -->
        <div
          id="iw3pdlkas_0"
          class="container container--u-iw3pdlkas container--s1-iw3pdlkas container--s2-iw3pdlkas">
          <BackButton to="/" label="← На главную" />
        </div>

        <!-- Белый центральный блок -->
        <div
          id="iryfzmc8c_0"
          class="container container--u-iryfzmc8c container--s1-iryfzmc8c container--s2-iryfzmc8c">
          <div class="lang-content">
            <h1 class="logo">Выберите язык</h1>
            <p class="subtitle">Language selection</p>

            <div class="cards-row">
              <!-- Карточка: Русский -->
              <div
                class="lang-card"
                role="button"
                tabindex="0"
                @click="selectLanguage('ru')"
                @keydown.enter="selectLanguage('ru')">
                <span class="flag-icon">🇷🇺</span>
                <span class="lang-name">Русский</span>
              </div>

              <!-- Карточка: English -->
              <div
                class="lang-card"
                role="button"
                tabindex="0"
                @click="selectLanguage('en')"
                @keydown.enter="selectLanguage('en')">
                <span class="flag-icon">🇬🇧</span>
                <span class="lang-name">English</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
  /* === Общие стили (как в HomeView и LoginView) === */

  .root--u-i8z5xo3zc {
    background-size: auto;
    background-image: linear-gradient(rgba(255, 253, 247, 1) 0%, rgba(255, 253, 247, 1) 100%);
    background-repeat: no-repeat;
    background-position: left 0px top 0px;
    background-attachment: scroll;
    display: flex;
    align-items: center;
    justify-content: center;
    min-height: 100vh;
  }

  .container--u-ikc0yythz {
    min-height: 100vh;
    width: 100%;
    max-width: 1200px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 2rem 0;
  }

  .container--u-iw3pdlkas {
    display: flex;
    justify-content: flex-start;
    align-items: center;
    width: 100%;
    max-width: 1000px;
    padding: 0 2rem;
    margin-bottom: 2rem;
  }

  /* Белый центральный блок */
  .container--u-iryfzmc8c {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 90%;
    max-width: 600px; /* немного шире для двух карточек */
    min-height: auto;
    margin: 0 auto;
    padding: 3rem 2.5rem;
    border-radius: 50px;
    background-size: auto;
    background-image: linear-gradient(rgba(255, 255, 255, 1) 0%, rgba(255, 255, 255, 1) 100%);
    background-repeat: no-repeat;
    background-position: left 0px top 0px;
    background-attachment: scroll;
    border: 1px solid rgba(88, 204, 2, 0.6);
    box-shadow:
      0px 20px 40px 0px rgba(88, 204, 2, 0.06),
      0px 10px 20px 0px rgba(0, 0, 0, 0.04);
  }

  /* Контент выбора языка */
  .lang-content {
    width: 100%;
    text-align: center;
  }

  .logo {
    font-size: 2.5rem;
    color: #2c3e50;
    margin-bottom: 0.5rem;
    font-weight: 700;
  }

  .subtitle {
    font-size: 1.1rem;
    color: #777;
    margin-bottom: 2.5rem;
    font-weight: 300;
  }

  /* Строка с карточками */
  .cards-row {
    display: flex;
    justify-content: center;
    gap: 2rem;
    flex-wrap: wrap;
  }

  /* Карточка языка */
  .lang-card {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 0.8rem;
    width: 180px;
    padding: 2rem 1.5rem;
    border-radius: 30px;
    background: #ffffff;
    border: 1px solid rgba(88, 204, 2, 0.25);
    box-shadow:
      0 10px 20px rgba(88, 204, 2, 0.04),
      0 5px 10px rgba(0, 0, 0, 0.06);
    cursor: pointer;
    transition:
      transform 0.2s,
      box-shadow 0.2s,
      border-color 0.2s;
    outline: none;
  }

  .lang-card:hover,
  .lang-card:focus-visible {
    transform: translateY(-4px);
    border-color: rgba(88, 204, 2, 0.6);
    box-shadow:
      0 12px 24px rgba(88, 204, 2, 0.08),
      0 6px 12px rgba(0, 0, 0, 0.1);
  }

  .flag-icon {
    font-size: 3rem;
    line-height: 1;
  }

  .lang-name {
    font-size: 1.2rem;
    font-weight: 600;
    color: #2c3e50;
  }
</style>
