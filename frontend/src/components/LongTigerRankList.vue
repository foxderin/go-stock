<script setup lang="ts">
import {onBeforeMount, ref} from 'vue'
import {LongTigerRank} from "../../wailsjs/go/main/App";
import _ from "lodash";
import {useMessage} from "naive-ui";
import RankSearchForm from "./rank/RankSearchForm.vue";
import RankTable from "./rank/RankTable.vue";

const message = useMessage()

const lhbList = ref<any[]>([])
const EXPLANATIONs = ref<any[]>([])

const today = new Date();
const year = today.getFullYear();
const month = String(today.getMonth() + 1).padStart(2, '0');
const day = String(today.getDate()).padStart(2, '0');
const formattedDate = `${year}-${month}-${day}`;

const searchForm = ref<{
  dateValue: string;
  EXPLANATION: string | null;
}>({
  dateValue: formattedDate,
  EXPLANATION: null,
})

onBeforeMount(() => {
  fetchLongTigerData(formattedDate);
})

function fetchLongTigerData(date: string) {
  if (date) {
    searchForm.value.dateValue = date;
  }

  let loading1 = message.loading("正在获取龙虎榜数据...", {
    duration: 0,
  });

  const fetchDate = (currentDate: string, retryCount = 0) => {
    if (retryCount > 7) { // 防止无限循环，最多尝试7次
      lhbList.value = [];
      EXPLANATIONs.value = [];
      loading1.destroy();
      message.info("暂无历史数据");
      return;
    }

    LongTigerRank(currentDate).then((res: any[]) => {
      if (res.length === 0) {
        const previousDate = new Date(currentDate);
        previousDate.setDate(previousDate.getDate() - 1);

        const year = previousDate.getFullYear();
        const month = String(previousDate.getMonth() + 1).padStart(2, '0');
        const day = String(previousDate.getDate()).padStart(2, '0');
        const prevFormattedDate = `${year}-${month}-${day}`;

        message.info(`当前日期 ${currentDate} 暂无数据，尝试查询前一日：${prevFormattedDate}`);

        searchForm.value.dateValue = prevFormattedDate;
        fetchDate(prevFormattedDate, retryCount + 1); // 递归调用
      } else {
        lhbList.value = res;
        loading1.destroy();
        EXPLANATIONs.value = _.uniqBy(_.map(lhbList.value, function (item) {
          return {
            label: item['EXPLANATION'],
            value: item['EXPLANATION'],
          };
        }), 'label');
      }
    }).catch(err => {
      loading1.destroy();
      message.error("获取数据失败，请重试");
      console.error(err);
    });
  };

  fetchDate(date || formattedDate);
}

function handleExplanationChange(value: string | null) {
  searchForm.value.EXPLANATION = value
  if (value) {
    LongTigerRank(searchForm.value.dateValue).then((res: any[]) => {
      lhbList.value = _.filter(res, function (o) {
        return o['EXPLANATION'] === value;
      });
      if (res.length === 0) {
        message.info("暂无数据,请切换日期")
      }
    })
  } else {
    fetchLongTigerData(searchForm.value.dateValue)
  }
}
</script>

<template>
  <RankSearchForm
      v-model="searchForm"
      :explanations="EXPLANATIONs"
      @date-change="fetchLongTigerData"
      @explanation-change="handleExplanationChange"
  />
  <RankTable :lhb-list="lhbList"/>
</template>

<style scoped>
</style>