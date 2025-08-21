<script setup>
import {defineProps} from 'vue'

const props = defineProps({
  indexes: Object,
})

function getAreaName(code) {
  switch (code) {
    case "america":
      return "美洲"
    case "europe":
      return "欧洲"
    case "asia":
      return "亚洲"
    case "common":
      return "常用"
    case "other":
      return "其他"
  }
}
</script>

<template>
  <n-grid :cols="5" :y-gap="0">
    <n-gi v-for="(val, key) in indexes" :key="key">
      <n-list bordered>
        <template #header>
          {{ getAreaName(key) }}
        </template>
        <n-list-item v-for="item in val" :key="item.code">
          <n-grid :cols="3" :y-gap="0">
            <n-gi>
              <n-text :type="item.zdf>0?'error':'success'">
                <n-image :src="item.img" width="20"/>
                &nbsp;{{ item.name }}
              </n-text>
            </n-gi>
            <n-gi>
              <n-text :type="item.zdf>0?'error':'success'">{{ item.zxj }}</n-text>&nbsp;
              <n-text :type="item.zdf>0?'error':'success'">
                <n-number-animation :precision="2" :from="0" :to="item.zdf"/>
                %
              </n-text>
            </n-gi>
            <n-gi>
              <n-text :type="item.state === 'open' ? 'success' : 'warning'">{{
                  item.state === 'open' ? '开市' : '休市'
                }}
              </n-text>
            </n-gi>
          </n-grid>
        </n-list-item>
      </n-list>
    </n-gi>
  </n-grid>
</template>
