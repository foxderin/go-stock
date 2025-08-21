<script setup lang="ts">
import {onMounted, ref} from "vue";
import {GetStockMoneyTrendByDay} from "../../wailsjs/go/main/App";
import MoneyChart from "./money/MoneyChart.vue";

const {code, name, darkTheme, days, chartHeight} = defineProps({
  code: {
    type: String,
    default: ''
  },
  name: {
    type: String,
    default: ''
  },
  days: {
    type: Number,
    default: 14
  },
  chartHeight: {
    type: Number,
    default: 500
  },
  darkTheme: {
    type: Boolean,
    default: false
  }
})

const chartData = ref<{
  categoryData: string[];
  netamount_values: string[];
  r0_net_values: string[];
  trades_values: number[];
  volume: string[];
  min: number;
  max: number;
} | null>(null);

const fetchData = (stockCode: string, numDays: number) => {
  GetStockMoneyTrendByDay(stockCode, numDays).then((result: any[]) => {
    const categoryData: string[] = [];
    const netamount_values: string[] = [];
    const r0_net_values: string[] = [];
    const trades_values: number[] = [];
    let volume: string[] = [];

    let min: number = 0;
    let max: number = 0;
    for (let i = 0; i < result.length; i++) {
      let resultElement = result[i]
      categoryData.push(resultElement.opendate)
      let netamount = (resultElement.netamount / 10000).toFixed(2);
      netamount_values.push(netamount)
      let price = Number(resultElement.trade);
      trades_values.push(price)
      r0_net_values.push((resultElement.r0_net / 10000).toFixed(2))

      if (min === 0 || min > price) {
        min = price
      }
      if (max < price) {
        max = price
      }

      if (i > 0) {
        let b = (Number(result[i].netamount) + Number(result[i - 1].netamount)) / 10000
        volume.push(b.toFixed(2))
      } else {
        volume.push((Number(result[i].netamount) / 10000).toFixed(2))
      }
    }
    chartData.value = {
      categoryData,
      netamount_values,
      r0_net_values,
      trades_values,
      volume,
      min,
      max
    };
  })
}

onMounted(() => {
  fetchData(code, days)
})
</script>

<template>
  <MoneyChart v-if="chartData" :chart-data="chartData" :dark-theme="darkTheme" :chart-height="chartHeight" :name="name" />
</template>

<style scoped>

</style>