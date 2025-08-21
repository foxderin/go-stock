<script setup>
import { NDivider, NSpace, NAvatar, NImage, NFlex } from 'naive-ui';

defineProps({
  wxgzh: String,
  alipay: String,
  wxpay: String,
});
</script>

<template>
  <div>
    <n-divider title-placement="center">关于作者</n-divider>
    <n-space vertical>
      <n-avatar width="100" src="https://avatars.githubusercontent.com/u/7401917?v=4" />
      <h2><a href="https://github.com/ArvinLovegood" target="_blank">@ArvinLovegood</a></h2>
      <p>一个热爱编程的小白，欢迎关注我的Github/微信公众号</p>
      <n-image width="300" :src="wxgzh" />
      <p>开源不易，如果觉得好用，可以请作者喝杯咖啡。</p>
      <n-flex justify="center">
        <n-image width="200" :src="alipay" />
        <n-image width="200" :src="wxpay" />
      </n-flex>
    </n-space>
  </div>
</template>
