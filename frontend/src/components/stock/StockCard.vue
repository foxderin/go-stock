<template>
  <n-card :bordered="false" class="stock-card" :class="{ 'blink-border': shouldBlink(stock) }">
    <template #header>
      <n-flex justify="space-between">
        <n-text class="stock-name" @click="showKLine(stock.code, stock.name)">
          {{ stock.name }}
        </n-text>
        <n-text :type="stock.changePercent > 0 ? 'error' : 'success'">
          {{ stock.changePercent > 0 ? '+' : '' }}{{ stock.changePercent }}%
        </n-text>
      </n-flex>
    </template>
    <template #header-extra>
      <n-text class="stock-code" @click="showFenShi(stock.code, stock.name, stock.changePercent)">
        {{ stock.code.replace('sh', '').replace('sz', '').replace('bj', '') }}
      </n-text>
    </template>
    <n-flex justify="space-between" class="stock-info">
      <n-text :type="stock.price > stock.open ? 'error' : 'success'" class="stock-price">
        {{ stock.price }}
      </n-text>
      <n-text :type="stock.change > 0 ? 'error' : 'success'" class="stock-change">
        {{ stock.change > 0 ? '+' : '' }}{{ stock.change }}
      </n-text>
    </n-flex>
    <n-flex justify="space-between" class="stock-extra-info">
      <n-text>高: {{ stock.high }}</n-text>
      <n-text>低: {{ stock.low }}</n-text>
    </n-flex>
    <n-flex justify="space-between" class="stock-extra-info">
      <n-text>量: {{ stock.volume }}</n-text>
      <n-text>额: {{ stock.amount }}</n-text>
    </n-flex>
    <template #action>
      <n-flex justify="space-between" class="stock-actions">
        <n-button type="error" size="tiny" @click="deleteStock(stock.code)">
          删除
        </n-button>
        <n-button type="primary" size="tiny" @click="editStock(stock)">
          编辑
        </n-button>
        <n-button type="info" size="tiny" @click="showMoneyTrend(stock.code, stock.name)">
          趋势
        </n-button>
        <n-button type="warning" size="tiny" @click="aiCheckStock(stock.name, stock.code)">
          AI分析
        </n-button>
        <n-dropdown trigger="click" :options="groupList" key-field="ID" label-field="name"
                    @select="(groupId) => addStockToGroup(groupId, stock.code, stock.name)">
          <n-button type="warning" size="tiny">分组</n-button>
        </n-dropdown>
      </n-flex>
    </template>
  </n-card>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue';

const props = defineProps({
  stock: Object,
  groupList: Array,
});

const emit = defineEmits([
  'delete', 'edit', 'show-k-line', 'show-fen-shi', 'show-money-trend',
  'ai-check', 'add-to-group'
]);

const shouldBlink = (stock) => {
  if (stock.alarmPrice > 0 && stock.price >= stock.alarmPrice) {
    return true;
  }
  return stock.alarm > 0 && Math.abs(stock.changePercent) >= stock.alarm;
};

const deleteStock = (code) => emit('delete', code);
const editStock = (stock) => emit('edit', stock);
const showKLine = (code, name) => emit('show-k-line', code, name);
const showFenShi = (code, name, changePercent) => emit('show-fen-shi', code, name, changePercent);
const showMoneyTrend = (code, name) => emit('show-money-trend', code, name);
const aiCheckStock = (name, code) => emit('ai-check', name, code);
const addStockToGroup = (groupId, code, name) => emit('add-to-group', groupId, code, name);

</script>

<style scoped>
.stock-card {
  text-align: left;
  margin-bottom: 10px;
  cursor: pointer;
}
.stock-name {
  font-weight: bold;
}
.stock-code {
  font-size: 0.8em;
  color: #999;
}
.stock-price {
  font-size: 1.5em;
  font-weight: bold;
}
.stock-change {
  font-size: 1.2em;
}
.stock-info, .stock-extra-info, .stock-actions {
  margin-top: 5px;
}
.blink-border {
  animation: blink-border 1s linear infinite;
  border: 2px solid red;
}

@keyframes blink-border {
  0% {
    border-color: red;
  }
  50% {
    border-color: transparent;
  }
  100% {
    border-color: red;
  }
}
</style>
