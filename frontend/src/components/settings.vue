<script setup>
import {h, onBeforeUnmount, onMounted, ref} from "vue";
import {
  AddPrompt, DelPrompt,
  ExportConfig,
  GetConfig,
  GetPromptTemplates,
  SendDingDingMessageByType,
  UpdateConfig, CheckSponsorCode
} from "../../wailsjs/go/main/App";
import {useMessage} from "naive-ui";
import {data} from "../../wailsjs/go/models";
import {EventsEmit} from "../../wailsjs/runtime";
import BasicSettings from "./settings/BasicSettings.vue";
import NotificationSettings from "./settings/NotificationSettings.vue";
import AISettings from "./settings/AISettings.vue";
import PromptModal from "./settings/PromptModal.vue";

const message = useMessage()

const formRef = ref(null)
const formValue = ref({
  ID: 1,
  tushareToken: '',
  dingPush: {
    enable: false,
    dingRobot: ''
  },
  localPush: {
    enable: true,
  },
  updateBasicInfoOnStart: false,
  refreshInterval: 1,
  openAI: {
    enable: false,
    aiConfigs: [], // AI配置列表
    prompt: "",
    questionTemplate: "{{stockName}}分析和总结",
    crawlTimeOut: 30,
    kDays: 30,
  },
  enableDanmu: false,
  browserPath: '',
  enableNews: false,
  darkTheme: true,
  enableFund: false,
  enablePushNews: false,
  enableOnlyPushRedNews: false,
  sponsorCode: "",
  httpProxy:"",
  httpProxyEnabled:false,
})

const promptTemplates = ref([])
onMounted(() => {
  GetConfig().then(res => {
    formValue.value.ID = res.ID
    formValue.value.tushareToken = res.tushareToken
    formValue.value.dingPush = {
      enable: res.dingPushEnable,
      dingRobot: res.dingRobot
    }
    formValue.value.localPush = {
      enable: res.localPushEnable,
    }
    formValue.value.updateBasicInfoOnStart = res.updateBasicInfoOnStart
    formValue.value.refreshInterval = res.refreshInterval
    // 加载AI配置
    formValue.value.openAI = {
      enable: res.openAiEnable,
      aiConfigs: res.aiConfigs || [],
      prompt: res.prompt,
      questionTemplate: res.questionTemplate ? res.questionTemplate : '{{stockName}}分析和总结',
      crawlTimeOut: res.crawlTimeOut,
      kDays: res.kDays,
    }


    formValue.value.enableDanmu = res.enableDanmu
    formValue.value.browserPath = res.browserPath
    formValue.value.enableNews = res.enableNews
    formValue.value.darkTheme = res.darkTheme
    formValue.value.enableFund = res.enableFund
    formValue.value.enablePushNews = res.enablePushNews
    formValue.value.enableOnlyPushRedNews = res.enableOnlyPushRedNews
    formValue.value.sponsorCode = res.sponsorCode
    formValue.value.httpProxy=res.httpProxy;
    formValue.value.httpProxyEnabled=res.httpProxyEnabled;

  })

  GetPromptTemplates("", "").then(res => {
    promptTemplates.value = res
  })
})
onBeforeUnmount(() => {
  message.destroyAll()
})

function saveConfig() {
  console.log('开始保存设置', formValue.value);
  // 构建配置时，包含aiConfigs列表
  let config = new data.SettingConfig({
    ID: formValue.value.ID,
    dingPushEnable: formValue.value.dingPush.enable,
    dingRobot: formValue.value.dingPush.dingRobot,
    localPushEnable: formValue.value.localPush.enable,
    updateBasicInfoOnStart: formValue.value.updateBasicInfoOnStart,
    refreshInterval: formValue.value.refreshInterval,
    openAiEnable: formValue.value.openAI.enable,
    aiConfigs: formValue.value.openAI.aiConfigs,
    // 序列化aiConfigs列表以传递给后端
    tushareToken: formValue.value.tushareToken,
    prompt: formValue.value.openAI.prompt,
    questionTemplate: formValue.value.openAI.questionTemplate,
    crawlTimeOut: formValue.value.openAI.crawlTimeOut,
    kDays: formValue.value.openAI.kDays,
    enableDanmu: formValue.value.enableDanmu,
    browserPath: formValue.value.browserPath,
    enableNews: formValue.value.enableNews,
    darkTheme: formValue.value.darkTheme,
    enableFund: formValue.value.enableFund,
    enablePushNews: formValue.value.enablePushNews,
    enableOnlyPushRedNews: formValue.value.enableOnlyPushRedNews,
    sponsorCode: formValue.value.sponsorCode,
    httpProxy:formValue.value.httpProxy,
    httpProxyEnabled:formValue.value.httpProxyEnabled,
  })

  if (config.sponsorCode) {
    CheckSponsorCode(config.sponsorCode).then(res => {
      if (res.code) {
        UpdateConfig(config).then(res => {
          message.success(res)
          EventsEmit("updateSettings", config);
        })
      } else {
        message.error(res.msg)
      }
    })
  } else {
    UpdateConfig(config).then(res => {
      message.success(res)
      EventsEmit("updateSettings", config);
    })
  }
}


function getHeight() {
  return document.documentElement.clientHeight
}

function sendTestNotice() {
  let markdown = "### go-stock test\n" + new Date()
  let msg = '{' +
      '     "msgtype": "markdown",' +
      '     "markdown": {' +
      '         "title":"go-stock' + new Date() + '",' +
      '         "text": "' + markdown + '"' +
      '     },' +
      '      "at": {' +
      '          "isAtAll": true' +
      '      }' +
      ' }'

  SendDingDingMessageByType(msg, "test-" + new Date().getTime(), 1).then(res => {
    message.info(res)
  })
}

function exportConfig() {
  ExportConfig().then(res => {
    message.info(res)
  })
}

function importConfig() {
  message.info("导入配置功能暂未实现");
}


window.onerror = function (event, source, lineno, colno, error) {
  EventsEmit("frontendError", {
    page: "settings.vue",
    message: event,
    source: source,
    lineno: lineno,
    colno: colno,
    error: error ? error.stack : null
  });
  return true;
};

const showManagePromptsModal = ref(false)
const promptTypeOptions = [
  {label: "模型系统Prompt", value: '模型系统Prompt'},
  {label: "模型用户Prompt", value: '模型用户Prompt'},]
const formPromptRef = ref(null)
const formPrompt = ref({
  ID: 0,
  Name: '',
  Content: '',
  Type: '',
})

function managePrompts() {
  formPrompt.value = {ID: 0, Name: '', Content: '', Type: ''};
  showManagePromptsModal.value = true
}

function savePrompt() {
  AddPrompt(formPrompt.value).then(res => {
    message.success(res)
    GetPromptTemplates("", "").then(res => {
      promptTemplates.value = res
    })
    showManagePromptsModal.value = false
  })
}

function editPrompt(prompt) {
  formPrompt.value.ID = prompt.ID
  formPrompt.value.Name = prompt.name
  formPrompt.value.Content = prompt.content
  formPrompt.value.Type = prompt.type
  showManagePromptsModal.value = true
}

function deletePrompt(ID) {
  DelPrompt(ID).then(res => {
    message.success(res)
    GetPromptTemplates("", "").then(res => {
      promptTemplates.value = res
    })
  })
}
</script>

<template>
  <n-flex justify="left" style="text-align: left; --wails-draggable:no-drag">
    <n-form ref="formRef" :label-placement="'left'" :label-align="'left'">
      <n-space vertical size="large">
        <BasicSettings v-model:form-value="formValue" @check-sponsor-code="code => CheckSponsorCode(code).then(res => message.warning(res.msg))"/>
        <NotificationSettings v-model:form-value="formValue" @send-test-notice="sendTestNotice"/>
        <AISettings v-model:form-value="formValue" />

        <n-card>
          <n-space vertical>
            <n-space justify="center">
              <n-button type="warning" @click="managePrompts">管理提示词模板</n-button>
              <n-button type="primary" strong @click="saveConfig">保存设置</n-button>
              <n-button type="info" @click="exportConfig">导出配置</n-button>
              <n-button type="error" @click="importConfig">导入配置</n-button>
            </n-space>

            <n-flex justify="start" style="margin-top: 10px" v-if="promptTemplates.length > 0">
              <n-tag :bordered="false" type="warning">提示词模板:</n-tag>
              <n-tag size="medium" secondary v-for="prompt in promptTemplates" closable
                     @close="deletePrompt(prompt.ID)" @click="editPrompt(prompt)" :title="prompt.content"
                     :type="prompt.type === '模型系统Prompt' ? 'success' : 'info'" :bordered="false">{{
                  prompt.name
                }}
              </n-tag>
            </n-flex>
          </n-space>
        </n-card>
      </n-space>
    </n-form>
  </n-flex>

  <PromptModal v-model:show="showManagePromptsModal" :prompt="formPrompt" :prompt-types="promptTypeOptions"
               @save="savePrompt" @cancel="showManagePromptsModal = false"/>
</template>

<style scoped>
.cardHeaderClass {
  font-size: 16px;
  font-weight: bold;
  color: red;
}
</style>
