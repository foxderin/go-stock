<script setup lang="ts">
import { NForm, NGrid, NFormItemGi, NDatePicker, NSelect, NText } from 'naive-ui';

const props = defineProps({
  modelValue: {
    type: Object,
    required: true
  },
  explanations: {
    type: Array,
    default: () => []
  }
});

const emit = defineEmits(['update:modelValue', 'date-change', 'explanation-change']);

const handleDateUpdate = (value, formattedValue) => {
  emit('date-change', formattedValue);
};

const handleExplanationUpdate = (value) => {
  emit('explanation-change', value);
};
</script>

<template>
  <n-form :model="modelValue">
    <n-grid :cols="24" :x-gap="24">
      <n-form-item-gi :span="4" label="日期" path="dateValue" label-placement="left">
        <n-date-picker
            :value="modelValue.dateValue"
            @update:formatted-value="handleDateUpdate"
            value-format="yyyy-MM-dd"
            type="date"
        />
      </n-form-item-gi>
      <n-form-item-gi :span="8" label="上榜原因" path="EXPLANATION" label-placement="left">
        <n-select
            clearable
            placeholder="上榜原因过滤"
            :value="modelValue.EXPLANATION"
            :options="explanations"
            @update:value="handleExplanationUpdate"
        />
      </n-form-item-gi>
      <n-form-item-gi :span="10" label="" label-placement="left">
        <n-text type="error">*当天的龙虎榜数据通常在收盘结束后一小时左右更新</n-text>
      </n-form-item-gi>
    </n-grid>
  </n-form>
</template>
