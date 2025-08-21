<template>
  <n-modal v-model:show="show" :title="title" style="width: 1000px"
           :preset="'card'" @after-enter="onAfterEnter" @after-leave="onAfterLeave">
    <div ref="chartRef" style="width: 100%; height: 500px;"></div>
  </n-modal>
</template>

<script setup>
import { ref, computed, defineProps, defineEmits, watch, onBeforeUnmount } from 'vue';
import * as echarts from 'echarts';
import { GetStockMinutePriceLineData } from '../../../wailsjs/go/main/App';

const props = defineProps({
  show: Boolean,
  code: String,
  name: String,
  changePercent: Number,
  darkTheme: Boolean,
});

const emit = defineEmits(['update:show']);

const chartRef = ref(null);
const feishiInterval = ref(null);

const show = computed({
  get: () => props.show,
  set: (value) => emit('update:show', value)
});

const title = computed(() => `${props.name} ${props.changePercent}%`);

const onAfterEnter = () => {
  if (props.code && props.name) {
    showFsChart();
    feishiInterval.value = setInterval(showFsChart, 1000 * 10);
  }
};

const onAfterLeave = () => {
  clearInterval(feishiInterval.value);
};

onBeforeUnmount(() => {
  clearInterval(feishiInterval.value);
});

const showFsChart = () => {
  if (!chartRef.value) return;
  const chart = echarts.init(chartRef.value);
  GetStockMinutePriceLineData(props.code, props.name).then(result => {
    const priceData = result.priceData;
    let category = [];
    let price = [];
    let openprice = 0;
    let closeprice = 0;
    let volume = [];
    let min = 0;
    let max = 0;

    if (priceData.length === 0) return;

    openprice = priceData[0].price;
    closeprice = priceData[priceData.length - 1].price;

    for (let i = 0; i < priceData.length; i++) {
      category.push(priceData[i].time);
      price.push(priceData[i].price);
      if (min === 0 || min > priceData[i].price) min = priceData[i].price;
      if (max < priceData[i].price) max = priceData[i].price;
      if (i > 0) {
        volume.push(priceData[i].volume - priceData[i - 1].volume);
      } else {
        volume.push(priceData[i].volume);
      }
    }

    let option = {
        title: {
        subtext: `[${result.date}] 开盘:${openprice} 最新:${closeprice} 最高:${max} 最低:${min}`,
        left: 'center',
        top: '10',
        textStyle: { color: props.darkTheme ? '#ccc' : '#456' }
      },
      legend: {
        data: ['股价', '成交量'],
        textStyle: { color: props.darkTheme ? '#ccc' : '#456' },
        right: 50,
      },
      darkMode: props.darkTheme,
      tooltip: { trigger: 'axis', axisPointer: { type: 'cross' } },
      xAxis: [
        { type: 'category', data: category, axisLabel: { show: false } },
        { gridIndex: 1, type: 'category', data: category }
      ],
      yAxis: [
        { name: "股价", type: 'value', scale: true, splitLine: { show: false } },
        { name: "成交量", type: 'value', scale: true, gridIndex: 1, splitLine: { show: false } }
      ],
      grid: [
        { left: '8%', right: '8%', height: '50%' },
        { left: '8%', right: '8%', top: '70%', height: '15%' }
      ],
      series: [
        { name: "股价", data: price, type: 'line', smooth: false, showSymbol: false },
        { name: "成交量", data: volume, type: 'bar', xAxisIndex: 1, yAxisIndex: 1 }
      ]
    };
    chart.setOption(option);
  });
};

watch(() => props.show, (newVal) => {
  if (!newVal) {
    clearInterval(feishiInterval.value);
  }
});
</script>
