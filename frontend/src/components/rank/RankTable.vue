<script setup lang="ts">
import { NTable, NThead, NTbody, NTr, NTh, NTd, NTag, NButton, NPopover, NText, NIcon } from 'naive-ui';
import { ArrowDownOutline } from "@vicons/ionicons5";
import KLineChart from "../KLineChart.vue";
import MoneyTrend from "../moneyTrend.vue";

defineProps({
  lhbList: {
    type: Array,
    default: () => []
  }
});
</script>

<template>
  <n-table :single-line="false" striped>
    <n-thead>
      <n-tr>
        <n-th>代码</n-th>
        <n-th width="60px">名称</n-th>
        <n-th>收盘价</n-th>
        <n-th width="60px">涨跌幅</n-th>
        <n-th>龙虎榜净买额(万)</n-th>
        <n-th>龙虎榜买入额(万)</n-th>
        <n-th>龙虎榜卖出额(万)</n-th>
        <n-th>龙虎榜成交额(万)</n-th>
        <n-th width="60px" data-field="TURNOVERRATE">换手率<n-icon :component="ArrowDownOutline" /></n-th>
        <n-th>流通市值(亿)</n-th>
        <n-th>上榜原因</n-th>
      </n-tr>
    </n-thead>
    <n-tbody>
      <n-tr v-for="(item, index) in lhbList" :key="index">
        <n-td>
          <n-tag :bordered=false type="info">{{ item.SECUCODE.split('.')[1].toLowerCase()+item.SECUCODE.split('.')[0] }}</n-tag>
        </n-td>
        <n-td>
          <n-popover trigger="hover" placement="right">
            <template #trigger>
              <n-button tag="a" text :type="item.CHANGE_RATE > 0 ? 'error' : 'success'" :bordered=false>{{ item.SECURITY_NAME_ABBR }}</n-button>
            </template>
            <k-line-chart style="width: 800px" :code="item.SECUCODE.split('.')[1].toLowerCase()+item.SECUCODE.split('.')[0]" :chart-height="500" :name="item.SECURITY_NAME_ABBR" :k-days="20" :dark-theme="true"></k-line-chart>
          </n-popover>
        </n-td>
        <n-td>
          <n-text :type="item.CHANGE_RATE > 0 ? 'error' : 'success'">{{ item.CLOSE_PRICE }}</n-text>
        </n-td>
        <n-td>
          <n-text :type="item.CHANGE_RATE > 0 ? 'error' : 'success'">{{ (item.CHANGE_RATE).toFixed(2) }}%</n-text>
        </n-td>
        <n-td>
          <n-popover trigger="hover" placement="right">
            <template #trigger>
              <n-button tag="a" text :type="item.BILLBOARD_NET_AMT > 0 ? 'error' : 'success'" :bordered=false>{{ (item.BILLBOARD_NET_AMT / 10000).toFixed(2) }}</n-button>
            </template>
            <money-trend :code="item.SECUCODE.split('.')[1].toLowerCase()+item.SECUCODE.split('.')[0]" :name="item.SECURITY_NAME_ABBR" :days="360" :dark-theme="true" :chart-height="500" style="width: 800px"></money-trend>
          </n-popover>
        </n-td>
        <n-td>
          <n-text :type="'error'">{{ (item.BILLBOARD_BUY_AMT / 10000).toFixed(2) }}</n-text>
        </n-td>
        <n-td>
          <n-text :type="'success'">{{ (item.BILLBOARD_SELL_AMT / 10000).toFixed(2) }}</n-text>
        </n-td>
        <n-td>
          <n-text :type="'info'">{{ (item.BILLBOARD_DEAL_AMT / 10000).toFixed(2) }}</n-text>
        </n-td>
        <n-td>
          <n-text :type="'info'">{{ (item.TURNOVERRATE).toFixed(2) }}%</n-text>
        </n-td>
        <n-td>
          <n-text :type="'info'">{{ (item.FREE_MARKET_CAP / 100000000).toFixed(2) }}</n-text>
        </n-td>
        <n-td>
          <n-text :type="'info'">{{ item.EXPLANATION }}</n-text>
        </n-td>
      </n-tr>
    </n-tbody>
  </n-table>
</template>
