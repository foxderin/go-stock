<template>
  <n-modal v-model:show="show" :title="title" style="width: 1000px" :preset="'card'" @after-enter="onAfterEnter">
    <div ref="kLineChartRef" style="width: 100%; height: 500px;"></div>
  </n-modal>
</template>

<script setup>
import { ref, defineProps, defineEmits, watch, computed } from 'vue';
import * as echarts from 'echarts';
import { GetStockKLine } from '../../../wailsjs/go/main/App';

const props = defineProps({
  show: Boolean,
  code: String,
  name: String,
  darkTheme: Boolean,
});

const emit = defineEmits(['update:show']);

const kLineChartRef = ref(null);

const show = computed({
  get: () => props.show,
  set: (value) => emit('update:show', value)
});

const title = computed(() => `${props.name} K线图`);

const onAfterEnter = () => {
  if (props.code && props.name) {
    handleKLine();
  }
};

const calculateMA = (dayCount, values) => {
  var result = [];
  for (var i = 0, len = values.length; i < len; i++) {
    if (i < dayCount) {
      result.push('-');
      continue;
    }
    var sum = 0;
    for (var j = 0; j < dayCount; j++) {
      sum += +values[i - j][1];
    }
    result.push((sum / dayCount).toFixed(2));
  }
  return result;
};

const handleKLine = () => {
  GetStockKLine(props.code, props.name, 365).then(result => {
    const chart = echarts.init(kLineChartRef.value);
    const categoryData = [];
    const values = [];
    const volumns = [];
    for (let i = 0; i < result.length; i++) {
      let resultElement = result[i];
      categoryData.push(resultElement.day);
      let flag = resultElement.close > resultElement.open ? 1 : -1;
      values.push([
        resultElement.open,
        resultElement.close,
        resultElement.low,
        resultElement.high
      ]);
      volumns.push([i, resultElement.volume / 10000, flag]);
    }

    const upColor = '#ec0000';
    const downColor = '#00da3c';

    let option = {
      darkMode: props.darkTheme,
      animation: false,
      legend: {
        bottom: 10,
        left: 'center',
        data: ['日K', 'MA5', 'MA10', 'MA20', 'MA30'],
        textStyle: {
          color: props.darkTheme ? '#ccc' : '#456'
        },
      },
      tooltip: {
        trigger: 'axis',
        axisPointer: {
          type: 'cross'
        },
      },
      axisPointer: {
        link: [{ xAxisIndex: 'all' }],
        label: { backgroundColor: '#888' }
      },
      visualMap: {
        show: false,
        seriesIndex: 5,
        dimension: 2,
        pieces: [
          { value: -1, color: downColor },
          { value: 1, color: upColor }
        ]
      },
      grid: [
        { left: '10%', right: '8%', height: '50%' },
        { left: '10%', right: '8%', top: '63%', height: '16%' }
      ],
      xAxis: [
        { type: 'category', data: categoryData, boundaryGap: false, axisLine: { onZero: false }, splitLine: { show: false }, min: 'dataMin', max: 'dataMax', axisPointer: { z: 100 } },
        { type: 'category', gridIndex: 1, data: categoryData, boundaryGap: false, axisLine: { onZero: false }, axisTick: { show: false }, splitLine: { show: false }, axisLabel: { show: false }, min: 'dataMin', max: 'dataMax' }
      ],
      yAxis: [
        { scale: true, splitArea: { show: true } },
        { scale: true, gridIndex: 1, splitNumber: 2, axisLabel: { show: false }, axisLine: { show: false }, axisTick: { show: false }, splitLine: { show: false } }
      ],
      dataZoom: [
        { type: 'inside', xAxisIndex: [0, 1], start: 86, end: 100 },
        { show: true, xAxisIndex: [0, 1], type: 'slider', top: '85%', start: 86, end: 100 }
      ],
      series: [
        { name: '日K', type: 'candlestick', data: values, itemStyle: { color: upColor, color0: downColor } },
        { name: 'MA5', type: 'line', data: calculateMA(5, values), smooth: true, showSymbol: false, lineStyle: { opacity: 0.6 } },
        { name: 'MA10', type: 'line', data: calculateMA(10, values), smooth: true, showSymbol: false, lineStyle: { opacity: 0.6 } },
        { name: 'MA20', type: 'line', data: calculateMA(20, values), smooth: true, showSymbol: false, lineStyle: { opacity: 0.6 } },
        { name: 'MA30', type: 'line', data: calculateMA(30, values), smooth: true, showSymbol: false, lineStyle: { opacity: 0.6 } },
        { name: '成交量(手)', type: 'bar', xAxisIndex: 1, yAxisIndex: 1, itemStyle: { color: '#7fbe9e' }, data: volumns }
      ]
    };
    chart.setOption(option);
  });
};

watch(() => props.show, (newVal) => {
  if (newVal) {
    // The modal is opened, but the chart DOM might not be ready yet.
    // `onAfterEnter` is a better place to initialize the chart.
  }
});
</script>
