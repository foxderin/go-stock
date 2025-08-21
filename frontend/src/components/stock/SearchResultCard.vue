<template>
  <n-card :bordered="false" style="text-align: left;margin-bottom: 10px;cursor: pointer">
    <template #header>
      <n-flex justify="space-between">
        <n-text>{{ stock['股票名称'] }}</n-text>
        <n-text>{{ stock['股票代码'] }}</n-text>
      </n-flex>
    </template>
    <template #action>
      <n-flex justify="space-between">
        <n-button type="primary" size="tiny" @click="emit('add', stock['股票代码'])">
          <n-icon :component="Add"/>
          关注
        </n-button>
        <n-dropdown trigger="click" :options="groupList" key-field="ID" label-field="name"
                    @select="(groupId) => emit('add-to-group', groupId, stock['股票代码'], stock['股票名称'])">
          <n-button type="warning" size="tiny">设置分组</n-button>
        </n-dropdown>
      </n-flex>
    </template>
  </n-card>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue';
import { Add } from '@vicons/ionicons5';

const props = defineProps({
  stock: Object,
  groupList: Array,
});

const emit = defineEmits(['add', 'add-to-group']);
</script>
