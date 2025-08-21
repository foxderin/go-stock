<script setup lang="ts">
import { defineProps } from 'vue';

const props = defineProps({
  item: Object,
  isStreamLoad: Boolean,
  icon: String,
});
</script>

<template>
  <template v-if="item.role === 'assistant'">
    <t-chat-reasoning expand-icon-placement="right">
      <t-chat-loading v-if="isStreamLoad" text="思考中..." />
      <t-chat-content v-if="item.reasoning" :content="item.reasoning" />
    </t-chat-reasoning>
    <t-chat-content v-if="item.content" :content="item.content" />
  </template>
  <t-chat-content v-else :content="item.content" />
</template>
