<template>
  <ul class="list">
    <div class="item">
      <button class="button-grad">Login</button>
    </div>

    <div class="item">
      <button class="button-grad">Register</button>
    </div>
  </ul>
</template>

<style scoped>
button {
  display: flex;
  width: 100%;
  height: 100%;
  align-items: center;
  justify-content: center;
}

.item {
  display: flex;
  width: 60%;
  height: 60%;
  align-items: center;
  justify-content: center;
  margin: 15px;
}

.list {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
}

.button-grad {
  background: rgb(2, 0, 36);
  background: linear-gradient(
    8deg,
    rgba(2, 0, 36, 1) 0%,
    rgba(9, 121, 32, 1) 54%,
    rgba(0, 212, 255, 1) 100%
  );
  border-radius: 8px;
  border-style: none;
  box-sizing: border-box;
  color: #ffffff;
  cursor: pointer;
  flex-shrink: 0;
  font-size: 16px;
  font-weight: 500;
  height: 4rem;
  padding: 0 1.6rem;
  text-align: center;
  text-shadow: rgba(0, 0, 0, 0.25) 0 3px 8px;
  transition: all 0.5s;
  user-select: none;
  -webkit-user-select: none;
  touch-action: manipulation;
}

.button-grad:hover {
  box-shadow: rgba(80, 63, 205, 0.5) 0 1px 30px;
  transition-duration: 0.1s;
}

@media (min-width: 768px) {
  .button-grad {
    padding: 0 2.6rem;
  }
}
</style>
