import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

import PrimeVue from "primevue/config";
import Button from "primevue/button";
import Ripple from "primevue/ripple"; //icons

import "primevue/resources/themes/saga-blue/theme.css"; //theme
import "primevue/resources/primevue.min.css"; //core css
import "primeicons/primeicons.css";

const app = createApp(App);

app.use(store);
app.use(router);
app.use(PrimeVue, { ripple: true });

// eslint-disable-next-line vue/multi-word-component-names
app.component("Button", Button);
app.directive("Ripple", Ripple);

app.mount("#app");
