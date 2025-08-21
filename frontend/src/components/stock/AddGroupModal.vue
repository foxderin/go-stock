<template>
  <n-modal v-model:show="show" title="添加分组" style="width: 400px;text-align: left" :preset="'card'">
    <n-form
        :model="formModel"
        size="medium"
        label-placement="left"
    >
      <n-grid :cols="2">
        <n-form-item-gi label="分组名称:" path="name" :span="5">
          <n-input v-model:value="formModel.name" style="width: 100%" placeholder="请输入分组名称"/>
        </n-form-item-gi>
        <n-form-item-gi label="分组排序:" path="sort" :span="5">
          <n-input-number v-model:value="formModel.sort" style="width: 100%" min="0"
                          placeholder="请输入分组排序值"></n-input-number>
        </n-form-item-gi>
      </n-grid>
    </n-form>
    <template #footer>
      <n-flex justify="end">
        <n-button type="primary" @click="handleSave">
          保存
        </n-button>
        <n-button type="warning" @click="show = false">
          取消
        </n-button>
      </n-flex>
    </template>
  </n-modal>
</template>

<script setup>
import { ref, computed, defineProps, defineEmits, watch } from 'vue';

const props = defineProps({
  show: Boolean,
});

const emit = defineEmits(['update:show', 'save']);

const formModel = ref({
  name: '',
  sort: 0
});

const show = computed({
  get: () => props.show,
  set: (value) => emit('update:show', value)
});

const handleSave = () => {
  emit('save', { ...formModel.value });
  show.value = false; // Close modal after save
};

watch(() => props.show, (newVal) => {
  if (newVal) {
    // Reset form when modal opens
    formModel.value = { name: '', sort: 0 };
  }
});
</script>
