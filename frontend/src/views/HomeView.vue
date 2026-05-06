<script setup></script>

<template>
  <div id="i8z5xo3zc_0" class="root root--u-i8z5xo3zc root--s1-i8z5xo3zc root--s2-i8z5xo3zc">
    <div
      id="ikc0yythz_0"
      class="container container--u-ikc0yythz container--s1-ikc0yythz container--s2-ikc0yythz">
      <div
        id="iw3pdlkas_0"
        class="container container--u-iw3pdlkas container--s1-iw3pdlkas container--s2-iw3pdlkas">
        <a
          id="isfns9y2c_0"
          role="button"
          href="/login"
          data-action-element=""
          class="button button--u-isfns9y2c button--s1-isfns9y2c button--s2-isfns9y2c">
          <span
            id="iuoln6uzk_0"
            class="text-button text-button--s1-iuoln6uzk text-button--s2-iuoln6uzk">
            <span class="text-block-wrap-div">
              <span>Вход</span>
            </span>
          </span>
        </a>
        <a
          id="irpkbfpw9_0"
          role="button"
          href="/register"
          data-action-element=""
          class="button button--u-irpkbfpw9 button--s1-irpkbfpw9 button--s2-irpkbfpw9">
          <span
            id="ik5zpz3n6_0"
            class="text-button text-button--s1-ik5zpz3n6 text-button--s2-ik5zpz3n6">
            <span class="text-block-wrap-div">
              <span>Регистрация</span>
            </span>
          </span>
        </a>
      </div>
      <div
        id="iryfzmc8c_0"
        class="container container--u-iryfzmc8c container--s1-iryfzmc8c container--s2-iryfzmc8c">
        <div class="welcome-content">
          <h1 class="logo">Добро пожаловать</h1>
          <p class="subtitle">Ваш сервис для изучения языков</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
  .root--u-i8z5xo3zc {
    background-size: auto;
    background-image: linear-gradient(rgba(255, 253, 247, 1) 0%, rgba(255, 253, 247, 1) 100%);
    background-repeat: no-repeat;
    background-position: left 0px top 0px;
    background-attachment: scroll;
    display: flex;
    align-items: center;
    justify-content: center;
    min-height: 100vh; /* на всю высоту экрана */
  }

  /* Главный контейнер */
  .container--u-ikc0yythz {
    min-height: 100vh; /* вместо фиксированной высоты 797px */
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center; /* центрирование по вертикали */
  }

  /* Панель с кнопками */
  .container--u-iw3pdlkas {
    display: flex;
    justify-content: flex-end; /* кнопки справа; если нужно по центру – заменить на center */
    align-items: center;
    gap: 1.5rem;
    margin-bottom: 2rem;
    width: 100%;
    max-width: 1000px;
  }

  /* Белый центральный блок */
  .container--u-iryfzmc8c {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100vh; /* ★ ШИРИНА УВЕЛИЧЕНА! Можете поменять на 95% или 85% */
    max-width: 1000px; /* чуть шире, чем раньше */
    min-height: 400px; /* или 65vh, если хотите по высоте */
    margin: 0 auto;
    padding: 3rem 2rem; /* внутренние отступы, чтобы текст не прилипал */
    border-radius: 50px; /* единое скругление */
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

  /* Кнопки */
  .button--u-isfns9y2c,
  .button--u-irpkbfpw9 {
    color: black;
    width: auto;
    padding: 0.8rem 2.5rem;
    font-size: 1.1rem;
    border-radius: 50px;
    background-image: linear-gradient(rgba(255, 255, 255, 1) 0%, rgba(255, 255, 255, 1) 100%);
    border: 1px solid rgba(88, 204, 2, 0.25);
    box-shadow:
      0px 10px 20px 0px rgba(88, 204, 2, 0.04),
      0px 5px 10px 0px rgba(0, 0, 0, 0.06);
    display: inline-flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    transition: transform 0.2s;
    text-decoration: none;
  }
  .button--u-isfns9y2c:hover,
  .button--u-irpkbfpw9:hover {
    transform: translateY(-2px);
    background: linear-gradient(135deg, #58cc02, #4cae01);
    color: white;
    box-shadow: 0 4px 12px rgba(88, 204, 2, 0.3);
    box-shadow: 0 6px 16px rgba(88, 204, 2, 0.4);
  }

  /* Текст внутри белого блока */
  .welcome-content {
    text-align: center;
  }
  .logo {
    font-size: 3.5rem;
    color: #2c3e50;
    margin-bottom: 1rem;
    font-weight: 700;
  }
  .subtitle {
    font-size: 1.5rem;
    color: #666;
    font-weight: 300;
  }
</style>
