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
import S3MaterialsMainTeacherView from "@/views/S3MaterialsMainTeacherView";
import S3ClassMaterialsView from "@/views/S3ClassMaterialsView";
import GradeUserView from "@/views/GradeUserView";
import StartLessonTeacherView from "@/views/StartLessonTeacherView";
import ConductLessonView from "@/views/ConductLessonView.vue";
import AttendLessonView from "@/views/AttendLessonView";
import CurrentLessonsView from "@/views/CurrentLessonsView";
import ManageUsersView from "@/views/ManageUsersView";
import ManageUserView from "@/views/ManageUserView";
import CreateTestsView from "@/views/CreateTestsView";
import CreateTestView from "@/views/CreateTestView";
import StudentTestListView from "@/views/StudentTestListView";
import StudentTakeTestView from "@/views/StudentTakeTestView";
import StudentListTestResultsView from "@/views/StudentListTestResultsView.vue";
import StudentTestResultsView from "@/views/StudentTestResultsView";
import TeacherTestResults from "@/views/TeacherTestResults";

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
  {
    path: "/classContent",
    name: "classContent",
    component: S3MaterialsMainTeacherView,
  },
  {
    path: "/classContent/:classId",
    name: "classContentFiles",
    component: S3ClassMaterialsView,
  },
  {
    path: "/gradeBook/:userId",
    name: "gradeUserView",
    component: GradeUserView,
  },
  {
    path: "/startLesson",
    name: "startLesson",
    component: StartLessonTeacherView,
  },
  {
    path: "/conductLesson/:presentationId",
    name: "conductLesson",
    component: ConductLessonView,
  },
  {
    path: "/lesson/:presentationId",
    name: "lesson",
    component: AttendLessonView,
  },
  {
    path: "/lessons",
    name: "lessons",
    component: CurrentLessonsView,
  },
  {
    path: "/manageUsers",
    name: "manageUsers",
    component: ManageUsersView,
  },
  {
    path: "/manageUsers/:userId",
    name: "manageUser",
    component: ManageUserView,
  },
  {
    path: "/createTests",
    name: "createTests",
    component: CreateTestsView,
  },
  {
    path: "/createTest/:classId",
    name: "createTest",
    component: CreateTestView,
  },
  {
    path: "/testsList/:classId",
    name: "testsList",
    component: StudentTestListView,
  },
  {
    path: "/testsList/takeTest/:testId",
    name: "takeTest",
    component: StudentTakeTestView,
  },
  {
    path: "/testResultList",
    name: "testResultList",
    component: StudentListTestResultsView,
  },
  {
    path: "/testResultList/:testId",
    name: "testResult",
    component: StudentTestResultsView,
  },
  {
    path: "/testResults",
    name: "testResults",
    component: TeacherTestResults,
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
