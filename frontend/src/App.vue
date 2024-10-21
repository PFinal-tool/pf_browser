<script setup>
import Header from './components/Header.vue'
import Sidebar from './components/Sidebar.vue'
import OpenBrowserContent from './components/OpenBrowserContent.vue'
import {ref} from "vue";
const selectedItems = ref([
  {
    name:"浏览器剪贴板",
    website:'data:text/html,%20%3Chtml%20contenteditable%3E',
  }
]);  // 保存点击的 item 数据

const handleLinkClick = (item) => {
  console.log('Clicked item:', item);
  // 检查 selectedItems 中是否已经存在该项
  const alreadySelected = selectedItems.value.some(selectedItem => selectedItem.website === item.website);

  if (!alreadySelected) {
    // 如果该项未被选中，添加到数组中
    selectedItems.value.push(item);
  } else {
    console.log('Item already selected:', item);
  }
}
</script>

<template>
  <div id="app">
  <!-- ========== HEADER ========== -->
<!--  <Header></Header>-->
  <!-- ========== END HEADER ========== -->
  <!-- Sidebar -->
  <Sidebar @link-clicked="handleLinkClick"></Sidebar>
  <!-- End Sidebar -->
  <!-- Content -->
  <OpenBrowserContent :items="selectedItems"></OpenBrowserContent>
  <!-- End Content -->
  <!-- ========== END MAIN CONTENT ========== -->
  </div>
</template>

<style>

</style>
