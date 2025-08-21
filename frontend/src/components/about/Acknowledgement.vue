<script setup>
import { NDivider, NGradientText } from 'naive-ui';
</script>

<template>
  <div>
    <n-divider title-placement="center">鸣谢</n-divider>
    <div style="justify-self: center;text-align: left">
      <p>
        感谢以下捐赠者：
        <n-gradient-text size="small" type="warning">*晨</n-gradient-text><n-divider vertical />
      </p>
      <p>
        感谢以下开发者：
        <a href="https://github.com/GiCo001" target="_blank">@Gico</a><n-divider vertical />
        <a href="https://github.com/CodeNoobLH" target="_blank">浓睡不消残酒</a><n-divider vertical />
        <a href="https://github.com/gnim2600" target="_blank">@gnim2600</a><n-divider vertical />
        <a href="https://github.com/XXXiaohuayanGGG" target="_blank">@XXXiaohuayanGGG</a><n-divider vertical />
        <a href="https://github.com/2lovecode" target="_blank">@2lovecode</a><n-divider vertical />
        <a href="https://github.com/JerryLookupU" target="_blank">@JerryLookupU</a><n-divider vertical />
      </p>
      <p>
        感谢以下开源项目：
        <a href="https://github.com/wailsapp/wails" target="_blank">Wails</a><n-divider vertical />
        <a href="https://github.com/vuejs" target="_blank">Vue</a><n-divider vertical />
        <a href="https://github.com/tusen-ai/naive-ui" target="_blank">NaiveUI</a><n-divider vertical />
      </p>
    </div>
  </div>
</template>
