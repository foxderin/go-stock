<script setup>
import { defineProps, defineEmits, ref } from 'vue';

const props = defineProps({
  show: Boolean,
  prompt: Object,
  promptTypes: Array,
});

const emits = defineEmits(['update:show', 'save', 'cancel']);

const formPromptRef = ref(null);

function savePrompt() {
  emits('save', props.prompt);
}

function cancel() {
  emits('cancel');
}
</script>

<template>
  <n-modal :show="show" closable :mask-closable="false" @update:show="$emit('update:show', $event)">
    <n-card style="width: 800px; height: 600px; text-align: left" :bordered="false"
            :title="(prompt.ID > 0 ? '修改' : '添加') + '提示词'" size="huge" role="dialog" aria-modal="true">
      <n-form ref="formPromptRef" :label-placement="'left'" :label-align="'left'">
        <n-form-item label="名称">
          <n-input :value="prompt.Name" @update:value="prompt.Name = $event" placeholder="请输入提示词名称"/>
        </n-form-item>
        <n-form-item label="类型">
          <n-select :value="prompt.Type" @update:value="prompt.Type = $event" :options="promptTypes"
                    placeholder="请选择提示词类型"/>
        </n-form-item>
        <n-form-item label="内容">
          <n-input :value="prompt.Content" @update:value="prompt.Content = $event" type="textarea" :show-count="true"
                   placeholder="请输入prompt" :autosize="{ minRows: 12, maxRows: 12, }"/>
        </n-form-item>
      </n-form>
      <template #footer>
        <n-flex justify="end">
          <n-button type="primary" @click="savePrompt">保存</n-button>
          <n-button type="warning" @click="cancel">取消</n-button>
        </n-flex>
      </template>
    </n-card>
  </n-modal>
</template>
