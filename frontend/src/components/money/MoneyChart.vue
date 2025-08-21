<script setup lang="ts">
import {onMounted, ref, watch} from "vue";
import * as echarts from "echarts";

const props = defineProps({
  chartData: {
    type: Object,
    required: true
  },
  darkTheme: {
    type: Boolean,
    default: false
  },
  chartHeight: {
    type: Number,
    default: 500
  },
  name: {
    type: String,
    default: ''
  }
});

const lineChartRef = ref(null);
let chartInstance = null;

const renderChart = () => {
  if (!lineChartRef.value || !props.chartData) return;

  const { categoryData, netamount_values, r0_net_values, trades_values, volume, min, max } = props.chartData;
  chartInstance = echarts.init(lineChartRef.value);

  const option = {
    title: {
      text: props.name,
      left: '20px',
      textStyle: {
        color: props.darkTheme ? '#ccc' : '#456'
      }
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'cross',
        lineStyle: {
          color: '#376df4',
          width: 1,
          opacity: 1
        }
      },
      borderWidth: 2,
      borderColor: props.darkTheme ? '#456' : '#ccc',
      backgroundColor: props.darkTheme ? '#456' : '#fff',
      padding: 10,
      textStyle: {
        color: props.darkTheme ? '#ccc' : '#456'
      },
    },
    axisPointer: {
      link: [{ xAxisIndex: 'all' }],
      label: { backgroundColor: '#888' }
    },
    legend: {
      show: true,
      data: ['当日净流入', '主力当日净流入', '累计净流入', '股价'],
      selected: {
        '当日净流入': true,
        '主力当日净流入': true,
        '累计净流入': true,
        '股价': true,
      },
      textStyle: {
        color: props.darkTheme ? 'rgb(253,252,252)' : '#456'
      },
      right: 150,
    },
    dataZoom: [
      {
        type: 'inside',
        xAxisIndex: [0, 1],
        start: 86,
        end: 100
      },
      {
        show: true,
        xAxisIndex: [0, 1],
        type: 'slider',
        top: '90%',
        start: 86,
        end: 100
      }
    ],
    grid: [
      { left: '8%', right: '8%', height: '50%' },
      { left: '8%', right: '8%', top: '74%', height: '15%' },
    ],
    xAxis: [
      {
        type: 'category',
        data: categoryData,
        axisPointer: { z: 100 },
        boundaryGap: false,
        axisLine: { onZero: false },
        splitLine: { show: false },
        min: 'dataMin',
        max: 'dataMax',
      },
      {
        gridIndex: 1,
        type: 'category',
        data: categoryData,
        axisLabel: { show: false },
      }
    ],
    yAxis: [
      {
        name: '当日净流入/万',
        type: 'value',
        axisLine: { show: true },
        splitLine: { show: false },
      },
      {
        name: '股价',
        type: 'value',
        min: min - 1,
        max: max + 1,
        minInterval: 0.01,
        axisLine: { show: true },
        splitLine: { show: false },
      },
      {
        gridIndex: 1,
        name: '累计净流入/万',
        type: 'value',
        axisLine: { show: true },
        splitLine: { show: false },
      },
    ],
    series: [
      {
        yAxisIndex: 0,
        name: '当日净流入',
        data: netamount_values,
        smooth: false,
        showSymbol: false,
        lineStyle: { width: 2 },
        markPoint: {
          symbol: 'arrow',
          symbolRotate: 90,
          symbolSize: [10, 20],
          symbolOffset: [10, 0],
          itemStyle: { color: '#0d7dfc' },
          label: { position: 'right' },
          data: [{ type: 'max', name: 'Max' }, { type: 'min', name: 'Min' }]
        },
        markLine: {
          data: [{ type: 'average', name: 'Average', lineStyle: { color: '#0077ff', width: 0.5 } }]
        },
        type: 'line'
      },
      {
        yAxisIndex: 0,
        name: '主力当日净流入',
        data: r0_net_values,
        smooth: false,
        showSymbol: false,
        lineStyle: { width: 2 },
        type: 'bar'
      },
      {
        yAxisIndex: 1,
        name: '股价',
        type: 'line',
        data: trades_values,
        smooth: true,
        showSymbol: false,
        lineStyle: { width: 3 },
        markPoint: {
          symbol: 'arrow',
          symbolRotate: 90,
          symbolSize: [10, 20],
          symbolOffset: [10, 0],
          itemStyle: { color: '#f39509' },
          label: { position: 'right' },
          data: [{ type: 'max', name: 'Max' }, { type: 'min', name: 'Min' }]
        },
        markLine: {
          data: [{ type: 'average', name: 'Average', lineStyle: { color: '#f39509', width: 0.5 } }]
        },
      },
      {
        type: 'bar',
        xAxisIndex: 1,
        yAxisIndex: 2,
        name: '累计净流入',
        data: volume,
        smooth: true,
        showSymbol: false,
        lineStyle: { width: 2 },
        markPoint: {
          symbol: 'arrow',
          symbolRotate: 90,
          symbolSize: [10, 20],
          symbolOffset: [10, 0],
          label: { position: 'right' },
          data: [{ type: 'max', name: 'Max' }, { type: 'min', name: 'Min' }]
        },
      },
    ]
  };
  chartInstance.setOption(option);
};

onMounted(() => {
  renderChart();
});

watch(() => props.chartData, () => {
  renderChart();
}, { deep: true });

watch(() => props.darkTheme, () => {
  if (chartInstance) {
    chartInstance.dispose();
    renderChart();
  }
});
</script>

<template>
  <div ref="lineChartRef" style="width: 100%;height: auto;" :style="{height: chartHeight + 'px'}"></div>
</template>

<style scoped>
</style>
