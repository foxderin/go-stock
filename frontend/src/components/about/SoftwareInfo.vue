<script setup>
import { NCard, NDivider, NSpace, NImage, NBadge, NGradientText, NTag, NButton } from 'naive-ui';

defineProps({
  versionInfo: String,
  vipLevel: String,
  vipEndTime: String,
  icon: String,
  updateLog: String,
});

const emit = defineEmits(['check-update']);
</script>

<template>
  <n-card size="large">
    <n-divider title-placement="center">关于软件</n-divider>
    <n-space vertical>
      <n-image width="100" :src="icon" />
      <h1>
        <n-badge v-if="!vipLevel" :value="versionInfo" :offset="[50, 10]" type="success">
          <n-gradient-text type="info" :size="50">go-stock</n-gradient-text>
        </n-badge>
        <n-badge v-if="vipLevel" :value="versionInfo" :offset="[50, 10]" type="success">
          <n-gradient-text type="warning" :size="50">go-stock</n-gradient-text>
          <n-tag :bordered="false" size="small" type="warning">VIP{{ vipLevel }}</n-tag>
        </n-badge>
      </h1>
      <n-gradient-text type="warning" v-if="vipLevel">vip到期时间：{{ vipEndTime }}</n-gradient-text>
      <n-button size="tiny" @click="emit('check-update')" type="info" tertiary>检查更新</n-button>
      <div style="justify-self: center;text-align: left">
        <p>自选股行情实时监控，基于Wails和NaiveUI构建的AI赋能股票分析工具</p>
        <p>目前已支持A股，港股，美股，未来计划加入基金，ETF等支持</p>
        <p>支持DeepSeek，OpenAI， Ollama，LMStudio，AnythingLLM，<a href="https://cloud.siliconflow.cn/i/foufCerk" target="_blank">硅基流动</a>，<a href="https://www.volcengine.com/experience/ark?utm_term=202502dsinvite&ac=DSASUQY5&rc=IJSE43PZ" target="_blank">火山方舟</a>，阿里云百炼等平台或模型</p>
        <p>
          <i style="color: crimson">本软件仅供学习研究目的，AI分析结果仅供参考，本软件不提供任何投资建议或决策，风险自担！</i>
        </p>
        <p>
          欢迎点赞GitHub：<a href="https://github.com/ArvinLovegood/go-stock" target="_blank">go-stock</a><n-divider vertical />
          <a href="https://github.com/ArvinLovegood/go-stock" target="_blank">GitHub</a><n-divider vertical />
          <a href="https://github.com/ArvinLovegood/go-stock/issues" target="_blank">Issues</a><n-divider vertical />
          <a href="https://github.com/ArvinLovegood/go-stock/releases" target="_blank">Releases</a><n-divider vertical />
        </p>
        <p v-if="updateLog">更新说明：{{ updateLog }}</p>
        <p>项目社区：<a href="https://go-stock.sparkmemory.top/" target="_blank">https://go-stock.sparkmemory.top/</a></p>
        <p>QQ交流群：<a href="http://qm.qq.com/cgi-bin/qm/qr?_wv=1027&k=0YQ8qD3exahsD4YLNhzQTWe5ssstWC89&authKey=usOMMRFtIQDC%2FYcatHYapcxQbJ7PwXPHK9OypTXWzNjAq%2FRVvQu9bj2lRgb%2BSZ3p&noverify=0&group_code=491605333" target="_blank">491605333</a></p>
      </div>
    </n-space>
  </n-card>
</template>
