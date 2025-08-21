<template>
  <div style="position: fixed; bottom: 18px; right: 5px; z-index: 10; width: 400px">
    <n-input-group>
      <n-auto-complete
          v-if="addBTN"
          :value="modelValue"
          @update:value="onUpdateValue"
          :input-props="{ autocomplete: 'disabled' }"
          :options="options"
          placeholder="股票指数名称/代码/弹幕"
          clearable
          @select="onSelect"
      />
      <n-popover trigger="manual" :show="showPopover">
        <template #trigger>
          <n-button type="primary" @click="emit('add-stock')" v-if="addBTN">
            <n-icon :component="Add" /> &nbsp;关注
          </n-button>
        </template>
        <span>输入股票名称/代码关键词开始吧~~~</span>
      </n-popover>
      <n-button type="info" @click="emit('send-danmu')" v-if="enableDanmu">
        <n-icon :component="ChatboxOutline" /> &nbsp;发送弹幕
      </n-button>
    </n-input-group>
  </div>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue';
import { Add, ChatboxOutline } from '@vicons/ionicons5';

const props = defineProps({
  modelValue: String,
  addBTN: Boolean,
  enableDanmu: Boolean,
  options: Array,
  showPopover: Boolean,
});

const emit = defineEmits(['update:modelValue', 'select', 'add-stock', 'send-danmu']);

const onUpdateValue = (value) => {
  emit('update:modelValue', value);
};

const onSelect = (value) => {
  emit('select', value);
};
</script>
