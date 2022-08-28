import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import LoginView from "@/views/LoginView";
import RegisterView from "@/views/RegisterView";
import MainView from "@/views/MainView";
import GradeBookView from "@/views/GradeBookView";
import ChatUserListView from "@/views/ChatUserListView";
import ChatWithUserView from "@/views/ChatWithUserView";
import ManagePresentations from "@/views/ManagePresentations";
import ManagePresentation from "@/views/ManagePresentation";
import SlideView from "@/views/SlideView";
import AddNewSlideView from "@/views/AddNewSlideView";
import ProfileView from "@/views/ProfileView";
import ClassesManage from "@/views/ClassesManageView";
import ClassManageView from "@/views/ClassManageView";

const routes = [
  {
    path: "/",
    name: "home",
    component: HomeView,
  },
  {
    path: "/login",
    name: "login",
    component: LoginView,
  },
  {
    path: "/register",
    name: "register",
    component: RegisterView,
  },
  {
    path: "/mainView",
    name: "mainView",
    component: MainView,
  },
  {
    path: "/gradeBook",
    name: "gradeBook",
    component: GradeBookView,
  },
  {
    path: "/chatList",
    name: "chatList",
    component: ChatUserListView,
  },
  {
    path: "/chatUser/:userId",
    name: "chatUser",
    component: ChatWithUserView,
  },
  {
    path: "/managePresentations",
    name: "managePresentations",
    component: ManagePresentations,
  },
  {
    path: "/managePresentation/:presentationId",
    name: "managePresentation",
    component: ManagePresentation,
  },
  {
    path: "/managePresentation/slide/:presentationId/:slideId",
    name: "slideView",
    component: SlideView,
  },
  {
    path: "/managePresentation/addNewSlide/:presentationId",
    name: "addNewSlideView",
    component: AddNewSlideView,
  },
  {
    path: "/profile",
    name: "profile",
    component: ProfileView,
  },
  {
    path: "/manageClasses",
    name: "manageClasses",
    component: ClassesManage,
  },
  {
    path: "/manageClass/:classId",
    name: "manageClass",
    component: ClassManageView,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
