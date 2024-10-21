<script setup>
import { onMounted, ref, watch } from "vue";

// 跟踪当前选中的 tab 索引
const selectedTabIndex = ref(0);
const props = defineProps({
  items: {
    type: Array,
    default: () => []
  }
});

// 在组件挂载后打印 items
onMounted(() => {
  console.table(props.items);
  
  // 获取 props.items 的最后一个索引
  if (props.items.length > 0) {
    selectedTabIndex.value = props.items.length - 1; // 设置为最后一个元素的索引
    loadTabContent(selectedTabIndex.value);
  }
});

// 监听 items 的变化
watch(
  () => props.items,
  (newItems) => {
    console.table(newItems); // 打印新 items

    // 检查是否有新项
    if (newItems.length > 0) {
      selectedTabIndex.value = newItems.length - 1; // 更新为最后一个元素的索引
      loadTabContent(selectedTabIndex.value); // 加载新的 tab 内容
    } else {
      // 如果没有选项卡，显示欢迎组件
      iframeSrc.value = ''; // 清空 iframe 源
    }
  },
  { deep: true } // 深度监听
);

// 异步加载选项卡内容的函数
const loadTabContent = async (index) => {
  try {
    const selectedItem = props.items[index];
    if (selectedItem) { // 确保 selectedItem 存在
      iframeSrc.value = selectedItem.website; // 存储 iframe 的 URL
    } else {
      console.warn(`Item at index ${index} does not exist.`);
    }
  } catch (error) {
    console.error('Error loading tab content:', error);
  }
};

// 函数：更新选中的 tab
const selectTab = (index) => {
  selectedTabIndex.value = index;
  loadTabContent(index); // 加载新选中的 tab 的内容
};

// 函数：关闭选中的 tab
const closeTab = (index) => {
  props.items.splice(index, 1); // 从 items 中移除关闭的 tab

  // 如果关闭的 tab 是当前选中的 tab，更新选中索引
  if (selectedTabIndex.value === index) {
    selectedTabIndex.value = Math.max(0, index - 1); // 选中前一个 tab，或选中第一个 tab
    if (props.items.length > 0) {
      loadTabContent(selectedTabIndex.value); // 加载新选中的 tab 的内容
    } else {
      iframeSrc.value = ''; // 清空 iframe 源，准备显示欢迎组件
    }
  }
};

// 存储 iframe 的 URL
const iframeSrc = ref('');

</script>

<template>
  <div class="w-full lg:ps-64 grid-container">
    <div class="dark:border-neutral-700 grid-item top-row">
      <nav class="flex gap-x-1 tab" aria-label="Tabs" role="tablist" aria-orientation="horizontal">
        <button
          v-for="(item, index) in props.items"
          :key="index"
          type="button"
          :class="[
            'tab-item',
            { 'active': index === selectedTabIndex },
          ]"
          :aria-selected="index === selectedTabIndex"
          :id="'card-type-tab-item-' + (index + 1)"
          @click="selectTab(index)"
          role="tab"
        >
          {{ item.name }}
          <span class="close-btn" @click.stop="closeTab(index)">✖</span> <!-- 关闭按钮 -->
        </button>
      </nav>
    </div>
    <div class="grid-item bottom-row">
      <div 
        :id="'card-type-tab-preview-' + (selectedTabIndex + 1)" 
        role="tabpanel"
        :aria-labelledby="'card-type-tab-item-' + (selectedTabIndex + 1)" 
        class="h-full"
      >
        <iframe v-if="iframeSrc" :src="iframeSrc" width="100%" height="100%" frameborder="0"></iframe>
        <div v-else class="welcome-message">欢迎使用！请添加选项卡以开始。</div> <!-- 欢迎组件 -->
      </div>
    </div>
  </div>
</template>

<style scoped>
.grid-container {
  display: grid;
  grid-template-rows: auto 1fr; /* 第一行自适应内容高度，第二行填满剩余空间 */
  height: 100vh; /* 设置容器高度为视口高度 */
}
.tab {
  display: flex;
  background-color: #060606;
  padding: 10px 15px 0 15px;
  font-size: 14px;
  color: #fff;
}
.tab-item {
  position: relative;
  background-color: transparent;
  padding: 10px 15px;
  border-radius: 12px 12px 0 0;
  cursor: pointer;
  transition: .2s;
}
.tab-item .close-btn {
  margin-left: 8px; /* 间隔 */
  color: #fff; /* 关闭按钮颜色 */
  cursor: pointer;
  font-size: 14px; /* 关闭按钮字体大小 */
  display: inline-block; /* 确保按钮在同一行 */
}
.tab-item:hover .close-btn {
  color: red; /* 鼠标悬停时关闭按钮变为红色 */
}
.tab-item.active {
  background-color: #34363e;
  z-index: 1;
}
.welcome-message {
  text-align: center;
  font-size: 24px;
  color: #aaa; /* 欢迎消息颜色 */
  margin-top: 20px;
}
</style>
