<template>
  <n-card title="板块信息" size="small">
    <n-table :single-line="false" size="small">
      <n-thead>
        <n-tr>
          <n-th>板块名称</n-th>
          <n-th>涨跌幅</n-th>
        </n-tr>
      </n-thead>
      <n-tbody>
        <n-tr v-for="(item, index) in plateInfo" :key="index">
          <n-td>{{ item.name }}</n-td>
          <n-td>
            <n-text :type="item.change_rate > 0 ? 'error' : 'success'">{{ item.change_rate }}%</n-text>
          </n-td>
        </n-tr>
      </n-tbody>
    </n-table>
  </n-card>
</template>

<script setup>
import { NCard, NTable, NThead, NTbody, NTr, NTh, NTd, NText } from 'naive-ui';

defineProps({
  plateInfo: {
    type: Array,
    default: () => []
  }
});
</script>
