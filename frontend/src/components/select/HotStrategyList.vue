<script setup lang="ts">
import { NList, NScrollbar, NListItem, NEllipsis, NTag, NText } from 'naive-ui';

defineProps({
  hotStrategy: {
    type: Array,
    default: () => []
  }
});

const emit = defineEmits(['search']);

const doSearch = (question) => {
  emit('search', question);
};
</script>

<template>
  <n-list bordered style="text-align: left;" hoverable clickable>
    <n-scrollbar style="max-height: calc(100vh - 170px);">
      <n-list-item v-for="item in hotStrategy" :key="item.rank" @click="doSearch(item.question)">
        <n-ellipsis line-clamp="1" :tooltip="true">
          <n-tag size="small" :bordered="false" type="info">#{{ item.rank }}</n-tag>
          <n-text type="warning">{{ item.question }}</n-text>
          <template #tooltip>
            <div style="text-align: center;max-width: 180px">
              <n-text type="warning">{{ item.question }}</n-text>
            </div>
          </template>
        </n-ellipsis>
      </n-list-item>
    </n-scrollbar>
  </n-list>
</template>
