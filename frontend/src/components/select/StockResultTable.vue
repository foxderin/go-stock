<script setup lang="ts">
import { h, ref } from 'vue';
import { NDataTable, NText, NButton, NTag } from 'naive-ui';
import { Environment } from "../../../wailsjs/runtime";
import { OpenURL } from "../../../wailsjs/go/main/App";

const props = defineProps({
  columns: {
    type: Array,
    default: () => []
  },
  dataList: {
    type: Array,
    default: () => []
  },
  tableScrollX: {
    type: Number,
    default: 2800
  }
});

const emit = defineEmits(['follow']);

const isNumeric = (value) => {
  return !isNaN(parseFloat(value)) && isFinite(value);
};

const openCenteredWindow = (url, width, height) => {
  const left = (window.screen.width - width) / 2;
  const top = (window.screen.height - height) / 2;

  Environment().then(env => {
    switch (env.platform) {
      case 'windows':
        window.open(
            url,
            'centeredWindow',
            `width=${width},height=${height},left=${left},top=${top},location=no,menubar=no,toolbar=no,display=standalone`
        );
        break;
      default:
        OpenURL(url);
    }
  });
};

const renderCell = (value, rowData, column) => {
  if (column.key === 'SECURITY_CODE' || column.key === 'SERIAL') {
    return h(NText, { type: 'info', border: false }, { default: () => `${value}` });
  }
  if (isNumeric(value)) {
    let type = 'info';
    if (Number(value) < 0) {
      type = 'success';
    }
    if (Number(value) >= 0 && Number(value) <= 5) {
      type = 'warning';
    }
    if (Number(value) > 5) {
      type = 'error';
    }
    return h(NText, { type: type }, { default: () => `${value}` });
  } else {
    if (column.key === 'SECURITY_SHORT_NAME') {
      return h(NButton, {
        type: 'info', bordered: false, size: 'small', onClick: () => {
          openCenteredWindow(`https://quote.eastmoney.com/${rowData.MARKET_SHORT_NAME}${rowData.SECURITY_CODE}.html#fullScreenChart`, 1240, 700);
        }
      }, { default: () => `${value}` });
    } else {
      return h(NText, { type: 'info' }, { default: () => `${value}` });
    }
  }
};

const pagination = { pageSize: 10 };

</script>

<template>
  <div>
    <n-data-table
        :striped="true"
        :max-height="'calc(100vh - 150px)'"
        size="medium"
        :columns="columns"
        :data="dataList"
        :pagination="pagination"
        :scroll-x="tableScrollX"
        :render-cell="renderCell"
    />
    <div style="margin-top: -25px">共找到
      <n-tag type="info" :bordered="false">{{ dataList.length }}</n-tag>
      只股
    </div>
  </div>
</template>
