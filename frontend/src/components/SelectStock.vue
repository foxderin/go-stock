<script setup lang="ts">
import {h, onBeforeMount, ref, VNode} from 'vue'
import {SearchStock, GetHotStrategy, Follow} from "../../wailsjs/go/main/App";
import {useMessage, NButton, NFlex, NGrid, NGi, NInputGroup, NInput, NEllipsis, NText} from 'naive-ui'
import HotStrategyList from "./select/HotStrategyList.vue";
import StockResultTable from "./select/StockResultTable.vue";

const message = useMessage()
const search = ref('')
const columns = ref<any[]>([])
const dataList = ref<any[]>([])
const hotStrategy = ref<any[]>([])
const traceInfo = ref('')
const tableScrollX = ref(2800) // 默认滚动宽度

// 计算表格总宽度
function calculateTableWidth(cols) {
  let totalWidth = 0;

  cols.forEach(col => {
    if (col.children && col.children.length > 0) {
      // 有子列的情况
      let childrenWidth = 0;
      col.children.forEach(child => {
        childrenWidth += child.width || child.minWidth || 100;
      });
      // 取标题列宽度和子列总宽度的较大值
      totalWidth += Math.max(col.width || col.minWidth || 200, childrenWidth);
    } else {
      // 没有子列的情况
      totalWidth += col.width || col.minWidth || 120;
    }
  });

  // 加上操作列的宽度
  totalWidth += 100;

  return Math.max(totalWidth, 1200); // 最小宽度1200
}

function performSearch() {
  if (!search.value) {
    message.warning('请输入选股指标或者要求')
    return
  }

  const loading = message.loading("正在获取选股数据...", {duration: 0});
  SearchStock(search.value).then(res => {
    loading.destroy()
    if (res.code == 100) {
      traceInfo.value = res.data.traceInfo.showText
      const newColumns: any[] = res.data.result.columns.filter(item => !item.hiddenNeed && (item.title != "市场码" && item.title != "市场简称")).map(item => {
        if (item.children) {
          return {
            title: item.title + (item.unit ? '[' + item.unit + ']' : ''),
            key: item.key,
            resizable: true,
            minWidth: 200,
            ellipsis: {
              tooltip: true
            },
            children: item.children.filter(child => !child.hiddenNeed).map(child => {
              return {
                title: child.dateMsg,
                key: child.key,
                minWidth: 100,
                resizable: true,
                ellipsis: {
                  tooltip: true
                },
                sorter: (row1, row2) => {
                  if (isNumeric(row1[child.key]) && isNumeric(row2[child.key])) {
                    return row1[child.key] - row2[child.key];
                  } else {
                    return 'default'
                  }
                },
              }
            })
          }
        } else {
          return {
            title: item.title + (item.unit ? '[' + item.unit + ']' : ''),
            key: item.key,
            resizable: true,
            minWidth: 120,
            ellipsis: {
              tooltip: true
            },
            sorter: (row1, row2) => {
              if (isNumeric(row1[item.key]) && isNumeric(row2[item.key])) {
                return row1[item.key] - row2[item.key];
              } else {
                return 'default'
              }
            },
          }
        }
      });
      newColumns.push({
        title: '操作',
        key: 'actions',
        width: 80,
        fixed: 'right', // 固定在右侧
        render: (row: any): VNode => {
          return h(
              NButton,
              {
                strong: true,
                tertiary: true,
                size: 'small',
                type: 'warning', // 橙色按钮
                style: 'font-size: 14px; padding: 0 10px;', // 稍微大一点的按钮
                onClick: () => handleFollow(row)
              },
              { default: () => '关注' }
          )
        }
      });
      columns.value = newColumns;
      dataList.value = res.data.result.dataList
      tableScrollX.value = calculateTableWidth(columns.value);
    } else {
      message.error(res.msg)
    }
  }).catch(err => {
    message.error(err)
  })
}

function handleFollow(row: any) {
  let code = row.MARKET_SHORT_NAME.toLowerCase() + row.SECURITY_CODE
  Follow(code).then(result => {
    if (result === "关注成功") {
      message.success(result)
    } else {
      message.error(result)
    }
  });
}

function isNumeric(value: any) {
  return !isNaN(parseFloat(value)) && isFinite(value);
}

onBeforeMount(() => {
  GetHotStrategy().then(res => {
    if (res.code == 1) {
      hotStrategy.value = res.data
      if (hotStrategy.value.length > 0) {
        search.value = hotStrategy.value[0].question
        performSearch()
      }
    }
  }).catch(err => {
    message.error(err)
  })
})

function handleHotStrategySearch(question: string) {
  search.value = question
  performSearch()
}
</script>

<template>
  <n-grid :cols="24" style="max-height: calc(100vh - 165px)">
    <n-gi :span="4">
      <HotStrategyList :hot-strategy="hotStrategy" @search="handleHotStrategySearch"/>
    </n-gi>
    <n-gi :span="20">
      <n-flex style="--wails-draggable:no-drag">
        <n-input-group style="text-align: left">
          <n-input :rows="1" clearable v-model:value="search" placeholder="请输入选股指标或者要求"/>
          <n-button type="primary" @click="performSearch">搜索A股</n-button>
        </n-input-group>
      </n-flex>
      <n-flex justify="start" v-if="traceInfo" style="margin: 5px 0;--wails-draggable:no-drag">
        <n-ellipsis line-clamp="1" :tooltip="true">
          <n-text type="info" :bordered="false">选股条件：</n-text>
          <n-text type="warning" :bordered="true">{{ traceInfo }}</n-text>
          <template #tooltip>
            <div style="text-align: center;max-width: 580px">
              <n-text type="warning">{{ traceInfo }}</n-text>
            </div>
          </template>
        </n-ellipsis>
      </n-flex>
      <StockResultTable :columns="columns" :data-list="dataList" :table-scroll-x="tableScrollX" @follow="handleFollow"/>
    </n-gi>
  </n-grid>
</template>

<style scoped>
</style>