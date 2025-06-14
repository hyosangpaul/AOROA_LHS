import { createRouter, createWebHistory } from "vue-router";
import IssueList from "@/pages/IssueList.vue";
import IssueDetail from "@/pages/IssueDetail.vue";
import IssueEdit from "@/pages/IssueEdit.vue"; // IssueForm을 IssueEdit으로 변경

const routes = [
  { path: "/", component: IssueList },
  { path: "/issue/:id", component: IssueDetail },
  { path: "/issue/edit/:id", component: IssueEdit }  // 수정 페이지 경로 변경
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;