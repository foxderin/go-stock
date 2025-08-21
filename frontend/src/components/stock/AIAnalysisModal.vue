<template>
  <n-modal transform-origin="center" v-model:show="show" preset="card" style="width: 800px;"
           :title="'['+stockName+']AI分析'">
    <n-spin size="small" :show="loading">
      <MdEditor v-if="enableEditor" :toolbars="toolbars" ref="mdEditorRef" style="height: 440px;text-align: left"
                :modelValue="aiResult" :theme="theme">
      </MdEditor>
      <MdPreview v-if="!enableEditor" ref="mdPreviewRef" style="height: 440px;text-align: left"
                 :modelValue="aiResult" :theme="theme"/>
    </n-spin>
    <template #footer>
      <n-flex justify="space-between" ref="tipsRef">
        <n-text type="info" v-if="time">
          <n-tag v-if="modelName" type="warning" round :title="chatId" :bordered="false">
            {{ modelName }}
          </n-tag>
          {{ time }}
        </n-text>
        <n-text type="error">*AI分析结果仅供参考，请以实际行情为准。投资需谨慎，风险自担。</n-text>
      </n-flex>
    </template>
    <template #action>
      <n-flex justify="left" style="margin-bottom: 10px">
        <n-switch v-model:value="enableTools" :round="false">
          <template #checked>
            启用AI函数工具调用
          </template>
          <template #unchecked>
            不启用AI函数工具调用
          </template>
        </n-switch>
        <n-gradient-text type="error" style="margin-left: 10px">
          *AI函数工具调用可以增强AI获取数据的能力,但会消耗更多tokens。
        </n-gradient-text>
      </n-flex>
      <n-flex justify="space-between" style="margin-bottom: 10px">
        <n-select style="width: 31%" v-model:value="aiConfigId" label-field="name" value-field="ID"
                  :options="aiConfigs" placeholder="请选择AI模型服务配置"/>
        <n-select style="width: 31%" v-model:value="sysPromptId" label-field="name" value-field="ID"
                  :options="sysPromptOptions" placeholder="请选择系统提示词"/>
        <n-select style="width: 31%" v-model:value="question" label-field="name" value-field="content"
                  :options="userPromptOptions" placeholder="请选择用户提示词"/>
      </n-flex>
      <n-flex justify="right">
        <n-input v-model:value="question" style="text-align: left" clearable
                 type="textarea"
                 :show-count="true"
                 placeholder="请输入您的问题:例如{{stockName}}[{{stockCode}}]分析和总结"
                 :autosize="{
              minRows: 2,
              maxRows: 5
            }"
        />
        <n-button size="tiny" type="warning" @click="startAnalysis">开始AI分析</n-button>
        <n-button size="tiny" type="info" @click="saveAsImage">保存为图片</n-button>
        <n-button size="tiny" type="success" @click="copyToClipboard">复制到剪切板</n-button>
        <n-button size="tiny" type="primary" @click="saveAsMarkdown">保存为Markdown文件</n-button>
        <n-button size="tiny" type="primary" @click="saveAsWord">保存为Word文件</n-button>
        <n-button size="tiny" type="error" @click="shareToCommunity">分享到项目社区</n-button>
      </n-flex>
    </template>
  </n-modal>
</template>

<script setup>
import { ref, defineProps, defineEmits, computed } from 'vue';
import { MdEditor, MdPreview } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';

const props = defineProps({
  show: Boolean,
  stockName: String,
  stockCode: String,
  aiResult: String,
  loading: Boolean,
  time: String,
  modelName: String,
  chatId: String,
  theme: String,
  aiConfigs: Array,
  sysPromptOptions: Array,
  userPromptOptions: Array,
});

const emit = defineEmits(['update:show', 'start-analysis', 'save-as-image', 'copy-to-clipboard', 'save-as-markdown', 'save-as-word', 'share-to-community']);

const show = computed({
  get: () => props.show,
  set: (value) => emit('update:show', value)
});

const enableEditor = ref(false);
const enableTools = ref(true);
const aiConfigId = ref(null);
const sysPromptId = ref(null);
const question = ref('');

const mdEditorRef = ref(null);
const mdPreviewRef = ref(null);
const tipsRef = ref(null);

const toolbars = [
  'revoke',
  'next',
  'save',
  'pageFullscreen',
  'fullscreen',
  'preview',
  'htmlPreview',
  'catalog',
];

const handleProgress = (progress) => {
  // console.log(progress);
};

const startAnalysis = () => {
  emit('start-analysis', {
    name: props.stockName,
    code: props.stockCode,
    aiConfigId: aiConfigId.value,
    sysPromptId: sysPromptId.value,
    question: question.value,
    enableTools: enableTools.value
  });
};

const saveAsImage = () => {
  emit('save-as-image', props.stockName, props.stockCode);
};

const copyToClipboard = () => {
  emit('copy-to-clipboard');
};

const saveAsMarkdown = () => {
  emit('save-as-markdown');
};

const saveAsWord = () => {
  emit('save-as-word');
};

const shareToCommunity = () => {
  emit('share-to-community', props.stockCode, props.stockName);
};

</script>

<style scoped>
</style>
