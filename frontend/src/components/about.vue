<script setup>
import 'md-editor-v3/lib/preview.css';
import {h, onBeforeUnmount, onMounted, ref} from 'vue';
import {CheckUpdate, GetVersionInfo, GetSponsorInfo, OpenURL} from "../../wailsjs/go/main/App";
import {EventsOff, EventsOn, Environment} from "../../wailsjs/runtime";
import {NAvatar, NButton, useNotification} from "naive-ui";
import SoftwareInfo from "./about/SoftwareInfo.vue";
import SponsorPlan from "./about/SponsorPlan.vue";
import AuthorInfo from "./about/AuthorInfo.vue";
import Acknowledgement from "./about/Acknowledgement.vue";
import SupportInfo from "./about/SupportInfo.vue";

const updateLog = ref('');
const versionInfo = ref('');
const icon = ref('https://raw.githubusercontent.com/ArvinLovegood/go-stock/master/build/appicon.png');
const alipay = ref('https://github.com/ArvinLovegood/go-stock/raw/master/build/screenshot/alipay.jpg');
const wxpay = ref('https://github.com/ArvinLovegood/go-stock/raw/master/build/screenshot/wxpay.jpg');
const wxgzh = ref('https://github.com/ArvinLovegood/go-stock/raw/dev/build/screenshot/%E6%89%AB%E7%A0%81_%E6%90%9C%E7%B4%A2%E8%81%94%E5%90%88%E4%BC%A0%E6%92%AD%E6%A0%B7%E5%BC%8F-%E7%99%BD%E8%89%B2%E7%89%88.png');
const notify = useNotification();
const vipLevel = ref("");
const vipStartTime = ref("");
const vipEndTime = ref("");

onMounted(() => {
  document.title = '关于软件';
  GetVersionInfo().then((res) => {
    updateLog.value = res.content;
    versionInfo.value = res.version;
    icon.value = res.icon;
    alipay.value = res.alipay;
    wxpay.value = res.wxpay;
    wxgzh.value = res.wxgzh;

    GetSponsorInfo().then((res) => {
      vipLevel.value = res.vipLevel;
      vipStartTime.value = res.vipStartTime;
      vipEndTime.value = res.vipEndTime;
    })
  });
});

onBeforeUnmount(() => {
  notify.destroyAll();
  EventsOff("updateVersion");
});

EventsOn("updateVersion", async (msg) => {
  const githubTimeStr = msg.published_at;
  const utcDate = new Date(githubTimeStr);
  const date = new Date(utcDate.getTime());
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');
  const seconds = String(date.getSeconds()).padStart(2, '0');
  const formattedDate = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;

  notify.info({
    avatar: () =>
        h(NAvatar, {
          size: 'small',
          round: false,
          src: icon.value
        }),
    title: '发现新版本: ' + msg.tag_name,
    content: () => {
      return h('div', {
        style: {
          'text-align': 'left',
          'font-size': '14px',
        }
      }, {default: () => msg.commit?.message})
    },
    duration: 5000,
    meta: "发布时间:" + formattedDate,
    action: () => {
      return h(NButton, {
        type: 'primary',
        size: 'small',
        onClick: () => {
          Environment().then(env => {
            switch (env.platform) {
              case 'windows':
                window.open(msg.html_url);
                break;
              default :
                OpenURL(msg.html_url);
                break;
            }
          })
        }
      }, {default: () => '查看'})
    }
  });
});

const handleCheckUpdate = () => {
  CheckUpdate(1);
};
</script>

<template>
  <n-space vertical size="large" style="--wails-draggable:no-drag">
    <SoftwareInfo
        :version-info="versionInfo"
        :vip-level="vipLevel"
        :vip-end-time="vipEndTime"
        :icon="icon"
        :update-log="updateLog"
        @check-update="handleCheckUpdate"
    />
    <SponsorPlan/>
    <AuthorInfo :wxgzh="wxgzh" :alipay="alipay" :wxpay="wxpay"/>
    <Acknowledgement/>
    <SupportInfo/>
  </n-space>
</template>

<style scoped>
h1, h2 {
  margin: 0;
  padding: 6px 0;
}

p {
  margin: 2px 0;
}

ul {
  list-style-type: disc;
  padding-left: 20px;
}

a {
  color: #18a058;
  text-decoration: none;
}

a:hover {
  text-decoration: underline;
}
</style>
