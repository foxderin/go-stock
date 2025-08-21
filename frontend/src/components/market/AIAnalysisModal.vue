<script setup>
import {h, ref, defineProps, defineEmits} from 'vue'
import {NAvatar, NButton, NFlex, NText, useMessage, useNotification} from "naive-ui";
import {MdPreview} from "md-editor-v3";

const props = defineProps({
  show: Boolean,
  darkTheme: Boolean,
  loading: Boolean,
  aiSummary: String,
  aiSummaryTime: String,
  modelName: String,
  question: String,
  aiConfigs: Array,
  sysPromptOptions: Array,
  userPromptOptions: Array,
  aiConfigId: [String, Number],
  sysPromptId: [String, Number],
  enableTools: Boolean,
})

const emits = defineEmits(['update:show', 're-summary', 'copy', 'save', 'share', 'update:aiConfigId', 'update:sysPromptId', 'update:enableTools'])

const message = useMessage()
const notify = useNotification()
const icon = ref('https://raw.githubusercontent.com/ArvinLovegood/go-stock/master/build/appicon.png');

function closeModal() {
  emits('update:show', false)
}

function reSummary() {
  emits('re-summary')
}

function copyToClipboard() {
  emits('copy')
}

function saveAsMarkdown() {
  emits('save')
}

function share() {
  emits('share')
}

</script>

<template>
  <n-modal :show="show" @update:show="closeModal" :mask-closable="false" style="width: 60%; --wails-draggable:drag">
    <n-card
        :title="modelName !== '' ? `由 ${modelName} 生成` : 'AI市场解读'"
        :bordered="false"
        size="huge"
        role="dialog"
        aria-modal="true"
    >
      <template #header-extra>
        <n-text depth="3" style="font-size: 12px" v-if="aiSummaryTime">
          {{ aiSummaryTime }}
        </n-text>
      </template>
      <n-spin :show="loading">
        <md-preview :theme="darkTheme ? 'dark' : 'light'" :model-value="aiSummary"/>
      </n-spin>
      <template #footer>
        <n-flex justify="space-between">
          <div>
            <n-button @click="reSummary" type="info" secondary>
              <template #icon>
                <n-icon>
                  <PulseOutline/>
                </n-icon>
              </template>
              重新生成
            </n-button>
          </div>
          <div>
            <n-button @click="copyToClipboard" type="primary" style="margin-left: 10px">复制</n-button>
            <n-button @click="saveAsMarkdown" type="primary" style="margin-left: 10px">保存</n-button>
            <n-button @click="share" type="primary" style="margin-left: 10px">分享</n-button>
          </div>
        </n-flex>
        <n-flex justify="start" style="margin-top: 10px">
          <n-form-item label="AI配置" label-placement="left">
            <n-select :value="aiConfigId"
                      @update:value="$emit('update:aiConfigId', $event)"
                      :options="aiConfigs.map(c => ({ label: c.name, value: c.ID }))"
                      style="width: 150px"/>
          </n-form-item>
          <n-form-item label="系统提示词" label-placement="left">
            <n-select :value="sysPromptId"
                      @update:value="$emit('update:sysPromptId', $event)"
                      :options="sysPromptOptions.map(p => ({ label: p.name, value: p.ID }))"
                      style="width: 200px"/>
          </n-form-item>
          <n-form-item label="启用Tools" label-placement="left">
            <n-switch :value="enableTools" @update:value="$emit('update:enableTools', $event)"/>
          </n-form-item>
        </n-flex>
      </template>
    </n-card>
  </n-modal>
</template>

<style scoped>

</style>
