<script setup>
import {computed, h, onBeforeMount, onBeforeUnmount, onMounted, ref} from 'vue'
import {
  GetAIResponseResult,
  GetConfig,
  GetIndustryRank,
  GetPromptTemplates,
  GetTelegraphList,
  GlobalStockIndexes,
  ReFleshTelegraphList,
  SaveAIResponseResult,
  SaveAsMarkdown,
  ShareAnalysis,
  SummaryStockNews,
  GetAiConfigs
} from "../../wailsjs/go/main/App";
import {EventsOff, EventsOn} from "../../wailsjs/runtime";
import StockPlateInfo from "./market/StockPlateInfo.vue";
import NewsList from "./newsList.vue";
import rankTable from "./rankTable.vue";
import industryMoneyRank from "./industryMoneyRank.vue";
import StockResearchReportList from "./StockResearchReportList.vue";
import StockNoticeList from "./StockNoticeList.vue";

const route = useRoute()
const icon = ref('https://raw.githubusercontent.com/ArvinLovegood/go-stock/master/build/appicon.png');

const message = useMessage()
const notify = useNotification()
const panelHeight = ref(window.innerHeight - 240)

const telegraphList = ref([])
const sinaNewsList = ref([])
const globalStockIndexes = ref(null)
const summaryModal = ref(false)
const summaryBTN = ref(true)
const darkTheme = ref(false)
const theme = computed(() => {
  return darkTheme ? 'dark' : 'light'
})
const aiSummary = ref(``)
const aiSummaryTime = ref("")
const modelName = ref("")
const chatId = ref("")
const question = ref(``)
const aiConfigId = ref(null)
const sysPromptId = ref(null)
const loading = ref(true)
const aiConfigs = ref([])
const sysPromptOptions = ref([])
const userPromptOptions = ref([])
const promptTemplates = ref([])
const industryRanks = ref([])
const sort = ref("0")
const nowTab = ref("市场快讯")
const indexInterval = ref(null)
const indexIndustryRank = ref(null)
const stockCode= ref('')
const enableTools= ref(true)

function getIndex() {
  GlobalStockIndexes().then((res) => {
    globalStockIndexes.value = res
  })
}

onBeforeMount(() => {
  nowTab.value = route.query.name
  stockCode.value = route.query.stockCode
  GetConfig().then(result => {
    summaryBTN.value = result.openAiEnable
    darkTheme.value = result.darkTheme
  })
  GetPromptTemplates("", "").then(res => {
    promptTemplates.value = res
    sysPromptOptions.value = promptTemplates.value.filter(item => item.type === '模型系统Prompt')
    userPromptOptions.value = promptTemplates.value.filter(item => item.type === '模型用户Prompt')
  })

  GetAiConfigs().then(res=>{
    aiConfigs.value = res
    aiConfigId.value = res[0].ID
  })

  GetTelegraphList("财联社电报").then((res) => {
    telegraphList.value = res
  })
  GetTelegraphList("新浪财经").then((res) => {
    sinaNewsList.value = res
  })
  getIndex();
  industryRank();
  indexInterval.value = setInterval(() => {
    getIndex()
  }, 3000)

  indexIndustryRank.value = setInterval(() => {
    industryRank()
  }, 1000 * 10)
})

onBeforeUnmount(() => {
  EventsOff("changeMarketTab")
  EventsOff("newTelegraph")
  EventsOff("newSinaNews")
  EventsOff("summaryStockNews")
  clearInterval(indexInterval.value)
  clearInterval(indexIndustryRank.value)
})

EventsOn("changeMarketTab", async (msg) => {
  //message.info(msg.name)
  updateTab(msg.name)
})

EventsOn("newTelegraph", (data) => {
  if (data!=null) {
    for (let i = 0; i < data.length; i++) {
      telegraphList.value.pop()
    }
    telegraphList.value.unshift(...data)
  }
})
EventsOn("newSinaNews", (data) => {
  if (data!=null) {
  for (let i = 0; i < data.length; i++) {
    sinaNewsList.value.pop()
  }
  sinaNewsList.value.unshift(...data)
  }
})

//获取页面高度
window.onresize = () => {
  panelHeight.value = window.innerHeight - 240
}

function changeIndustryRankSort() {
  if (sort.value === "0") {
    sort.value = "1"
  } else {
    sort.value = "0"
  }
  industryRank()
}

function industryRank() {

  GetIndustryRank(sort.value, 150).then(result => {
    if (result.length > 0) {
      //console.log(result)
      industryRanks.value = result
    } else {
      message.info("暂无数据")
    }
  })
}

function reAiSummary() {
  aiSummary.value = ""
  summaryModal.value = true
  loading.value = true
  SummaryStockNews(question.value,aiConfigId.value, sysPromptId.value,enableTools.value)
}

function getAiSummary() {
  summaryModal.value = true
  loading.value = true
  GetAIResponseResult("市场资讯").then(result => {
    loading.value = false
    if (result.content) {
      aiSummary.value = result.content
      question.value = result.question
      loading.value = false

      const date = new Date(result.CreatedAt);
      const year = date.getFullYear();
      const month = String(date.getMonth() + 1).padStart(2, '0');
      const day = String(date.getDate()).padStart(2, '0');
      const hours = String(date.getHours()).padStart(2, '0');
      const minutes = String(date.getMinutes()).padStart(2, '0');
      const seconds = String(date.getSeconds()).padStart(2, '0');
      aiSummaryTime.value = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
      modelName.value = result.modelName
    } else {
      aiSummaryTime.value = ""
      aiSummary.value = ""
      modelName.value = ""
      //SummaryStockNews(question.value, sysPromptId.value,enableTools.value)
    }
  })
}

function updateTab(name) {
  summaryBTN.value = (name === "市场快讯");
  nowTab.value = name
}

EventsOn("summaryStockNews", async (msg) => {
  loading.value = false
  ////console.log(msg)
  if (msg === "DONE") {
    await SaveAIResponseResult("市场资讯", "市场资讯", aiSummary.value, chatId.value, question.value,aiConfigId.value)
    message.info("AI分析完成！")
    message.destroyAll()

  } else {
    if (msg.chatId) {
      chatId.value = msg.chatId
    }
    if (msg.question) {
      question.value = msg.question
    }
    if (msg.content) {
      aiSummary.value = aiSummary.value + msg.content
    }
    if (msg.extraContent) {
      aiSummary.value = aiSummary.value + msg.extraContent
    }
    if (msg.model) {
      modelName.value = msg.model
    }
    if (msg.time) {
      aiSummaryTime.value = msg.time
    }
  }
})

async function copyToClipboard() {
  try {
    await navigator.clipboard.writeText(aiSummary.value);
    message.success('分析结果已复制到剪切板');
  } catch (err) {
    message.error('复制失败: ' + err);
  }
}

function saveAsMarkdown() {
  SaveAsMarkdown('市场资讯', '市场资讯').then(result => {
    message.success(result)
  })
}

function share() {
  ShareAnalysis('市场资讯', '市场资讯').then(msg => {
    //message.info(msg)
    notify.info({
      avatar: () =>
          h(NAvatar, {
            size: 'small',
            round: false,
            src: icon.value
          }),
      title: '分享到社区',
      duration: 1000 * 30,
      content: () => {
        return h('div', {
          style: {
            'text-align': 'left',
            'font-size': '14px',
          }
        }, {default: () => msg})
      },
    })
  })
}

function ReFlesh(source) {
  //console.log("ReFlesh:", source)
  ReFleshTelegraphList(source).then(res => {
    if (source === "财联社电报") {
      telegraphList.value = res
    }
    if (source === "新浪财经") {
      sinaNewsList.value = res
    }
  })
}
</script>

<template>
  <n-card>
    <n-tabs type="line" animated @update-value="updateTab" :value="nowTab" style="--wails-draggable:no-drag">
      <n-tab-pane name="市场快讯" tab="市场快讯">
        <n-grid :cols="2" :y-gap="0">
          <n-gi>
            <news-list :newsList="telegraphList" :header-title="'财联社电报'" @update:message="ReFlesh"></news-list>
          </n-gi>
          <n-gi>
            <news-list :newsList="sinaNewsList" :header-title="'新浪财经'" @update:message="ReFlesh"></news-list>
          </n-gi>
        </n-grid>
      </n-tab-pane>
      <n-tab-pane name="全球股指" tab="全球股指">
        <n-tabs type="segment" animated>
          <n-tab-pane name="全球指数" tab="全球指数">
            <GlobalIndexList :indexes="globalStockIndexes"/>
          </n-tab-pane>
          <MajorIndexCharts :panel-height="panelHeight" :dark-theme="darkTheme"/>
        </n-tabs>
      </n-tab-pane>
      <n-tab-pane name="重大指数" tab="重大指数">
        <OtherIndexCharts :panel-height="panelHeight" :dark-theme="darkTheme"/>
      </n-tab-pane>
      <n-tab-pane name="行业排名" tab="行业排名">
        <n-tabs type="card" animated>
          <n-tab-pane name="行业涨幅排名" tab="行业涨幅排名">
            <IndustryRankTable :industry-ranks="industryRanks" :sort="sort" @change-sort="changeIndustryRankSort"/>
          </n-tab-pane>
          <n-tab-pane name="行业资金流向" tab="行业资金流向">
            <industry-money-rank :chart-height="panelHeight"></industry-money-rank>
          </n-tab-pane>
        </n-tabs>
      </n-tab-pane>
      <n-tab-pane name="个股公告" tab="个股公告">
        <StockNoticeList :stock-code="stockCode" :panel-height="panelHeight"></StockNoticeList>
      </n-tab-pane>
      <n-tab-pane name="个股研报" tab="个股研报">
        <StockResearchReportList :stock-code="stockCode" :panel-height="panelHeight"></StockResearchReportList>
      </n-tab-pane>
      <n-tab-pane name="龙虎榜" tab="龙虎榜">
        <LongTigerRankList :panel-height="panelHeight"></LongTigerRankList>
      </n-tab-pane>
      <n-tab-pane name="行业研报" tab="行业研报">
        <IndustryResearchReportList :panel-height="panelHeight"></IndustryResearchReportList>
      </n-tab-pane>
      <n-tab-pane name="热门股票" tab="热门股票">
        <HotStockList :panel-height="panelHeight"></HotStockList>
      </n-tab-pane>
      <n-tab-pane name="热门事件" tab="热门事件">
        <HotEvents :panel-height="panelHeight"></HotEvents>
      </n-tab-pane>
      <n-tab-pane name="热门题材" tab="热门题材">
        <HotTopics :panel-height="panelHeight"></HotTopics>
      </n-tab-pane>
      <n-tab-pane name="投资日历" tab="投资日历">
        <InvestCalendarTimeLine :panel-height="panelHeight"></InvestCalendarTimeLine>
      </n-tab-pane>
      <n-tab-pane name="财联社日历" tab="财联社日历">
        <ClsCalendarTimeLine :panel-height="panelHeight"></ClsCalendarTimeLine>
      </n-tab-pane>
      <n-tab-pane name="自选股" tab="自选股">
        <SelectStock :panel-height="panelHeight"></SelectStock>
      </n-tab-pane>
      <n-tab-pane name="热力图" tab="热力图">
        <stockhotmap :chart-height="panelHeight"></stockhotmap>
      </n-tab-pane>
    </n-tabs>
    <n-float-button v-if="summaryBTN" @click="getAiSummary" :right="40" :bottom="100">
      <n-icon>
        <PulseOutline/>
      </n-icon>
    </n-float-button>
    <AIAnalysisModal
        v-model:show="summaryModal"
        :dark-theme="darkTheme"
        :loading="loading"
        :ai-summary="aiSummary"
        :ai-summary-time="aiSummaryTime"
        :model-name="modelName"
        :question="question"
        :ai-configs="aiConfigs"
        :sys-prompt-options="sysPromptOptions"
        :user-prompt-options="userPromptOptions"
        v-model:ai-config-id="aiConfigId"
        v-model:sys-prompt-id="sysPromptId"
        v-model:enable-tools="enableTools"
        @re-summary="reAiSummary"
        @copy="copyToClipboard"
        @save="saveAsMarkdown"
        @share="share"
    />
  </n-card>
</template>

<style scoped>
.n-card {
  border-radius: 20px;
  --wails-draggable: drag;
}
</style>
