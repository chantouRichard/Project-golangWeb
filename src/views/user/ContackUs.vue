<template>
  <div @click="handleClick" class="HomePage">
    <!-- 页面点击后显示文字的容器 -->
    <span
      v-for="(text, index) in displayedTexts"
      :key="text.id"
      :style="{ left: text.x + 'px', top: text.y + 'px' }"
      class="animated-text"
    >
      {{ text.content }}
    </span>
    <div>
      <!-- 插入的原始 HTML 代码 -->
      <div class="a">
        <div class="b">
          <a href="#">黑白羽衣</a>
          <h2>WX:ct13692969989</h2>
          <span>
            I'm a computer enthusiast,<br />I'm learning new knowledge every
            day,<br />I expect the FUTURE <br /><br /><b>Keep going!!</b>
          </span>
        </div>
        <div class="c">
          <div class="d" style="--i: 1; --w: 1.5"></div>
          <div class="d" style="--i: 2; --w: 1.6"></div>
          <div class="d" style="--i: 3; --w: 1.4"></div>
          <div class="d" style="--i: 4; --w: 1.7"></div>
          <div class="e" style="--i: 1"></div>
        </div>
        <!-- 设置二维码 -->
        <div class="f"></div>
      </div>


    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

// 设置随机出现的文字数组
const arr = [
  "在这披霜戴雪的冬日，一头扎进这沸腾的生命",
  "命运如同手中的掌纹，无论多么曲折，始终掌握在你的手里",
  "少言自寡，胜过千言万语",
  "千万不要在奋斗的年纪选择安逸",
  "除了你自己，没人会时刻在意你",
  "命运的每一个玩去，都是你走向成功的一个转折",
  "年少的期待，都会在日后被一一兑现，哪怕它们换了形态，且姗姗来迟",
  "当你不能拥有的时候，唯一能做的便是不能忘记",
  "如果命运是条孤独的河流，那么你就是你的灵魂摆渡人",
  "燃烧，使你获得最终的宁静",
  "生命的价值，在于始终坚持一个目标",
  "没办法，我喜欢她，我对得起自己的喜欢",
  "不求苍天俯就我的美意，但求永远恣意挥洒",
  "每个人心底都有一座坟墓，是用来埋葬所爱的人",
  "黑夜无论多么漫长，白昼总会到来",
  "无论谁，领先一步，都是暂时的",
  "这是黄昏的太阳，我却当做是黎明的曙光",
  "人生处处有诱惑，贪欲者自上钩",
  "认知自己的无知是最大的智慧",
  "一路曲折，换来的是生命的成长",
  "既然已经伤害了过去，就不要再辜负将来",
  "人类最大的勇气就是豁出去的心",
  "人成熟的标志在于，该动脑时，不在动情",
  "对美好的追求，对残缺的接纳",
  "青春时光转瞬即逝",
  "生活不是一种刁难，而是一种雕刻",
  "永远不要在别人面前调侃你的理想，你为他付出的是生命",
  "凡是都有偶然的凑巧，结果却如宿命般的必然",
  "总之岁月漫长，所有值得等待",
];

// 存储显示的文字信息
const displayedTexts = ref([]);

// 点击处理函数
const handleClick = (event) => {
  const randomText = arr[Math.floor(Math.random() * arr.length)];

  // 添加到显示文字数组中
  displayedTexts.value.push({
    content: randomText,
    x: event.clientX - 240,
    y: event.clientY - 100,
    id: Date.now(), // 生成唯一的ID
  });

  // 2秒后移除显示的文字
  setTimeout(() => {
    displayedTexts.value.shift();
  }, 2000);
};

import { onMounted, onBeforeUnmount } from "vue";

// 生命周期钩子：组件挂载时
onMounted(() => {
  //   document.addEventListener('click', handleClick);
});

// 生命周期钩子：组件卸载前
onBeforeUnmount(() => {
  document.removeEventListener("click", handleClick);
});
</script>

<style scoped>
.HomePage {
  width: 100%;
  height: 650px;
  background-color: rgb(35, 35, 35);
  position: relative;
  overflow: hidden; /* 防止溢出 */
}

.animated-text {
  position: absolute;
  font-size: 20px;
  color: blanchedalmond;
  user-select: none;
  cursor: default;
  animation: moveUp 2s ease forwards; /* 动画 */
}

/* 定义文字移动和消失的动画 */
@keyframes moveUp {
  0% {
    transform: translateY(0);
    opacity: 1;
  }
  50% {
    transform: translateY(-50px);
    opacity: 1;
  }
  100% {
    transform: translateY(-100px);
    opacity: 0;
  }
}

.a {
  position: relative;
  width: 700px;
  height: 400px;
  /* border: #fff 10px solid;  */
  /* background-color: rgb(120,140,200); */
  margin-left: auto;
  margin-right: 100px;
  top: 120px;
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0px 0px 20px rgba(255, 204, 0, 0.445);
  filter: drop-shadow(4px 4px 5px rgba(0, 0, 0, 0.164));
}
.b {
  position: absolute;
  width: 200px;
  height: 300px;
  left: 0;
  margin: 75px 50px;
  transition: 1s;
}
.b a {
  text-decoration: none;
  color: #fff;
  font: 900 28px "";
}
.b h2 {
  /* 鼠标放开时的动画，第一个值是动画的过渡时间
            第二个值是延迟一秒后执行动画 */
  transition: 0.5s 1s;
  opacity: 0;
  color: rgb(30, 210, 200);
}
.b span {
  transition: 0.5s 1s;
  color: #fff;
  font: 500 15px "";
  position: absolute;
  top: 70px;
}
.c {
  position: absolute;
  top: -130px;
  right: -240px;
}
.d,
.e {
  position: absolute;
  right: calc(var(--i) * 100px);
  width: calc(var(--w) * 40px);
  height: 500px;
  overflow: hidden;
  border-radius: 100px;
  transform: rotateZ(220deg) translate(0, 0);
  background: rgb(240, 220, 150);
  transition: 0.5s 0.5s;
}
.d:nth-child(2) {
  background: rgb(240, 190, 230);
}
.e {
  left: -470px;
  top: -140px;
  width: 70px;
  background-color: rgb(90, 220, 150);
}
.a:hover .c div {
  /* 设置延迟动画 */
  transition: 0.5s calc(var(--i) * 0.1s);
  /* 移动的轨迹 */
  transform: rotateZ(220deg) translate(-200px, 400px);
}
.a:hover .b {
  transition: 1s 0.5s;
  left: 370px;
}
.a:hover .b span {
  transition: 1s 0.5s;
  top: 105px;
}
.a:hover .b h2 {
  transition: 1s 0.5s;
  opacity: 1;
}
.f {
  width: 250px;
  height: 250px;
  position: absolute;
  background-image: url("/images/MyCard.png");
  background-size: cover;
  margin: 70px;
  opacity: 0;
  transition: 0.5s;
}
.a:hover .f {
  transition: 1s 1.3s;
  opacity: 1;
}

</style>