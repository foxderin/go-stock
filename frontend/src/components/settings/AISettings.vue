<script setup>
import { defineProps, defineEmits } from 'vue';
import { NTag } from "naive-ui";
import { h } from "vue";
import { data } from "../../../wailsjs/go/models";

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

function addAiConfig() {
  const newConfigs = [...props.formValue.openAI.aiConfigs, new data.AIConfig({
    name: '',
    baseUrl: 'https://api.deepseek.com',
    apiKey: '',
    modelName: 'deepseek-chat',
    temperature: 0.1,
    maxTokens: 1024,
    timeOut: 60,
  })];
  updateValue('openAI.aiConfigs', newConfigs);
}

function removeAiConfig(index) {
  const newConfigs = props.formValue.openAI.aiConfigs.filter((_, i) => i !== index);
  updateValue('openAI.aiConfigs', newConfigs);
}
</script>

<template>
  <n-card :title="() => h(NTag, { type: 'primary', bordered: false }, () => 'AI设置')" size="small">
    <n-grid :cols="24" :x-gap="24" style="text-align: left;">
      <n-form-item-gi :span="24" label="AI诊股：" path="openAI.enable">
        <n-switch :value="formValue.openAI.enable" @update:value="updateValue('openAI.enable', $event)"/>
      </n-form-item-gi>

      <template v-if="formValue.openAI.enable">
        <n-form-item-gi :span="6" label="Crawler Timeout(秒)" title="资讯采集超时时间(秒)" path="openAI.crawlTimeOut">
          <n-input-number min="30" step="1" :value="formValue.openAI.crawlTimeOut"
                          @update:value="updateValue('openAI.crawlTimeOut', $event)"/>
        </n-form-item-gi>
        <n-form-item-gi :span="4" title="天数越多消耗tokens越多" label="日K线数据(天)" path="openAI.kDays">
          <n-input-number min="30" step="1" max="365" :value="formValue.openAI.kDays"
                          @update:value="updateValue('openAI.kDays', $event)"/>
        </n-form-item-gi>
        <n-form-item-gi :span="2" label="http代理" path="httpProxyEnabled">
          <n-switch :value="formValue.httpProxyEnabled" @update:value="updateValue('httpProxyEnabled', $event)"/>
        </n-form-item-gi>
        <n-form-item-gi :span="10" v-if="formValue.httpProxyEnabled" title="http代理地址" label="http代理地址"
                        path="httpProxy">
          <n-input type="text" placeholder="http代理地址" :value="formValue.httpProxy"
                   @update:value="updateValue('httpProxy', $event)" clearable/>
        </n-form-item-gi>

        <n-gi :span="24">
          <n-divider title-placement="left">Prompt 内容设置</n-divider>
        </n-gi>
        <n-form-item-gi :span="12" label="模型系统 Prompt" path="openAI.prompt">
          <n-input :value="formValue.openAI.prompt" @update:value="updateValue('openAI.prompt', $event)"
                   type="textarea" :show-count="true"
                   placeholder="请输入系统prompt" :autosize="{ minRows: 4, maxRows: 8 }"/>
        </n-form-item-gi>
        <n-form-item-gi :span="12" label="模型用户 Prompt" path="openAI.questionTemplate">
          <n-input :value="formValue.openAI.questionTemplate"
                   @update:value="updateValue('openAI.questionTemplate', $event)" type="textarea" :show-count="true"
                   placeholder="请输入用户prompt:例如{{stockName}}[{{stockCode}}]分析和总结"
                   :autosize="{ minRows: 4, maxRows: 8 }"/>
        </n-form-item-gi>

        <n-gi :span="24">
          <n-divider title-placement="left">AI模型服务配置</n-divider>
        </n-gi>
        <n-gi :span="24">
          <n-space vertical>
            <n-card v-for="(aiConfig, index) in formValue.openAI.aiConfigs" :key="index" :bordered="true"
                    size="small">
              <template #header>
                <n-flex justify="space-between" align="center">
                  <n-text depth="3">AI 配置 #{{ index + 1 }}</n-text>
                  <n-button type="error" size="tiny" ghost @click="removeAiConfig(index)">删除</n-button>
                </n-flex>
              </template>
              <n-grid :cols="24" :x-gap="24">
                <n-form-item-gi :span="12" label="配置名称" :path="`openAI.aiConfigs[${index}].name`">
                  <n-input type="text" placeholder="配置名称" :value="aiConfig.name"
                           @update:value="updateValue(`openAI.aiConfigs[${index}].name`, $event)" clearable/>
                </n-form-item-gi>
                <n-form-item-gi :span="12" label="接口地址" :path="`openAI.aiConfigs[${index}].baseUrl`">
                  <n-input type="text" placeholder="AI接口地址" :value="aiConfig.baseUrl"
                           @update:value="updateValue(`openAI.aiConfigs[${index}].baseUrl`, $event)" clearable/>
                </n-form-item-gi>
                <n-form-item-gi :span="12" label="令牌(apiKey)" :path="`openAI.aiConfigs[${index}].apiKey`">
                  <n-input type="password" placeholder="apiKey" :value="aiConfig.apiKey"
                           @update:value="updateValue(`openAI.aiConfigs[${index}].apiKey`, $event)" clearable
                           show-password-on="click"/>
                </n-form-item-gi>
                <n-form-item-gi :span="8" label="模型名称" :path="`openAI.aiConfigs[${index}].modelName`">
                  <n-input type="text" placeholder="AI模型名称" :value="aiConfig.modelName"
                           @update:value="updateValue(`openAI.aiConfigs[${index}].modelName`, $event)" clearable/>
                </n-form-item-gi>
                <n-form-item-gi :span="5" label="Temperature" :path="`openAI.aiConfigs[${index}].temperature`">
                  <n-input-number placeholder="temperature" :value="aiConfig.temperature"
                                  @update:value="updateValue(`openAI.aiConfigs[${index}].temperature`, $event)"
                                  :step="0.1"/>
                </n-form-item-gi>
                <n-form-item-gi :span="5" label="MaxTokens" :path="`openAI.aiConfigs[${index}].maxTokens`">
                  <n-input-number placeholder="maxTokens" :value="aiConfig.maxTokens"
                                  @update:value="updateValue(`openAI.aiConfigs[${index}].maxTokens`, $event)"/>
                </n-form-item-gi>
                <n-form-item-gi :span="5" label="Timeout(秒)" :path="`openAI.aiConfigs[${index}].timeOut`">
                  <n-input-number min="60" step="1" placeholder="超时(秒)" :value="aiConfig.timeOut"
                                  @update:value="updateValue(`openAI.aiConfigs[${index}].timeOut`, $event)"/>
                </n-form-item-gi>
              </n-grid>
            </n-card>
            <n-button type="primary" dashed @click="addAiConfig" style="width: 100%;">+ 添加AI配置</n-button>
          </n-space>
        </n-gi>
      </template>
    </n-grid>
  </n-card>
</template>
