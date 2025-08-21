<script setup>
import { defineProps, defineEmits } from 'vue';
import { Add, ChatboxOutline } from "@vicons/ionicons5";

const props = defineProps({
  modelValue: String,
  options: Array,
  enableDanmu: Boolean,
});

const emits = defineEmits(['update:modelValue', 'select', 'search', 'add', 'send-danmu']);

function handleUpdate(value) {
  emits('update:modelValue', value);
  emits('search', value);
}

function handleSelect(value) {
  emits('select', value);
}

function handleAdd() {
  emits('add');
}

function handleSendDanmu() {
  emits('send-danmu');
}
</script>

<template>
  <div style="position: fixed;bottom: 18px;right:5px;z-index: 10;width: 400px">
    <n-input-group>
      <n-auto-complete
          :value="modelValue"
          @update:value="handleUpdate"
          :input-props="{ autocomplete: 'disabled' }"
          :options="options"
          placeholder="基金名称/代码/弹幕"
          clearable
          @select="handleSelect"
      />
      <n-button type="primary" @click="handleAdd">
        <n-icon :component="Add"/>
        关注
      </n-button>
      <n-button type="info" @click="handleSendDanmu" v-if="enableDanmu">
        <n-icon :component="ChatboxOutline"/>
        发送弹幕
      </n-button>
    </n-input-group>
  </div>
</template>
