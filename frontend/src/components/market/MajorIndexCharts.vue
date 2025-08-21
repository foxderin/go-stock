<script setup>
import {defineProps} from 'vue'
import KLineChart from "../KLineChart.vue";

const props = defineProps({
  panelHeight: Number,
  darkTheme: Boolean,
})

const indexes = [
  {code: 'sh000001', name: '上证指数'},
  {code: 'sz399001', name: '深证成指'},
  {code: 'sz399006', name: '创业板指'},
  {code: 'hkHSI', name: '恒生指数'},
  {code: 'us.IXIC', name: '纳斯达克'},
  {code: 'us.DJI', name: '道琼斯'},
  {code: 'us.INX', name: '标普500'},
]
</script>

<template>
  <n-tabs type="segment" animated>
    <n-tab-pane v-for="index in indexes" :key="index.code" :name="index.name" :tab="index.name">
      <k-line-chart :code="index.code" :chart-height="panelHeight" :name="index.name" :k-days="20"
                    :dark-theme="darkTheme"></k-line-chart>
    </n-tab-pane>
  </n-tabs>
</template>
