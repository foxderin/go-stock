<script setup>
import { NDivider, NFlex, NTable, NThead, NTr, NTh, NTbody, NTd } from 'naive-ui';
</script>

<template>
  <div>
    <n-divider title-placement="center">支持💕开源</n-divider>
    <n-flex justify="center">
      <n-table size="small" style="width: 820px">
        <n-thead>
          <n-tr>
            <n-th>赞助计划</n-th>
            <n-th>赞助等级</n-th>
            <n-th>权益说明</n-th>
          </n-tr>
        </n-thead>
        <n-tbody>
          <n-tr>
            <n-td>每月 0 RMB</n-td>
            <n-td>vip0</n-td>
            <n-td>🌟 全部功能,软件自动更新(从GitHub下载),自行解决github平台网络问题。</n-td>
          </n-tr>
          <n-tr>
            <n-td>赞助 18.8 RMB/月<br>赞助 120 RMB/年</n-td>
            <n-td>vip1</n-td>
            <n-td>💕 全部功能,软件自动更新(从CDN下载),更新快速便捷。AI配置指导，提示词参考等</n-td>
          </n-tr>
          <n-tr>
            <n-td>赞助 28.8 RMB/月<br>赞助 240 RMB/年</n-td>
            <n-td>vip2</n-td>
            <n-td>💕 vip1全部功能,赠送硅基流动AI分析服务💕</n-td>
          </n-tr>
          <n-tr>
            <n-td>每月赞助 X RMB</n-td>
            <n-td>vipX</n-td>
            <n-td>🧩 更多计划，视go-stock开源项目发展情况而定...(承接GitHub项目README广告推广💖)</n-td>
          </n-tr>
        </n-tbody>
      </n-table>
    </n-flex>
  </div>
</template>
