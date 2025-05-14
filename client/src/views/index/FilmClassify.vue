<template>
  <div class="container"  v-if="d.content.news.length > 0">
    <div class="title">
      <a :href="`/filmClassify/${d.title.id}`" class="h_active">{{ d.title.name }}</a>
      <span class="line"/>
      <router-link :to="{name:'FilmClassifySearch',params:{Pid:d.title.id}}" class="">{{ `${d.title.name}库` }}</router-link>
    </div>

    <!--影片列表展示-->
    <div class="content">
      <div class="news">
        <div class="c_nav">
          <span class="c_nav_text silver">最新上映</span>
          <router-link :to="{name:'FilmClassifySearch',params:{Pid:d.title.id,Sort:'release_stamp'}}" class="c_nav_more ">更多<b class="iconfont icon-more"/></router-link>
        </div>
        <FilmList :col="7" :list="d.content.news"/>
      </div>
      <div class="news">
        <div class="c_nav">
          <span class="c_nav_text silver">排行榜</span>
          <router-link :to="{name:'FilmClassifySearch',params:{Pid:d.title.id,Sort:'hits'}}" class="c_nav_more ">更多<b class="iconfont icon-more"/></router-link>
        </div>
        <FilmList :col="7" :list="d.content.top"/>
      </div>
      <div class="news">
        <div class="c_nav">
          <span class="c_nav_text silver">最近更新</span>
          <router-link :to="{name:'FilmClassifySearch',params:{Pid:d.title.id,Sort:'update_stamp'}}" class="c_nav_more ">更多<b class="iconfont icon-more"/></router-link>
        </div>
        <FilmList :col="7" :list="d.content.recent"/>
      </div>
    </div>

    <!--分页展示区域-->
  </div>
</template>

<script setup lang="ts">
import {ApiGet} from "../../utils/request";
import {ElMessage} from "element-plus";
import {onMounted, reactive} from "vue";
import {useRouter} from "vue-router";
import FilmList from "../../components/index/FilmList.vue";
import {Bottom} from "@element-plus/icons-vue";
import {useSiteStore} from "../../store/site";

const d = reactive({
  title: {},
  content: {
    news: [],
    top: [],
    recent: [],
  }

})

const router = useRouter()
const getFilmData = () => {
  let query = router.currentRoute.value.params
  ApiGet(`/filmClassify`, {Pid: query.pid}).then((resp: any) => {
    const siteStore=useSiteStore();
    if (resp.code === 0 ) {
      d.title = resp.data.title
      d.content = resp.data.content
      siteStore.setTitle(d.title)
      siteStore.setDes(d.content)
    } else {
      ElMessage.error({message: resp.msg, duration: 1000})
    }
  })
}

onMounted(() => {
  getFilmData()
})
</script>

<style scoped>
@import "/src/assets/css/classify.css";

.c_nav {
  display: flex;
  justify-content: space-between;
  padding: 6px;
  border-bottom: 2px solid rgba(255, 255, 255, 0.1);
}

.c_nav_text {
  font-weight: 700;
  line-height: 1.1;
}

.c_nav_more {
  border-radius: 5px;
  background: linear-gradient(#ffffff15, #ffffff1a);
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.c_nav_more:hover {
  color: var(--active--text-color);
  background-color: var(--active--bg-color);
}

.content > div {
  padding-bottom: 20px;
}

/**/
@media (min-width: 768px) {
  .c_nav{
    margin-bottom: 15px;
  }
  .c_nav_text {
    font-size: 28px;
  }
  .c_nav_more {
    font-size: 14px;
    padding: 0 15px;
    line-height: 32px;
  }

}

@media (max-width: 768px) {
  .c_nav{
    margin-bottom: 10px;
  }
  .c_nav_text {
    font-size: 20px;
    line-height: 28px;
  }
  .c_nav_more {
    font-size: 12px;
    padding: 0 10px;
    line-height: 28px;
  }
  .c_nav_more b{
    font-size: 12px;
  }
}

</style>