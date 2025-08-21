<script setup>
import {h, onBeforeMount, onBeforeUnmount, onMounted, reactive, ref} from "vue";
import {ChatboxOutline} from "@vicons/ionicons5";
import {useMessage} from "naive-ui";
import {
  FollowFund,
  GetConfig,
  GetFollowedFund,
  GetfundList,
  GetVersionInfo, OpenURL,
  UnFollowFund,
} from "../../wailsjs/go/main/App";
import {Environment} from "../../wailsjs/runtime";
import FundSearch from "./fund/FundSearch.vue";
import FundCard from "./fund/FundCard.vue";

const ws = ref(null)
const icon = ref(null)
const message = useMessage()
const modalShow = ref(false)
const data = reactive({
  modelName:"",
  chatId: "",
  question:"",
  name: "",
  code: "",
  fenshiURL:"",
  kURL:"",
  fullscreen: false,
  airesult: "",
  openAiEnable: false,
  loading: true,
  enableDanmu: false,
})

const followList=ref([])
const options=ref([])
const ticker=ref({})

onBeforeMount(()=>{
  GetConfig().then(result => {
    if (result.openAiEnable) {
      data.openAiEnable = true
    }
    if (result.enableDanmu) {
      data.enableDanmu = true
    }
  })
  GetFollowedFund().then(result => {
    followList.value = result
    //console.log("followList",followList.value)
  })
})

onMounted(() => {
  GetVersionInfo().then((res) => {
    icon.value = res.icon;
  });

  ticker.value=setInterval(() => {
    GetFollowedFund().then(result => {
      followList.value = result
      //console.log("followList",followList.value)
    })
  }, 1000*60)

})

onBeforeUnmount(() => {
  clearInterval(ticker.value)
  message.destroyAll()
})



function SendDanmu(){
  ws.value.send(data.name)
}
function AddFund(){
  FollowFund(data.code).then(result=>{
    if(result){
      message.success("关注成功")
      GetFollowedFund().then(result => {
        followList.value = result
        //console.log("followList",followList.value)
      })
    }
  })
}
function unFollow(code){
  UnFollowFund(code).then(result=>{
    if(result){
      message.success("取消关注成功")
      GetFollowedFund().then(result => {
        followList.value = result
        //console.log("followList",followList.value)
      })
    }
  })
}

function getFundList(value){
  GetfundList(value).then(result=>{
    options.value=[]
    result.forEach(item=>{
      options.value.push({
        label: item.name+" ["+item.code+"]",
        value: item.code,
      })
    })
  })
}
function onSelectFund(value){
  data.code=value
  blinkBorder(value)
}

function search(code,name){
  setTimeout(() => {
    //window.open("https://fund.eastmoney.com/"+code+".html","_blank","noreferrer,width=1000,top=100,left=100,status=no,toolbar=no,location=no,scrollbars=no")
    //window.open("https://finance.sina.com.cn/fund/quotes/"+code+"/bc.shtml","_blank","width=1000,height=800,top=100,left=100,toolbar=no,location=no")

    Environment().then(env => {
      switch (env.platform) {
        case 'windows':
          window.open("https://fund.eastmoney.com/"+code+".html","_blank","noreferrer,width=1000,top=100,left=100,status=no,toolbar=no,location=no,scrollbars=no")
          break
        default :
          OpenURL("https://fund.eastmoney.com/"+code+".html")
      }
    })

  }, 500)
}

function newchart(code,name){
  modalShow.value=true
  data.name=name
  data.code=code
  data.fenshiURL='https://image.sinajs.cn/newchart/v5/fund/nav/ss/'+code+'.gif'+"?t="+Date.now()
}

function blinkBorder(findId){
  // 获取要滚动到的元素
  const element = document.getElementById(findId);
  if (element) {
    // 滚动到该元素
    element.scrollIntoView({ behavior: 'smooth' });
    const pelement = document.getElementById(findId +'_gi');
    if(pelement){
      // 添加闪烁效果
      pelement.classList.add('blink-border');
      // 3秒后移除闪烁效果
      setTimeout(() => {
        pelement.classList.remove('blink-border');
      }, 1000*5);
    }else{
      console.error(`Element with ID ${findId}_gi not found`);
    }
  }
}
</script>

<template>
  <n-flex justify="start" >
    <n-grid :x-gap="8" :cols="3"  :y-gap="8" >
      <FundCard v-for="info in followList" :key="info.code" :fund="info"
                @unfollow="unFollow"
                @show-chart="newchart"
                @show-details="search"
      />
    </n-grid>
  </n-flex>

  <n-modal v-model:show="modalShow" :title="data.name" style="width: 400px" :preset="'card'">
    <n-image :src="data.fenshiURL"   />
  </n-modal>

  <FundSearch v-model="data.name" :options="options" :enable-danmu="data.enableDanmu"
              @search="getFundList"
              @select="onSelectFund"
              @add="AddFund"
              @send-danmu="SendDanmu"
  />
</template>

<style scoped>
/* 添加闪烁效果的CSS类 */
.blink-border {
  animation: blink-border 1s linear infinite;
  border: 4px  solid transparent;
}

@keyframes blink-border {
  0% {
    border-color: red;
  }
  50% {
    border-color: transparent;
  }
  100% {
    border-color: red;
  }
}
</style>
