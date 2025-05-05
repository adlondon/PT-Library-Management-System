<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>
<template>
  <div>
    <nav class="nav-bar">
      <RouterLink to="/">Library</RouterLink>
      <RouterLink to="/checkout">Checkout</RouterLink>
    </nav>
  </div>
  <RouterView />
</template>

<style scoped>
.nav-bar {
  display: flex;
  gap: 30px;
  height: 50px;
  padding: 20px;

  a {
    color: white;
    text-decoration: none;
    font-size: 20px;
    font-weight: bold;

    &.router-link-active {
      text-decoration: underline;
    }
  }
}
</style>
