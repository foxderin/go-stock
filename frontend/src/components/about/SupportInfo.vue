3456789<script setup>
import { NDivider, NTable, NThead, NTr, NTh, NTbody, NTd } from 'naive-ui';
</script>

<template>
  <div>
    <n-divider title-placement="center">关于版权和技术支持申明</n-divider>
    <div style="justify-self: center;text-align: left">
      <p style="color: #FAA04A">如有问题，请先查看项目文档，如果问题依然存在，请优先加群（491605333）咨询。</p>
      <p>
        如需软件商业授权或定制开发，请联系作者微信(备注 商业咨询)：ArvinLovegood
      </p>
      <n-divider/>
      <p>
        本软件基于开源技术构建，使用Wails、NaiveUI、Vue等开源项目。技术上如有问题，可以先向对应的开源社区请求帮助。
      </p>
      <p>
        开源不易，本人精力和时间有限，如确实需要一对一技术支持，<i style="color: crimson">请先赞助！</i>联系微信(备注 技术支持)：ArvinLovegood
      </p>
      <p style="color: #FAA04A">*加微信或者QQ时，请先备注或留言需求(如：<a href="#support">技术支持</a>，功能建议，商业咨询等，否则会被忽略)</p>
      <n-table id="support">
        <n-thead>
          <n-tr>
            <n-th>技术支持方式</n-th>
            <n-th>赞助(元)</n-th>
          </n-tr>
        </n-thead>
        <n-tbody>
          <n-tr>
            <n-td>
              加 QQ：506808970，微信：ArvinLovegood
            </n-td>
            <n-td>
              100/次
            </n-td>
          </n-tr>
          <n-tr>
            <n-td>
              长期技术支持（不限次数，新功能优先体验等）
            </n-td>
            <n-td>
              5000
            </n-td>
          </n-tr>
        </n-tbody>
      </n-table>
    </div>
  </div>
</template>
