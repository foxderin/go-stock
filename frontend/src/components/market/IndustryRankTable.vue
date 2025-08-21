<script setup>
import {defineProps} from 'vue'
import {CaretDown, CaretUp} from "@vicons/ionicons5";

const props = defineProps({
  industryRanks: Array,
  sort: String,
})

const emit = defineEmits(['change-sort'])

function changeIndustryRankSort() {
  emit('change-sort')
}
</script>

<template>
  <n-table striped>
    <n-thead>
    <n-tr>
      <n-th>行业名称</n-th>
      <n-th @click="changeIndustryRankSort">行业涨幅
        <n-icon v-if="sort==='0'">
          <CaretDown/>
        </n-icon>
        <n-icon v-if="sort==='1'">
          <CaretUp/>
        </n-icon>
      </n-th>
      <n-th>行业5日涨幅</n-th>
      <n-th>行业20日涨幅</n-th>
      <n-th>领涨股</n-th>
      <n-th>涨幅</n-th>
      <n-th>最新价</n-th>
    </n-tr>
    </n-thead>
    <n-tbody>
    <n-tr v-for="item in industryRanks" :key="item.bd_code">
      <n-td>
        <n-tag :bordered=false type="info">{{ item.bd_name }}</n-tag>
      </n-td>
      <n-td>
        <n-text :type="item.bd_zdf>0?'error':'success'">{{ item.bd_zdf }}%</n-text>
      </n-td>
      <n-td>
        <n-text :type="item.bd_zdf5>0?'error':'success'">{{ item.bd_zdf5 }}%</n-text>
      </n-td>
      <n-td>
        <n-text :type="item.bd_zdf20>0?'error':'success'">{{ item.bd_zdf20 }}%</n-text>
      </n-td>
      <n-td>
        <n-text :type="item.nzg_zdf>0?'error':'success'"> {{ item.nzg_name }}
          <n-text type="info">{{ item.nzg_code }}</n-text>
        </n-text>
      </n-td>
      <n-td>
        <n-text :type="item.nzg_zdf>0?'error':'success'"> {{ item.nzg_zdf }}%</n-text>
      </n-td>
      <n-td>
        <n-text :type="item.nzg_zdf>0?'error':'success'">{{ item.nzg_zxj }}</n-text>
      </n-td>
    </n-tr>
    </n-tbody>
  </n-table>
</template>
