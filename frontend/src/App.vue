<script setup>
import HelloWorld from './components/HelloWorld.vue'</script>

<template>
<!--  <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>-->
  <HelloWorld class="main-container"/>
</template>

<style>
.main-container {
  padding: 10px 20px;
}
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
