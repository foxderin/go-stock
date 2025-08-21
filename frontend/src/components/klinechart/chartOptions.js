import _ from 'lodash';

const upColor = '#ec0000';
const downColor = '#00da3c';

function calculateMA(dayCount, data) {
  let result = [];
  for (let i = 0, len = data.length; i < len; i++) {
    if (i < dayCount) {
      result.push('-');
      continue;
    }
    let sum = 0;
    for (let j = 0; j < dayCount; j++) {
      sum += +data[i - j][1];
    }
    result.push((sum / dayCount).toFixed(2));
  }
  return result;
}

export function getChartOption({darkTheme, name, code, values, volumns, kDays}) {
  return {
    title: {
      text: `${name} ${code}`,
      left: '20px',
      textStyle: {
        color: darkTheme ? '#ccc' : '#456',
      },
    },
    darkMode: darkTheme,
    animation: false,
    legend: {
      right: 20,
      top: 0,
      data: ['日K', 'MA5', 'MA10', 'MA20', 'MA30'],
      textStyle: {
        color: darkTheme ? '#ccc' : '#456',
      },
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'cross',
        lineStyle: {
          color: '#376df4',
          width: 2,
          opacity: 1,
        },
      },
      borderWidth: 2,
      borderColor: darkTheme ? '#456' : '#ccc',
      backgroundColor: darkTheme ? '#456' : '#fff',
      padding: 10,
      textStyle: {
        color: darkTheme ? '#ccc' : '#456',
      },
      formatter: (params) => {
        const currentItemData = _.filter(params, (param) => param.seriesIndex === 0)[0].data;
        const ma5 = _.filter(params, (param) => param.seriesIndex === 1)[0].data;
        const ma10 = _.filter(params, (param) => param.seriesIndex === 2)[0].data;
        const ma20 = _.filter(params, (param) => param.seriesIndex === 3)[0].data;
        const ma30 = _.filter(params, (param) => param.seriesIndex === 4)[0].data;
        const volum = _.filter(params, (param) => param.seriesIndex === 5)[0].data;
        return `${_.filter(params, (param) => param.seriesIndex === 0)[0].name}<br/>
              开盘: ${currentItemData[1]}<br/>
              收盘: ${currentItemData[2]}<br/>
              最低: ${currentItemData[3]}<br/>
              最高: ${currentItemData[4]}<br/>
              成交量(万手): ${volum[1]}<br/>
              MA5日均线: ${ma5}<br/>
              MA10日均线: ${ma10}<br/>
              MA20日均线: ${ma20}<br/>
              MA30日均线: ${ma30}`;
      },
    },
    axisPointer: {
      link: [{xAxisIndex: 'all'}],
      label: {backgroundColor: '#888'},
    },
    visualMap: {
      show: false,
      seriesIndex: 5,
      dimension: 2,
      pieces: [
        {value: -1, color: downColor},
        {value: 1, color: upColor},
      ],
    },
    grid: [
      {left: '8%', right: '8%', height: '50%'},
      {left: '8%', right: '8%', top: '66%', height: '15%'},
    ],
    xAxis: [
      {
        type: 'category',
        data: values.map(item => item.day),
        boundaryGap: false,
        axisLine: {onZero: false},
        splitLine: {show: false},
        min: 'dataMin',
        max: 'dataMax',
        axisPointer: {z: 100},
      },
      {
        type: 'category',
        gridIndex: 1,
        data: values.map(item => item.day),
        boundaryGap: false,
        axisLine: {onZero: false},
        axisTick: {show: false},
        splitLine: {show: false},
        axisLabel: {show: false},
        min: 'dataMin',
        max: 'dataMax',
      },
    ],
    yAxis: [
      {scale: true, splitArea: {show: true}},
      {
        scale: true,
        gridIndex: 1,
        splitNumber: 2,
        axisLabel: {show: false},
        axisLine: {show: false},
        axisTick: {show: false},
        splitLine: {show: false},
      },
    ],
    dataZoom: [
      {
        type: 'inside',
        xAxisIndex: [0, 1],
        start: 100 - kDays,
        end: 100,
      },
      {
        show: true,
        xAxisIndex: [0, 1],
        type: 'slider',
        top: '85%',
        start: 100 - kDays,
        end: 100,
      },
    ],
    series: [
      {
        name: '日K',
        type: 'candlestick',
        data: values.map(item => [item.open, item.close, item.low, item.high]),
        itemStyle: {
          color: upColor,
          color0: downColor,
        },
        markPoint: {
          label: {
            formatter: (param) => (param != null ? `${param.value}` : ''),
          },
          data: [
            {name: '最高', type: 'max', valueDim: 'highest'},
            {name: '最低', type: 'min', valueDim: 'lowest'},
            {name: '平均收盘价', type: 'average', valueDim: 'close'},
          ],
          tooltip: {
            formatter: (param) => `${param.name}<br/>${param.data.coord || ''}`,
          },
        },
        markLine: {
          symbol: ['none', 'none'],
          data: [
            [
              {
                name: 'from lowest to highest',
                type: 'min',
                valueDim: 'lowest',
                symbol: 'circle',
                symbolSize: 10,
                label: {show: false},
                emphasis: {label: {show: false}},
              },
              {
                type: 'max',
                valueDim: 'highest',
                symbol: 'circle',
                symbolSize: 10,
                label: {show: false},
                emphasis: {label: {show: false}},
              },
            ],
            {name: 'min line on close', type: 'min', valueDim: 'close'},
            {name: 'max line on close', type: 'max', valueDim: 'close'},
          ],
        },
      },
      {
        name: 'MA5',
        type: 'line',
        data: calculateMA(5, values.map(item => [item.open, item.close, item.low, item.high])),
        smooth: true,
        showSymbol: false,
        lineStyle: {opacity: 0.6},
      },
      {
        name: 'MA10',
        type: 'line',
        data: calculateMA(10, values.map(item => [item.open, item.close, item.low, item.high])),
        smooth: true,
        showSymbol: false,
        lineStyle: {opacity: 0.6},
      },
      {
        name: 'MA20',
        type: 'line',
        data: calculateMA(20, values.map(item => [item.open, item.close, item.low, item.high])),
        smooth: true,
        showSymbol: false,
        lineStyle: {opacity: 0.6},
      },
      {
        name: 'MA30',
        type: 'line',
        data: calculateMA(30, values.map(item => [item.open, item.close, item.low, item.high])),
        smooth: true,
        showSymbol: false,
        lineStyle: {opacity: 0.6},
      },
      {
        name: '成交量(手)',
        type: 'bar',
        xAxisIndex: 1,
        yAxisIndex: 1,
        itemStyle: {color: '#7fbe9e'},
        data: volumns,
      },
    ],
  }
}
