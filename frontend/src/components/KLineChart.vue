<script setup>

import {ref, onMounted, watch} from 'vue';
import {GetStockKLine} from "../../wailsjs/go/main/App";
import * as echarts from "echarts";
import {getChartOption} from './klinechart/chartOptions.js';

const props = defineProps({
  code: String,
  name: String,
  darkTheme: Boolean,
  kDays: {
    type: Number,
    default: 14
  },
  chartHeight: {
    type: Number,
    default: 500
  }
})
const kLineChartRef = ref(null);
let chartInstance = null;

onMounted(() => {
  chartInstance = echarts.init(kLineChartRef.value);
  handleKLine(props.code,props.name)

  chartInstance.on('click',{seriesName:'日K'}, function(params) {
    //console.log("click:",params);
  });
})

watch([() => props.code, () => props.name, () => props.darkTheme], () => {
  handleKLine(props.code, props.name);
});


function  handleKLine(code,name){
  GetStockKLine(code,name,365).then(result => {
    const categoryData = [];
    const values = [];
    const volumns=[];
    for (let i = 0; i < result.length; i++) {
      let resultElement=result[i]
      categoryData.push(resultElement.day)
      let flag=resultElement.close>resultElement.open?1:-1
      values.push([
        resultElement.open,
        resultElement.close,
        resultElement.low,
        resultElement.high
      ])
      volumns.push([i,resultElement.volume/10000,flag])
    }

    drawChart(values, volumns);
  })
}

function drawChart(values, volumns) {
  if (!chartInstance) {
    chartInstance = echarts.init(kLineChartRef.value);
  }
  const option = getChartOption({
    darkTheme: props.darkTheme,
    name: props.name,
    code: props.code,
    values: values,
    volumns: volumns,
    kDays: props.kDays,
  });
  chartInstance.setOption(option);
}
</script>
<template>
  <div ref="kLineChartRef" style="width: 100%;height: auto;--wails-draggable:no-drag" :style="{height:chartHeight+'px'}" ></div>
</template>
<style>
</style>
