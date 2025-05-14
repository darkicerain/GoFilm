// store/title.js
import { defineStore } from 'pinia';

export const useSiteStore = defineStore('site', {
    state: () => ({
        title: '(╥﹏╥)',
        siteName: "FreeFilm",
        domain: "https://filmfree.email",
        logo: "https://s2.loli.net/2023/12/05/O2SEiUcMx5aWlv4.jpg",
        keyword: "在线视频, 免费观影",
        describe: "在线观影网站",
        hint: "网站升级中, 暂时无法访问 !!!"
    }),
    actions: {
        setSite(newSite){
            this.siteName=newSite.siteName
            this.domain=newSite.domain
            this.logo=newSite.logo
            this.keyword=newSite.keyword
            this.describe=newSite.describe
            this.hint=newSite.hint
            document.title=this.siteName+"-"+this.title;
            document.querySelector("meta[name='keywords']").setAttribute('content', this.keyword);
            document.querySelector("meta[name='description']").setAttribute('content', this.describe);
        },
        setTitle(newTitle) {
            this.title = newTitle;
            document.title = this.siteName+"-"+newTitle;
        },
        setKeyWords(newKey){
            this.keyword=newKey
            document.querySelector("meta[name='keywords']").setAttribute('content', this.keyword);

        },
        setDes(newDes){
            this.describe=newDes
            document.querySelector("meta[name='description']").setAttribute('content', this.describe);
        }

    }
});