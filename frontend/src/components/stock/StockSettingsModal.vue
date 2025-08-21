<template>
  <n-modal v-model:show="show" :title="title" style="width: 400px" :preset="'card'">
    <n-form :model="formModel" :rules="rules" label-placement="left" label-width="80px">
      <n-form-item label="股票成本" path="costPrice">
        <n-input-number v-model:value="formModel.costPrice" min="0" placeholder="请输入股票成本">
          <template #suffix>
            {{ currencySymbol }}
          </template>
        </n-input-number>
      </n-form-item>
      <n-form-item label="股票数量" path="volume">
        <n-input-number v-model:value="formModel.volume" min="0" step="100" placeholder="请输入股票数量">
          <template #suffix>
            股
          </template>
        </n-input-number>
      </n-form-item>
      <n-form-item label="涨跌提醒" path="alarm">
        <n-input-number v-model:value="formModel.alarm" min="0" placeholder="请输入涨跌报警值(%)">
          <template #suffix>
            %
          </template>
        </n-input-number>
      </n-form-item>
      <n-form-item label="股价提醒" path="alarmPrice">
        <n-input-number v-model:value="formModel.alarmPrice" min="0" placeholder="请输入股价报警值(¥)">
          <template #suffix>
            {{ currencySymbol }}
          </template>
        </n-input-number>
      </n-form-item>
      <n-form-item label="股票排序" path="sort">
        <n-input-number v-model:value="formModel.sort" min="0" placeholder="请输入股价排序值">
        </n-input-number>
      </n-form-item>
      <n-form-item label="AI cron" path="cron">
        <n-input v-model:value="formModel.cron" placeholder="请输入cron表达式"/>
      </n-form-item>
    </n-form>
    <template #footer>
      <n-button type="primary" @click="handleSave">
        保存
      </n-button>
    </template>
  </n-modal>
</template>

<script setup>
import { ref, computed, defineProps, defineEmits, watch } from 'vue';

const props = defineProps({
  show: Boolean,
  title: String,
  initialFormModel: Object,
});

const emit = defineEmits(['update:show', 'save']);

const formModel = ref({});

watch(() => props.initialFormModel, (newVal) => {
  formModel.value = { ...newVal };
}, { immediate: true, deep: true });


const show = computed({
  get: () => props.show,
  set: (value) => emit('update:show', value)
});

const currencySymbol = computed(() => {
  return formModel.value.code && formModel.value.code.includes("hk") ? "HK$" : "¥";
});

const rules = {
  costPrice: { required: true, message: '请输入成本'},
  volume: { required: true, message: '请输入数量'},
  alarm:{required: true, message: '涨跌报警值'} ,
  alarmPrice: { required: true, message: '请输入报警价格'},
  sort: { required: true, message: '请输入排序值'},
};

const handleSave = () => {
  emit('save', formModel.value);
};

</script>
