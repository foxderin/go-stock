<script setup lang="ts">
import { defineProps, defineEmits } from 'vue';
import { NFlex, NSelect } from "naive-ui";

const props = defineProps({
  loading: Boolean,
  isStreamLoad: Boolean,
  inputValue: String,
  selectValue: [String, Number],
  selectOptions: Array,
});

const emits = defineEmits(['update:inputValue', 'send', 'stop', 'update:selectValue']);
</script>

<template>
  <t-chat-sender
      :value="inputValue"
      @update:value="emits('update:inputValue', $event)"
      class="chat-sender"
      :textarea-props="{
      placeholder: '请输入消息...',
    }"
      :loading="loading"
      :stop-disabled="isStreamLoad"
      @send="emits('send')"
      @stop="emits('stop')"
  >
    <template #suffix>
      <t-button theme="default" variant="text" size="large" class="btn" @click="emits('send')"> 发送 </t-button>
    </template>
    <template #prefix>
      <NFlex>
        <NSelect
            :value="selectValue"
            @update:value="emits('update:selectValue', $event)"
            :options="selectOptions"
            label-field="name"
            value-field="ID"
            size="tiny"
            style="width: 200px;"
        />
      </NFlex>
    </template>
  </t-chat-sender>
</template>

<style lang="less" scoped>
.chat-sender {
  .btn {
    color: var(--td-text-color-disabled);
    border: none;
    &:hover {
      color: var(--td-brand-color-hover);
      border: none;
      background: none;
    }
  }
  .btn.t-button {
    height: var(--td-comp-size-m);
    padding: 0;
  }
}
</style>
