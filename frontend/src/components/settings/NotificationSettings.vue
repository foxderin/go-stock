<script setup>
import { defineProps, defineEmits } from 'vue';
import { NTag } from "naive-ui";
import { h } from "vue";

const props = defineProps({
  formValue: Object,
});

const emits = defineEmits(['update:formValue', 'send-test-notice']);

function updateValue(path, value) {
  const keys = path.split('.');
  let obj = props.formValue;
  for (let i = 0; i < keys.length - 1; i++) {
    obj = obj[keys[i]];
  }
  obj[keys[keys.length - 1]] = value;
  emits('update:formValue', props.formValue);
}
</script>

<template>
  <n-card :title="() => h(NTag, { type: 'primary', bordered: false }, () => '通知设置')" size="small">
    <n-grid :cols="24" :x-gap="24" style="text-align: left">
      <n-form-item-gi :span="3" label="钉钉推送：" path="dingPush.enable">
        <n-switch :value="formValue.dingPush.enable" @update:value="updateValue('dingPush.enable', $event)"/>
      </n-form-item-gi>
      <n-form-item-gi :span="3" label="本地推送：" path="localPush.enable">
        <n-switch :value="formValue.localPush.enable" @update:value="updateValue('localPush.enable', $event)"/>
      </n-form-item-gi>
      <n-form-item-gi :span="3" label="弹幕功能：" path="enableDanmu">
        <n-switch :value="formValue.enableDanmu" @update:value="updateValue('enableDanmu', $event)"/>
      </n-form-item-gi>
      <n-form-item-gi :span="3" label="显示滚动快讯：" path="enableNews">
        <n-switch :value="formValue.enableNews" @update:value="updateValue('enableNews', $event)"/>
      </n-form-item-gi>
      <n-form-item-gi :span="3" label="市场资讯提醒：" path="enablePushNews">
        <n-switch :value="formValue.enablePushNews" @update:value="updateValue('enablePushNews', $event)"/>
      </n-form-item-gi>
      <n-form-item-gi v-if="formValue.enablePushNews" :span="4" label="只提醒红字或关注个股的新闻："
                      path="enableOnlyPushRedNews">
        <n-switch :value="formValue.enableOnlyPushRedNews"
                  @update:value="updateValue('enableOnlyPushRedNews', $event)"/>
      </n-form-item-gi>

      <n-form-item-gi :span="22" v-if="formValue.dingPush.enable" label="钉钉机器人接口地址："
                      path="dingPush.dingRobot">
        <n-input placeholder="请输入钉钉机器人接口地址" :value="formValue.dingPush.dingRobot"
                 @update:value="updateValue('dingPush.dingRobot', $event)"/>
        <n-button type="primary" @click="$emit('send-test-notice')">发送测试通知</n-button>
      </n-form-item-gi>
    </n-grid>
  </n-card>
</template>
