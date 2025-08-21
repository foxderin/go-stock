<script setup>
import { defineProps, defineEmits } from 'vue';
import { NTag } from "naive-ui";
import { h } from "vue";

const props = defineProps({
  formValue: Object,
});

const emits = defineEmits(['update:formValue']);

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
  <n-card :title="() => h(NTag, { type: 'primary', bordered: false }, () => '基础设置')" size="small">
    <n-grid :cols="24" :x-gap="24" style="text-align: left">
      <n-form-item-gi :span="10" label="Tushare Token：" path="tushareToken">
        <n-input type="text" placeholder="Tushare api token" :value="formValue.tushareToken"
                 @update:value="updateValue('tushareToken', $event)" clearable/>
      </n-form-item-gi>
      <n-form-item-gi :span="4" label="启动时更新基础信息：" path="updateBasicInfoOnStart">
        <n-switch :value="formValue.updateBasicInfoOnStart"
                  @update:value="updateValue('updateBasicInfoOnStart', $event)"/>
      </n-form-item-gi>
      <n-form-item-gi :span="4" label="数据刷新间隔：" path="refreshInterval">
        <n-input-number :value="formValue.refreshInterval" @update:value="updateValue('refreshInterval', $event)"
                        placeholder="请输入数据刷新间隔(秒)">
          <template #suffix>秒</template>
        </n-input-number>
      </n-form-item-gi>
      <n-form-item-gi :span="6" label="暗黑主题：" path="darkTheme">
        <n-switch :value="formValue.darkTheme" @update:value="updateValue('darkTheme', $event)"/>
      </n-form-item-gi>
      <n-form-item-gi :span="10" label="浏览器安装路径：" path="browserPath">
        <n-input type="text" placeholder="浏览器安装路径" :value="formValue.browserPath"
                 @update:value="updateValue('browserPath', $event)" clearable/>
      </n-form-item-gi>
      <n-form-item-gi :span="3" label="指数基金：" path="enableFund">
        <n-switch :value="formValue.enableFund" @update:value="updateValue('enableFund', $event)"/>
      </n-form-item-gi>
      <n-form-item-gi :span="11" label="赞助码：" path="sponsorCode">
        <n-input-group>
          <n-input :show-count="true" placeholder="赞助码" :value="formValue.sponsorCode"
                   @update:value="updateValue('sponsorCode', $event)"/>
          <n-button type="success" secondary strong
                    @click="$emit('check-sponsor-code', formValue.sponsorCode)">验证
          </n-button>
        </n-input-group>
      </n-form-item-gi>
    </n-grid>
  </n-card>
</template>
