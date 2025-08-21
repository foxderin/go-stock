<script setup>
import { defineProps, defineEmits } from 'vue';
import { NEllipsis, NText } from "naive-ui";
import { h } from "vue";

const props = defineProps({
  fund: Object,
});

const emits = defineEmits(['unfollow', 'show-chart', 'show-details']);

function unfollow() {
  emits('unfollow', props.fund.code);
}

function showChart() {
  emits('show-chart', props.fund.code, props.fund.name);
}

function showDetails() {
  emits('show-details', props.fund.code, props.fund.name);
}

function formatterTitle(title) {
  return () => h(NEllipsis, {
    style: {
      'font-size': '16px',
      'max-width': '180px',
    },
  }, { default: () => title });
}
</script>

<template>
  <n-gi :id="fund.code + '_gi'" style="margin-left: 2px">
    <n-card :id="fund.code" :title="formatterTitle(fund.name)">
      <template #header-extra>
        <n-tag size="small" :bordered="false" type="info">{{ fund.code }}</n-tag>&nbsp;
        <n-tag size="small" :bordered="false" type="success" @click="unfollow"> 取消关注</n-tag>
      </template>
      <n-flex>
        <n-text size="small" :type="fund.netEstimatedRate > 0 ? 'error' : 'success'" :bordered="false" v-if="fund.netEstimatedUnit">
          估算净值：{{ fund.netEstimatedUnit }}&nbsp;
          {{ fund.netEstimatedRate }} %&nbsp;&nbsp;&nbsp;
          ({{ fund.netEstimatedUnitTime }})
        </n-text>
        <br>
        <n-text size="small" :type="fund.netEstimatedRate > 0 ? 'error' : 'success'" :bordered="false" v-if="fund.netUnitValue">
          单位净值：{{ fund.netUnitValue }}&nbsp;&nbsp;&nbsp; ({{ fund.netUnitValueDate }})
        </n-text>
      </n-flex>
      <n-flex justify="start" style="margin-top: 10px">
        <n-tag size="small" :type="fund.fundBasic.netGrowth1 > 0 ? 'error' : 'success'" :bordered="false" v-if="fund.fundBasic.netGrowth1">近一月：{{ fund.fundBasic.netGrowth1 }}%</n-tag>
        <n-tag size="small" :type="fund.fundBasic.netGrowth3 > 0 ? 'error' : 'success'" :bordered="false" v-if="fund.fundBasic.netGrowth3">近三月：{{ fund.fundBasic.netGrowth3 }}%</n-tag>
        <n-tag size="small" :type="fund.fundBasic.netGrowth6 > 0 ? 'error' : 'success'" :bordered="false" v-if="fund.fundBasic.netGrowth6">近六月：{{ fund.fundBasic.netGrowth6 }}%</n-tag>
        <n-tag size="small" :type="fund.fundBasic.netGrowth12 > 0 ? 'error' : 'success'" :bordered="false" v-if="fund.fundBasic.netGrowth12">近一年：{{ fund.fundBasic.netGrowth12 }}%</n-tag>
        <n-tag size="small" :type="fund.fundBasic.netGrowth36 > 0 ? 'error' : 'success'" :bordered="false" v-if="fund.fundBasic.netGrowth36">近三年：{{ fund.fundBasic.netGrowth36 }}%</n-tag>
        <n-tag size="small" :type="fund.fundBasic.netGrowth60 > 0 ? 'error' : 'success'" :bordered="false" v-if="fund.fundBasic.netGrowth60">近五年：{{ fund.fundBasic.netGrowth60 }}%</n-tag>
        <n-tag size="small" :type="fund.fundBasic.netGrowthYTD > 0 ? 'error' : 'success'" :bordered="false" v-if="fund.fundBasic.netGrowthYTD">今年来：{{ fund.fundBasic.netGrowthYTD }}%</n-tag>
        <n-tag size="small" :type="fund.fundBasic.netGrowthAll > 0 ? 'error' : 'success'" :bordered="false">成立来：{{ fund.fundBasic.netGrowthAll }}%</n-tag>
      </n-flex>
      <template #footer>
        <n-flex justify="space-between">
          <n-tag size="small" :bordered="false" type="warning"> {{ fund.fundBasic.type }}</n-tag>
          <n-tag size="small" :bordered="false" type="info"> {{ fund.fundBasic.company }}：{{ fund.fundBasic.manager }}</n-tag>
        </n-flex>
      </template>
      <template #action>
        <n-flex justify="end">
          <n-button size="tiny" type="error" @click="showChart"> 走势 </n-button>
          <n-button size="tiny" type="warning" @click="showDetails"> 详情 </n-button>
        </n-flex>
      </template>
    </n-card>
  </n-gi>
</template>

<style scoped>
.blink-border {
  animation: blink-border 1s linear infinite;
  border: 4px solid transparent;
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
