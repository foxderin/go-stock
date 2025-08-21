<script setup>
import {defineProps} from 'vue'
import KLineChart from "../KLineChart.vue";

const props = defineProps({
  panelHeight: Number,
  darkTheme: Boolean,
})

const indexes = [
  {code: 'hkHSTECH', name: '恒生科技指数'},
  {code: 'sh000688', name: '科创50'},
  {code: 'sh000685', name: '科创芯片'},
  {code: 'sz399437', name: '证券龙头'},
  {code: 'sz399437', name: '高端装备'},
  {code: 'sz399986', name: '中证银行'},
  {code: 'sh000037', name: '上证医药'},
  {code: 'sh000300', name: '沪深300'},
  {code: 'sh000016', name: '上证50'},
  {code: 'sh000510', name: '中证A500'},
  {code: 'sh000852', name: '中证1000'},
  {code: 'sz399997', name: '中证白酒'},
  {code: 'usYINN.AM', name: '富时中国三倍做多'},
  {code: 'usUVXY.AM', name: 'VIX恐慌指数'},
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
